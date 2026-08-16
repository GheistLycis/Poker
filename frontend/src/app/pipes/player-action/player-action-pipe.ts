import type { PipeTransform } from '@angular/core';
import { Pipe } from '@angular/core';
import type { PlayerAction } from '@app-types/PlayerAction';

const LABELS: Record<PlayerAction, string> = {
  BET: 'Apostar',
  CALL: 'Cobrir',
  CHECK: 'Passar',
  FOLD: 'Desistir',
};

@Pipe({ name: 'playerAction' })
export class PlayerActionPipe implements PipeTransform {
  transform(value: PlayerAction): string {
    return LABELS[value];
  }
}
