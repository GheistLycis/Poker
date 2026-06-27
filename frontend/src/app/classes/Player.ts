export interface CreatePlayer {
  id: string;
  name: string;
  score: number;
}

export class Player {
  id: string;
  name: string;
  score: number;

  constructor(args: CreatePlayer) {
    this.id = args.id;
    this.name = args.name;
    this.score = args.score;
  }
}
