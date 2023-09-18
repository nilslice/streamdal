"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Internal = exports.SendSchemaRequest = exports.GetAttachCommandsByServiceResponse = exports.GetAttachCommandsByServiceRequest = exports.DeregisterRequest = exports.RegisterRequest = exports.MetricsRequest = exports.NotifyRequest = exports.HeartbeatRequest = exports.NewAudienceRequest = void 0;
// @generated by protobuf-ts 2.9.0 with parameter long_type_string
// @generated from protobuf file "sp_internal.proto" (package "protos", syntax proto3)
// tslint:disable
const sp_common_1 = require("./sp_common");
const sp_common_2 = require("./sp_common");
const runtime_rpc_1 = require("@protobuf-ts/runtime-rpc");
const runtime_1 = require("@protobuf-ts/runtime");
const runtime_2 = require("@protobuf-ts/runtime");
const runtime_3 = require("@protobuf-ts/runtime");
const runtime_4 = require("@protobuf-ts/runtime");
const runtime_5 = require("@protobuf-ts/runtime");
const sp_common_3 = require("./sp_common");
const sp_command_1 = require("./sp_command");
const sp_info_1 = require("./sp_info");
const sp_common_4 = require("./sp_common");
const sp_common_5 = require("./sp_common");
// @generated message type with reflection information, may provide speed optimized methods
class NewAudienceRequest$Type extends runtime_5.MessageType {
    constructor() {
        super("protos.NewAudienceRequest", [
            { no: 1, name: "session_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "audience", kind: "message", T: () => sp_common_5.Audience }
        ]);
    }
    create(value) {
        const message = { sessionId: "" };
        globalThis.Object.defineProperty(message, runtime_4.MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            (0, runtime_3.reflectionMergePartial)(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string session_id */ 1:
                    message.sessionId = reader.string();
                    break;
                case /* protos.Audience audience */ 2:
                    message.audience = sp_common_5.Audience.internalBinaryRead(reader, reader.uint32(), options, message.audience);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? runtime_2.UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* string session_id = 1; */
        if (message.sessionId !== "")
            writer.tag(1, runtime_1.WireType.LengthDelimited).string(message.sessionId);
        /* protos.Audience audience = 2; */
        if (message.audience)
            sp_common_5.Audience.internalBinaryWrite(message.audience, writer.tag(2, runtime_1.WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? runtime_2.UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.NewAudienceRequest
 */
exports.NewAudienceRequest = new NewAudienceRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class HeartbeatRequest$Type extends runtime_5.MessageType {
    constructor() {
        super("protos.HeartbeatRequest", [
            { no: 1, name: "session_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value) {
        const message = { sessionId: "" };
        globalThis.Object.defineProperty(message, runtime_4.MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            (0, runtime_3.reflectionMergePartial)(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string session_id */ 1:
                    message.sessionId = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? runtime_2.UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* string session_id = 1; */
        if (message.sessionId !== "")
            writer.tag(1, runtime_1.WireType.LengthDelimited).string(message.sessionId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? runtime_2.UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.HeartbeatRequest
 */
exports.HeartbeatRequest = new HeartbeatRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class NotifyRequest$Type extends runtime_5.MessageType {
    constructor() {
        super("protos.NotifyRequest", [
            { no: 1, name: "pipeline_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "step_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "audience", kind: "message", T: () => sp_common_5.Audience },
            { no: 4, name: "occurred_at_unix_ts_utc", kind: "scalar", T: 3 /*ScalarType.INT64*/ }
        ]);
    }
    create(value) {
        const message = { pipelineId: "", stepName: "", occurredAtUnixTsUtc: "0" };
        globalThis.Object.defineProperty(message, runtime_4.MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            (0, runtime_3.reflectionMergePartial)(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string pipeline_id */ 1:
                    message.pipelineId = reader.string();
                    break;
                case /* string step_name */ 2:
                    message.stepName = reader.string();
                    break;
                case /* protos.Audience audience */ 3:
                    message.audience = sp_common_5.Audience.internalBinaryRead(reader, reader.uint32(), options, message.audience);
                    break;
                case /* int64 occurred_at_unix_ts_utc */ 4:
                    message.occurredAtUnixTsUtc = reader.int64().toString();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? runtime_2.UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* string pipeline_id = 1; */
        if (message.pipelineId !== "")
            writer.tag(1, runtime_1.WireType.LengthDelimited).string(message.pipelineId);
        /* string step_name = 2; */
        if (message.stepName !== "")
            writer.tag(2, runtime_1.WireType.LengthDelimited).string(message.stepName);
        /* protos.Audience audience = 3; */
        if (message.audience)
            sp_common_5.Audience.internalBinaryWrite(message.audience, writer.tag(3, runtime_1.WireType.LengthDelimited).fork(), options).join();
        /* int64 occurred_at_unix_ts_utc = 4; */
        if (message.occurredAtUnixTsUtc !== "0")
            writer.tag(4, runtime_1.WireType.Varint).int64(message.occurredAtUnixTsUtc);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? runtime_2.UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.NotifyRequest
 */
exports.NotifyRequest = new NotifyRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MetricsRequest$Type extends runtime_5.MessageType {
    constructor() {
        super("protos.MetricsRequest", [
            { no: 1, name: "metrics", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => sp_common_4.Metric }
        ]);
    }
    create(value) {
        const message = { metrics: [] };
        globalThis.Object.defineProperty(message, runtime_4.MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            (0, runtime_3.reflectionMergePartial)(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated protos.Metric metrics */ 1:
                    message.metrics.push(sp_common_4.Metric.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? runtime_2.UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* repeated protos.Metric metrics = 1; */
        for (let i = 0; i < message.metrics.length; i++)
            sp_common_4.Metric.internalBinaryWrite(message.metrics[i], writer.tag(1, runtime_1.WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? runtime_2.UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.MetricsRequest
 */
exports.MetricsRequest = new MetricsRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RegisterRequest$Type extends runtime_5.MessageType {
    constructor() {
        super("protos.RegisterRequest", [
            { no: 1, name: "service_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "session_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "client_info", kind: "message", T: () => sp_info_1.ClientInfo },
            { no: 4, name: "audiences", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => sp_common_5.Audience },
            { no: 5, name: "dry_run", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
    create(value) {
        const message = { serviceName: "", sessionId: "", audiences: [], dryRun: false };
        globalThis.Object.defineProperty(message, runtime_4.MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            (0, runtime_3.reflectionMergePartial)(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string service_name */ 1:
                    message.serviceName = reader.string();
                    break;
                case /* string session_id */ 2:
                    message.sessionId = reader.string();
                    break;
                case /* protos.ClientInfo client_info */ 3:
                    message.clientInfo = sp_info_1.ClientInfo.internalBinaryRead(reader, reader.uint32(), options, message.clientInfo);
                    break;
                case /* repeated protos.Audience audiences */ 4:
                    message.audiences.push(sp_common_5.Audience.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* bool dry_run */ 5:
                    message.dryRun = reader.bool();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? runtime_2.UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* string service_name = 1; */
        if (message.serviceName !== "")
            writer.tag(1, runtime_1.WireType.LengthDelimited).string(message.serviceName);
        /* string session_id = 2; */
        if (message.sessionId !== "")
            writer.tag(2, runtime_1.WireType.LengthDelimited).string(message.sessionId);
        /* protos.ClientInfo client_info = 3; */
        if (message.clientInfo)
            sp_info_1.ClientInfo.internalBinaryWrite(message.clientInfo, writer.tag(3, runtime_1.WireType.LengthDelimited).fork(), options).join();
        /* repeated protos.Audience audiences = 4; */
        for (let i = 0; i < message.audiences.length; i++)
            sp_common_5.Audience.internalBinaryWrite(message.audiences[i], writer.tag(4, runtime_1.WireType.LengthDelimited).fork(), options).join();
        /* bool dry_run = 5; */
        if (message.dryRun !== false)
            writer.tag(5, runtime_1.WireType.Varint).bool(message.dryRun);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? runtime_2.UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.RegisterRequest
 */
exports.RegisterRequest = new RegisterRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class DeregisterRequest$Type extends runtime_5.MessageType {
    constructor() {
        super("protos.DeregisterRequest", [
            { no: 1, name: "service_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "session_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value) {
        const message = { serviceName: "", sessionId: "" };
        globalThis.Object.defineProperty(message, runtime_4.MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            (0, runtime_3.reflectionMergePartial)(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string service_name */ 1:
                    message.serviceName = reader.string();
                    break;
                case /* string session_id */ 2:
                    message.sessionId = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? runtime_2.UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* string service_name = 1; */
        if (message.serviceName !== "")
            writer.tag(1, runtime_1.WireType.LengthDelimited).string(message.serviceName);
        /* string session_id = 2; */
        if (message.sessionId !== "")
            writer.tag(2, runtime_1.WireType.LengthDelimited).string(message.sessionId);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? runtime_2.UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.DeregisterRequest
 */
exports.DeregisterRequest = new DeregisterRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetAttachCommandsByServiceRequest$Type extends runtime_5.MessageType {
    constructor() {
        super("protos.GetAttachCommandsByServiceRequest", [
            { no: 1, name: "service_name", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value) {
        const message = { serviceName: "" };
        globalThis.Object.defineProperty(message, runtime_4.MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            (0, runtime_3.reflectionMergePartial)(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string service_name */ 1:
                    message.serviceName = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? runtime_2.UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* string service_name = 1; */
        if (message.serviceName !== "")
            writer.tag(1, runtime_1.WireType.LengthDelimited).string(message.serviceName);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? runtime_2.UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.GetAttachCommandsByServiceRequest
 */
exports.GetAttachCommandsByServiceRequest = new GetAttachCommandsByServiceRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class GetAttachCommandsByServiceResponse$Type extends runtime_5.MessageType {
    constructor() {
        super("protos.GetAttachCommandsByServiceResponse", [
            { no: 1, name: "active", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => sp_command_1.Command },
            { no: 2, name: "paused", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => sp_command_1.Command }
        ]);
    }
    create(value) {
        const message = { active: [], paused: [] };
        globalThis.Object.defineProperty(message, runtime_4.MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            (0, runtime_3.reflectionMergePartial)(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* repeated protos.Command active */ 1:
                    message.active.push(sp_command_1.Command.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* repeated protos.Command paused */ 2:
                    message.paused.push(sp_command_1.Command.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? runtime_2.UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* repeated protos.Command active = 1; */
        for (let i = 0; i < message.active.length; i++)
            sp_command_1.Command.internalBinaryWrite(message.active[i], writer.tag(1, runtime_1.WireType.LengthDelimited).fork(), options).join();
        /* repeated protos.Command paused = 2; */
        for (let i = 0; i < message.paused.length; i++)
            sp_command_1.Command.internalBinaryWrite(message.paused[i], writer.tag(2, runtime_1.WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? runtime_2.UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.GetAttachCommandsByServiceResponse
 */
exports.GetAttachCommandsByServiceResponse = new GetAttachCommandsByServiceResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SendSchemaRequest$Type extends runtime_5.MessageType {
    constructor() {
        super("protos.SendSchemaRequest", [
            { no: 1, name: "schema", kind: "message", T: () => sp_common_3.Schema }
        ]);
    }
    create(value) {
        const message = {};
        globalThis.Object.defineProperty(message, runtime_4.MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            (0, runtime_3.reflectionMergePartial)(this, message, value);
        return message;
    }
    internalBinaryRead(reader, length, options, target) {
        let message = target !== null && target !== void 0 ? target : this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* protos.Schema schema */ 1:
                    message.schema = sp_common_3.Schema.internalBinaryRead(reader, reader.uint32(), options, message.schema);
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? runtime_2.UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message, writer, options) {
        /* protos.Schema schema = 1; */
        if (message.schema)
            sp_common_3.Schema.internalBinaryWrite(message.schema, writer.tag(1, runtime_1.WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? runtime_2.UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message protos.SendSchemaRequest
 */
exports.SendSchemaRequest = new SendSchemaRequest$Type();
/**
 * @generated ServiceType for protobuf service protos.Internal
 */
exports.Internal = new runtime_rpc_1.ServiceType("protos.Internal", [
    { name: "Register", serverStreaming: true, options: {}, I: exports.RegisterRequest, O: sp_command_1.Command },
    { name: "NewAudience", options: {}, I: exports.NewAudienceRequest, O: sp_common_2.StandardResponse },
    { name: "Heartbeat", options: {}, I: exports.HeartbeatRequest, O: sp_common_2.StandardResponse },
    { name: "Notify", options: {}, I: exports.NotifyRequest, O: sp_common_2.StandardResponse },
    { name: "Metrics", options: {}, I: exports.MetricsRequest, O: sp_common_2.StandardResponse },
    { name: "GetAttachCommandsByService", options: {}, I: exports.GetAttachCommandsByServiceRequest, O: exports.GetAttachCommandsByServiceResponse },
    { name: "SendTail", clientStreaming: true, options: {}, I: sp_common_1.TailResponse, O: sp_common_2.StandardResponse },
    { name: "SendSchema", options: {}, I: exports.SendSchemaRequest, O: sp_common_2.StandardResponse }
]);
