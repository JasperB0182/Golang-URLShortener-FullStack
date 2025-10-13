import { Component } from '@angular/core';
import {ShortenerComponent} from "./shortener/shortener.component";

@Component({
  selector: 'app-homepage',
  standalone: true,
  imports: [
    ShortenerComponent
  ],
  templateUrl: './homepage.component.html',
  styleUrl: './homepage.component.scss'
})
export class HomepageComponent {

}
