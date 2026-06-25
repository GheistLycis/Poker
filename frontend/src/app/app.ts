import { AsyncPipe } from '@angular/common';
import { Component, inject } from '@angular/core';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { Opponent } from '@components/opponent/opponent';
import { User } from '@components/user/user';
import { MatchService } from '@services/match/match';
import { map } from 'rxjs';

interface OpponentSeat {
  seat: number;
  left: string;
  top: string;
}

@Component({
  selector: 'app-root',
  imports: [Opponent, AsyncPipe, MatProgressSpinnerModule, User],
  templateUrl: './app.html',
})
export class App {
  matchService = inject(MatchService);

  opponentsSeats$ = this.matchService.seats$.pipe(
    map((seatsMap) => {
      const opponentsCount = Object.keys(seatsMap).length - 1;

      return Array.from({ length: opponentsCount }, (_, i) => {
        const angleDeg = opponentsCount === 1 ? 90 : 180 - i * (180 / (opponentsCount - 1));
        const rad = (angleDeg * Math.PI) / 180;

        return {
          seat: i + 1,
          left: `${50 + 44 * Math.cos(rad)}%`,
          top: `${50 - 40 * Math.sin(rad)}%`,
        };
      }) as OpponentSeat[];
    }),
  );
}
