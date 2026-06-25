import { AsyncPipe, CurrencyPipe } from '@angular/common';
import { Component, inject, input } from '@angular/core';
import { toObservable } from '@angular/core/rxjs-interop';
import { CardsHand } from '@components/cards-hand/cards-hand';
import { MatchService } from '@services/match/match';
import { combineLatest, map } from 'rxjs';

@Component({
  selector: 'app-opponent',
  imports: [CardsHand, AsyncPipe, CurrencyPipe],
  templateUrl: './opponent.html',
})
export class Opponent {
  matchService = inject(MatchService);

  seat = input.required<number>();

  opponent$ = combineLatest([
    this.matchService.opponents$,
    this.matchService.seats$,
    toObservable(this.seat),
  ]).pipe(map(([opponents, seats, seat]) => opponents.find(({ id }) => id === seats[seat])));
}
