import {inject, Injectable} from '@angular/core';
import {UrlModel} from "../models/url-model";
import { HttpClient } from "@angular/common/http";
import {URLResponse} from "../models/url-response";
import {Observable} from "rxjs";
import {URLListResponse} from "../models/URLlist-model";
import {UsersResponse} from "../models/user-model";
import {environment} from "../../environments/environment";


@Injectable({
  providedIn: 'root'
})

export class ShortenerService {

  constructor() { }

  protected httpClient = inject(HttpClient)

  disableAdminMultipleURL(Urls: string[]): Observable<any> {
    return this.httpClient.put<URLResponse>(`${environment.apiUrl}/admindisablemultipleurl`, { codes: Urls },
      { withCredentials: true });
  }

  shorten(urlData: UrlModel): Observable<URLResponse> {
    return this.httpClient.post<URLResponse>(`${environment.apiUrl}/shorten`, urlData,
      { withCredentials: true });
  }

  getMyURLS(): Observable<URLListResponse>{
    return this.httpClient.get<URLListResponse>(`${environment.apiUrl}/getmyurls`,
      { withCredentials: true });
  }

  getAdminURLS(): Observable<URLListResponse>{
    return this.httpClient.get<URLListResponse>(`${environment.apiUrl}/getactive`,
      { withCredentials: true });
  }

  getAdminAllAccounts(): Observable<UsersResponse>{
    return this.httpClient.get<UsersResponse>(`${environment.apiUrl}/getusers`,
      { withCredentials: true });
  }

  disableURL(id: string) {
    var APIlink = `${environment.apiUrl}/disable/` + id
    return this.httpClient.put<any>(APIlink, {},
      { withCredentials: true });
  }

  enableURL(id: string) {
    var APIlink = `${environment.apiUrl}/enable/` + id
    return this.httpClient.put<any>(APIlink, {},
      { withCredentials: true });
  }

  disableAdminURL(id: string) {
    var APIlink = `${environment.apiUrl}/admindisableurl/` + id
    return this.httpClient.put<any>(APIlink, {},
      { withCredentials: true });
  }

  disableAdminAccount(id: string) {
    var APIlink = `${environment.apiUrl}/deleteaccountadmin/` + id
    return this.httpClient.delete<any>(APIlink,
      { withCredentials: true });
  }
}
