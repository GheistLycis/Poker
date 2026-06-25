import { AsyncPipe } from '@angular/common';
import { Component, inject, input } from '@angular/core';
import { toObservable } from '@angular/core/rxjs-interop';
import { CardEnum, Card as CardType } from '@app-types/Card';
import { CardOwnerEnum } from '@app-types/CardOwner';
import { MatchService } from '@services/match/match';
import { UserService } from '@services/user/user';
import { combineLatest, interval, map, startWith, switchMap, takeWhile } from 'rxjs';
import { Card } from './components/card/card';

@Component({
  selector: 'app-cards-hand',
  imports: [Card, AsyncPipe],
  templateUrl: './cards-hand.html',
  styleUrl: './cards-hand.css',
})
export class CardsHand {
  CARD_ENUM = CardEnum;
  CARD_OWNER_ENUM = CardOwnerEnum;

  userService = inject(UserService);
  matchService = inject(MatchService);

  seat = input.required<number>();

  hand$ = combineLatest([this.userService.user$, toObservable(this.seat)]).pipe(
    map(([user, seat]) => (seat === 0 ? user.cards : [null, null])),
    // PUSH CARD TO HAND EFFECT
    switchMap((fullHand) =>
      interval(200).pipe(
        startWith(0),
        takeWhile((cardCount) => cardCount < 2, true),
        map((cardCount) => fullHand.slice(0, cardCount) as [CardType, CardType] | [null, null]),
      ),
    ),
  );
}
