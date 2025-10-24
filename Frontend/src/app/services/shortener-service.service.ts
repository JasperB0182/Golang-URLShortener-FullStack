import {inject, Injectable} from '@angular/core';
import {UrlModel} from "../models/url-model";
import {HttpClient} from "@angular/common/http";
import {URLResponse} from "../models/url-response";
import {Observable} from "rxjs";
import {URLListResponse} from "../models/URLlist-model";
import {UsersResponse} from "../models/user-model";


@Injectable({
  providedIn: 'root'
})

export class ShortenerService {

  constructor() { }

  protected httpClient = inject(HttpClient)



  shorten(urlData: UrlModel): Observable<URLResponse> {
    return this.httpClient.post<URLResponse>("http://localhost:8080/shorten", urlData,
      { withCredentials: true });
  }

  getMyURLS(): Observable<URLListResponse>{
    return this.httpClient.get<URLListResponse>("http://localhost:8080/getmyurls",
      { withCredentials: true });
  }

  getAdminURLS(): Observable<URLListResponse>{
    return this.httpClient.get<URLListResponse>("http://localhost:8080/getactive",
      { withCredentials: true });
  }

  getAdminAllAccounts(): Observable<UsersResponse>{
    return this.httpClient.get<UsersResponse>("http://localhost:8080/getusers",
      { withCredentials: true });
  }

  disableURL(id: string) {
    var APIlink = "http://localhost:8080/disable/" + id
    return this.httpClient.put<any>(APIlink, {},
      { withCredentials: true });
  }

  disableAdminURL(id: string) {
    var APIlink = "http://localhost:8080/admindisableurl/" + id
    return this.httpClient.put<any>(APIlink, {},
      { withCredentials: true });
  }

  disableAdminAccount(id: string) {
    var APIlink = "http://localhost:8080/deleteaccountadmin/" + id
    return this.httpClient.delete<any>(APIlink,
      { withCredentials: true });
  }
}
