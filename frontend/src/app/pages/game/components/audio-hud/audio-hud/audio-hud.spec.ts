import { ComponentFixture, TestBed } from '@angular/core/testing';

import { AudioHud } from './audio-hud';

describe('AudioHud', () => {
  let component: AudioHud;
  let fixture: ComponentFixture<AudioHud>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      imports: [AudioHud],
    }).compileComponents();

    fixture = TestBed.createComponent(AudioHud);
    component = fixture.componentInstance;
    await fixture.whenStable();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
