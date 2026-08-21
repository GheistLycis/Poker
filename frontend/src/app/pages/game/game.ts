import { Component, computed, inject } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import type { SeatIndex } from '@app-types/SeatIndex';
import { Opponent } from '@components/opponent/opponent';
import { Table } from '@components/table/table';
import { User } from '@components/user/user';
import { MatchService } from '@services/match/match';
import { UserService } from '@services/user/user';
import { TOTAL_SEATS } from './consts';

@Component({
  selector: 'app-game',
  imports: [Opponent, MatProgressSpinnerModule, User, Table],
  templateUrl: './game.html',
})
export class Game {
  matchService = inject(MatchService);
  userService = inject(UserService);

  private seats = toSignal(this.matchService.seats$);
  opponentsSeats = computed(() => {
    const user = this.userService.user();

    if (!user) return [];

    const seats = Object.entries(this.seats() ?? {}).map<[SeatIndex, string | null]>(
      ([seatIndex, playerId]) => [+seatIndex as SeatIndex, playerId],
    );

    if (!seats.length) return [];

    const [userSeat] = seats.find(([_, userId]) => userId === user.id)!;
    const opponentsSeats = seats
      .filter(([seatIndex]) => seatIndex !== userSeat)
      .sort(
        ([seatA], [seatB]) =>
          ((seatA - userSeat + TOTAL_SEATS) % TOTAL_SEATS) -
          ((seatB - userSeat + TOTAL_SEATS) % TOTAL_SEATS),
      );

    return opponentsSeats.map(([seat], i) => {
      const angleDeg = 180 - i * (180 / (TOTAL_SEATS - 1));
      const rad = (angleDeg * Math.PI) / 180;

      return {
        seat,
        left: `${50 + 44 * Math.cos(rad)}%`,
        top: `${50 - 40 * Math.sin(rad)}%`,
      };
    });
  });
}
