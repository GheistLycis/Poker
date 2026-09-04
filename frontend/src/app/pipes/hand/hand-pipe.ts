import type { PipeTransform } from '@angular/core';
import { Pipe } from '@angular/core';
import type { Hand } from '@app-types/Hand';

const LABELS: Record<Hand, string> = {
  HIGH_CARD: 'Carta Alta',
  ONE_PAIR: 'Par',
  TWO_PAIRS: 'Dois Pares',
  THREE_OF_A_KIND: 'Trinca',
  STRAIGHT: 'Sequência',
  FLUSH: 'Flush',
  FULL_HOUSE: 'Full House',
  FOUR_OF_A_KIND: 'Quadra',
  STRAIGHT_FLUSH: 'Straight Flush',
  ROYAL_FLUSH: 'Royal Flush',
};

@Pipe({ name: 'hand' })
export class HandPipe implements PipeTransform {
  transform(value: Hand): string {
    return LABELS[value];
  }
}
