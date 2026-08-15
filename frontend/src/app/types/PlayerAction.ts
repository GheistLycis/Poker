export const PlayerActionEnum = {
  BET: 'BET',
  CALL: 'CALL',
  CHECK: 'CHECK',
  FOLD: 'FOLD',
} as const;

export type PlayerAction = (typeof PlayerActionEnum)[keyof typeof PlayerActionEnum];
