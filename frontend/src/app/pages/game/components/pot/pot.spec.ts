import { ComponentFixture, TestBed } from '@angular/core/testing';

import { Pot } from './pot';

describe('Pot', () => {
  let component: Pot;
  let fixture: ComponentFixture<Pot>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [Pot],
    }).compileComponents();

    fixture = TestBed.createComponent(Pot);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
