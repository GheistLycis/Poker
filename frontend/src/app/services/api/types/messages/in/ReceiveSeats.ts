import { WebSocketIncomingMessage } from './WebSocketIncomingMessage';

export interface ReceiveSeats extends WebSocketIncomingMessage {
  type: 'match.seats';
  payload: Record<number, string | null>;
}
