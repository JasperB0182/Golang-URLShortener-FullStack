import { Component } from '@angular/core';
import {ShortenerComponent} from "./shortener/shortener.component";
import {NavbarComponent} from "../navbar/navbar.component";

@Component({
    selector: 'app-homepage',
    imports: [
        ShortenerComponent
    ],
    templateUrl: './homepage.component.html',
    styleUrl: './homepage.component.scss'
})
export class HomepageComponent {

}
