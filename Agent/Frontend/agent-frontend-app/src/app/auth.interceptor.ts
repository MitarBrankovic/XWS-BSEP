import {
    HttpEvent,
    HttpHandler,
    HttpInterceptor,
    HttpRequest,
    HTTP_INTERCEPTORS,
  } from '@angular/common/http';
  import { Injectable } from '@angular/core';
  import { Observable } from 'rxjs';
  
  @Injectable()
  export class AuthInterceptor implements HttpInterceptor {
    constructor() {}
  
    intercept(
      request: HttpRequest<unknown>,
      next: HttpHandler
    ): Observable<HttpEvent<unknown>> {
      request = request.clone({
        headers: request.headers.set('Authorization', 'Bearer ' + localStorage.getItem('jwtToken')?.slice(1, -1)),
      });
  
      return next.handle(request);
    }
  }
  
  export const AuthInterceptorProvider = {
    provide: HTTP_INTERCEPTORS,
    useClass: AuthInterceptor,
    multi: true,
  };