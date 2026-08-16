import type { WebSocketIncomingMessage } from './WebSocketIncomingMessage';

export interface ReceiveSeatTurn extends WebSocketIncomingMessage {
  type: 'match.seat-turn';
  payload: { seatIndex: number };
}
