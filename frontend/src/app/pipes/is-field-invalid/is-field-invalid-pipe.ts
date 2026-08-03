import { Pipe, PipeTransform } from '@angular/core';
import { FieldState } from '@angular/forms/signals';

@Pipe({ name: 'isFieldInvalid', pure: false })
export class IsFieldInvalidPipe implements PipeTransform {
  transform<T>(field: FieldState<T>): boolean {
    return field.touched() && field.errors().length > 0;
  }
}
