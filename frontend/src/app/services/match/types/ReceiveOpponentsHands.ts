import { Card } from '@app-types/Card';
import { WebSocketIncomingMessage } from '@services/api/types/WebSocketIncomingMessage';

export interface ReceiveOpponentsHands extends WebSocketIncomingMessage {
  type: 'opponents.reveal-hands';
  payload: Record<string, [Card, Card]>;
}
