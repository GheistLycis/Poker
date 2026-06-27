import { Service } from '@angular/core';
import { of, shareReplay } from 'rxjs';
import { USER } from './consts';

@Service()
export class UserService {
  user$ = of(USER).pipe(shareReplay());
}
