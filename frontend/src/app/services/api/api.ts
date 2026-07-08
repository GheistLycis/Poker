import { Service, signal } from '@angular/core';
import { environment } from '@env';
import { ConnMessage, InConnMessage, OutConnMessage } from '@services/match/types/ConnMessage';
import { filter, Observable, retry, shareReplay, timer } from 'rxjs';
import { webSocket, WebSocketSubject } from 'rxjs/webSocket';
import { WebSocketConnState, WebSocketConnStateEnum } from './types/ConnState';

@Service()
export class ApiService {
  private API_URL = environment.apiUrl;

  private connection$: WebSocketSubject<ConnMessage>;
  receivedMessages$: Observable<InConnMessage>;
  connState = signal<WebSocketConnState>(WebSocketConnStateEnum.CLOSE);

  constructor() {
    this.connState.set(WebSocketConnStateEnum.CONNECTING);
    this.connection$ = webSocket<ConnMessage>({
      url: this.API_URL,
      openObserver: { next: () => this.connState.set(WebSocketConnStateEnum.OPEN) },
      closeObserver: { next: () => this.connState.set(WebSocketConnStateEnum.CLOSE) },
    });
    this.receivedMessages$ = this.connection$.pipe(
      retry({ delay: (_, count) => timer(Math.min(1000 * 2 ** count, 30_000)) }),
      filter((msg) => msg.origin === 'SERVER'),
      shareReplay(),
    );
  }

  send(msg: OutConnMessage): void {
    this.connection$.next(msg);
  }

  close(): void {
    this.connection$.complete();
  }
}
