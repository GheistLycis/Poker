import type { SeatIndex } from '@app-types/SeatIndex';
import type { WebSocketIncomingMessage } from './WebSocketIncomingMessage';

export interface ReceiveSeats extends WebSocketIncomingMessage {
  type: 'match.seats';
  payload: Record<SeatIndex, string | null>;
}
