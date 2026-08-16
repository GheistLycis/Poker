import type { WebSocketMessage } from '../WebSocketMessage';

export interface WebSocketIncomingMessage extends WebSocketMessage {
  origin: 'SERVER';
  requestId: string | null;
  error: {
    message: string;
    details: object | null;
  } | null;
}
