import { WebSocketIncomingMessage } from '@services/api/types/WebSocketIncomingMessage';

export interface ReceiveWinner extends WebSocketIncomingMessage {
  type: 'match.winner';
  payload: {
    player: string;
  };
}
