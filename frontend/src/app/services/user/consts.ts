import { User } from '@classes/User';
import { faker } from '@faker-js/faker';
import { CardEnum } from '../../types/Card';

export const USER = new User(
  faker.string.uuid(),
  faker.person.fullName(),
  faker.number.int(0),
  Object.values(CardEnum).slice(1, 6),
);
