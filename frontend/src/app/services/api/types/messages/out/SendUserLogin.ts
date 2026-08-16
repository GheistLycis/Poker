import type { WebSocketOutgoingMessage } from './WebSocketOutgoingMessage';

export interface SendUserLogin extends WebSocketOutgoingMessage {
  type: 'user.login';
  payload: {
    userName: string;
  };
}
