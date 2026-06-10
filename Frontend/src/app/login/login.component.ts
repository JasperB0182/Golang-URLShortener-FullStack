import {Component, inject, OnInit} from '@angular/core';
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {Router} from "@angular/router";
import {AuthService} from "../services/auth.service";

@Component({
    selector: 'app-login',
    standalone: true,
    imports: [
        FormsModule,
        ReactiveFormsModule
    ],
    templateUrl: './login.component.html',
    styleUrl: './login.component.scss'
})
export class LoginComponent implements OnInit {

  ngOnInit(): void {
      if (this.loginService.isLoggedIn$) {
        this.router.navigate(['/'])
      }
  }

  protected loginService = inject(AuthService)
  protected Email = ""
  protected Password = ""

  protected ErrorLogin = ""
  protected SuccessLogin = ""
  protected router = inject(Router)

  private pause(ms: number) {
    return new Promise<void>((resolve) => setTimeout(resolve, ms));
  }

  protected login(){
    const loginData = {Email: this.Email, Password: this.Password}
    this.loginService.login(loginData).subscribe({
      next: async () => {
        this.SuccessLogin = "Succesfully logged in!"
        await this.pause(1000)
        this.router.navigate(['/'])
      },
      error: (error) => {
        this.ErrorLogin = error.error.error;
      }
    })
  }
}
