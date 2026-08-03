import { WebSocketMessage } from './WebSocketMessage';

export interface WebSocketIncomingMessage extends WebSocketMessage {
  origin: 'SERVER';
  request_id?: string;
}
