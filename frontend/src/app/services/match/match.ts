import { computed, inject, Service } from '@angular/core';
import { PlayerAction, PlayerActionEnum } from '@app-types/PlayerAction';
import { ApiService } from '@services/api/api';
import { WebSocketConnStateEnum } from '@services/api/types/ConnState';
import { filter, map, shareReplay, startWith } from 'rxjs';
import { combineLatest } from 'rxjs/internal/observable/combineLatest';
import { InConnMessage } from './types/ConnMessage';
import { ReceiveOpponentsHands } from './types/ReceiveOpponentsHands';

@Service()
export class MatchService {
  private apiService = inject(ApiService);

  isLoading = computed(() => this.apiService.connState() === WebSocketConnStateEnum.CONNECTING);
  seats$ = this.getMessages('match.seats').pipe(shareReplay());
  opponents$ = combineLatest([
    this.getMessages('opponents.info'),
    this.getMessages('opponents.reveal-hands').pipe(
      startWith<ReceiveOpponentsHands['payload']>({}),
    ),
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
  seatTurn$ = this.getMessages('match.seat-turn').pipe(
    map((msg) => msg.seatIndex),
    shareReplay(),
  );
  revealedCards$ = this.getMessages('match.table-cards').pipe(shareReplay());

  private getMessages<T extends InConnMessage['type']>(type: T) {
    type Message = Extract<InConnMessage, { type: T }>;

    return this.apiService.receivedMessages$.pipe(
      filter((msg): msg is Message => msg.type === type),
      map((msg) => msg.payload as Message['payload']),
    );
  }

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
