import { NgClass, NgOptimizedImage, NgStyle } from '@angular/common';
import { Component, computed, inject, input } from '@angular/core';
import type { Card as CardType } from '@app-types/Card';
import { CardEnum } from '@app-types/Card';
import type { CardOwner } from '@app-types/CardOwner';
import { CardOwnerEnum } from '@app-types/CardOwner';
import { AudioService } from '@services/audio/audio';
import { BIGGER_CARD_PROPORTION, CARD_HEIGHT_PX, CARD_SFX, CARD_WIDTH_PX } from './consts';

@Component({
  selector: 'app-card',
  imports: [NgOptimizedImage, NgClass, NgStyle],
  templateUrl: './card.html',
  styleUrl: './card.css',
})
export class Card {
  private audioService = inject(AudioService);

  variant = input.required<CardType | null>();
  owner = input.required<CardOwner>();
  class = input('');
  highlight = input(false);

  imgFileType = computed(() => (this.variant() === CardEnum.BACK ? '.png' : '.svg'));
  size = computed(() => {
    let height = CARD_HEIGHT_PX;
    let width = CARD_WIDTH_PX;

    if (this.owner() === CardOwnerEnum.USER) {
      height *= BIGGER_CARD_PROPORTION;
      width *= BIGGER_CARD_PROPORTION;
    }

    return { width, height };
  });
  isCardBack = computed(() => this.variant() === CardEnum.BACK);

  constructor() {
    this.audioService.play(CARD_SFX);
  }
}
