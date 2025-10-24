import {Component, inject} from '@angular/core';
import {DatePipe, NgForOf} from "@angular/common";
import {ShortenerService} from "../services/shortener-service.service";
import {URLItem, URLListResponse} from "../models/URLlist-model";
import {UsersResponse} from "../models/user-model";

@Component({
  selector: 'app-admin-panel',
  standalone: true,
    imports: [
        DatePipe,
        NgForOf
    ],
  templateUrl: './admin-panel.component.html',
  styleUrl: './admin-panel.component.scss'
})
export class AdminPanelComponent {
  protected shortenerService = inject(ShortenerService)

  protected myURLS: URLItem[] = [];
  protected myAccounts!: UsersResponse;

  constructor() {
    this.getUrls()
    this.getAccounts()
  }

  protected getUrls(){
    this.shortenerService.getAdminURLS().subscribe({
      next: (res: URLListResponse) => {
        this.myURLS = res.Code;
      }
    });
  }

  protected disableURL(id: string){
    this.shortenerService.disableAdminURL(id).subscribe({
      next: (res : any)=> {
        this.getUrls()
      }
    })
  }

  protected disableAccountAdmin(id: string){
    this.shortenerService.disableAdminAccount(id).subscribe({
      next: (res : any)=> {
        this.getAccounts()
      }
    })
  }

  protected getAccounts() {
    this.shortenerService.getAdminAllAccounts().subscribe({
      next: (res: UsersResponse) => {
        this.myAccounts = res;
      }
    });
  }

  protected readonly String = String;
}
