import { Opponent } from '@classes/Opponent';
import { faker } from '@faker-js/faker';
import { USER } from '@services/user/consts';

const SEATS_NUMBER = faker.number.int({ min: 2, max: 8 });

export const OPPONENTS = Array.from({ length: SEATS_NUMBER - 1 }).map(
  () => new Opponent(faker.string.uuid(), faker.person.fullName(), faker.number.int(1_000_000)),
);

export const PUSH_CARD_TO_HAND_DELAY_MS = 200;

const shuffledOpponents = faker.helpers.shuffle(OPPONENTS);

export const SEATS: Record<number, string | null> = Array.from({
  length: SEATS_NUMBER,
}).reduce<Record<number, string | null>>(
  (acc, _, index) => ({
    ...acc,
    [index]: index === 0 ? USER.id : shuffledOpponents[index - 1].id,
  }),
  {},
);
