import { computed, inject, Service } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { ApiService } from '@services/api/api';

@Service()
export class UserService {
  private apiService = inject(ApiService);

  user = toSignal(this.apiService.getMessages('user.info'));
  isLoggedIn = computed(() => !!this.user());
}
