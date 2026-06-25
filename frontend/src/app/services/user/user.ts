import { Service } from '@angular/core';
import { of } from 'rxjs';
import { USER } from './consts';

@Service()
export class UserService {
  user$ = of(USER);
}
