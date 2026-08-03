import { InConnMessage } from './messages/ConnMessage';
import { ReceiveException } from './messages/in/ReceiveException';

export interface PendingRequest {
  resolve: (msg: Exclude<InConnMessage, ReceiveException>) => void;
  reject: (msg: ReceiveException) => void;
}
