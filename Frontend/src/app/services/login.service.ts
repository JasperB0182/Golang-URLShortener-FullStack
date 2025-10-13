import {inject, Injectable} from '@angular/core';
import {HttpClient} from "@angular/common/http";
import {LoginModel} from "../models/login-model";
import {Observable} from "rxjs";

@Injectable({
  providedIn: 'root'
})
export class LoginService {

  constructor() { }

  protected httpClient = inject(HttpClient)

  login(logindata: LoginModel) : Observable<any>{
    return this.httpClient.post<any>("http://localhost:8080/login", logindata,
    { withCredentials: true })
  }
}
