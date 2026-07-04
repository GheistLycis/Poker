import { Opponent } from '@classes/Opponent';
import { WebSocketIncomingMessage } from '@services/api/types/WebSocketIncomingMessage';

export interface ReceiveOpponentsInfo extends WebSocketIncomingMessage {
  type: 'opponents.info';
  payload: Opponent[];
}
