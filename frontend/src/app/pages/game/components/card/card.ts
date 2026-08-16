import { NgClass, NgOptimizedImage } from '@angular/common';
import { Component, computed, input } from '@angular/core';
import type { Card as CardType } from '@app-types/Card';
import { CardEnum } from '@app-types/Card';
import type { CardOwner} from '@app-types/CardOwner';
import { CardOwnerEnum } from '@app-types/CardOwner';

@Component({
  selector: 'app-card',
  imports: [NgOptimizedImage, NgClass],
  templateUrl: './card.html',
})
export class Card {
  variant = input.required<CardType>();
  owner = input.required<CardOwner>();
  class = input('');

  imgFileType = computed(() => (this.variant() === CardEnum.BACK ? '.png' : '.svg'));
  size = computed(() => {
    let height = 140;
    let width = height * 0.71;

    if (this.owner() !== CardOwnerEnum.OPPONENT) {
      height *= 1.5;
      width *= 1.5;
    }

    return { width, height };
  });
  isCardBack = computed(() => this.variant() === CardEnum.BACK);
}
