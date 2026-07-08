import { inject, Service } from '@angular/core';
import { ApiService } from '@services/api/api';

@Service()
export class UserService {
  private apiService = inject(ApiService);

  user$ = this.apiService.getMessages('user.info');
}
