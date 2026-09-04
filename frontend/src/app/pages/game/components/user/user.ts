import { CurrencyPipe, UpperCasePipe } from '@angular/common';
import { Component, inject, model } from '@angular/core';
import { toObservable, toSignal } from '@angular/core/rxjs-interop';
import type { Card } from '@app-types/Card';
import { CardEnum } from '@app-types/Card';
import { HandEnum, type Hand } from '@app-types/Hand';
import type { PlayerAction } from '@app-types/PlayerAction';
import { PlayerActionEnum } from '@app-types/PlayerAction';
import { faker } from '@faker-js/faker';
import { HandPipe } from '@pipes/hand/hand-pipe';
import { MatchService } from '@services/match/match';
import { UserService } from '@services/user/user';
import { HlmButtonImports } from '@ui/button';
import { HlmLabel } from '@ui/label';
import { HlmSliderImports } from '@ui/slider';
import { concat, filter, map, of, Subject, switchMap, tap, timer } from 'rxjs';
import { PlayerActionPipe } from '../../../../pipes/player-action/player-action-pipe';
import { CardsHand } from '../cards-hand/cards-hand';
import { USER_WON_POPUP_DURATION_SEC } from './consts';

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
  ],
  templateUrl: './user.html',
  styleUrl: './user.css',
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
  // private userWon$ = this.user$.pipe(switchMap((user) => this.matchService.playerWon(user.id)));
  private userWon$ = new Subject<{ winningHand: Hand; winningCards: Card[] }>();
  userWon = toSignal(
    this.userWon$.pipe(
      switchMap((userWon) =>
        concat(of(userWon), timer(USER_WON_POPUP_DURATION_SEC * 1000).pipe(map(() => undefined))),
      ),
    ),
    { initialValue: undefined },
  );

  bet = model<[number]>([0]);

  constructor() {
    timer(2_000, 7_000)
      .pipe(
        tap(() =>
          this.userWon$.next({
            winningHand: faker.helpers.enumValue(HandEnum),
            winningCards: faker.helpers.arrayElements(Object.values(CardEnum)),
          }),
        ),
      )
      .subscribe();
  }

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
