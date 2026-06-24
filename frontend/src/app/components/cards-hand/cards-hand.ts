import { AsyncPipe } from '@angular/common';
import { Component } from '@angular/core';
import { concatMap, delay, from, of, scan } from 'rxjs';
import { Card } from './components/card/card';
import { Card as CardType } from './components/card/types/Card';
import { CARDS, PUSH_TO_HAND_DELAY_MS } from './consts';

@Component({
  selector: 'app-cards-hand',
  imports: [Card, AsyncPipe],
  templateUrl: './cards-hand.html',
  styleUrl: './cards-hand.css',
})
export class CardsHand {
  cards = from(CARDS).pipe(
    concatMap((card, i) => of(card).pipe(delay(i ? PUSH_TO_HAND_DELAY_MS : 0))),
    scan((cards, card) => [...cards, card], [] as CardType[]),
  );
}
