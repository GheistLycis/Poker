import { computed, DestroyRef, effect, inject, Service } from '@angular/core';
import { takeUntilDestroyed, toSignal } from '@angular/core/rxjs-interop';
import { Router } from '@angular/router';
import { HasPlayerChangedPipe } from '@pipes/has-player-changed/has-player-changed-pipe';
import { ApiService } from '@services/api/api';
import { MatchService } from '@services/match/match';
import { distinctUntilChanged, take, tap } from 'rxjs';
import { USER_STORAGE_KEY } from './consts';
import type { LoginPayload } from './types/LoginPayload';

@Service()
export class UserService {
  private destroyRef = inject(DestroyRef);
  private router = inject(Router);
  private apiService = inject(ApiService);
  private matchService = inject(MatchService);
  private hasPlayerChangedPipe = inject(HasPlayerChangedPipe);

  private user$ = this.apiService
    .getMessages('user.info')
    .pipe(distinctUntilChanged((prev, curr) => !this.hasPlayerChangedPipe.transform([prev, curr])));
  user = toSignal(this.user$);
  isLoggedIn = computed(() => !!this.user());

  constructor() {
    effect(() => {
      const storedUserName = sessionStorage.getItem(USER_STORAGE_KEY);

      if (storedUserName && !this.matchService.isLoading()) {
        this.logIn({ userName: storedUserName })
          .pipe(takeUntilDestroyed(this.destroyRef))
          .subscribe();
      }
    });
  }

  logIn(payload: LoginPayload) {
    this.apiService.send({ type: 'user.login', payload });

    return this.user$.pipe(
      tap((user) => sessionStorage.setItem(USER_STORAGE_KEY, user.name)),
      tap(() => this.router.navigate([''])),
      take(1),
    );
  }
}
