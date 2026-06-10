import {Component, inject} from '@angular/core';
import { DatePipe } from "@angular/common";
import {ShortenerService} from "../services/shortener-service.service";
import {URLItem} from "../models/URLlist-model";
import {AuthService} from "../services/auth.service";
import {environment} from "../../environments/environment";

@Component({
    selector: 'app-profile',
    standalone: true,
    imports: [
    DatePipe
],
    templateUrl: './profile.component.html',
    styleUrl: './profile.component.scss'
})
export class ProfileComponent {
  protected shortenerService = inject(ShortenerService)
  protected authService = inject(AuthService)

  protected myURLS: URLItem[] = [];
  protected myDisabledURLS: URLItem[] = [];
  protected myAmountOfURLs : any
  protected userinfo : any


  constructor() {
    this.getUrls()
    this.getUserInfo()
  }

  changeName() {
    var newName = prompt("Please enter a new name here! (Min 2 characters, Max 15 characters.)")

    if (newName === null) {
      return;
    }

    if (newName!.length < 2 || newName!.length > 15){
      alert("Name requirements not met.")
      return
    }

    const completeNewName = {
      "Name": newName
    }


    this.authService.changeName(completeNewName).subscribe({
      next: () => {
        alert("Succesfully changed name!")
        this.getUserInfo()
      },
      error: (err : any)=> {
        alert(err.error.error)
      }
    })
  }

  changeEmail() {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    var newEmail = prompt("Please enter a new email here!")


    if (newEmail === null) {
      return;
    }
    newEmail = newEmail.trim()

    if(!emailRegex.test(newEmail)){
      alert("Not a valid email!")
      return
    }

    const completeNewEmail = {
      "Email": newEmail
    }


    this.authService.changeEmail(completeNewEmail).subscribe({
      next: () => {
        alert("Succesfully changed Email!")
        this.getUserInfo()
      },
      error: (err : any)=> {
        alert(err.error.error)
      }
    })
  }

  changePassword() {
    var oldPassword = prompt("Please enter your original password here!")

    if (oldPassword === null) {
      return;
    }

    var newPassword = prompt("Please enter your new password here! (Min 4 characters, Max 20)")

    if (newPassword === null) {
      return;
    }

    if (newPassword!.length < 4 || newPassword!.length > 20){
      alert("Password requirements not met.")
      return
    }

    const completeNewPassword = {
      "OldPassword": oldPassword,
      "NewPassword": newPassword
    }


    this.authService.changePassword(completeNewPassword).subscribe({
      next: () => {
        alert("Succesfully changed password!")
        this.getUserInfo()
      },
      error: (err : any)=> {
        alert(err.error.error)
      }
    })
  }



  protected getUserInfo(){
    this.authService.getuserInfo().subscribe({
      next: (res: any) => {
        this.userinfo = res.user
      }
    })
  }

  protected getUrls(){
    this.shortenerService.getMyURLS().subscribe({
      next: (res: any) => {
        this.myURLS = res.Code;
        this.myDisabledURLS = res.Disabledcode;
        this.myAmountOfURLs = res.AmountOfURLs;
        console.log(this.myDisabledURLS)
      }
    });
  }

  protected disableURL(id: string){
    this.shortenerService.disableURL(id).subscribe({
      next: ()=> {
        this.getUrls()
      },
      error: (error : any) => {
        alert(error.error.error)
      }
    })
  }

  protected enableURL(id: string){
    this.shortenerService.enableURL(id).subscribe({
      next: ()=> {
        this.getUrls()
      },
      error: (err : any) => {
        alert(err.error.error)
      }
    })
  }

  protected readonly environment = environment;
}
