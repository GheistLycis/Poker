import { Player } from '@app-types/Player';
import { WebSocketIncomingMessage } from './WebSocketIncomingMessage';

export interface ReceiveOpponentsInfo extends WebSocketIncomingMessage {
  type: 'opponents.info';
  payload: Player[];
}
