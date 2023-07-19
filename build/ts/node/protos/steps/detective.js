// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size
// @generated from protobuf file "steps/detective.proto" (package "protos.steps", syntax proto3)
// tslint:disable
function _assert_this_initialized(self) {
    if (self === void 0) {
        throw new ReferenceError("this hasn't been initialised - super() hasn't been called");
    }
    return self;
}
function _class_call_check(instance, Constructor) {
    if (!(instance instanceof Constructor)) {
        throw new TypeError("Cannot call a class as a function");
    }
}
function _get_prototype_of(o) {
    _get_prototype_of = Object.setPrototypeOf ? Object.getPrototypeOf : function getPrototypeOf(o) {
        return o.__proto__ || Object.getPrototypeOf(o);
    };
    return _get_prototype_of(o);
}
function _inherits(subClass, superClass) {
    if (typeof superClass !== "function" && superClass !== null) {
        throw new TypeError("Super expression must either be null or a function");
    }
    subClass.prototype = Object.create(superClass && superClass.prototype, {
        constructor: {
            value: subClass,
            writable: true,
            configurable: true
        }
    });
    if (superClass) _set_prototype_of(subClass, superClass);
}
function _possible_constructor_return(self, call) {
    if (call && (_type_of(call) === "object" || typeof call === "function")) {
        return call;
    }
    return _assert_this_initialized(self);
}
function _set_prototype_of(o, p) {
    _set_prototype_of = Object.setPrototypeOf || function setPrototypeOf(o, p) {
        o.__proto__ = p;
        return o;
    };
    return _set_prototype_of(o, p);
}
function _type_of(obj) {
    "@swc/helpers - typeof";
    return obj && typeof Symbol !== "undefined" && obj.constructor === Symbol ? "symbol" : typeof obj;
}
function _is_native_reflect_construct() {
    if (typeof Reflect === "undefined" || !Reflect.construct) return false;
    if (Reflect.construct.sham) return false;
    if (typeof Proxy === "function") return true;
    try {
        Boolean.prototype.valueOf.call(Reflect.construct(Boolean, [], function() {}));
        return true;
    } catch (e) {
        return false;
    }
}
function _create_super(Derived) {
    var hasNativeReflectConstruct = _is_native_reflect_construct();
    return function _createSuperInternal() {
        var Super = _get_prototype_of(Derived), result;
        if (hasNativeReflectConstruct) {
            var NewTarget = _get_prototype_of(this).constructor;
            result = Reflect.construct(Super, arguments, NewTarget);
        } else {
            result = Super.apply(this, arguments);
        }
        return _possible_constructor_return(this, result);
    };
}
import { MessageType } from "@protobuf-ts/runtime";
export var DetectiveType;
(function(DetectiveType) {
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_UNKNOWN = 0;
     */ "UNKNOWN"] = 0] = "UNKNOWN";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_IS_EMPTY = 1000;
     */ "IS_EMPTY"] = 1000] = "IS_EMPTY";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_HAS_FIELD = 1001;
     */ "HAS_FIELD"] = 1001] = "HAS_FIELD";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_IS_TYPE = 1002;
     */ "IS_TYPE"] = 1002] = "IS_TYPE";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_STRING_CONTAINS_ANY = 1003;
     */ "STRING_CONTAINS_ANY"] = 1003] = "STRING_CONTAINS_ANY";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_STRING_CONTAINS_ALL = 1004;
     */ "STRING_CONTAINS_ALL"] = 1004] = "STRING_CONTAINS_ALL";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_STRING_EQUAL = 1005;
     */ "STRING_EQUAL"] = 1005] = "STRING_EQUAL";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_IPV4_ADDRESS = 1006;
     */ "IPV4_ADDRESS"] = 1006] = "IPV4_ADDRESS";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_IPV6_ADDRESS = 1007;
     */ "IPV6_ADDRESS"] = 1007] = "IPV6_ADDRESS";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_MAC_ADDRESS = 1008;
     */ "MAC_ADDRESS"] = 1008] = "MAC_ADDRESS";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_REGEX = 1009;
     */ "REGEX"] = 1009] = "REGEX";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_TIMESTAMP_RFC3339 = 1010;
     */ "TIMESTAMP_RFC3339"] = 1010] = "TIMESTAMP_RFC3339";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_TIMESTAMP_UNIX_NANO = 1011;
     */ "TIMESTAMP_UNIX_NANO"] = 1011] = "TIMESTAMP_UNIX_NANO";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_TIMESTAMP_UNIX = 1012;
     */ "TIMESTAMP_UNIX"] = 1012] = "TIMESTAMP_UNIX";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_BOOLEAN_TRUE = 1013;
     */ "BOOLEAN_TRUE"] = 1013] = "BOOLEAN_TRUE";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_BOOLEAN_FALSE = 1014;
     */ "BOOLEAN_FALSE"] = 1014] = "BOOLEAN_FALSE";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_UUID = 1015;
     */ "UUID"] = 1015] = "UUID";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_URL = 1016;
     */ "URL"] = 1016] = "URL";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_HOSTNAME = 1017;
     */ "HOSTNAME"] = 1017] = "HOSTNAME";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_STRING_LENGTH_MIN = 1018;
     */ "STRING_LENGTH_MIN"] = 1018] = "STRING_LENGTH_MIN";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_STRING_LENGTH_MAX = 1019;
     */ "STRING_LENGTH_MAX"] = 1019] = "STRING_LENGTH_MAX";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_STRING_LENGTH_RANGE = 1020;
     */ "STRING_LENGTH_RANGE"] = 1020] = "STRING_LENGTH_RANGE";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_SEMVER = 2021;
     */ "SEMVER"] = 2021] = "SEMVER";
    DetectiveType[DetectiveType[/**
     * / Payloads containing values with any PII - runs all PII matchers
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_ANY = 2000;
     */ "PII_ANY"] = 2000] = "PII_ANY";
    DetectiveType[DetectiveType[/**
     * Payloads containing values with a credit card number
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_CREDIT_CARD = 2001;
     */ "PII_CREDIT_CARD"] = 2001] = "PII_CREDIT_CARD";
    DetectiveType[DetectiveType[/**
     * Payloads containing values with a social security number
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_SSN = 2002;
     */ "PII_SSN"] = 2002] = "PII_SSN";
    DetectiveType[DetectiveType[/**
     * Payloads containing values with an email address
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_EMAIL = 2003;
     */ "PII_EMAIL"] = 2003] = "PII_EMAIL";
    DetectiveType[DetectiveType[/**
     * Payloads containing values with a phone number
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_PHONE = 2004;
     */ "PII_PHONE"] = 2004] = "PII_PHONE";
    DetectiveType[DetectiveType[/**
     * Payloads containing values with a driver's license
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_DRIVER_LICENSE = 2005;
     */ "PII_DRIVER_LICENSE"] = 2005] = "PII_DRIVER_LICENSE";
    DetectiveType[DetectiveType[/**
     * Payloads containing values with a passport ID
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_PASSPORT_ID = 2006;
     */ "PII_PASSPORT_ID"] = 2006] = "PII_PASSPORT_ID";
    DetectiveType[DetectiveType[/**
     * Payloads containing values with a VIN number
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_VIN_NUMBER = 2007;
     */ "PII_VIN_NUMBER"] = 2007] = "PII_VIN_NUMBER";
    DetectiveType[DetectiveType[/**
     * Payloads containing values with various serial number formats
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_SERIAL_NUMBER = 2008;
     */ "PII_SERIAL_NUMBER"] = 2008] = "PII_SERIAL_NUMBER";
    DetectiveType[DetectiveType[/**
     * Payloads containing fields named "login", "username", "user", "userid", "user_id", "user", "password", "pass", "passwd", "pwd"
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_LOGIN = 2009;
     */ "PII_LOGIN"] = 2009] = "PII_LOGIN";
    DetectiveType[DetectiveType[/**
     * Payloads containing fields named "taxpayer_id", "tax_id", "taxpayerid", "taxid"
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_TAXPAYER_ID = 2010;
     */ "PII_TAXPAYER_ID"] = 2010] = "PII_TAXPAYER_ID";
    DetectiveType[DetectiveType[/**
     * Payloads containing fields named "address", "street", "city", "state", "zip", "zipcode", "zip_code", "country"
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_ADDRESS = 2011;
     */ "PII_ADDRESS"] = 2011] = "PII_ADDRESS";
    DetectiveType[DetectiveType[/**
     * Payloads containing fields named "signature", "signature_image", "signature_image_url", "signature_image_uri"
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_SIGNATURE = 2012;
     */ "PII_SIGNATURE"] = 2012] = "PII_SIGNATURE";
    DetectiveType[DetectiveType[/**
     * Payloads containing values that contain GPS data or coordinates like "lat", "lon", "latitude", "longitude"
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_GEOLOCATION = 2013;
     */ "PII_GEOLOCATION"] = 2013] = "PII_GEOLOCATION";
    DetectiveType[DetectiveType[/**
     * Payloads containing fields like "school", "university", "college", "education"
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_EDUCATION = 2014;
     */ "PII_EDUCATION"] = 2014] = "PII_EDUCATION";
    DetectiveType[DetectiveType[/**
     * Payloads containing fields like "account", "bank", "credit", "debit", "financial", "finance"
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_FINANCIAL = 2015;
     */ "PII_FINANCIAL"] = 2015] = "PII_FINANCIAL";
    DetectiveType[DetectiveType[/**
     * Payloads containing fields like "patient", "health", "healthcare", "health care", "medical"
     *
     * @generated from protobuf enum value: DETECTIVE_TYPE_PII_HEALTH = 2016;
     */ "PII_HEALTH"] = 2016] = "PII_HEALTH";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_NUMERIC_EQUAL_TO = 3000;
     */ "NUMERIC_EQUAL_TO"] = 3000] = "NUMERIC_EQUAL_TO";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_NUMERIC_GREATER_THAN = 3001;
     */ "NUMERIC_GREATER_THAN"] = 3001] = "NUMERIC_GREATER_THAN";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_NUMERIC_GREATER_EQUAL = 3002;
     */ "NUMERIC_GREATER_EQUAL"] = 3002] = "NUMERIC_GREATER_EQUAL";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_NUMERIC_LESS_THAN = 3003;
     */ "NUMERIC_LESS_THAN"] = 3003] = "NUMERIC_LESS_THAN";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_NUMERIC_LESS_EQUAL = 3004;
     */ "NUMERIC_LESS_EQUAL"] = 3004] = "NUMERIC_LESS_EQUAL";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_NUMERIC_RANGE = 3005;
     */ "NUMERIC_RANGE"] = 3005] = "NUMERIC_RANGE";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_NUMERIC_MIN = 3006;
     */ "NUMERIC_MIN"] = 3006] = "NUMERIC_MIN";
    DetectiveType[DetectiveType[/**
     * @generated from protobuf enum value: DETECTIVE_TYPE_NUMERIC_MAX = 3007;
     */ "NUMERIC_MAX"] = 3007] = "NUMERIC_MAX";
})(DetectiveType || (DetectiveType = {}));
// @generated message type with reflection information, may provide speed optimized methods
var DetectiveStep$Type = /*#__PURE__*/ function(MessageType1) {
    "use strict";
    _inherits(DetectiveStep$Type, MessageType1);
    var _super = _create_super(DetectiveStep$Type);
    function DetectiveStep$Type() {
        _class_call_check(this, DetectiveStep$Type);
        return _super.call(this, "protos.steps.DetectiveStep", [
            {
                no: 1,
                name: "path",
                kind: "scalar",
                T: 9 /*ScalarType.STRING*/ 
            },
            {
                no: 2,
                name: "args",
                kind: "scalar",
                repeat: 2 /*RepeatType.UNPACKED*/ ,
                T: 9 /*ScalarType.STRING*/ 
            },
            {
                no: 3,
                name: "negate",
                kind: "scalar",
                T: 8 /*ScalarType.BOOL*/ 
            },
            {
                no: 4,
                name: "type",
                kind: "enum",
                T: function() {
                    return [
                        "protos.steps.DetectiveType",
                        DetectiveType,
                        "DETECTIVE_TYPE_"
                    ];
                }
            }
        ]);
    }
    return DetectiveStep$Type;
}(MessageType);
/**
 * @generated MessageType for protobuf message protos.steps.DetectiveStep
 */ export var DetectiveStep = new DetectiveStep$Type();
