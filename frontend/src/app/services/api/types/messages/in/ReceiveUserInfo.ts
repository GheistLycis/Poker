import { User } from '@app-types/User';
import { WebSocketIncomingMessage } from '../WebSocketIncomingMessage';

export interface ReceiveUserInfo extends WebSocketIncomingMessage {
  type: 'user.info';
  payload: User;
}
