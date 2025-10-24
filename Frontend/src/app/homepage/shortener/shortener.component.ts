import {Component, inject} from '@angular/core';
import {FormsModule} from "@angular/forms";
import {ShortenerService} from "../../services/shortener-service.service";
import {AuthService} from "../../services/auth.service";

@Component({
  selector: 'app-shortener',
  standalone: true,
  imports: [
    FormsModule
  ],
  templateUrl: './shortener.component.html',
  styleUrls: ['./shortener.component.scss']
})
export class ShortenerComponent {

  protected inputURL = ""
  protected expiryDate = ""

  protected newURL = ""
  protected Error = ""

  protected shortenService = inject(ShortenerService)
  protected authService = inject(AuthService)

  protected tomorrow : any;

  constructor() {
    const tomorrow = new Date();
    tomorrow.setDate(tomorrow.getDate() + 1);
    this.tomorrow = tomorrow.toISOString().split('T')[0];
  }

  protected shortenURL(){
    this.Error = "";
    if (this.expiryDate){
      var expiry = new Date(this.expiryDate)
      const shortenData = {URL: this.inputURL, ExpiryDate: expiry.toISOString()}
      this.shortenService.shorten(shortenData).subscribe({
        next: (res) =>{
          this.newURL = "New URL: " + "http://localhost:4200/rd/" +  res.Code
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
          this.newURL = "New URL code: " + "http://localhost:4200/rd/" +  res.Code
        },
        error: (error) => {
          this.Error = error.error.error; //Ik weet het.... heel prachtig
          console.log(this.Error)
        }
      })
    }



  }

}
