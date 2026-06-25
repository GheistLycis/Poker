import { AsyncPipe } from '@angular/common';
import { Component, inject, input } from '@angular/core';
import { toObservable } from '@angular/core/rxjs-interop';
import { CardEnum, Card as CardType } from '@app-types/Card';
import { CardOwnerEnum } from '@app-types/CardOwner';
import { MatchService } from '@services/match/match';
import { UserService } from '@services/user/user';
import { combineLatest, interval, map, Observable, startWith, switchMap, takeWhile } from 'rxjs';
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

  hand$: Observable<(CardType | null)[]> = combineLatest([
    this.userService.user$,
    this.matchService.opponents$,
    this.matchService.seats$,
    toObservable(this.seat),
  ]).pipe(
    map(([user, opponents, seats, seat]) => {
      if (seat == 0) return user.cards;

      const opponent = opponents.find((opponent) => opponent.id == seats[seat]);

      if (!opponent) return [];

      return Array.from({ length: opponent.handSize }).fill(null) as null[];
    }),
    switchMap((fullHand) =>
      interval(200).pipe(
        startWith(0),
        map((_, tick) => tick + 1),
        takeWhile((cardCount) => cardCount < fullHand.length, true),
        map((cardCount) => fullHand.slice(0, cardCount)),
      ),
    ),
  );
}
