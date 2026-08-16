import type { Card } from './Card';

export interface Player {
  id: string;
  name: string;
  score: number;
  seatIndex: number;
  cards: [] | [Card, Card];
}
