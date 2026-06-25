import { Player } from './Player';

export class Opponent extends Player {
  handSize: number;

  constructor(id: string, name: string, score: number, handSize: number) {
    super(id, name, score);
    this.handSize = handSize;
  }
}
