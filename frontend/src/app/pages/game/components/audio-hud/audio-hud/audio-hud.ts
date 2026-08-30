import { NgClass } from '@angular/common';
import { Component, computed, effect, inject, signal } from '@angular/core';
import { MatIconModule } from '@angular/material/icon';
import { MatSliderModule } from '@angular/material/slider';
import { CARD_SFX } from '@components/card/consts';
import { AudioService } from '@services/audio/audio';

@Component({
  selector: 'app-audio-hud',
  imports: [MatIconModule, MatSliderModule, NgClass],
  templateUrl: './audio-hud.html',
  host: { '(document:click)': 'hasInteractedWithDom.set(true)' },
})
export class AudioHud {
  private audioService = inject(AudioService);

  hasInteractedWithDom = signal(false);
  audioIsEnabled = computed(() => this.audioService.settings().isEnabled);
  volume = computed(() => this.audioService.settings().volume);

  constructor() {
    this.audioService.preload(CARD_SFX);

    effect(() => {
      if (this.hasInteractedWithDom()) this.toggleAudio(true);
    });
  }

  toggleAudio(enable?: boolean) {
    this.audioService.toggleIsEnabled(enable);
  }

  changeVolume(volume: number) {
    this.audioService.changeVolume(volume);
  }
}
