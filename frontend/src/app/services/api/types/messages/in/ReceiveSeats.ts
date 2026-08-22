import type { Player } from '@app-types/Player';
import type { SeatIndex } from '@app-types/SeatIndex';
import type { WebSocketIncomingMessage } from './WebSocketIncomingMessage';

export interface ReceiveSeats extends WebSocketIncomingMessage {
  type: 'match.seats';
  payload: Record<SeatIndex, Player['id'] | null>;
}
