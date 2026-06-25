export const CardOwnerEnum = {
  USER: 'USER',
  OPPONENT: 'OPPONENT',
  TABLE: 'TABLE',
  DECK: 'DECK',
} as const;

export type CardOwner = (typeof CardOwnerEnum)[keyof typeof CardOwnerEnum];
