import { Component, signal } from '@angular/core';
import { CardsHand } from './components/cards-hand/cards-hand';

@Component({
  selector: 'app-root',
  imports: [CardsHand],
  templateUrl: './app.html',
})
export class App {
  protected readonly title = signal('frontend');
}
