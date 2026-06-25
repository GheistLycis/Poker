import { AsyncPipe } from '@angular/common';
import { Component, inject } from '@angular/core';
import { MatProgressSpinner } from '@angular/material/progress-spinner';
import { CardsHand } from '@components/cards-hand/cards-hand';
import { MatchService } from '@services/match/match';
import { map } from 'rxjs';

interface Opponent {
  seat: number;
  left: string;
  top: string;
}

@Component({
  selector: 'app-root',
  imports: [CardsHand, AsyncPipe, MatProgressSpinner],
  templateUrl: './app.html',
})
export class App {
  matchService = inject(MatchService);

  opponents$ = this.matchService.seats$.pipe(
    map((seatsMap) => {
      const opponentsCount = Object.keys(seatsMap).length - 1;

      return Array.from({ length: opponentsCount }, (_, i) => {
        const angleDeg = opponentsCount == 1 ? 90 : 180 - i * (180 / (opponentsCount - 1));
        const rad = (angleDeg * Math.PI) / 180;

        return {
          seat: i + 1,
          left: `${50 + 45 * Math.cos(rad)}%`,
          top: `${50 - 40 * Math.sin(rad)}%`,
        };
      }) as Opponent[];
    }),
  );
}
