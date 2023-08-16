// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size
// @generated from protobuf file "internal.proto" (package "protos", syntax proto3)
// tslint:disable
import { StandardResponse } from "./common.js";
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { MessageType } from "@protobuf-ts/runtime";
import { Command } from "./command.js";
import { ClientInfo } from "./info.js";
import { Audience } from "./common.js";
// @generated message type with reflection information, may provide speed optimized methods
class NewAudienceRequest$Type extends MessageType {
    constructor() {
        super("protos.NewAudienceRequest", [
            { no: 1, name: "session_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "audience", kind: "message", T: () => Audience }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.NewAudienceRequest
 */
export const NewAudienceRequest = new NewAudienceRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class HeartbeatRequest$Type extends MessageType {
    constructor() {
        super("protos.HeartbeatRequest", [
            { no: 1, name: "session_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.HeartbeatRequest
 */
export const HeartbeatRequest = new HeartbeatRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class NotifyRequest$Type extends MessageType {
    constructor() {
        super("protos.NotifyRequest", [
            { no: 1, name: "pipeline_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "step_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "audience", kind: "message", T: () => Audience },
            { no: 4, name: "occurred_at_unix_ts_utc", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.NotifyRequest
 */
export const NotifyRequest = new NotifyRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Metrics$Type extends MessageType {
    constructor() {
        super("protos.Metrics", [
            { no: 1, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "labels", kind: "map", K: 9 /*ScalarType.STRING*/, V: { kind: "scalar", T: 9 /*ScalarType.STRING*/ } },
            { no: 3, name: "value", kind: "scalar", T: 1 /*ScalarType.DOUBLE*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.Metrics
 */
export const Metrics = new Metrics$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MetricsRequest$Type extends MessageType {
    constructor() {
        super("protos.MetricsRequest", [
            { no: 1, name: "metrics", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Metrics }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.MetricsRequest
 */
export const MetricsRequest = new MetricsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RegisterRequest$Type extends MessageType {
    constructor() {
        super("protos.RegisterRequest", [
            { no: 1, name: "service_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "session_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "client_info", kind: "message", T: () => ClientInfo },
            { no: 4, name: "audiences", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Audience },
            { no: 5, name: "dry_run", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.RegisterRequest
 */
export const RegisterRequest = new RegisterRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeregisterRequest$Type extends MessageType {
    constructor() {
        super("protos.DeregisterRequest", [
            { no: 1, name: "service_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "session_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.DeregisterRequest
 */
export const DeregisterRequest = new DeregisterRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetAttachCommandsByServiceRequest$Type extends MessageType {
    constructor() {
        super("protos.GetAttachCommandsByServiceRequest", [
            { no: 1, name: "service_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.GetAttachCommandsByServiceRequest
 */
export const GetAttachCommandsByServiceRequest = new GetAttachCommandsByServiceRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetAttachCommandsByServiceResponse$Type extends MessageType {
    constructor() {
        super("protos.GetAttachCommandsByServiceResponse", [
            { no: 1, name: "active", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Command },
            { no: 2, name: "paused", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Command }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.GetAttachCommandsByServiceResponse
 */
export const GetAttachCommandsByServiceResponse = new GetAttachCommandsByServiceResponse$Type();
/**
 * @generated ServiceType for protobuf service protos.Internal
 */
export const Internal = new ServiceType("protos.Internal", [
    { name: "Register", serverStreaming: true, options: {}, I: RegisterRequest, O: Command },
    { name: "NewAudience", options: {}, I: NewAudienceRequest, O: StandardResponse },
    { name: "Heartbeat", options: {}, I: HeartbeatRequest, O: StandardResponse },
    { name: "Notify", options: {}, I: NotifyRequest, O: StandardResponse },
    { name: "Metrics", options: {}, I: MetricsRequest, O: StandardResponse },
    { name: "GetAttachCommandsByService", options: {}, I: GetAttachCommandsByServiceRequest, O: GetAttachCommandsByServiceResponse }
]);
//# sourceMappingURL=internal.js.map