import {inject, Injectable} from '@angular/core';
import { HttpClient } from "@angular/common/http";
import {environment} from "../../environments/environment";

@Injectable({
  providedIn: 'root'
})
export class RedirectService {

  constructor() { }

  protected httpClient = inject(HttpClient)

  redirect(id : string){
    return this.httpClient.get<any>(`${environment.apiUrl}/link/` + id)
  }
}
