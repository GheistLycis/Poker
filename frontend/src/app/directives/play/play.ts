import { afterNextRender, DestroyRef, Directive, inject, input } from '@angular/core';
import { AudioService } from '@services/audio/audio';
import type { AudioInput } from './types/AudioInput';

@Directive({ selector: '[appPlay]' })
export class Play {
  private destroyRef = inject(DestroyRef);
  private audioService = inject(AudioService);

  playOnEnter = input<string | AudioInput | undefined>(undefined);
  playOnLeave = input<string | AudioInput | undefined>(undefined);

  constructor() {
    afterNextRender(() => {
      const mountSfx = this.playOnEnter();

      if (!mountSfx) return;
      if (typeof mountSfx === 'string') this.audioService.play(mountSfx);
      else this.audioService.play(mountSfx.src, mountSfx.start, mountSfx.end);
    });
    this.destroyRef.onDestroy(() => {
      const unmountSfx = this.playOnLeave();

      if (!unmountSfx) return;
      if (typeof unmountSfx === 'string') this.audioService.play(unmountSfx);
      else this.audioService.play(unmountSfx.src, unmountSfx.start, unmountSfx.end);
    });
  }
}
