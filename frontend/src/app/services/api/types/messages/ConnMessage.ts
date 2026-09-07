import type { ReceiveLastBet } from './in/ReceiveLastBet';
import type { ReceiveOpponentAction } from './in/ReceiveOpponentAction';
import type { ReceiveOpponentsInfo } from './in/ReceiveOpponentsInfo';
import type { ReceivePotAmount } from './in/ReceivePotAmount';
import type { ReceiveSeats } from './in/ReceiveSeats';
import type { ReceiveSeatTurn } from './in/ReceiveSeatTurn';
import type { ReceiveTableCards } from './in/ReceiveTableCards';
import type { ReceiveUserInfo } from './in/ReceiveUserInfo';
import type { ReceiveWinners } from './in/ReceiveWinners';
import type { SendUserAction } from './out/SendUserAction';
import type { SendUserLogin } from './out/SendUserLogin';

export type InConnMessage =
  | ReceiveOpponentAction
  | ReceiveOpponentsInfo
  | ReceivePotAmount
  | ReceiveTableCards
  | ReceiveWinners
  | ReceiveSeats
  | ReceiveSeatTurn
  | ReceiveUserInfo
  | ReceiveLastBet;

export type OutConnMessage = SendUserAction | SendUserLogin;

export type ConnMessage = InConnMessage | OutConnMessage;
