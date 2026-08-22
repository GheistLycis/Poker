import type { Player } from '@app-types/Player';
import type { WebSocketIncomingMessage } from './WebSocketIncomingMessage';

export interface ReceiveWinner extends WebSocketIncomingMessage {
  type: 'match.winner';
  payload: {
    player: Player['id'];
  };
}
