export type Context = 'user' | 'opponents' | 'match';

export interface WebSocketMessage<T extends Context> {
  type: `${Extract<Context, T>}.${string}`;
  payload: object;
}
