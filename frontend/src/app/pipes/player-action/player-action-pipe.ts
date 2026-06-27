import { Pipe, PipeTransform } from '@angular/core';
import { PlayerAction } from '@app-types/PlayerAction';

const LABELS: Record<PlayerAction, string> = {
  CHECK: 'Passar',
  BET: 'Apostar',
  CALL: 'Cobrir',
  RAISE: 'Aumentar',
  FOLD: 'Desistir',
};

@Pipe({ name: 'playerAction' })
export class PlayerActionPipe implements PipeTransform {
  transform(value: PlayerAction): string {
    return LABELS[value];
  }
}
