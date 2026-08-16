import { computed, inject, Service } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { Router } from '@angular/router';
import { ApiService } from '@services/api/api';
import { take, tap } from 'rxjs';
import type { LoginPayload } from './types/LoginPayload';

@Service()
export class UserService {
  private router = inject(Router);
  private apiService = inject(ApiService);

  private user$ = this.apiService.getMessages('user.info');
  user = toSignal(this.user$);
  isLoggedIn = computed(() => !!this.user());

  logIn(payload: LoginPayload) {
    this.apiService.send({ type: 'user.login', payload });

    return this.user$.pipe(
      tap(() => this.router.navigate([''])),
      take(1),
    );
  }
}
