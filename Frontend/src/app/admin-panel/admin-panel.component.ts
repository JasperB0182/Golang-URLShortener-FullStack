import {Component, inject, OnDestroy} from '@angular/core';
import {DatePipe, NgForOf} from "@angular/common";
import {ShortenerService} from "../services/shortener-service.service";
import {URLItem, URLListResponse} from "../models/URLlist-model";
import {UsersResponse} from "../models/user-model";
import {interval, switchMap} from "rxjs";

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
export class AdminPanelComponent implements OnDestroy{
  protected shortenerService = inject(ShortenerService)

  protected myURLS: URLItem[] = [];
  protected myAccounts!: UsersResponse;

  protected urlDisableMessage = ""
  protected userDisableMessage = ""

  private urlSub! : any;

  constructor() {
    this.getUrls()
    this.getAccounts()
  }

  ngOnDestroy(){
    this.urlSub.unsubscribe();
  }

  protected refreshUrls(){
    this.shortenerService.getAdminURLS().subscribe({
      next: (res: URLListResponse) => (this.myURLS = res.Code)
    });
  }

  protected getUrls(){
    this.shortenerService.getAdminURLS().subscribe({
      next: (res: URLListResponse) => (this.myURLS = res.Code)
    });

    this.urlSub = interval(5000)
      .pipe(switchMap(() => this.shortenerService.getAdminURLS()))
      .subscribe({
        next: (res: URLListResponse) => {
          this.myURLS = res.Code;
        }
      })
  }

  protected disableURL(id: string){
    this.shortenerService.disableAdminURL(id).subscribe({
      next: (res : any)=> {
        this.refreshUrls()
        this.urlDisableMessage = res.Message
      }
    })
  }

  protected disableAccountAdmin(id: string){
    this.shortenerService.disableAdminAccount(id).subscribe({
      next: (res : any)=> {
        this.getAccounts()
        this.userDisableMessage = res.Message
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
