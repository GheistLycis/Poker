export const PlayerActionEnum = {
  CHECK: 'CHECK',
  BET: 'BET',
  CALL: 'CALL',
  RAISE: 'RAISE',
  FOLD: 'FOLD',
} as const;

export type PlayerAction = (typeof PlayerActionEnum)[keyof typeof PlayerActionEnum];
