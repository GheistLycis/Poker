import { WebSocketIncomingMessage } from '../WebSocketIncomingMessage';

export interface ReceiveException extends WebSocketIncomingMessage {
  type: 'exception.';
}
