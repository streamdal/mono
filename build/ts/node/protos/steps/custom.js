"use strict";
var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (Object.prototype.hasOwnProperty.call(b, p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        if (typeof b !== "function" && b !== null)
            throw new TypeError("Class extends value " + String(b) + " is not a constructor or null");
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
Object.defineProperty(exports, "__esModule", { value: true });
exports.CustomStep = void 0;
// @generated by protobuf-ts 2.9.0 with parameter optimize_code_size,client_grpc1
// @generated from protobuf file "steps/custom.proto" (package "protos.steps", syntax proto3)
// tslint:disable
var runtime_1 = require("@protobuf-ts/runtime");
// @generated message type with reflection information, may provide speed optimized methods
var CustomStep$Type = /** @class */ (function (_super) {
    __extends(CustomStep$Type, _super);
    function CustomStep$Type() {
        return _super.call(this, "protos.steps.CustomStep", [
            { no: 1, name: "id", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]) || this;
    }
    return CustomStep$Type;
}(runtime_1.MessageType));
/**
 * @generated MessageType for protobuf message protos.steps.CustomStep
 */
exports.CustomStep = new CustomStep$Type();
//# sourceMappingURL=custom.js.map