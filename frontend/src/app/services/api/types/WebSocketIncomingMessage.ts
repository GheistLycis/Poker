import { WebSocketMessage } from './WebSocketMessage';

export interface WebSocketIncomingMessage extends WebSocketMessage<'match' | 'opponents' | 'user'> {
  origin: 'SERVER';
}
