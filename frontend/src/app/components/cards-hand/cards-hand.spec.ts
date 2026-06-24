import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CardsHand } from './cards-hand';

describe('CardsHand', () => {
  let component: CardsHand;
  let fixture: ComponentFixture<CardsHand>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [CardsHand],
    }).compileComponents();

    fixture = TestBed.createComponent(CardsHand);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
