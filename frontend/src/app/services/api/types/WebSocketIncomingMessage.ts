import { WebSocketMessage } from './WebSocketMessage';

export type WebSocketIncomingMessage = WebSocketMessage<'match' | 'opponents' | 'user'>;
