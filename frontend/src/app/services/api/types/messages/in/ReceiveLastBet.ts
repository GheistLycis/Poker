import type { WebSocketIncomingMessage } from './WebSocketIncomingMessage';

export interface ReceiveLastBet extends WebSocketIncomingMessage {
  type: 'match.last-bet';
  payload: { amount: number };
}
