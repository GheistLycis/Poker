import { AsyncPipe, CurrencyPipe } from '@angular/common';
import { Component, inject } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { MatButtonModule } from '@angular/material/button';
import { CardsHand } from '@components/cards-hand/cards-hand';
import { UserService } from '@services/user/user';
import { interval, scan } from 'rxjs';

@Component({
  selector: 'app-user',
  imports: [CardsHand, AsyncPipe, CurrencyPipe, MatButtonModule],
  templateUrl: './user.html',
})
export class User {
  userService = inject(UserService);

  isUserTurn = toSignal(interval(2000).pipe(scan((prev) => !prev, false)), { initialValue: false });
}
