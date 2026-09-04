import { CurrencyPipe } from '@angular/common';
import { Component, inject, model } from '@angular/core';
import { outputFromObservable, toObservable, toSignal } from '@angular/core/rxjs-interop';
import type { PlayerAction } from '@app-types/PlayerAction';
import { PlayerActionEnum } from '@app-types/PlayerAction';
import { MatchService } from '@services/match/match';
import { UserService } from '@services/user/user';
import { HlmButtonImports } from '@ui/button';
import { HlmLabel } from '@ui/label';
import { HlmSliderImports } from '@ui/slider';
import { filter, map, switchMap } from 'rxjs';
import { PlayerActionPipe } from '../../../../pipes/player-action/player-action-pipe';
import { CardsHand } from '../cards-hand/cards-hand';

@Component({
  selector: 'app-user',
  imports: [
    PlayerActionPipe,
    CardsHand,
    CurrencyPipe,
    HlmButtonImports,
    HlmSliderImports,
    HlmLabel,
  ],
  templateUrl: './user.html',
})
export class User {
  PLAYER_ACTION_ENUM = PlayerActionEnum;
  PLAYER_ACTIONS = Object.values(PlayerActionEnum);

  private userService = inject(UserService);
  private matchService = inject(MatchService);

  user = this.userService.user;
  private user$ = toObservable(this.userService.user).pipe(filter((user) => !!user));
  isUserTurn = toSignal(
    this.user$.pipe(switchMap((user) => this.matchService.isPlayerTurn(user.seatIndex))),
  );
  userWon = outputFromObservable(
    this.user$.pipe(
      switchMap((user) => this.matchService.playerWon(user.id).pipe(map(() => user))),
    ),
  );
  bet = model<[number]>([0]);

  incrementBet(amount: number, operation: '-' | '+') {
    let final = this.bet()[0];

    if (operation === '-') {
      final -= amount;
      if (final < 0) final = 0;
    } else {
      const curr = this.user()!.score;

      final += amount;
      if (final > curr) final = curr;
    }

    this.bet.set([final]);
  }

  sendAction(action: PlayerAction) {
    this.matchService.registerUserAction(action, this.bet()[0] || undefined);
  }
}
