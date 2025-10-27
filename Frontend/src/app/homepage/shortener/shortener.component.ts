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
  }

  protected shortenURL(){
    this.Error = "";
    if (this.expiryDate){
      var expiry = new Date(this.expiryDate)
      const shortenData = {URL: this.inputURL, ExpiryDate: expiry.toISOString()}
      this.shortenService.shorten(shortenData).subscribe({
        next: (res) =>{
          this.newURLString = "New URL: " + "http://localhost:4200/rd/" +  res.Code
          this.urlCode = "http://localhost:4200/rd/" + res.Code
        },
        error: (error) => {
          this.Error = error.statusText;
          console.log(this.Error)
        }
      })
    } else {
      const shortenData = {URL: this.inputURL}
      this.shortenService.shorten(shortenData).subscribe({
        next: (res) =>{
          this.newURLString = "New URL code: " + "http://localhost:4200/rd/" +  res.Code
          this.urlCode = "http://localhost:4200/rd/" + res.Code
        },
        error: (error) => {
          this.Error = error.error.error; //Ik weet het.... heel prachtig
          console.log(this.Error)
        }
      })
    }



  }

}
