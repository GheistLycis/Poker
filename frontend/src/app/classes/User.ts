import { Card } from '@app-types/Card';
import { CreatePlayer, Player } from './Player';

export interface CreateUser extends CreatePlayer {
  cards: [Card, Card];
}

export class User extends Player {
  cards: [Card, Card];

  constructor(args: CreateUser) {
    super(args);
    this.cards = args.cards;
  }
}
