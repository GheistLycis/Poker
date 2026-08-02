import { inject } from '@angular/core';
import { CanActivateFn, Router } from '@angular/router';
import { UserService } from '@services/user/user';

export const isNotLoggedInGuard: CanActivateFn = () => {
  const router = inject(Router);
  const userService = inject(UserService);

  return !userService.isLoggedIn() || router.createUrlTree(['/']);
};
