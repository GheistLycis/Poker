import { Card } from '@app-types/Card';
import { WebSocketIncomingMessage } from '@services/api/types/WebSocketIncomingMessage';

export interface ReceiveTableCards extends WebSocketIncomingMessage {
  type: 'match.table-cards';
  payload: [Card, Card, Card, Card, Card];
}
