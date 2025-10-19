import {inject, Injectable} from '@angular/core';
import {UrlModel} from "../models/url-model";
import {HttpClient} from "@angular/common/http";
import {URLResponse} from "../models/url-response";
import {Observable} from "rxjs";
import {URLListResponse} from "../models/URLlist-model";


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

  disableURL(id: string) {
    var APIlink = "http://localhost:8080/disable/" + id
    return this.httpClient.put<any>(APIlink, {},
      { withCredentials: true });
  }
}
