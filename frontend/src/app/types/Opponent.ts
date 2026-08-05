import { Card } from '@app-types/Card';
import { Player } from './Player';

export interface Opponent extends Player {
  cards: [Card, Card] | [null, null];
}
