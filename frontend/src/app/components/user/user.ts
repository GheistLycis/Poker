import { AsyncPipe, CurrencyPipe } from '@angular/common';
import { Component, inject } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { CardsHand } from '@components/cards-hand/cards-hand';
import { MatchService } from '@services/match/match';
import { UserService } from '@services/user/user';

@Component({
  selector: 'app-user',
  imports: [CardsHand, AsyncPipe, CurrencyPipe, MatButtonModule],
  templateUrl: './user.html',
})
export class User {
  userService = inject(UserService);
  matchService = inject(MatchService);

  isUserTurn$ = this.matchService.isPlayerTurn(0);
}
