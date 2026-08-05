import { Card } from '@app-types/Card';
import { WebSocketIncomingMessage } from './WebSocketIncomingMessage';

export interface ReceiveTableCards extends WebSocketIncomingMessage {
  type: 'match.table-cards';
  payload: [Card, Card, Card, Card, Card];
}
