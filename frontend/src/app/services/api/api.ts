import { HttpClient } from '@angular/common/http';
import { inject, Service } from '@angular/core';
import { environment } from '@env';

@Service()
export class ApiService {
  API_URL = environment.apiUrl;

  http = inject(HttpClient);

  private queryMaker(query?: object): string {
    if (!query) return '';

    return (
      '&' +
      Object.entries(query)
        .filter(([, val]) => {
          if (Array.isArray(val)) return !!val.length;
          if (val === 0) return true;
          else return !!val;
        })
        .map(([key, val]) => key + '=' + val)
        .join('&')
    );
  }

  get<T>(endpoint: string, id: string, extras?: object) {
    return this.http.get<T>(`${this.API_URL}/${endpoint}/${id}?${this.queryMaker(extras)}`);
  }

  list<T>(endpoint: string, extras?: object) {
    return this.http.get<T[]>(`${this.API_URL}/${endpoint}?${this.queryMaker(extras)}`);
  }
}
