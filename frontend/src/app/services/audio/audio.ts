import { Service, signal } from '@angular/core';
import { AUDIO_VOLUME_STORAGE_KEY } from './consts';
import type { AudioSettings } from './types/AudioSettings';

@Service()
export class AudioService {
  private audioContext = new AudioContext();
  private bufferCache = new Map<string, Promise<AudioBuffer>>();
  private _settings = signal<AudioSettings>({ isEnabled: false, volume: 1 });

  settings = this._settings.asReadonly();

  constructor() {
    const storedAudioVol = localStorage.getItem(AUDIO_VOLUME_STORAGE_KEY);

    if (storedAudioVol && !isNaN(parseFloat(storedAudioVol))) {
      this.changeVolume(+storedAudioVol);
    }
  }

  toggleIsEnabled(enable?: boolean) {
    this._settings.update((prev) => ({
      ...prev,
      isEnabled: enable !== undefined ? enable : !prev.isEnabled,
    }));
  }

  changeVolume(volume: number) {
    if (volume < 0) volume = 0;
    if (volume > 1) volume = 1;
    this._settings.update((prev) => ({ ...prev, volume }));
    localStorage.setItem(AUDIO_VOLUME_STORAGE_KEY, volume.toString());
  }

  preload(url: string) {
    return this.loadBuffer(url).then(() => undefined);
  }

  async play(url: string) {
    const { isEnabled, volume } = this._settings();

    if (!isEnabled || !volume) return;

    const buffer = await this.loadBuffer(url);
    const source = this.audioContext.createBufferSource();
    const gain = this.audioContext.createGain();

    source.buffer = buffer;
    gain.gain.value = this._settings().volume;
    source.connect(gain).connect(this.audioContext.destination);
    source.start(0);
  }

  private loadBuffer(url: string) {
    let cached = this.bufferCache.get(url);

    if (!cached) {
      cached = fetch(url)
        .then((res) => res.arrayBuffer())
        .then((data) => this.audioContext.decodeAudioData(data));

      this.bufferCache.set(url, cached);
    }

    return cached;
  }
}
