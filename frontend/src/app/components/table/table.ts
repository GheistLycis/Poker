import { AsyncPipe } from '@angular/common';
import { Component, inject } from '@angular/core';
import { CardEnum } from '@app-types/Card';
import { CardOwnerEnum } from '@app-types/CardOwner';
import { Card } from '@components/card/card';
import { MatchService } from '@services/match/match';
import { map } from 'rxjs';

@Component({
  selector: 'app-table',
  imports: [Card, AsyncPipe],
  templateUrl: './table.html',
})
export class Table {
  CARD_ENUM = CardEnum;
  CARD_OWNER_ENUM = CardOwnerEnum;

  matchService = inject(MatchService);

  cards$ = this.matchService.revealedCards$.pipe(
    map((cards) => Array.from({ length: 5 }, (_, i) => cards[i] ?? CardEnum.BACK)),
  );
}
