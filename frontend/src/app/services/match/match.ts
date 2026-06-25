import { inject, Service } from '@angular/core';
import { Card } from '@app-types/Card';
import { UserService } from '@services/user/user';
import { concatMap, delay, from, of, scan, switchMap } from 'rxjs';
import { OPPONENTS, PUSH_CARD_TO_HAND_DELAY_MS, SEATS } from './consts';

@Service()
export class MatchService {
  userService = inject(UserService);

  seats$ = of<Record<number, string | null>>(SEATS);
  opponents$ = of(OPPONENTS);
  userHand$ = this.userService.user$.pipe(
    switchMap((user) => from(user.cards)),
    concatMap((card, i) => of(card).pipe(delay(i ? PUSH_CARD_TO_HAND_DELAY_MS : 0))),
    scan((cards, card) => [...cards, card], [] as Card[]),
  );
}
