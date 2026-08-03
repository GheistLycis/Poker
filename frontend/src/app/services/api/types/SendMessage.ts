import { OutConnMessage } from './messages/ConnMessage';

type DistributiveOmit<T, K extends PropertyKey> = T extends unknown ? Omit<T, K> : never;

export type SendMessage = DistributiveOmit<OutConnMessage, 'origin' | 'request_id'>;
