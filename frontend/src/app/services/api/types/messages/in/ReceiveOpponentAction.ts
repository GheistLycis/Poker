import type { PlayerAction } from '@app-types/PlayerAction';
import type { WebSocketIncomingMessage } from './WebSocketIncomingMessage';

export interface ReceiveOpponentAction extends WebSocketIncomingMessage {
  type: 'opponents.action';
  payload: PurePayload | AmountPayload;
}

interface Payload {
  player: string;
  action: PlayerAction;
}

interface PurePayload extends Payload {
  action: 'CHECK' | 'CALL' | 'FOLD';
  amount?: undefined;
}

interface AmountPayload extends Payload {
  action: 'BET';
  amount: number;
}
