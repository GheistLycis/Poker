import { inject, Service } from '@angular/core';
import { faker } from '@faker-js/faker';
import { UserService } from '@services/user/user';
import { interval, map, of, shareReplay, switchMap } from 'rxjs';
import { OPPONENTS, SEATS } from './consts';

@Service()
export class MatchService {
  userService = inject(UserService);

  seats$ = of<Record<number, string | null>>(SEATS);
  opponents$ = of(OPPONENTS);
  seatTurn$ = this.seats$.pipe(
    map((seats) =>
      Object.entries(seats)
        .filter(([, opponentId]) => !!opponentId)
        .map(([seat]) => +seat),
    ),
    switchMap((seats) => interval(2000).pipe(map(() => faker.helpers.arrayElement(seats)))),
    shareReplay(),
  );

  isPlayerTurn(playerSeat: number) {
    return this.seatTurn$.pipe(map((seatTurn) => seatTurn === playerSeat));
  }
}
