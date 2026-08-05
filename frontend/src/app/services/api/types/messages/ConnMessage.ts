import { ReceiveOpponentAction } from './in/ReceiveOpponentAction';
import { ReceiveOpponentsHands } from './in/ReceiveOpponentsHands';
import { ReceiveOpponentsInfo } from './in/ReceiveOpponentsInfo';
import { ReceivePotAmount } from './in/ReceivePotAmount';
import { ReceiveSeats } from './in/ReceiveSeats';
import { ReceiveSeatTurn } from './in/ReceiveSeatTurn';
import { ReceiveTableCards } from './in/ReceiveTableCards';
import { ReceiveUserInfo } from './in/ReceiveUserInfo';
import { ReceiveWinner } from './in/ReceiveWinner';
import { SendUserAction } from './out/SendUserAction';
import { SendUserEmote } from './out/SendUserEmote';
import { SendUserLogin } from './out/SendUserLogin';

export type InConnMessage =
  | ReceiveOpponentAction
  | ReceiveOpponentsHands
  | ReceiveOpponentsInfo
  | ReceivePotAmount
  | ReceiveTableCards
  | ReceiveWinner
  | ReceiveSeats
  | ReceiveSeatTurn
  | ReceiveUserInfo;

export type OutConnMessage = SendUserAction | SendUserEmote | SendUserLogin;

export type ConnMessage = InConnMessage | OutConnMessage;
