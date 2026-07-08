import { computed, inject, Service } from '@angular/core';
import { PlayerAction, PlayerActionEnum } from '@app-types/PlayerAction';
import { ApiService } from '@services/api/api';
import { WebSocketConnStateEnum } from '@services/api/types/ConnState';
import { ReceiveOpponentsHands } from '@services/api/types/messages/ReceiveOpponentsHands';
import { map, shareReplay, startWith } from 'rxjs';
import { combineLatest } from 'rxjs/internal/observable/combineLatest';

@Service()
export class MatchService {
  private apiService = inject(ApiService);

  isLoading = computed(() => this.apiService.connState() === WebSocketConnStateEnum.CONNECTING);
  seats$ = this.apiService.getMessages('match.seats').pipe(shareReplay());
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
    shareReplay(),
  );
  seatTurn$ = this.apiService.getMessages('match.seat-turn').pipe(
    map((msg) => msg.seatIndex),
    shareReplay(),
  );
  revealedCards$ = this.apiService.getMessages('match.table-cards').pipe(shareReplay());

  isPlayerTurn(playerSeat: number) {
    return this.seatTurn$.pipe(map((seatTurn) => seatTurn === playerSeat));
  }

  registerUserAction(action: PlayerAction, amount?: number) {
    const baseMessage = { origin: 'CLIENT', type: 'user.action' } as const;

    if (action === PlayerActionEnum.BET || action === PlayerActionEnum.RAISE)
      return this.apiService.send({
        ...baseMessage,
        payload: { action, amount: amount ?? 0 },
      });
    this.apiService.send({
      ...baseMessage,
      payload: { action },
    });
  }
}
