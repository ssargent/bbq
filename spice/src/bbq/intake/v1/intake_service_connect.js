// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=js+dts"
// @generated from file bbq/intake/v1/intake_service.proto (package bbq.intake.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { RecordRequest, RecordResponse, SessionRequest, SessionResponse } from "./intake_service_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service bbq.intake.v1.IntakeService
 */
export const IntakeService = {
  typeName: "bbq.intake.v1.IntakeService",
  methods: {
    /**
     * @generated from rpc bbq.intake.v1.IntakeService.Record
     */
    record: {
      name: "Record",
      I: RecordRequest,
      O: RecordResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bbq.intake.v1.IntakeService.Session
     */
    session: {
      name: "Session",
      I: SessionRequest,
      O: SessionResponse,
      kind: MethodKind.Unary,
    },
  }
};
