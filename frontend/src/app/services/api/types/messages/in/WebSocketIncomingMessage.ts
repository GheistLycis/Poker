import type { WebSocketOutgoingMessage } from '../out/WebSocketOutgoingMessage';
import type { WebSocketMessage } from '../WebSocketMessage';

export interface WebSocketIncomingMessage extends WebSocketMessage {
  origin: 'SERVER';
  requestId: WebSocketOutgoingMessage['requestId'] | null;
  error: {
    message: string;
    details: object | null;
  } | null;
}
