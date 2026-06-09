import {Component, inject} from '@angular/core';
import {FormsModule} from "@angular/forms";
import {NgForOf} from "@angular/common";
import {AuthService} from "../services/auth.service";

@Component({
    selector: 'app-credit',
    standalone: true,
    imports: [
        FormsModule,
        NgForOf
    ],
    templateUrl: './credit.component.html',
    styleUrl: './credit.component.scss'
})
export class CreditComponent {

  credits = [10, 50, 100, 500, 2000];
  selectedCredits = 10;

  protected authService = inject(AuthService)

  protected addCredit(){
    this.authService.addCredit(this.selectedCredits).subscribe({
      next: (res : any) => {
        alert("Woohoo added credits!")
      },
      error: (err : any) => {
        alert(":(")
      }
    })
  }

}
