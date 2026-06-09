import {Component, inject} from '@angular/core';
import {RouterLink} from "@angular/router";
import {AuthService} from "../services/auth.service";
import { AsyncPipe } from "@angular/common";

@Component({
  selector: 'app-navbar',
  imports: [
    RouterLink,
    AsyncPipe
],
  templateUrl: './navbar.component.html',
  standalone: true,
  styleUrl: './navbar.component.scss'
})
export class NavbarComponent {
  protected authService = inject(AuthService)
}
