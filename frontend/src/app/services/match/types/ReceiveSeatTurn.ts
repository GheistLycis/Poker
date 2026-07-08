import { WebSocketIncomingMessage } from '@services/api/types/WebSocketIncomingMessage';

export interface ReceiveSeatTurn extends WebSocketIncomingMessage {
  type: 'match.seat-turn';
  payload: { seatIndex: number };
}
