import {inject, Injectable} from '@angular/core';
import {UrlModel} from "../models/url-model";
import {HttpClient} from "@angular/common/http";
import {URLResponse} from "../models/url-response";
import {Observable} from "rxjs";

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
}
