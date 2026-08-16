import { Component, inject, signal } from '@angular/core';
import { form, FormField, FormRoot, minLength, required, validate } from '@angular/forms/signals';
import { MatButton } from '@angular/material/button';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { IsFieldInvalidPipe } from '@pipes/is-field-invalid/is-field-invalid-pipe';
import { UserService } from '@services/user/user';
import { firstValueFrom, map } from 'rxjs';
import type { LoginFormData } from './types/LoginFormData';

@Component({
  selector: 'app-login',
  imports: [FormRoot, MatFormFieldModule, FormField, MatInputModule, MatButton, IsFieldInvalidPipe],
  templateUrl: './login.html',
})
export class Login {
  private userService = inject(UserService);

  model = signal<LoginFormData>({ userName: '' });
  form = form(
    this.model,
    ({ userName }) => {
      required(userName);
      minLength(userName, 4, { message: 'Mínimo de 4 caracteres' });
      validate(userName, ({ value }) =>
        !/^[a-z]+$/i.test(value())
          ? { kind: 'alphabet-only', message: 'Apenas letras são permitidas' }
          : undefined,
      );
    },
    {
      submission: {
        action: (form) => {
          const val = form().value();

          return firstValueFrom(this.userService.logIn(val).pipe(map(() => undefined)));
        },
      },
    },
  );
}
