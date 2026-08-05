import { WebSocketMessage } from './WebSocketMessage';

export interface WebSocketOutgoingMessage extends WebSocketMessage<'user'> {
  origin: 'CLIENT';
  requestId: string;
}
