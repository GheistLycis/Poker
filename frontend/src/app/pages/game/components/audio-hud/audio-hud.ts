import { NgClass } from '@angular/common';
import { Component, computed, effect, inject, signal } from '@angular/core';
import { CARD_SFX } from '@components/card/consts';
import { NgIcon, provideIcons } from '@ng-icons/core';
import { heroSpeakerWave, heroSpeakerXMark } from '@ng-icons/heroicons/outline';
import { AudioService } from '@services/audio/audio';
import { HlmSliderImports } from '@ui/slider';

@Component({
  selector: 'app-audio-hud',
  imports: [NgIcon, NgClass, HlmSliderImports],
  templateUrl: './audio-hud.html',
  providers: [provideIcons({ heroSpeakerWave, heroSpeakerXMark })],
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
