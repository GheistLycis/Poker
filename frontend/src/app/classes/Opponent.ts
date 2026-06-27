import { Card } from '@app-types/Card';
import { CreatePlayer, Player } from './Player';

interface CreateOpponent extends CreatePlayer {
  cards?: [Card, Card] | [null, null];
}

export class Opponent extends Player {
  cards: [Card, Card] | [null, null];

  constructor({ cards = [null, null], ...args }: CreateOpponent) {
    super(args);
    this.cards = cards;
  }
}
