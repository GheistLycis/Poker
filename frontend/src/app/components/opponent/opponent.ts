import { AsyncPipe, CurrencyPipe, NgClass } from '@angular/common';
import { Component, inject, input } from '@angular/core';
import { toObservable } from '@angular/core/rxjs-interop';
import { CardsHand } from '@components/cards-hand/cards-hand';
import { MatchService } from '@services/match/match';
import { combineLatest, map, switchMap } from 'rxjs';

@Component({
  selector: 'app-opponent',
  imports: [CardsHand, AsyncPipe, CurrencyPipe, NgClass],
  templateUrl: './opponent.html',
})
export class Opponent {
  matchService = inject(MatchService);

  seat = input.required<number>();

  seat$ = toObservable(this.seat);
  opponent$ = combineLatest([
    this.matchService.opponents$,
    this.matchService.seats$,
    this.seat$,
  ]).pipe(map(([opponents, seats, seat]) => opponents.find(({ id }) => id === seats[seat])));
  isOpponentTurn$ = this.seat$.pipe(switchMap((seat) => this.matchService.isPlayerTurn(seat)));
}
