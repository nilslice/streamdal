#![allow(non_snake_case)]

use anyhow::anyhow;
use extism_pdk::*;
use protobuf::EnumOrUnknown;
use protos::sp_shared::KVAction;
use protos::sp_steps_kv::{KVMode, KVStatus, KVStep, KVStepResponse};
use protos::sp_wsm::{WASMExitCode, WASMRequest, WASMResponse};
use streamdal_wasm_detective::detective;

#[host_fn]
extern "ExtismHost" {
    fn kvExists(step: Protobuf<KVStep>) -> Protobuf<KVStepResponse>;
}

#[plugin_fn]
pub fn f(Protobuf(req): Protobuf<WASMRequest>) -> FnResult<Protobuf<WASMResponse>> {
    // Validate request
    if let Err(err) = validate_wasm_request(&req) {
        return Err(WithReturnCode::new(
            anyhow!("invalid wasm request: {}", err.to_string()),
            WASMExitCode::WASM_EXIT_CODE_ERROR as i32,
        ));
    }

    // Determine which host func we'll execute (OK to unwrap since we validated)
    let kv_func = match req.step.kv().action.unwrap() {
        KVAction::KV_ACTION_EXISTS => kvExists,
        _ => {
            return Err(WithReturnCode::new(
                anyhow!("invalid action: {:?}", req.step.kv().action),
                WASMExitCode::WASM_EXIT_CODE_FALSE as i32,
            ));
        }
    };

    // Maybe update step key (if dynamic mode is specified)
    let mut step = req.step.kv().clone();

    // Handle "static" or "dynamic" mode - ie. how do we use the provided "key"?
    if req.step.kv().mode == EnumOrUnknown::from(KVMode::KV_MODE_DYNAMIC) {
        // Lookup what the actual key will be by looking at the value in the JSON path
        match detective::parse_field(req.input_payload.as_slice(), &step.key) {
            Ok(key_contents) => step.key = key_contents.to_string(),
            Err(err) => {
                return Err(WithReturnCode::new(
                    anyhow!("unable to complete dynamic key lookup: {}", err.to_string()),
                    WASMExitCode::WASM_EXIT_CODE_FALSE as i32, // TODO: This should be exit code from parse call
                ));
            }
        }
    }

    let Protobuf(kv_step_response): Protobuf<KVStepResponse> = unsafe { kv_func(Protobuf(step))? };

    // Validate KVResp
    if let Err(err) = validate_kv_step_response(&kv_step_response) {
        return Err(WithReturnCode::new(
            anyhow!("unable to validate kv step response: {}", err.to_string()),
            WASMExitCode::WASM_EXIT_CODE_FALSE as i32,
        ));
    }

    // Write + return response
    let wasm_exit_code = match kv_step_response.status.unwrap() {
        KVStatus::KV_STATUS_SUCCESS => WASMExitCode::WASM_EXIT_CODE_TRUE,
        KVStatus::KV_STATUS_FAILURE => WASMExitCode::WASM_EXIT_CODE_FALSE,
        KVStatus::KV_STATUS_ERROR => WASMExitCode::WASM_EXIT_CODE_ERROR,
        _ => WASMExitCode::WASM_EXIT_CODE_ERROR,
    };

    // This may have been a KVGet that might produce a result - we should populate it (if it's there)
    // NOTE: Doing this ugly thing to avoid "value does not live long enough" errors
    if kv_step_response.value.is_some() {
        Ok(Protobuf(common::into_response(
            Some(req.input_payload.as_slice()),
            Some(kv_step_response.value.unwrap().as_slice()),
            None,
            wasm_exit_code,
            format!("kv step response: {:?}", kv_step_response.message),
        )))
    } else {
        Ok(Protobuf(common::into_response(
            Some(req.input_payload.as_slice()),
            None,
            None,
            wasm_exit_code,
            format!("kv step response: {:?}", kv_step_response.message),
        )))
    }
}

fn validate_kv_step_response(kv_step_response: &KVStepResponse) -> Result<(), String> {
    match kv_step_response.status.enum_value() {
        Ok(status) => {
            if status == KVStatus::KV_STATUS_UNSET {
                return Err("KV status must be set".to_string());
            }
        }
        Err(_) => return Err("unable to get KV status".to_string()),
    };

    Ok(())
}

fn validate_wasm_request(req: &WASMRequest) -> Result<(), String> {
    if !req.step.has_kv() {
        return Err("kv step must be set".to_string());
    }

    // Key must be set for all actions
    if req.step.kv().key.is_empty() {
        return Err("key must be set".to_string());
    }

    // Mode must be set for all
    if req.step.kv().mode == EnumOrUnknown::from(KVMode::KV_MODE_UNSET) {
        return Err("mode must be set".to_string());
    }

    // Action must be set for all
    if req.step.kv().action == EnumOrUnknown::from(KVAction::KV_ACTION_UNSET) {
        return Err("action must be set".to_string());
    }

    // Value must be set for CREATE action
    if req.step.kv().action == EnumOrUnknown::from(KVAction::KV_ACTION_CREATE)
        && req.step.kv().value.is_none()
    {
        return Err("value must be set for create action".to_string());
    }

    Ok(())
}
