import { Opponent } from '@app-types/Opponent';
import { WebSocketIncomingMessage } from '../WebSocketIncomingMessage';

export interface ReceiveOpponentsInfo extends WebSocketIncomingMessage {
  type: 'opponents.info';
  payload: Opponent[];
}
