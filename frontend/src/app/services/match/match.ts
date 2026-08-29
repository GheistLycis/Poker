import { computed, inject, Service } from '@angular/core';
import type { Player } from '@app-types/Player';
import type { PlayerAction } from '@app-types/PlayerAction';
import { PlayerActionEnum } from '@app-types/PlayerAction';
import type { SeatIndex } from '@app-types/SeatIndex';
import { HasPlayerChangedPipe } from '@pipes/has-player-changed/has-player-changed-pipe';
import { ApiService } from '@services/api/api';
import { WebSocketConnStateEnum } from '@services/api/types/ConnState';
import { combineLatest, distinctUntilChanged, filter, map, shareReplay } from 'rxjs';

@Service()
export class MatchService {
  private apiService = inject(ApiService);
  private hasPlayerChangedPipe = inject(HasPlayerChangedPipe);

  isLoading = computed(() => this.apiService.connState() === WebSocketConnStateEnum.CONNECTING);
  seats$ = this.apiService
    .getMessages('match.seats')
    .pipe(shareReplay({ bufferSize: 1, refCount: false }));
  opponents$ = this.apiService.getMessages('opponents.info').pipe(
    distinctUntilChanged((prev, curr) => {
      if (prev.length !== curr.length) return false;

      const sortedPrev = prev.toSorted((a, b) => a.id.localeCompare(b.id));
      const sortedCurr = curr.toSorted((a, b) => a.id.localeCompare(b.id));

      for (let i = 0; i < sortedPrev.length; i++) {
        if (this.hasPlayerChangedPipe.transform([sortedPrev[i], sortedCurr[i]])) {
          return false;
        }
      }
      return true;
    }),
    shareReplay({ bufferSize: 1, refCount: false }),
  );
  seatTurn$ = this.apiService.getMessages('match.seat-turn').pipe(
    map((msg) => msg.seatIndex),
    shareReplay({ bufferSize: 1, refCount: false }),
  );
  tableCards$ = this.apiService
    .getMessages('match.table-cards')
    .pipe(shareReplay({ bufferSize: 1, refCount: false }));
  roundWInners$ = this.apiService
    .getMessages('match.winners')
    .pipe(shareReplay({ bufferSize: 1, refCount: false }));
  pot$ = this.apiService.getMessages('match.pot-amount').pipe(
    map((msg) => msg.amount),
    shareReplay({ bufferSize: 1, refCount: false }),
  );

  isPlayerTurn(playerSeat: SeatIndex) {
    return this.seatTurn$.pipe(map((seatTurn) => seatTurn === playerSeat));
  }

  getPlayerAtSeat(seat: SeatIndex) {
    return combineLatest([this.opponents$, this.seats$]).pipe(
      map(([opponents, seats]) => opponents.find(({ id }) => id === seats[seat])),
    );
  }

  playerWon(playerId: Player['id']) {
    return this.roundWInners$.pipe(
      filter((winnersIds) => winnersIds.includes(playerId)),
      map(() => {}),
    );
  }

  registerUserAction(action: PlayerAction, amount?: number) {
    const type = 'user.action';

    if (action === PlayerActionEnum.BET) {
      if (amount === undefined)
        throw new Error(`Unexpected value: "amount" cannot be undefined in a ${action} action`);

      return this.apiService.send({ type, payload: { action, amount } });
    }
    return this.apiService.send({ type, payload: { action } });
  }
}
