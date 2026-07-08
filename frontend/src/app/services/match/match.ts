import { inject, Service } from '@angular/core';
import { CardEnum } from '@app-types/Card';
import { PlayerAction, PlayerActionEnum } from '@app-types/PlayerAction';
import { Opponent } from '@classes/Opponent';
import { faker } from '@faker-js/faker';
import { ApiService } from '@services/api/api';
import { WebSocketConnStateEnum } from '@services/api/types/ConnState';
import { filter, interval, map, merge, of, scan, shareReplay, switchMap } from 'rxjs';
import { SEATS } from './consts';
import { InConnMessage } from './types/ConnMessage';

@Service()
export class MatchService {
  private apiService = inject(ApiService);

  seats$ = of<Record<number, string | null>>(SEATS).pipe(shareReplay());
  opponents$ = merge(
    this.getMessages('opponents.info').pipe(map((msg) => (): Opponent[] => msg.payload)),
    this.getMessages('opponents.reveal-hands').pipe(
      map(
        (msg) =>
          (opponents: Opponent[]): Opponent[] =>
            opponents.map(
              (opponent) =>
                new Opponent({
                  ...opponent,
                  cards: [faker.helpers.enumValue(CardEnum), faker.helpers.enumValue(CardEnum)],
                }),
            ),
      ),
    ),
  ).pipe(
    scan((opponents, reduce) => reduce(opponents), [] as Opponent[]),
    shareReplay(),
  );
  seatTurn$ = this.seats$.pipe(
    map((seats) =>
      Object.entries(seats)
        .filter(([, opponentId]) => !!opponentId)
        .map(([seat]) => +seat),
    ),
    switchMap((seats) => interval(2000).pipe(map(() => faker.helpers.arrayElement(seats)))),
    shareReplay(),
  );
  revealedCards$ = this.getMessages('match.table-cards').pipe(map((msg) => msg.payload));

  constructor() {
    if (this.apiService.connState() === WebSocketConnStateEnum.CLOSE) this.apiService.connect();
  }

  private getMessages<T extends InConnMessage['type']>(type: T) {
    return this.apiService.receivedMessages$.pipe(
      filter((msg): msg is Extract<InConnMessage, { type: T }> => msg.type === type),
      shareReplay(),
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
