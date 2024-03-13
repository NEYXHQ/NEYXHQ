/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";
import { Params } from "../token/params";
import { Token } from "../token/token";

export const protobufPackage = "Palo_Alt0.vanillaneyx.token";

/** GenesisState defines the token module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  tokenList: Token[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  tokenCount: number;
}

const baseGenesisState: object = { tokenCount: 0 };

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.tokenList) {
      Token.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    if (message.tokenCount !== 0) {
      writer.uint32(24).uint64(message.tokenCount);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.tokenList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.tokenList.push(Token.decode(reader, reader.uint32()));
          break;
        case 3:
          message.tokenCount = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.tokenList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.tokenList !== undefined && object.tokenList !== null) {
      for (const e of object.tokenList) {
        message.tokenList.push(Token.fromJSON(e));
      }
    }
    if (object.tokenCount !== undefined && object.tokenCount !== null) {
      message.tokenCount = Number(object.tokenCount);
    } else {
      message.tokenCount = 0;
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.tokenList) {
      obj.tokenList = message.tokenList.map((e) =>
        e ? Token.toJSON(e) : undefined
      );
    } else {
      obj.tokenList = [];
    }
    message.tokenCount !== undefined && (obj.tokenCount = message.tokenCount);
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.tokenList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.tokenList !== undefined && object.tokenList !== null) {
      for (const e of object.tokenList) {
        message.tokenList.push(Token.fromPartial(e));
      }
    }
    if (object.tokenCount !== undefined && object.tokenCount !== null) {
      message.tokenCount = object.tokenCount;
    } else {
      message.tokenCount = 0;
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
