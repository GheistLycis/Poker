import type { ReceiveOpponentAction } from './in/ReceiveOpponentAction';
import type { ReceiveOpponentsHands } from './in/ReceiveOpponentsHands';
import type { ReceiveOpponentsInfo } from './in/ReceiveOpponentsInfo';
import type { ReceivePotAmount } from './in/ReceivePotAmount';
import type { ReceiveSeats } from './in/ReceiveSeats';
import type { ReceiveSeatTurn } from './in/ReceiveSeatTurn';
import type { ReceiveTableCards } from './in/ReceiveTableCards';
import type { ReceiveUserInfo } from './in/ReceiveUserInfo';
import type { ReceiveWinner } from './in/ReceiveWinner';
import type { SendUserAction } from './out/SendUserAction';
import type { SendUserEmote } from './out/SendUserEmote';
import type { SendUserLogin } from './out/SendUserLogin';

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
