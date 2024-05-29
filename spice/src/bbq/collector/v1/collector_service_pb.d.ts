// @generated by protoc-gen-es v1.9.0 with parameter "target=js+dts"
// @generated from file bbq/collector/v1/collector_service.proto (package bbq.collector.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage, Timestamp } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import type { Reading, Session } from "./bbq_pb.js";

/**
 * @generated from message bbq.collector.v1.RecordRequest
 */
export declare class RecordRequest extends Message<RecordRequest> {
  /**
   * @generated from field: bbq.collector.v1.Reading reading = 1;
   */
  reading?: Reading;

  constructor(data?: PartialMessage<RecordRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "bbq.collector.v1.RecordRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RecordRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RecordRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RecordRequest;

  static equals(a: RecordRequest | PlainMessage<RecordRequest> | undefined, b: RecordRequest | PlainMessage<RecordRequest> | undefined): boolean;
}

/**
 * @generated from message bbq.collector.v1.RecordResponse
 */
export declare class RecordResponse extends Message<RecordResponse> {
  /**
   * @generated from field: string session_id = 1;
   */
  sessionId: string;

  /**
   * @generated from field: google.protobuf.Timestamp recorded_at = 2;
   */
  recordedAt?: Timestamp;

  constructor(data?: PartialMessage<RecordResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "bbq.collector.v1.RecordResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): RecordResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): RecordResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): RecordResponse;

  static equals(a: RecordResponse | PlainMessage<RecordResponse> | undefined, b: RecordResponse | PlainMessage<RecordResponse> | undefined): boolean;
}

/**
 * @generated from message bbq.collector.v1.SessionRequest
 */
export declare class SessionRequest extends Message<SessionRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name: string;

  constructor(data?: PartialMessage<SessionRequest>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "bbq.collector.v1.SessionRequest";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SessionRequest;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SessionRequest;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SessionRequest;

  static equals(a: SessionRequest | PlainMessage<SessionRequest> | undefined, b: SessionRequest | PlainMessage<SessionRequest> | undefined): boolean;
}

/**
 * @generated from message bbq.collector.v1.SessionResponse
 */
export declare class SessionResponse extends Message<SessionResponse> {
  /**
   * @generated from field: bbq.collector.v1.Session session = 1;
   */
  session?: Session;

  constructor(data?: PartialMessage<SessionResponse>);

  static readonly runtime: typeof proto3;
  static readonly typeName = "bbq.collector.v1.SessionResponse";
  static readonly fields: FieldList;

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SessionResponse;

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SessionResponse;

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SessionResponse;

  static equals(a: SessionResponse | PlainMessage<SessionResponse> | undefined, b: SessionResponse | PlainMessage<SessionResponse> | undefined): boolean;
}
