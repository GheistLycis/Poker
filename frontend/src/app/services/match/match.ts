import { inject, Service } from '@angular/core';
import { UserService } from '@services/user/user';
import { of } from 'rxjs';
import { OPPONENTS, SEATS } from './consts';

@Service()
export class MatchService {
  userService = inject(UserService);

  seats$ = of<Record<number, string | null>>(SEATS);
  opponents$ = of(OPPONENTS);
}
