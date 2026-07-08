import { WebSocketOutgoingMessage } from '@services/api/types/WebSocketOutgoingMessage';

export interface SendUserEmote extends WebSocketOutgoingMessage {
  type: 'user.emote';
  payload: {
    emote: unknown; // TODO
  };
}
