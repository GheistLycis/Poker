import type { Card } from '@app-types/Card';
import type { Hand } from '@app-types/Hand';
import type { Player } from '@app-types/Player';
import type { WebSocketIncomingMessage } from './WebSocketIncomingMessage';

export interface ReceiveWinners extends WebSocketIncomingMessage {
  type: 'match.winners';
  payload: {
    id: Player['id'];
    winningHand: Hand;
    winningCards: Card[];
  }[];
}
