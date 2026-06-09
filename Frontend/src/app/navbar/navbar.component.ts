import {Component, inject} from '@angular/core';
import {RouterLink} from "@angular/router";
import {AuthService} from "../services/auth.service";
import {AsyncPipe, NgIf} from "@angular/common";

@Component({
    selector: 'app-navbar',
    imports: [
        RouterLink,
        NgIf,
        AsyncPipe
    ],
    templateUrl: './navbar.component.html',
    styleUrl: './navbar.component.scss'
})
export class NavbarComponent {
  protected authService = inject(AuthService)
}
