export const CardOwnerEnum = {
  USER: 'USER',
  OPPONENT: 'OPPONENT',
  TABLE: 'TABLE',
} as const;

export type CardOwner = (typeof CardOwnerEnum)[keyof typeof CardOwnerEnum];
