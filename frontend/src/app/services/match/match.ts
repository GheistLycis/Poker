import { computed, inject, Service } from '@angular/core';
import type { PlayerAction } from '@app-types/PlayerAction';
import { PlayerActionEnum } from '@app-types/PlayerAction';
import { ApiService } from '@services/api/api';
import { WebSocketConnStateEnum } from '@services/api/types/ConnState';
import type { ReceiveOpponentsHands } from '@services/api/types/messages/in/ReceiveOpponentsHands';
import { map, shareReplay, startWith } from 'rxjs';
import { combineLatest } from 'rxjs/internal/observable/combineLatest';

@Service()
export class MatchService {
  private apiService = inject(ApiService);

  isLoading = computed(() => this.apiService.connState() === WebSocketConnStateEnum.CONNECTING);
  seats$ = this.apiService
    .getMessages('match.seats')
    .pipe(shareReplay({ bufferSize: 1, refCount: false }));
  opponents$ = combineLatest([
    this.apiService.getMessages('opponents.info'),
    this.apiService
      .getMessages('opponents.reveal-hands')
      .pipe(startWith<ReceiveOpponentsHands['payload']>({})),
  ]).pipe(
    map(([opponents, hands]) => {
      Object.entries(hands).forEach(([id, hand]) => {
        const opponent = opponents.find((opponent) => opponent.id === id);

        if (opponent) opponent.cards = hand;
      });

      return opponents;
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

  isPlayerTurn(playerSeat: number) {
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
