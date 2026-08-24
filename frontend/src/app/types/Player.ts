import type { Card } from './Card';
import type { SeatIndex } from './SeatIndex';

export interface Player {
  id: string;
  name: string;
  score: number;
  seatIndex: SeatIndex;
  cards: [] | [Card, Card];
}
