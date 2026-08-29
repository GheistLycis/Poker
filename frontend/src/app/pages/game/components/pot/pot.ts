import { CurrencyPipe, NgOptimizedImage, NgStyle } from '@angular/common';
import { Component, computed, inject } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { RangePipe } from '@pipes/range/range-pipe';
import { MatchService } from '@services/match/match';
import type { Stack } from './types/Stack';

@Component({
  selector: 'app-pot',
  imports: [NgOptimizedImage, CurrencyPipe, NgStyle],
  providers: [RangePipe],
  templateUrl: './pot.html',
})
export class Pot {
  private matchService = inject(MatchService);
  private rangePipe = inject(RangePipe);

  amount = toSignal(this.matchService.pot$, { initialValue: 0 });
  private stacksCount = computed(() => {
    const amount = this.amount();

    if (!amount) return 0;
    if (amount < 100) return 1;
    if (amount < 500) return 2;
    if (amount < 1_000) return 3;
    if (amount < 10_000) return 4;
    if (amount < 50_000) return 5;
    return 6;
  });
  stacks = computed<Stack[]>(() =>
    this.rangePipe.transform(this.stacksCount()).map((i) => ({
      variant: i % 2 ? 'black' : 'red',
      position: i % 2 ? 'back' : 'front',
    })),
  );
}
