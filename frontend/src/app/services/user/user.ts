import { Service } from '@angular/core';
import { BehaviorSubject } from 'rxjs';
import { USER } from './consts';

@Service()
export class UserService {
  user$ = new BehaviorSubject(USER);
}
