import { WebSocketIncomingMessage } from './WebSocketIncomingMessage';

export interface ReceiveWinner extends WebSocketIncomingMessage {
  type: 'match.winner';
  payload: {
    player: string;
  };
}
