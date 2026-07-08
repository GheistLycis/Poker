import { WebSocketIncomingMessage } from '@services/api/types/WebSocketIncomingMessage';

export interface ReceivePotAmount extends WebSocketIncomingMessage {
  type: 'match.pot-amount';
  payload: {
    amount: number;
  };
}
