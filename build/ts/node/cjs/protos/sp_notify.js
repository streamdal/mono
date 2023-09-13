"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.NotificationPagerDuty = exports.NotificationEmailSES = exports.NotificationEmailSMTP = exports.NotificationEmail = exports.NotificationSlack = exports.NotificationConfig = exports.NotificationType = exports.NotificationPagerDuty_Urgency = exports.NotificationEmail_Type = void 0;
// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size
// @generated from protobuf file "sp_notify.proto" (package "protos", syntax proto3)
// tslint:disable
const runtime_1 = require("@protobuf-ts/runtime");
/**
 * @generated from protobuf enum protos.NotificationEmail.Type
 */
var NotificationEmail_Type;
(function (NotificationEmail_Type) {
    /**
     * @generated from protobuf enum value: TYPE_UNSET = 0;
     */
    NotificationEmail_Type[NotificationEmail_Type["UNSET"] = 0] = "UNSET";
    /**
     * @generated from protobuf enum value: TYPE_SMTP = 1;
     */
    NotificationEmail_Type[NotificationEmail_Type["SMTP"] = 1] = "SMTP";
    /**
     * @generated from protobuf enum value: TYPE_SES = 2;
     */
    NotificationEmail_Type[NotificationEmail_Type["SES"] = 2] = "SES";
})(NotificationEmail_Type || (exports.NotificationEmail_Type = NotificationEmail_Type = {}));
/**
 * @generated from protobuf enum protos.NotificationPagerDuty.Urgency
 */
var NotificationPagerDuty_Urgency;
(function (NotificationPagerDuty_Urgency) {
    /**
     * @generated from protobuf enum value: URGENCY_UNSET = 0;
     */
    NotificationPagerDuty_Urgency[NotificationPagerDuty_Urgency["UNSET"] = 0] = "UNSET";
    /**
     * @generated from protobuf enum value: URGENCY_LOW = 1;
     */
    NotificationPagerDuty_Urgency[NotificationPagerDuty_Urgency["LOW"] = 1] = "LOW";
    /**
     * @generated from protobuf enum value: URGENCY_HIGH = 2;
     */
    NotificationPagerDuty_Urgency[NotificationPagerDuty_Urgency["HIGH"] = 2] = "HIGH";
})(NotificationPagerDuty_Urgency || (exports.NotificationPagerDuty_Urgency = NotificationPagerDuty_Urgency = {}));
/**
 * @generated from protobuf enum protos.NotificationType
 */
var NotificationType;
(function (NotificationType) {
    /**
     * @generated from protobuf enum value: NOTIFICATION_TYPE_UNSET = 0;
     */
    NotificationType[NotificationType["UNSET"] = 0] = "UNSET";
    /**
     * @generated from protobuf enum value: NOTIFICATION_TYPE_SLACK = 1;
     */
    NotificationType[NotificationType["SLACK"] = 1] = "SLACK";
    /**
     * @generated from protobuf enum value: NOTIFICATION_TYPE_EMAIL = 2;
     */
    NotificationType[NotificationType["EMAIL"] = 2] = "EMAIL";
    /**
     * @generated from protobuf enum value: NOTIFICATION_TYPE_PAGERDUTY = 3;
     */
    NotificationType[NotificationType["PAGERDUTY"] = 3] = "PAGERDUTY";
})(NotificationType || (exports.NotificationType = NotificationType = {}));
// @generated message type with reflection information, may provide speed optimized methods
class NotificationConfig$Type extends runtime_1.MessageType {
    constructor() {
        super("protos.NotificationConfig", [
            { no: 1, name: "id", kind: "scalar", opt: true, T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "name", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "type", kind: "enum", T: () => ["protos.NotificationType", NotificationType, "NOTIFICATION_TYPE_"] },
            { no: 1000, name: "slack", kind: "message", oneof: "config", T: () => exports.NotificationSlack },
            { no: 1001, name: "email", kind: "message", oneof: "config", T: () => exports.NotificationEmail },
            { no: 1002, name: "pagerduty", kind: "message", oneof: "config", T: () => exports.NotificationPagerDuty }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.NotificationConfig
 */
exports.NotificationConfig = new NotificationConfig$Type();
// @generated message type with reflection information, may provide speed optimized methods
class NotificationSlack$Type extends runtime_1.MessageType {
    constructor() {
        super("protos.NotificationSlack", [
            { no: 1, name: "bot_token", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "channel", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.NotificationSlack
 */
exports.NotificationSlack = new NotificationSlack$Type();
// @generated message type with reflection information, may provide speed optimized methods
class NotificationEmail$Type extends runtime_1.MessageType {
    constructor() {
        super("protos.NotificationEmail", [
            { no: 1, name: "type", kind: "enum", T: () => ["protos.NotificationEmail.Type", NotificationEmail_Type, "TYPE_"] },
            { no: 2, name: "recipients", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "from_address", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 1000, name: "smtp", kind: "message", oneof: "config", T: () => exports.NotificationEmailSMTP },
            { no: 1001, name: "ses", kind: "message", oneof: "config", T: () => exports.NotificationEmailSES }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.NotificationEmail
 */
exports.NotificationEmail = new NotificationEmail$Type();
// @generated message type with reflection information, may provide speed optimized methods
class NotificationEmailSMTP$Type extends runtime_1.MessageType {
    constructor() {
        super("protos.NotificationEmailSMTP", [
            { no: 1, name: "host", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "port", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "user", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "password", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 5, name: "use_tls", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.NotificationEmailSMTP
 */
exports.NotificationEmailSMTP = new NotificationEmailSMTP$Type();
// @generated message type with reflection information, may provide speed optimized methods
class NotificationEmailSES$Type extends runtime_1.MessageType {
    constructor() {
        super("protos.NotificationEmailSES", [
            { no: 1, name: "ses_region", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "ses_access_key_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "ses_secret_access_key", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.NotificationEmailSES
 */
exports.NotificationEmailSES = new NotificationEmailSES$Type();
// @generated message type with reflection information, may provide speed optimized methods
class NotificationPagerDuty$Type extends runtime_1.MessageType {
    constructor() {
        super("protos.NotificationPagerDuty", [
            { no: 1, name: "token", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "email", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "service_id", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "urgency", kind: "enum", T: () => ["protos.NotificationPagerDuty.Urgency", NotificationPagerDuty_Urgency, "URGENCY_"] }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message protos.NotificationPagerDuty
 */
exports.NotificationPagerDuty = new NotificationPagerDuty$Type();
