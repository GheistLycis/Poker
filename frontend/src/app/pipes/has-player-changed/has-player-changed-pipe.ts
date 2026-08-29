import type { PipeTransform } from '@angular/core';
import { Pipe } from '@angular/core';
import type { Player } from '@app-types/Player';

@Pipe({ name: 'hasPlayerChanged' })
export class HasPlayerChangedPipe implements PipeTransform {
  transform([prev, next]: [Player, Player]): boolean {
    if (prev.score !== next.score) return true;
    if (prev.cards.toString() !== next.cards.toString()) return true;
    return false;
  }
}
