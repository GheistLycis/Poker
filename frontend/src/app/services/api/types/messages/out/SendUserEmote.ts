import { WebSocketOutgoingMessage } from '../WebSocketOutgoingMessage';

export interface SendUserEmote extends WebSocketOutgoingMessage {
  type: 'user.emote';
  payload: {
    emote: unknown; // TODO
  };
}
