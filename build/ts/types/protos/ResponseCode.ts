// Original file: ../../protos/common.proto

export const ResponseCode = {
  RESPONSE_CODE_UNSET: 'RESPONSE_CODE_UNSET',
  RESPONSE_CODE_OK: 'RESPONSE_CODE_OK',
  RESPONSE_CODE_BAD_REQUEST: 'RESPONSE_CODE_BAD_REQUEST',
  RESPONSE_CODE_NOT_FOUND: 'RESPONSE_CODE_NOT_FOUND',
  RESPONSE_CODE_INTERNAL_SERVER_ERROR: 'RESPONSE_CODE_INTERNAL_SERVER_ERROR',
  RESPONSE_CODE_GENERIC_ERROR: 'RESPONSE_CODE_GENERIC_ERROR',
} as const;

export type ResponseCode =
  | 'RESPONSE_CODE_UNSET'
  | 0
  | 'RESPONSE_CODE_OK'
  | 1
  | 'RESPONSE_CODE_BAD_REQUEST'
  | 2
  | 'RESPONSE_CODE_NOT_FOUND'
  | 3
  | 'RESPONSE_CODE_INTERNAL_SERVER_ERROR'
  | 4
  | 'RESPONSE_CODE_GENERIC_ERROR'
  | 5

export type ResponseCode__Output = typeof ResponseCode[keyof typeof ResponseCode]
