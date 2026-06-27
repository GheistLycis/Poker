import { NgClass, NgOptimizedImage } from '@angular/common';
import { Component, computed, input } from '@angular/core';
import { CardEnum, Card as CardType } from '@app-types/Card';
import { CardOwner, CardOwnerEnum } from '@app-types/CardOwner';

@Component({
  selector: 'app-card',
  imports: [NgOptimizedImage, NgClass],
  templateUrl: './card.html',
})
export class Card {
  CARD_ENUM = CardEnum;
  CARD_OWNER_ENUM = CardOwnerEnum;

  variant = input.required<CardType>();
  owner = input.required<CardOwner>();
  class = input('');

  cardImgFileType = computed(() => (this.variant() === CardEnum.BACK ? '.png' : '.svg'));
}
