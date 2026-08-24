import { computed, inject, Service } from '@angular/core';
import type { PlayerAction } from '@app-types/PlayerAction';
import { PlayerActionEnum } from '@app-types/PlayerAction';
import type { SeatIndex } from '@app-types/SeatIndex';
import { ApiService } from '@services/api/api';
import { WebSocketConnStateEnum } from '@services/api/types/ConnState';
import { map, shareReplay } from 'rxjs';

@Service()
export class MatchService {
  private apiService = inject(ApiService);

  isLoading = computed(() => this.apiService.connState() === WebSocketConnStateEnum.CONNECTING);
  seats$ = this.apiService
    .getMessages('match.seats')
    .pipe(shareReplay({ bufferSize: 1, refCount: false }));
  opponents$ = this.apiService
    .getMessages('opponents.info')
    .pipe(shareReplay({ bufferSize: 1, refCount: false }));
  seatTurn$ = this.apiService.getMessages('match.seat-turn').pipe(
    map((msg) => msg.seatIndex),
    shareReplay({ bufferSize: 1, refCount: false }),
  );
  tableCards$ = this.apiService
    .getMessages('match.table-cards')
    .pipe(shareReplay({ bufferSize: 1, refCount: false }));

  isPlayerTurn(playerSeat: SeatIndex) {
    return this.seatTurn$.pipe(map((seatTurn) => seatTurn === playerSeat));
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
