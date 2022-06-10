/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "smarshallspitzbart.ride.ride";

export interface StoredRide {
  index: string;
  destination: string;
  driver: string;
  passenger: string;
  mutualStake: number;
  payPerHour: number;
  distanceTip: number;
}

const baseStoredRide: object = {
  index: "",
  destination: "",
  driver: "",
  passenger: "",
  mutualStake: 0,
  payPerHour: 0,
  distanceTip: 0,
};

export const StoredRide = {
  encode(message: StoredRide, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.destination !== "") {
      writer.uint32(18).string(message.destination);
    }
    if (message.driver !== "") {
      writer.uint32(26).string(message.driver);
    }
    if (message.passenger !== "") {
      writer.uint32(34).string(message.passenger);
    }
    if (message.mutualStake !== 0) {
      writer.uint32(40).uint64(message.mutualStake);
    }
    if (message.payPerHour !== 0) {
      writer.uint32(48).uint64(message.payPerHour);
    }
    if (message.distanceTip !== 0) {
      writer.uint32(56).uint64(message.distanceTip);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): StoredRide {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseStoredRide } as StoredRide;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.destination = reader.string();
          break;
        case 3:
          message.driver = reader.string();
          break;
        case 4:
          message.passenger = reader.string();
          break;
        case 5:
          message.mutualStake = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.payPerHour = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.distanceTip = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): StoredRide {
    const message = { ...baseStoredRide } as StoredRide;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.destination !== undefined && object.destination !== null) {
      message.destination = String(object.destination);
    } else {
      message.destination = "";
    }
    if (object.driver !== undefined && object.driver !== null) {
      message.driver = String(object.driver);
    } else {
      message.driver = "";
    }
    if (object.passenger !== undefined && object.passenger !== null) {
      message.passenger = String(object.passenger);
    } else {
      message.passenger = "";
    }
    if (object.mutualStake !== undefined && object.mutualStake !== null) {
      message.mutualStake = Number(object.mutualStake);
    } else {
      message.mutualStake = 0;
    }
    if (object.payPerHour !== undefined && object.payPerHour !== null) {
      message.payPerHour = Number(object.payPerHour);
    } else {
      message.payPerHour = 0;
    }
    if (object.distanceTip !== undefined && object.distanceTip !== null) {
      message.distanceTip = Number(object.distanceTip);
    } else {
      message.distanceTip = 0;
    }
    return message;
  },

  toJSON(message: StoredRide): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.destination !== undefined &&
      (obj.destination = message.destination);
    message.driver !== undefined && (obj.driver = message.driver);
    message.passenger !== undefined && (obj.passenger = message.passenger);
    message.mutualStake !== undefined &&
      (obj.mutualStake = message.mutualStake);
    message.payPerHour !== undefined && (obj.payPerHour = message.payPerHour);
    message.distanceTip !== undefined &&
      (obj.distanceTip = message.distanceTip);
    return obj;
  },

  fromPartial(object: DeepPartial<StoredRide>): StoredRide {
    const message = { ...baseStoredRide } as StoredRide;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.destination !== undefined && object.destination !== null) {
      message.destination = object.destination;
    } else {
      message.destination = "";
    }
    if (object.driver !== undefined && object.driver !== null) {
      message.driver = object.driver;
    } else {
      message.driver = "";
    }
    if (object.passenger !== undefined && object.passenger !== null) {
      message.passenger = object.passenger;
    } else {
      message.passenger = "";
    }
    if (object.mutualStake !== undefined && object.mutualStake !== null) {
      message.mutualStake = object.mutualStake;
    } else {
      message.mutualStake = 0;
    }
    if (object.payPerHour !== undefined && object.payPerHour !== null) {
      message.payPerHour = object.payPerHour;
    } else {
      message.payPerHour = 0;
    }
    if (object.distanceTip !== undefined && object.distanceTip !== null) {
      message.distanceTip = object.distanceTip;
    } else {
      message.distanceTip = 0;
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}