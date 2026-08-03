export interface CreatePlayer {
  name: string;
  score: number;
}

export class Player {
  name: string;
  score: number;

  constructor(args: CreatePlayer) {
    this.name = args.name;
    this.score = args.score;
  }
}
