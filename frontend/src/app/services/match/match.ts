import { inject, Service } from '@angular/core';
import { Card, CardEnum } from '@app-types/Card';
import { PlayerAction } from '@app-types/PlayerAction';
import { Opponent } from '@classes/Opponent';
import { faker } from '@faker-js/faker';
import { ApiService } from '@services/api/api';
import { BehaviorSubject, interval, map, of, shareReplay, switchMap, tap } from 'rxjs';
import { takeWhile } from 'rxjs/operators';
import { OPPONENTS, SEATS } from './consts';

@Service()
export class MatchService {
  private apiService = inject(ApiService);

  seats$ = of<Record<number, string | null>>(SEATS).pipe(shareReplay());
  opponents$ = new BehaviorSubject(OPPONENTS);
  seatTurn$ = this.seats$.pipe(
    map((seats) =>
      Object.entries(seats)
        .filter(([, opponentId]) => !!opponentId)
        .map(([seat]) => +seat),
    ),
    switchMap((seats) => interval(2000).pipe(map(() => faker.helpers.arrayElement(seats)))),
    shareReplay(),
  );
  revealedCards$ = new BehaviorSubject<Card[]>([]);

  constructor() {
    interval(2000)
      .pipe(
        takeWhile(() => this.revealedCards$.value.length < 5),
        tap(() => this.revealTableCards()),
      )
      .subscribe();
  }

  private revealOpponentsCards() {
    const opponents = this.opponents$.value;
    const updatedOpponents = opponents.map(
      (opponent) =>
        new Opponent({
          ...opponent,
          cards: [faker.helpers.enumValue(CardEnum), faker.helpers.enumValue(CardEnum)],
        }),
    );

    this.opponents$.next(updatedOpponents);
  }

  private revealTableCards() {
    const newCard = faker.helpers.enumValue(CardEnum);
    const current = this.revealedCards$.value;

    this.revealedCards$.next([...current, newCard]);
  }

  isPlayerTurn(playerSeat: number) {
    return this.seatTurn$.pipe(map((seatTurn) => seatTurn === playerSeat));
  }

  registerUserAction(_action: PlayerAction) {}
}
