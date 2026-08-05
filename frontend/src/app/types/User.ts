import { Card } from '@app-types/Card';
import { Player } from './Player';

export interface User extends Player {
  cards: [Card, Card];
}
