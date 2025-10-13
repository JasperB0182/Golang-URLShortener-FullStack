import {Component, inject} from '@angular/core';
import {FormsModule} from "@angular/forms";
import {ShortenerService} from "../../services/shortener-service.service";
import {LoginService} from "../../services/login.service";

@Component({
  selector: 'app-shortener',
  standalone: true,
  imports: [
    FormsModule
  ],
  templateUrl: './shortener.component.html',
  styleUrl: './shortener.component.scss'
})
export class ShortenerComponent {

  protected inputURL = ""
  protected expiryDate = ""

  protected newURL = ""
  protected Error = ""

  protected shortenService = inject(ShortenerService)
  protected loginService = inject(LoginService)

  protected Email = ""
  protected Password = ""

  protected ErrorLogin = ""
  protected SuccessLogin = ""

  protected login(){
    const loginData = {Email: this.Email, Password: this.Password}
    this.loginService.login(loginData).subscribe({
      next: (res) => {
        this.SuccessLogin = "Succesfully logged in!"
      },
      error: (error) => {
        this.ErrorLogin = error.message;
      }
    })
  }

  protected shortenURL(){
    const expiry = new Date(this.expiryDate)
    const shortenData = {URL: this.inputURL, ExpiryDate: expiry.toISOString()}
    this.shortenService.shorten(shortenData).subscribe({
      next: (res) =>{
        this.newURL = "New URL code: " + "http://localhost:4200/rd/" +  res.Code
      },
      error: (error) => {
        this.Error = error.error;
      }
    })

  }

}
