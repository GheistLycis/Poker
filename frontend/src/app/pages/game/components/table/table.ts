import { AsyncPipe, NgTemplateOutlet } from '@angular/common';
import { Component, inject, input } from '@angular/core';
import { toObservable } from '@angular/core/rxjs-interop';
import { CardEnum } from '@app-types/Card';
import { CardOwnerEnum } from '@app-types/CardOwner';
import { RangePipe } from '@pipes/range/range-pipe';
import type { ReceiveWinners } from '@services/api/types/messages/in/ReceiveWinners';
import { MatchService } from '@services/match/match';
import { concat, map, of, switchMap, timer } from 'rxjs';
import { WINNING_FX_DUR_SEC } from '../../consts';
import { Card } from '../card/card';

@Component({
  selector: 'app-table',
  imports: [Card, AsyncPipe, NgTemplateOutlet, RangePipe],
  templateUrl: './table.html',
})
export class Table {
  CARD_ENUM = CardEnum;
  CARD_OWNER_ENUM = CardOwnerEnum;

  private matchService = inject(MatchService);

  roundWinners = input<ReceiveWinners['payload']>();

  cards$ = this.matchService.tableCards$;
  winningCards$ = toObservable(this.roundWinners).pipe(
    map((winners) => winners?.flatMap(({ winningCards }) => winningCards)),
    switchMap((cards) =>
      concat(of(cards), timer(WINNING_FX_DUR_SEC * 1000).pipe(map(() => undefined))),
    ),
  );
}
