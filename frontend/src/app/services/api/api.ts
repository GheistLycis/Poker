import { DestroyRef, inject, Service, signal } from '@angular/core';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { environment } from '@env';
import { ConnMessage, InConnMessage, OutConnMessage } from '@services/match/types/ConnMessage';
import { filter, retry, Subject, tap, timer } from 'rxjs';
import { webSocket, WebSocketSubject } from 'rxjs/webSocket';
import { WebSocketConnState, WebSocketConnStateEnum } from './types/ConnState';

@Service()
export class ApiService {
  private API_URL = environment.apiUrl;

  private destroyRef = inject(DestroyRef);

  private connection$?: WebSocketSubject<ConnMessage>;
  connState = signal<WebSocketConnState>(WebSocketConnStateEnum.CLOSE);
  receivedMessages$ = new Subject<InConnMessage>();

  connect() {
    if (this.connection$) return;

    this.connState.set(WebSocketConnStateEnum.CONNECTING);
    this.connection$ = webSocket<ConnMessage>({
      url: this.API_URL,
      openObserver: { next: () => this.connState.set(WebSocketConnStateEnum.OPEN) },
      closeObserver: { next: () => this.connState.set(WebSocketConnStateEnum.CLOSE) },
    });
    this.connection$
      .pipe(
        retry({ delay: (_, count) => timer(Math.min(1000 * 2 ** count, 30_000)) }),
        filter((msg) => msg.origin === 'SERVER'),
        tap((msg) => this.receivedMessages$.next(msg)),
        takeUntilDestroyed(this.destroyRef),
      )
      .subscribe();
  }

  send(msg: OutConnMessage): void {
    if (!this.connection$) return;

    this.connection$.next(msg);
  }

  close(): void {
    this.connection$?.complete();
    this.connection$ = undefined;
  }
}
