import { AsyncPipe, CurrencyPipe, UpperCasePipe } from '@angular/common';
import { Component, computed, inject, input, linkedSignal } from '@angular/core';
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
import { WINNING_FX_MS } from '../../consts';
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
      concat(of(winningUser), timer(WINNING_FX_MS).pipe(map(() => undefined))),
    ),
  );
  roundLastBet = toSignal(this.matchService.lastBet$, { initialValue: 0 });
  minBet = computed(() => this.roundLastBet() + 1);
  bet = linkedSignal(() => this.minBet());

  incrementBet(amount: number, operation: '-' | '+') {
    let final = this.bet();

    if (operation === '-') {
      const minBet = this.minBet();

      final -= amount;
      if (final < minBet) final = minBet;
    } else {
      const maxBet = this.user()!.score;

      final += amount;
      if (final > maxBet) final = maxBet;
    }

    this.bet.set(final);
  }

  sendAction(action: PlayerAction) {
    const bet = action === PlayerActionEnum.BET ? this.bet() : undefined;

    this.matchService.registerUserAction(action, bet);
    this.bet.set(0);
  }
}
