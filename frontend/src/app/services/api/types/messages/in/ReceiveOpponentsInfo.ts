import type { Player } from '@app-types/Player';
import type { WebSocketIncomingMessage } from './WebSocketIncomingMessage';

export interface ReceiveOpponentsInfo extends WebSocketIncomingMessage {
  type: 'opponents.info';
  payload: Player[];
}
