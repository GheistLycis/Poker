export const WebSocketConnStateEnum = {
  OPEN: 'OPEN',
  CLOSE: 'CLOSE',
  CONNECTING: 'CONNECTING',
} as const;

export type WebSocketConnState =
  (typeof WebSocketConnStateEnum)[keyof typeof WebSocketConnStateEnum];
