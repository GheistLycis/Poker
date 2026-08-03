import { Service, signal } from '@angular/core';
import { environment } from '@env';
import { filter, map, Observable, retry, shareReplay, timer } from 'rxjs';
import { webSocket, WebSocketSubject } from 'rxjs/webSocket';
import { WebSocketConnState, WebSocketConnStateEnum } from './types/ConnState';
import { ConnMessage, InConnMessage } from './types/messages/ConnMessage';
import { SendMessage } from './types/SendMessage';

@Service()
export class ApiService {
  private API_URL = environment.apiUrl;

  private connection$: WebSocketSubject<ConnMessage>;
  readonly receivedMessages$: Observable<InConnMessage>;
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

  getMessages<T extends InConnMessage['type']>(type: T) {
    type Message = Extract<InConnMessage, { type: T }>;

    return this.receivedMessages$.pipe(
      filter((msg): msg is Message => msg.type === type),
      map((msg) => msg.payload as Message['payload']),
    );
  }

  send(msg: SendMessage): void {
    this.connection$.next({ ...msg, origin: 'CLIENT' });
  }

  close(): void {
    this.connection$.complete();
  }
}
