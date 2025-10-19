import {inject, Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {LoginModel} from "../models/login-model";
import {BehaviorSubject, Observable, tap} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  protected httpClient = inject(HttpClient)
  private loggedIn = new BehaviorSubject<boolean>(false);
  public isLoggedIn$ = this.loggedIn.asObservable();

  constructor() {
    this.checkLoginStatus();
  }

  login(logindata: LoginModel): Observable<any> {
    return this.httpClient.post<any>(
      "http://localhost:8080/login",
      logindata,
      { withCredentials: true }
    ).pipe(
      tap(() => this.loggedIn.next(true))
    );
  }

  register(registerdata: LoginModel): Observable<any> {
    return this.httpClient.post<any>(
      "http://localhost:8080/signup",
      registerdata,
      { withCredentials: true }
    ).pipe(
      tap(() => this.loggedIn.next(true))
    );
  }

  checkLoginStatus(): void {
    this.httpClient.get<any>(
      "http://localhost:8080/validate",
      { withCredentials: true }
    ).subscribe({
      next: () => this.loggedIn.next(true),
      error: () => this.loggedIn.next(false)
    });
  }

  logout(): void {
    this.httpClient.post<any>(
      "http://localhost:8080/logout", {},
      { withCredentials: true }
    ).subscribe({
      next: () => this.loggedIn.next(false),
      error: () => this.loggedIn.next(false)
    })
  }

}


