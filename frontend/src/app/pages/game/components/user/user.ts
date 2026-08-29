import { AsyncPipe, CurrencyPipe } from '@angular/common';
import { Component, inject } from '@angular/core';
import { outputFromObservable, toObservable } from '@angular/core/rxjs-interop';
import { MatButtonModule } from '@angular/material/button';
import type { PlayerAction } from '@app-types/PlayerAction';
import { PlayerActionEnum } from '@app-types/PlayerAction';
import { MatchService } from '@services/match/match';
import { UserService } from '@services/user/user';
import { filter, map, switchMap } from 'rxjs';
import { PlayerActionPipe } from '../../../../pipes/player-action/player-action-pipe';
import { CardsHand } from '../cards-hand/cards-hand';

@Component({
  selector: 'app-user',
  imports: [PlayerActionPipe, CardsHand, AsyncPipe, CurrencyPipe, MatButtonModule],
  templateUrl: './user.html',
})
export class User {
  PLAYER_ACTIONS = Object.values(PlayerActionEnum);

  private userService = inject(UserService);
  private matchService = inject(MatchService);

  user$ = toObservable(this.userService.user).pipe(filter((user) => !!user));
  isUserTurn$ = this.user$.pipe(
    switchMap((user) => this.matchService.isPlayerTurn(user.seatIndex)),
  );
  userWon = outputFromObservable(
    this.user$.pipe(
      switchMap((user) => this.matchService.playerWon(user.id).pipe(map(() => user))),
    ),
  );

  sendAction(action: PlayerAction, amount?: number) {
    this.matchService.registerUserAction(action, amount);
  }
}
