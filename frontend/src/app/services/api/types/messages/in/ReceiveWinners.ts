import type { Player } from '@app-types/Player';
import type { WebSocketIncomingMessage } from './WebSocketIncomingMessage';

export interface ReceiveWinners extends WebSocketIncomingMessage {
  type: 'match.winners';
  payload: Player['id'][];
}
