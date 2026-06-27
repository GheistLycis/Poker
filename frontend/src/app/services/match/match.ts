import { effect, inject, Service } from '@angular/core';
import { CardEnum } from '@app-types/Card';
import { Opponent } from '@classes/Opponent';
import { faker } from '@faker-js/faker';
import { UserService } from '@services/user/user';
import { BehaviorSubject, interval, map, of, shareReplay, switchMap, tap, timer } from 'rxjs';
import { OPPONENTS, SEATS } from './consts';
@Service()
export class MatchService {
  userService = inject(UserService);

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

  constructor() {
    effect(() => {
      timer(3000, 3000)
        .pipe(
          tap(() => {
            const newOpponents: Opponent[] = [];

            this.opponents$.getValue().forEach((opponent) => {
              const newOpponent = new Opponent(opponent.id, opponent.name, opponent.score);

              newOpponent.cards = [
                faker.helpers.enumValue(CardEnum),
                faker.helpers.enumValue(CardEnum),
              ];
              newOpponents.push(newOpponent);
            });
            this.opponents$.next(newOpponents);
          }),
        )
        .subscribe();
    });
  }

  isPlayerTurn(playerSeat: number) {
    return this.seatTurn$.pipe(map((seatTurn) => seatTurn === playerSeat));
  }
}
