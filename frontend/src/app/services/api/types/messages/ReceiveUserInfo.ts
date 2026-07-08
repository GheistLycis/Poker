import { User } from '@classes/User';
import { WebSocketIncomingMessage } from '@services/api/types/WebSocketIncomingMessage';

export interface ReceiveUserInfo extends WebSocketIncomingMessage {
  type: 'user.info';
  payload: User;
}
