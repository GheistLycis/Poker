import type { Player } from '@app-types/Player';
import type { WebSocketIncomingMessage } from './WebSocketIncomingMessage';

export interface ReceiveUserInfo extends WebSocketIncomingMessage {
  type: 'user.info';
  payload: Player;
}
