import {Component, inject, OnDestroy} from '@angular/core';
import {DatePipe, NgForOf, NgIf} from "@angular/common";
import {ShortenerService} from "../services/shortener-service.service";
import {URLItem, URLListResponse} from "../models/URLlist-model";
import {UsersResponse} from "../models/user-model";
import {interval, switchMap} from "rxjs";

@Component({
  selector: 'app-admin-panel',
  standalone: true,
  imports: [
    DatePipe,
    NgForOf,
    NgIf
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

  protected urlsToDisable : string[] = []

  private urlSub! : any;

  protected onCheckboxClick(event : any, shortcode: string) {
    if (event.target.checked){
      this.urlsToDisable.push(shortcode)
    } else {
      this.urlsToDisable = this.urlsToDisable.filter(code => code !== shortcode);
    }
    console.log(this.urlsToDisable)
  }

  protected disableMultipleURLS() {
    this.shortenerService.disableAdminMultipleURL(this.urlsToDisable).subscribe({
      next: (res : any)=> {
        this.refreshUrls()
        this.urlsToDisable = []
        alert("Successfully disabled urls!")
      },
      error: (err : any) => {
        alert(err.error.error)
      }
    })
  }

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
    this.refreshUrls()

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
        this.urlsToDisable = []
      }
    })
  }

  protected disableAccountAdmin(id: string){
    this.shortenerService.disableAdminAccount(id).subscribe({
      next: (res : any)=> {
        this.getAccounts()
        this.getUrls()
        this.userDisableMessage = res.Message
      },
      error: (err : any) => {
        alert(err.error.Error)
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
