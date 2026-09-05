import { AsyncPipe, CurrencyPipe, NgClass, UpperCasePipe } from '@angular/common';
import { Component, inject, input } from '@angular/core';
import { toObservable } from '@angular/core/rxjs-interop';
import type { SeatIndex } from '@app-types/SeatIndex';
import { HandPipe } from '@pipes/hand/hand-pipe';
import type { ReceiveWinners } from '@services/api/types/messages/in/ReceiveWinners';
import { MatchService } from '@services/match/match';
import { combineLatest, concat, filter, map, of, switchMap, timer } from 'rxjs';
import { WINNING_FX_DUR_SEC } from '../../consts';
import { CardsHand } from '../cards-hand/cards-hand';

@Component({
  selector: 'app-opponent',
  imports: [CardsHand, AsyncPipe, CurrencyPipe, NgClass, HandPipe, UpperCasePipe],
  templateUrl: './opponent.html',
})
export class Opponent {
  private matchService = inject(MatchService);

  seat = input.required<SeatIndex>();
  roundWinners = input<ReceiveWinners['payload']>();

  private seat$ = toObservable(this.seat);
  opponent$ = this.seat$.pipe(
    switchMap((seat) => this.matchService.getPlayerAtSeat(seat)),
    filter((opponent) => !!opponent),
  );
  isOpponentTurn$ = this.seat$.pipe(switchMap((seat) => this.matchService.isPlayerTurn(seat)));
  private roundWinners$ = toObservable(this.roundWinners);
  opponentWon$ = combineLatest([this.roundWinners$, this.opponent$]).pipe(
    map(([winners, opponent]) => winners?.find(({ id }) => id === opponent.id)),
    filter((winningOpponent) => !!winningOpponent),
    switchMap((winningOpponent) =>
      concat(of(winningOpponent), timer(WINNING_FX_DUR_SEC * 1000).pipe(map(() => undefined))),
    ),
  );
}
