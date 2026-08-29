import type { PipeTransform } from '@angular/core';
import { Pipe } from '@angular/core';

@Pipe({ name: 'range' })
export class RangePipe implements PipeTransform {
  transform(value: number): number[] {
    return Array.from({ length: value }).map((_, i) => i);
  }
}
