import {Component, inject} from '@angular/core';
import {FormsModule} from "@angular/forms";
import {ShortenerService} from "../../services/shortener-service.service";
import {AuthService} from "../../services/auth.service";
import {AsyncPipe, NgIf} from "@angular/common";
import {QRCodeModule} from "angularx-qrcode";

@Component({
  selector: 'app-shortener',
  standalone: true,
  imports: [
    FormsModule,
    NgIf,
    AsyncPipe,
    QRCodeModule
  ],
  templateUrl: './shortener.component.html',
  styleUrls: ['./shortener.component.scss']
})
export class ShortenerComponent {

  protected inputURL = ""
  protected expiryDate = ""

  protected newURLString = ""
  protected urlCode = ""
  protected Error = ""

  protected credit : any
  protected activeUrls : any
  protected usedCredit = false;

  protected shortenService = inject(ShortenerService)
  protected authService = inject(AuthService)

  protected tomorrow : any;
  protected inamonth: any;

  constructor() {
    const tomorrow = new Date();
    tomorrow.setDate(tomorrow.getDate() + 1);
    this.tomorrow = tomorrow.toISOString().split('T')[0];

    const inamonth = new Date();
    inamonth.setMonth(inamonth.getMonth() + 1);
    this.inamonth = inamonth.toISOString().split('T')[0];

    this.getValues()
  }

  protected onCheckBox(){

  }

  protected getValues(){
    this.authService.getCreditAndUrls().subscribe({
      next: (res) => {
        this.credit = res.Credit
        var active_url = 10 - res.activeUrls
        if (active_url < 0){
          active_url = 0
        }
        this.activeUrls = active_url
      }
    })
  }

  protected shortenURL(){
    this.Error = "";
    if (this.expiryDate){
      var expiry = new Date(this.expiryDate)
      const shortenData = {URL: this.inputURL, ExpiryDate: expiry.toISOString(), usedCredits: this.usedCredit}
      this.shortenService.shorten(shortenData).subscribe({
        next: (res) =>{
          this.newURLString = "New URL: " + "http://localhost:4200/rd/" +  res.Code
          this.urlCode = "http://localhost:4200/rd/" + res.Code
          this.getValues()
        },
        error: (error) => {
          this.Error = error.error.error;
          this.newURLString = ""
          this.urlCode = ""
          console.log(this.Error)
        }
      })
    } else {
      const shortenData = {URL: this.inputURL}
      this.shortenService.shorten(shortenData).subscribe({
        next: (res) =>{
          this.newURLString = "New URL code: " + "http://localhost:4200/rd/" +  res.Code
          this.urlCode = "http://localhost:4200/rd/" + res.Code
          this.getValues()
        },
        error: (error) => {
          this.Error = error.error.error; //Ik weet het.... heel prachtig
          this.newURLString = ""
          this.urlCode = ""
          console.log(this.Error)
        }
      })
    }



  }

}
