import { Component, inject, signal } from '@angular/core';
import { form, FormField, FormRoot, minLength, required, validate } from '@angular/forms/signals';
import { IsFieldInvalidPipe } from '@pipes/is-field-invalid/is-field-invalid-pipe';
import { UserService } from '@services/user/user';
import { HlmButtonImports } from '@ui/button';
import { HlmFieldImports } from '@ui/field';
import { HlmInputImports } from '@ui/input';
import { firstValueFrom, map } from 'rxjs';
import { IS_VALID_NAME, MIN_NAME_LEN } from './consts';
import type { LoginFormData } from './types/LoginFormData';

@Component({
  selector: 'app-login',
  imports: [
    FormRoot,
    FormField,
    HlmInputImports,
    HlmButtonImports,
    IsFieldInvalidPipe,
    HlmFieldImports,
  ],
  templateUrl: './login.html',
})
export class Login {
  private userService = inject(UserService);

  model = signal<LoginFormData>({ userName: '' });
  form = form(
    this.model,
    ({ userName }) => {
      required(userName);
      minLength(userName, MIN_NAME_LEN, { message: `Mínimo de ${MIN_NAME_LEN} caracteres` });
      validate(userName, ({ value }) =>
        IS_VALID_NAME.test(value())
          ? undefined
          : { kind: 'invalid-chars', message: 'Caracteres especiais não são permitidos' },
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
