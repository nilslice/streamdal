// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size
// @generated from protobuf file "sp_bus.proto" (package "protos", syntax proto3)
// tslint:disable
import { MessageType } from "@protobuf-ts/runtime";
import { TailRequest } from "./sp_common.ts";
import { DeleteAudienceRequest } from "./sp_external.ts";
import { KVRequest } from "./sp_kv.ts";
import { MetricsRequest } from "./sp_internal.ts";
import { ResumePipelineRequest } from "./sp_external.ts";
import { PausePipelineRequest } from "./sp_external.ts";
import { DetachPipelineRequest } from "./sp_external.ts";
import { AttachPipelineRequest } from "./sp_external.ts";
import { UpdatePipelineRequest } from "./sp_external.ts";
import { DeletePipelineRequest } from "./sp_external.ts";
import { CreatePipelineRequest } from "./sp_external.ts";
import { DeregisterRequest } from "./sp_internal.ts";
import { RegisterRequest } from "./sp_internal.ts";
/**
 * Type used by `snitch-server` for broadcasting events to other snitch nodes
 *
 * @generated from protobuf message protos.BusEvent
 */
export interface BusEvent {
    /**
     * @generated from protobuf field: string source = 1;
     */
    source: string;
    /**
     * @generated from protobuf oneof: event
     */
    event: {
        oneofKind: "registerRequest";
        /**
         * @generated from protobuf field: protos.RegisterRequest register_request = 100;
         */
        registerRequest: RegisterRequest;
    } | {
        oneofKind: "deregisterRequest";
        /**
         * @generated from protobuf field: protos.DeregisterRequest deregister_request = 101;
         */
        deregisterRequest: DeregisterRequest;
    } | {
        oneofKind: "createPipelineRequest";
        /**
         * @generated from protobuf field: protos.CreatePipelineRequest create_pipeline_request = 102;
         */
        createPipelineRequest: CreatePipelineRequest;
    } | {
        oneofKind: "deletePipelineRequest";
        /**
         * @generated from protobuf field: protos.DeletePipelineRequest delete_pipeline_request = 103;
         */
        deletePipelineRequest: DeletePipelineRequest;
    } | {
        oneofKind: "updatePipelineRequest";
        /**
         * @generated from protobuf field: protos.UpdatePipelineRequest update_pipeline_request = 104;
         */
        updatePipelineRequest: UpdatePipelineRequest;
    } | {
        oneofKind: "attachPipelineRequest";
        /**
         * @generated from protobuf field: protos.AttachPipelineRequest attach_pipeline_request = 105;
         */
        attachPipelineRequest: AttachPipelineRequest;
    } | {
        oneofKind: "detachPipelineRequest";
        /**
         * @generated from protobuf field: protos.DetachPipelineRequest detach_pipeline_request = 106;
         */
        detachPipelineRequest: DetachPipelineRequest;
    } | {
        oneofKind: "pausePipelineRequest";
        /**
         * @generated from protobuf field: protos.PausePipelineRequest pause_pipeline_request = 107;
         */
        pausePipelineRequest: PausePipelineRequest;
    } | {
        oneofKind: "resumePipelineRequest";
        /**
         * @generated from protobuf field: protos.ResumePipelineRequest resume_pipeline_request = 108;
         */
        resumePipelineRequest: ResumePipelineRequest;
    } | {
        oneofKind: "metricsRequest";
        /**
         * @generated from protobuf field: protos.MetricsRequest metrics_request = 109;
         */
        metricsRequest: MetricsRequest;
    } | {
        oneofKind: "kvRequest";
        /**
         * @generated from protobuf field: protos.KVRequest kv_request = 110;
         */
        kvRequest: KVRequest;
    } | {
        oneofKind: "deleteAudienceRequest";
        /**
         * @generated from protobuf field: protos.DeleteAudienceRequest delete_audience_request = 111;
         */
        deleteAudienceRequest: DeleteAudienceRequest;
    } | {
        oneofKind: "tailRequest";
        /**
         * @generated from protobuf field: protos.TailRequest tail_request = 112;
         */
        tailRequest: TailRequest;
    } | {
        oneofKind: undefined;
    };
    /**
     * All gRPC metadata is stored in ctx; when request goes outside of gRPC
     * bounds, we will translate ctx metadata into this field.
     *
     * Example:
     * 1. Request comes into snitch-server via external gRPC to set new pipeline
     * 2. snitch-server has to send SetPipeline cmd to SDK via gRPC - it passes
     *    on original metadata in request.
     * 3. snitch-server has to broadcast SetPipeline cmd to other services via bus
     * 4. Since this is not a gRPC call, snitch-server translates ctx metadata to
     *    this field and includes it in the bus event.
     *
     * @generated from protobuf field: map<string, string> _metadata = 1000;
     */
    Metadata: {
        [key: string]: string;
    }; // protolint:disable:this FIELD_NAMES_LOWER_SNAKE_CASE
}
// @generated message type with reflection information, may provide speed optimized methods
class BusEvent$Type extends MessageType<BusEvent> {
    constructor() {
        super("protos.BusEvent", [
            { no: 1, name: "source", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 100, name: "register_request", kind: "message", oneof: "event", T: () => RegisterRequest },
            { no: 101, name: "deregister_request", kind: "message", oneof: "event", T: () => DeregisterRequest },
            { no: 102, name: "create_pipeline_request", kind: "message", oneof: "event", T: () => CreatePipelineRequest },
            { no: 103, name: "delete_pipeline_request", kind: "message", oneof: "event", T: () => DeletePipelineRequest },
            { no: 104, name: "update_pipeline_request", kind: "message", oneof: "event", T: () => UpdatePipelineRequest },
            { no: 105, name: "attach_pipeline_request", kind: "message", oneof: "event", T: () => AttachPipelineRequest },
            { no: 106, name: "detach_pipeline_request", kind: "message", oneof: "event", T: () => DetachPipelineRequest },
            { no: 107, name: "pause_pipeline_request", kind: "message", oneof: "event", T: () => PausePipelineRequest },
            { no: 108, name: "resume_pipeline_request", kind: "message", oneof: "event", T: () => ResumePipelineRequest },
            { no: 109, name: "metrics_request", kind: "message", oneof: "event", T: () => MetricsRequest },
            { no: 110, name: "kv_request", kind: "message", oneof: "event", T: () => KVRequest },
            { no: 111, name: "delete_audience_request", kind: "message", oneof: "event", T: () => DeleteAudienceRequest },
            { no: 112, name: "tail_request", kind: "message", oneof: "event", T: () => TailRequest },
            { no: 1000, name: "_metadata", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.BusEvent
 */
export const BusEvent = new BusEvent$Type();
