import { Player } from '@app-types/Player';
import { WebSocketIncomingMessage } from './WebSocketIncomingMessage';

export interface ReceiveUserInfo extends WebSocketIncomingMessage {
  type: 'user.info';
  payload: Player;
}
