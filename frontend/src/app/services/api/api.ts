import { Service, signal } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { environment } from '@env';
import type { Observable } from 'rxjs';
import { filter, map, of, retry, shareReplay, switchMap, take, throwError, timer } from 'rxjs';
import { webSocket } from 'rxjs/webSocket';
import type { WebSocketConnState } from './types/ConnState';
import { WebSocketConnStateEnum } from './types/ConnState';
import type { ConnMessage, InConnMessage } from './types/messages/ConnMessage';
import type { SendMessage } from './types/SendMessage';

@Service()
export class ApiService {
  private API_URL = environment.apiUrl;

  connState = signal<WebSocketConnState>(WebSocketConnStateEnum.CONNECTING);
  private connection$ = webSocket<ConnMessage>({
    url: this.API_URL,
    openObserver: { next: () => this.connState.set(WebSocketConnStateEnum.OPEN) },
    closeObserver: { next: () => this.connState.set(WebSocketConnStateEnum.CLOSE) },
  });
  readonly receivedMessages$ = this.connection$.pipe(
    retry({ delay: (_, count) => timer(Math.min(1000 * 2 ** count, 30_000)) }),
    filter((msg) => msg.origin === 'SERVER'),
    shareReplay({ bufferSize: 1, refCount: false }),
  );

  constructor() {
    this.receivedMessages$.pipe(takeUntilDestroyed()).subscribe();
  }

  getMessages<T extends InConnMessage['type']>(type: T) {
    type Message = Extract<InConnMessage, { type: T }>;

    return this.receivedMessages$.pipe(
      filter((msg): msg is Message => msg.type === type),
      map((msg) => msg.payload as Message['payload']),
    );
  }

  send(msg: SendMessage): Observable<InConnMessage> {
    const requestId = crypto.randomUUID();

    this.connection$.next({ ...msg, origin: 'CLIENT', requestId });

    return this.receivedMessages$.pipe(
      filter((msg) => msg.requestId === requestId),
      take(1),
      switchMap((msg) => (msg.error ? throwError(() => msg.error) : of(msg))),
    );
  }
}
