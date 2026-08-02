import { Routes } from '@angular/router';

import { isLoggedInGuard } from '@guards/is-logged-in/is-logged-in-guard';
import { isNotLoggedInGuard } from '@guards/is-not-logged-in/is-not-logged-in-guard';
import { Game } from './pages/game/game';
import { Login } from './pages/login/login';

export const routes: Routes = [
  { path: '', redirectTo: 'game', pathMatch: 'full' },
  { path: 'game', loadComponent: () => Game, canActivate: [isLoggedInGuard] },
  { path: 'login', loadComponent: () => Login, canActivate: [isNotLoggedInGuard] },
  { path: '**', redirectTo: '' },
];
