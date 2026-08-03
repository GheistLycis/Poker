export type Context = 'exception' | 'user' | 'opponents' | 'match';

export interface WebSocketMessage<T extends Context = Context> {
  origin: 'CLIENT' | 'SERVER';
  type: `${Extract<Context, T>}.${string}`;
  payload: object;
}
