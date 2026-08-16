import type { Card } from '@app-types/Card';
import type { WebSocketIncomingMessage } from './WebSocketIncomingMessage';

export interface ReceiveOpponentsHands extends WebSocketIncomingMessage {
  type: 'opponents.reveal-hands';
  payload: Record<string, [Card, Card]>;
}
