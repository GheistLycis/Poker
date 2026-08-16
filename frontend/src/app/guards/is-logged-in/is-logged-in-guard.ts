import { inject } from '@angular/core';
import type { CanActivateFn } from '@angular/router';
import { Router } from '@angular/router';
import { UserService } from '@services/user/user';

export const isLoggedInGuard: CanActivateFn = () => {
  const router = inject(Router);
  const userService = inject(UserService);

  return userService.isLoggedIn() || router.createUrlTree(['/login']);
};
