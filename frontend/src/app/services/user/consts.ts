import { User } from '@classes/User';
import { faker } from '@faker-js/faker';
import { CardEnum } from '../../types/Card';

export const USER = new User({
  id: faker.string.uuid(),
  name: faker.person.fullName(),
  score: faker.number.int(1_000_000),
  cards: [faker.helpers.enumValue(CardEnum), faker.helpers.enumValue(CardEnum)],
});
