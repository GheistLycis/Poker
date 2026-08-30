import { NgOptimizedImage } from '@angular/common';
import { Component, computed, inject } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import type { Player } from '@app-types/Player';
import type { SeatIndex } from '@app-types/SeatIndex';
import { AudioHud } from '@components/audio-hud/audio-hud/audio-hud';
import { Opponent } from '@components/opponent/opponent';
import { Pot } from '@components/pot/pot';
import { Table } from '@components/table/table';
import { User } from '@components/user/user';
import { MatchService } from '@services/match/match';
import { UserService } from '@services/user/user';

@Component({
  selector: 'app-game',
  imports: [Opponent, User, Table, NgOptimizedImage, Pot, AudioHud],
  templateUrl: './game.html',
})
export class Game {
  private matchService = inject(MatchService);
  private userService = inject(UserService);

  private seats = toSignal(this.matchService.seats$);
  opponentsSeats = computed(() => {
    const seats = Object.entries(this.seats() ?? {}).map<[SeatIndex, string | null]>(
      ([seatIndex, playerId]) => [+seatIndex as SeatIndex, playerId],
    );
    const user = this.userService.user();

    if (!user || !seats.length) return [];

    const [userSeat] = seats.find(([_, userId]) => userId === user.id)!;
    const opponentsSeats = seats.filter(([seatIndex]) => seatIndex !== userSeat);
    const opponentsSeatsCount = opponentsSeats.length;
    const tableAngleStep = 180 / (opponentsSeatsCount - 1);

    return opponentsSeats
      .sort(
        ([seatA], [seatB]) =>
          ((seatA - userSeat + opponentsSeatsCount) % opponentsSeatsCount) -
          ((seatB - userSeat + opponentsSeatsCount) % opponentsSeatsCount),
      )
      .map(([seat], i) => {
        const angleDeg = 180 - i * tableAngleStep;
        const rad = (angleDeg * Math.PI) / 180;

        return {
          seat,
          left: `${50 + 40 * Math.cos(rad)}%`,
          top: `${50 - 40 * Math.sin(rad)}%`,
        };
      });
  });

  onPlayerWon(player: Player) {
    // TODO
    console.log(player);
  }
}
