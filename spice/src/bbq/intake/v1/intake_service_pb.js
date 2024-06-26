// @generated by protoc-gen-es v1.9.0 with parameter "target=js+dts"
// @generated from file bbq/intake/v1/intake_service.proto (package bbq.intake.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { proto3, Timestamp } from "@bufbuild/protobuf";
import { Reading, Session } from "./bbq_pb.js";

/**
 * @generated from message bbq.intake.v1.RecordRequest
 */
export const RecordRequest = /*@__PURE__*/ proto3.makeMessageType(
  "bbq.intake.v1.RecordRequest",
  () => [
    { no: 1, name: "reading", kind: "message", T: Reading, repeated: true },
  ],
);

/**
 * @generated from message bbq.intake.v1.RecordResponse
 */
export const RecordResponse = /*@__PURE__*/ proto3.makeMessageType(
  "bbq.intake.v1.RecordResponse",
  () => [
    { no: 1, name: "session_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "recorded_at", kind: "message", T: Timestamp },
  ],
);

/**
 * @generated from message bbq.intake.v1.SessionRequest
 */
export const SessionRequest = /*@__PURE__*/ proto3.makeMessageType(
  "bbq.intake.v1.SessionRequest",
  () => [
    { no: 1, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "device_name", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 3, name: "sensor_name", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 4, name: "subject_id", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 5, name: "desired_state", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
  ],
);

/**
 * @generated from message bbq.intake.v1.SessionResponse
 */
export const SessionResponse = /*@__PURE__*/ proto3.makeMessageType(
  "bbq.intake.v1.SessionResponse",
  () => [
    { no: 1, name: "session", kind: "message", T: Session },
  ],
);

