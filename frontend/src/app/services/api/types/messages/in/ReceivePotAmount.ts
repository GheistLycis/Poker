import type { WebSocketIncomingMessage } from './WebSocketIncomingMessage';

export interface ReceivePotAmount extends WebSocketIncomingMessage {
  type: 'match.pot-amount';
  payload: {
    amount: number;
  };
}
