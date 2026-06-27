import { Card } from '@app-types/Card';
import { Player } from './Player';

export class Opponent extends Player {
  cards: [null, null] | [Card, Card] = [null, null];

  constructor(id: string, name: string, score: number) {
    super(id, name, score);
  }
}
