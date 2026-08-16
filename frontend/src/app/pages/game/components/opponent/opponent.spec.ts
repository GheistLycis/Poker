import type { ComponentFixture} from '@angular/core/testing';
import { TestBed } from '@angular/core/testing';

import { Opponent } from './opponent';

describe('Opponent', () => {
  let component: Opponent;
  let fixture: ComponentFixture<Opponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Opponent],
    }).compileComponents();

    fixture = TestBed.createComponent(Opponent);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
