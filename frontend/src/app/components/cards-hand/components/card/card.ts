import { NgOptimizedImage } from '@angular/common';
import { Component, computed, input } from '@angular/core';
import { CardEnum, Card as CardType } from './types/Card';

@Component({
  selector: 'app-card',
  imports: [NgOptimizedImage],
  templateUrl: './card.html',
})
export class Card {
  variant = input<CardType>(CardEnum.BACK);
  fileType = computed<string>(() => (this.variant() == CardEnum.BACK ? '.png' : '.svg'));
}
