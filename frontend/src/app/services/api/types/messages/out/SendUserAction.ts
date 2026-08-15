import { PlayerAction } from '@app-types/PlayerAction';
import { WebSocketOutgoingMessage } from './WebSocketOutgoingMessage';

export interface SendUserAction extends WebSocketOutgoingMessage {
  type: 'user.action';
  payload: PurePayload | AmountPayload;
}

interface Payload {
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
