import {inject, Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {LoginModel} from "../models/login-model";
import {BehaviorSubject, catchError, Observable, of, tap} from "rxjs";
import {map} from "rxjs/operators";

@Injectable({
  providedIn: 'root'
})
export class AuthService {

  protected httpClient = inject(HttpClient)
  private loggedIn = new BehaviorSubject<boolean>(false);
  public isLoggedIn$ = this.loggedIn.asObservable();

  private admin = new BehaviorSubject<boolean>(false);
  public isAdmin$ = this.admin.asObservable();

  constructor() {
    this.checkLoginStatus();
  }

  getCreditAndUrls(): Observable<any> {
    return this.httpClient.get<any>(
      "http://localhost:8080/getcredit",
      { withCredentials: true }
    )
  }

  changeName(data : any): Observable<any> {
    return this.httpClient.put<any>(
      "http://localhost:8080/changename",
      data,
      { withCredentials: true }
    )
  }

  changeEmail(data : any): Observable<any> {
    return this.httpClient.put<any>(
      "http://localhost:8080/changeemail",
      data,
      { withCredentials: true }
    )
  }

  changePassword(data : any): Observable<any> {
    return this.httpClient.put<any>(
      "http://localhost:8080/changepassword",
      data,
      { withCredentials: true }
    )
  }

  login(logindata: LoginModel): Observable<any> {
    return this.httpClient.post<any>(
      "http://localhost:8080/login",
      logindata,
      { withCredentials: true }
    ).pipe(
      tap(() => { this.loggedIn.next(true)
      this.checkLoginStatus()
      })
    );
  }

  register(registerdata: LoginModel): Observable<any> {
    return this.httpClient.post<any>(
      "http://localhost:8080/signup",
      registerdata,
      { withCredentials: true }
    ).pipe(
      tap(() => { this.loggedIn.next(true)
        this.checkLoginStatus()
      })
    );
  }

  getuserInfo(): Observable<any> {
    return this.httpClient.get<any>(
      "http://localhost:8080/validate",
      { withCredentials: true }
    )
  }

  checkLoginStatus(): void {
    this.httpClient.get<any>(
      "http://localhost:8080/validate",
      { withCredentials: true }
    ).subscribe({
      next: () => this.loggedIn.next(true),
      error: () => this.loggedIn.next(false)
    });

    this.httpClient.get<any>(
      "http://localhost:8080/admincheck",
      { withCredentials: true }
    ).subscribe({
      next: () => this.admin.next(true),
      error: () => this.admin.next(false)
    });
  }

  logout(): void {
    this.httpClient.post<any>(
      "http://localhost:8080/logout", {},
      { withCredentials: true }
    ).subscribe({
      next: () => {
        this.loggedIn.next(false)
        this.admin.next(false)
      },
      error: () => {
        this.loggedIn.next(false)
        this.admin.next(false)
      }
    })
  }

  checkAdminStatus(): Observable<boolean> {
    return this.httpClient.get<any>("http://localhost:8080/admincheck", { withCredentials: true }).pipe(
      map(() => true),
      catchError(() => of(false))
    );
  }

  addCredit(credit: number): Observable<any> {
    return this.httpClient.put<any>(
      'http://localhost:8080/addtocredit',
      { addedCredit: credit },
      { withCredentials: true }
    );
  }
}


