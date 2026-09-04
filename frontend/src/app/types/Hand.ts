export const HandEnum = {
  HIGH_CARD: 'HIGH_CARD',
  ONE_PAIR: 'ONE_PAIR',
  TWO_PAIRS: 'TWO_PAIRS',
  THREE_OF_A_KIND: 'THREE_OF_A_KIND',
  STRAIGHT: 'STRAIGHT',
  FLUSH: 'FLUSH',
  FULL_HOUSE: 'FULL_HOUSE',
  FOUR_OF_A_KIND: 'FOUR_OF_A_KIND',
  STRAIGHT_FLUSH: 'STRAIGHT_FLUSH',
  ROYAL_FLUSH: 'ROYAL_FLUSH',
} as const;

export type Hand = (typeof HandEnum)[keyof typeof HandEnum];
