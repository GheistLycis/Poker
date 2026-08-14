import { computed, effect, inject, linkedSignal, Service } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { Router } from '@angular/router';
import { User } from '@app-types/User';
import { ApiService } from '@services/api/api';
import { take, tap } from 'rxjs';
import { USER_STORAGE_KEY } from './consts';
import { LoginPayload } from './types/LoginPayload';

@Service()
export class UserService {
  private apiService = inject(ApiService);
  private router = inject(Router);

  private user$ = this.apiService.getMessages('user.info');
  private userSig = toSignal(this.user$);
  user = linkedSignal(() => this.userSig());
  isLoggedIn = computed(() => !!this.user());

  logIn(payload: LoginPayload) {
    this.apiService.send({ type: 'user.login', payload });

    return this.user$.pipe(
      take(1),
      tap(() => this.router.navigate([''])),
    );
  }

  constructor() {
    const storedUser = localStorage.getItem(USER_STORAGE_KEY);

    if (storedUser) {
      const user: User = JSON.parse(storedUser);

      this.user.set(user);
    }

    effect(() => {
      const user = this.user();

      localStorage.setItem(USER_STORAGE_KEY, JSON.stringify(user));
    });
  }
}
