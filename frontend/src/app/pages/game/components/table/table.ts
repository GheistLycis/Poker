import { AsyncPipe, NgTemplateOutlet } from '@angular/common';
import { Component, inject } from '@angular/core';
import { CardEnum } from '@app-types/Card';
import { CardOwnerEnum } from '@app-types/CardOwner';
import { RangePipe } from '@pipes/range/range-pipe';
import { MatchService } from '@services/match/match';
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

  cards$ = this.matchService.tableCards$;
}
