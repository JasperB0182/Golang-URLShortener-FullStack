import {Component, inject} from '@angular/core';
import {DatePipe, NgForOf} from "@angular/common";
import {ShortenerService} from "../services/shortener-service.service";
import {URLItem, URLListResponse} from "../models/URLlist-model";

@Component({
  selector: 'app-profile',
  standalone: true,
  imports: [
    NgForOf,
    DatePipe
  ],
  templateUrl: './profile.component.html',
  styleUrl: './profile.component.scss'
})
export class ProfileComponent {
  protected shortenerService = inject(ShortenerService)

  protected myURLS: URLItem[] = [];

  constructor() {
    this.getUrls()
  }

  protected getUrls(){
    this.shortenerService.getMyURLS().subscribe({
      next: (res: URLListResponse) => {
        this.myURLS = res.Code;
      }
    });
  }

  protected disableURL(id: string){
    this.shortenerService.disableURL(id).subscribe({
      next: (res : any)=> {
        this.getUrls()
      }
    })
  }
}
