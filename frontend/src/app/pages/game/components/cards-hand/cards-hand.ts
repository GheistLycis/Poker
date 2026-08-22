import { AsyncPipe } from '@angular/common';
import { Component, input } from '@angular/core';
import { toObservable } from '@angular/core/rxjs-interop';
import type { Card as CardType } from '@app-types/Card';
import { CardEnum } from '@app-types/Card';
import { CardOwnerEnum } from '@app-types/CardOwner';
import type { Player } from '@app-types/Player';
import { distinctUntilChanged, filter, interval, map, switchMap, takeWhile } from 'rxjs';
import { Card } from '../card/card';

@Component({
  selector: 'app-cards-hand',
  imports: [Card, AsyncPipe],
  templateUrl: './cards-hand.html',
})
export class CardsHand {
  CARD_ENUM = CardEnum;
  CARD_OWNER_ENUM = CardOwnerEnum;

  player = input.required<Player>();
  isUser = input.required<boolean>();

  cards$ = toObservable(this.player).pipe(
    map(({ cards }) => cards),
    filter((cards) => !!cards.length),
    distinctUntilChanged((prev, curr) => prev.toString() === curr.toString()),
    switchMap((cards) =>
      // * "push each card to hand" effect
      interval(200).pipe(
        takeWhile((cardCount) => cardCount < 2, true),
        map((cardCount) => cards.slice(0, cardCount) as (CardType | null)[]),
      ),
    ),
  );
}
