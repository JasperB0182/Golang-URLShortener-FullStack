import {inject, Injectable} from '@angular/core';
import { HttpClient } from "@angular/common/http";

@Injectable({
  providedIn: 'root'
})
export class RedirectService {

  constructor() { }

  protected httpClient = inject(HttpClient)

  redirect(id : string){
    return this.httpClient.get<any>("http://localhost:8080/link/" + id)
  }
}
