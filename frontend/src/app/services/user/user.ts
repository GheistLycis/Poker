import { computed, inject, Service, signal } from '@angular/core';
import { Router } from '@angular/router';
import { CreateUser, User } from '@classes/User';
import { ApiService } from '@services/api/api';
import { tap } from 'rxjs';
import { USER_STORAGE_KEY } from './consts';
import { LoginPayload } from './types/LoginPayload';

@Service()
export class UserService {
  private apiService = inject(ApiService);
  private router = inject(Router);

  user = signal<User | undefined>(undefined);
  isLoggedIn = computed(() => !!this.user());

  logIn(payload: LoginPayload) {
    this.apiService.send({ type: 'user.login', payload });

    return this.apiService.getMessages('user.info').pipe(
      tap((user) => {
        this.user.set(user);
        localStorage.setItem(USER_STORAGE_KEY, JSON.stringify(user));
        this.router.navigate(['']);
      }),
    );
  }

  constructor() {
    const storedUser = localStorage.getItem(USER_STORAGE_KEY);

    if (storedUser) {
      const user: CreateUser = JSON.parse(storedUser);

      this.user.set(new User(user));
    }
  }
}
