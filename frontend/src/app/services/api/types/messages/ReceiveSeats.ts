import { WebSocketIncomingMessage } from '@services/api/types/WebSocketIncomingMessage';

export interface ReceiveSeats extends WebSocketIncomingMessage {
  type: 'match.seats';
  payload: Record<number, string | null>;
}
