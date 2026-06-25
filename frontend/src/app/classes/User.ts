import { Card } from '@app-types/Card';
import { Player } from './Player';

export class User extends Player {
  cards: Card[];

  constructor(id: string, name: string, score: number, cards: Card[]) {
    super(id, name, score);
    this.cards = cards;
  }
}
