import { AsyncPipe, CurrencyPipe, NgClass } from '@angular/common';
import { Component, inject, input } from '@angular/core';
import { outputFromObservable, toObservable } from '@angular/core/rxjs-interop';
import type { SeatIndex } from '@app-types/SeatIndex';
import { MatchService } from '@services/match/match';
import { filter, map, switchMap } from 'rxjs';
import { CardsHand } from '../cards-hand/cards-hand';

@Component({
  selector: 'app-opponent',
  imports: [CardsHand, AsyncPipe, CurrencyPipe, NgClass],
  templateUrl: './opponent.html',
})
export class Opponent {
  private matchService = inject(MatchService);

  seat = input.required<SeatIndex>();

  private seat$ = toObservable(this.seat);
  opponent$ = this.seat$.pipe(
    switchMap((seat) => this.matchService.getPlayerAtSeat(seat)),
    filter((opponent) => !!opponent),
  );
  isOpponentTurn$ = this.seat$.pipe(switchMap((seat) => this.matchService.isPlayerTurn(seat)));
  opponentWon = outputFromObservable(
    this.opponent$.pipe(
      switchMap((opponent) => this.matchService.playerWon(opponent.id).pipe(map(() => opponent))),
    ),
  );
}
