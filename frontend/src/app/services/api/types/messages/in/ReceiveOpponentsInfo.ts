import { Opponent } from '@classes/Opponent';
import { WebSocketIncomingMessage } from '../WebSocketIncomingMessage';

export interface ReceiveOpponentsInfo extends WebSocketIncomingMessage {
  type: 'opponents.info';
  payload: Opponent[];
}
