import { AsyncPipe } from '@angular/common';
import { Component, computed, inject, input } from '@angular/core';
import { toObservable } from '@angular/core/rxjs-interop';
import { CardEnum, Card as CardType } from '@app-types/Card';
import { CardOwnerEnum } from '@app-types/CardOwner';
import { Opponent } from '@classes/Opponent';
import { User } from '@classes/User';
import { MatchService } from '@services/match/match';
import { UserService } from '@services/user/user';
import { interval, map, switchMap, takeWhile } from 'rxjs';
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

  player = input.required<User | Opponent>();

  isUser = computed(() => this.player() instanceof User);
  // APPLYING "PUSH EACH CARD TO HAND" EFFECT
  cards$ = toObservable(this.player).pipe(
    switchMap(({ cards }) =>
      interval(200).pipe(
        takeWhile((cardCount) => cardCount < 2, true),
        map((cardCount) => cards.slice(0, cardCount) as (CardType | null)[]),
      ),
    ),
  );
}
