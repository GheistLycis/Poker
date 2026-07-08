import { ReceiveOpponentAction } from './ReceiveOpponentAction';
import { ReceiveOpponentsHands } from './ReceiveOpponentsHands';
import { ReceiveOpponentsInfo } from './ReceiveOpponentsInfo';
import { ReceivePotAmount } from './ReceivePotAmount';
import { ReceiveSeats } from './ReceiveSeats';
import { ReceiveSeatTurn } from './ReceiveSeatTurn';
import { ReceiveTableCards } from './ReceiveTableCards';
import { ReceiveWinner } from './ReceiveWinner';
import { SendUserAction } from './SendUserAction';
import { SendUserEmote } from './SendUserEmote';

export type InConnMessage =
  | ReceiveOpponentAction
  | ReceiveOpponentsHands
  | ReceiveOpponentsInfo
  | ReceivePotAmount
  | ReceiveTableCards
  | ReceiveWinner
  | ReceiveSeats
  | ReceiveSeatTurn;

export type OutConnMessage = SendUserAction | SendUserEmote;

export type ConnMessage = InConnMessage | OutConnMessage;
