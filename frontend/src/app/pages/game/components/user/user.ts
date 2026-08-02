import { AsyncPipe, CurrencyPipe } from '@angular/common';
import { Component, inject } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { PlayerActionEnum } from '@app-types/PlayerAction';
import { MatchService } from '@services/match/match';
import { UserService } from '@services/user/user';
import { PlayerActionPipe } from '../../../../pipes/player-action/player-action-pipe';
import { CardsHand } from '../cards-hand/cards-hand';

@Component({
  selector: 'app-user',
  imports: [PlayerActionPipe, CardsHand, AsyncPipe, CurrencyPipe, MatButtonModule],
  templateUrl: './user.html',
})
export class User {
  PLAYER_ACTIONS = Object.values(PlayerActionEnum);

  userService = inject(UserService);
  matchService = inject(MatchService);

  isUserTurn$ = this.matchService.isPlayerTurn(0);
}
