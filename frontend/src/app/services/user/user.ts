import { computed, inject, Service } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { Router } from '@angular/router';
import { User } from '@app-types/User';
import { ApiService } from '@services/api/api';
import { injectLocalStorage } from 'ngxtension/inject-local-storage';
import { filter, take, tap } from 'rxjs';
import { USER_STORAGE_KEY } from './consts';
import { LoginPayload } from './types/LoginPayload';

@Service()
export class UserService {
  private router = inject(Router);
  private apiService = inject(ApiService);

  private user$ = this.apiService.getMessages('user.info');
  user = injectLocalStorage<User>(USER_STORAGE_KEY);
  isLoggedIn = computed(() => !!this.user());

  constructor() {
    this.user$
      .pipe(
        tap((user) => this.user.set(user)),
        takeUntilDestroyed(),
      )
      .subscribe();
  }

  logIn(payload: LoginPayload) {
    this.apiService.send({ type: 'user.login', payload });

    return this.user$.pipe(
      filter((user) => !!user),
      tap(() => this.router.navigate([''])),
      take(1),
    );
  }
}
