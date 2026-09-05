import { AsyncPipe, CurrencyPipe, UpperCasePipe } from '@angular/common';
import { Component, inject, input, model } from '@angular/core';
import { toObservable, toSignal } from '@angular/core/rxjs-interop';
import type { PlayerAction } from '@app-types/PlayerAction';
import { PlayerActionEnum } from '@app-types/PlayerAction';
import { HandPipe } from '@pipes/hand/hand-pipe';
import type { ReceiveWinners } from '@services/api/types/messages/in/ReceiveWinners';
import { MatchService } from '@services/match/match';
import { UserService } from '@services/user/user';
import { HlmButtonImports } from '@ui/button';
import { HlmLabel } from '@ui/label';
import { HlmSliderImports } from '@ui/slider';
import { combineLatest, concat, filter, map, of, switchMap, timer } from 'rxjs';
import { PlayerActionPipe } from '../../../../pipes/player-action/player-action-pipe';
import { WINNING_FX_DUR_SEC } from '../../consts';
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
    HandPipe,
    UpperCasePipe,
    AsyncPipe,
  ],
  templateUrl: './user.html',
})
export class User {
  PLAYER_ACTION_ENUM = PlayerActionEnum;
  PLAYER_ACTIONS = Object.values(PlayerActionEnum);

  private userService = inject(UserService);
  private matchService = inject(MatchService);

  roundWinners = input<ReceiveWinners['payload']>();

  user = this.userService.user;
  private user$ = toObservable(this.userService.user).pipe(filter((user) => !!user));
  isUserTurn = toSignal(
    this.user$.pipe(switchMap((user) => this.matchService.isPlayerTurn(user.seatIndex))),
  );
  private roundWinners$ = toObservable(this.roundWinners);
  userWon$ = combineLatest([this.roundWinners$, this.user$]).pipe(
    map(([winners, user]) => winners?.find(({ id }) => id === user.id)),
    filter((winningUser) => !!winningUser),
    switchMap((winningUser) =>
      concat(of(winningUser), timer(WINNING_FX_DUR_SEC * 1000).pipe(map(() => undefined))),
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
    this.bet.set([0]);
  }
}
