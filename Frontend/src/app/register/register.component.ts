import {Component, inject, OnInit} from '@angular/core';
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {Router} from "@angular/router";
import {AuthService} from "../services/auth.service";

@Component({
    selector: 'app-register',
    standalone: true,
    imports: [
        FormsModule,
        ReactiveFormsModule
    ],
    templateUrl: './register.component.html',
    styleUrl: './register.component.scss'
})
export class RegisterComponent{
  protected registerService = inject(AuthService)
  protected Email = ""
  protected Password = ""
  protected Name = ""

  protected ErrorLogin = ""
  protected SuccessLogin = ""
  protected router = inject(Router)

  private pause(ms: number) {
    return new Promise<void>((resolve) => setTimeout(resolve, ms));
  }



  protected register(){
    const register = {Name: this.Name, Email: this.Email, Password: this.Password}
    this.registerService.register(register).subscribe({
      next: async () => {
        this.SuccessLogin = "Succesfully signed up and logged in!"
        await this.pause(1000)
        this.router.navigate(['/'])
      },
      error: (error) => {
        this.ErrorLogin = error.error.error;
      }
    })
  }
}
