import { computed, DestroyRef, effect, inject, Service } from '@angular/core';
import { takeUntilDestroyed, toSignal } from '@angular/core/rxjs-interop';
import { Router } from '@angular/router';
import { ApiService } from '@services/api/api';
import { MatchService } from '@services/match/match';
import { take, tap } from 'rxjs';
import { USER_STORAGE_KEY } from './consts';
import type { LoginPayload } from './types/LoginPayload';

@Service()
export class UserService {
  private destroyRef = inject(DestroyRef);
  private router = inject(Router);
  private apiService = inject(ApiService);
  private matchService = inject(MatchService);

  private user$ = this.apiService.getMessages('user.info');
  user = toSignal(this.user$);
  isLoggedIn = computed(() => !!this.user());

  logIn(payload: LoginPayload) {
    this.apiService.send({ type: 'user.login', payload });

    return this.user$.pipe(
      tap((user) => sessionStorage.setItem(USER_STORAGE_KEY, user.name)),
      tap(() => this.router.navigate([''])),
      take(1),
    );
  }

  constructor() {
    effect(() => {
      if (this.matchService.isLoading()) return;

      const storedUserName = sessionStorage.getItem(USER_STORAGE_KEY);

      if (storedUserName)
        this.logIn({ userName: storedUserName })
          .pipe(takeUntilDestroyed(this.destroyRef))
          .subscribe();
    });
  }
}
