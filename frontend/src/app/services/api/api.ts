import { Service } from '@angular/core';
import { environment } from '@env';
import { webSocket, WebSocketSubject } from 'rxjs/webSocket';

@Service()
export class ApiService {
  API_URL = environment.apiUrl;

  private connection$: WebSocketSubject<>

  connect() {
    if(this.connection$) return
    
    this.connection$ = webSocket()
  }
}
