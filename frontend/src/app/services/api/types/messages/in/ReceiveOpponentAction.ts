import type { Player } from '@app-types/Player';
import type { PlayerAction } from '@app-types/PlayerAction';
import type { WebSocketIncomingMessage } from './WebSocketIncomingMessage';

export interface ReceiveOpponentAction extends WebSocketIncomingMessage {
  type: 'opponents.action';
  payload: PurePayload | AmountPayload;
}

interface Payload {
  player: Player['id'];
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
