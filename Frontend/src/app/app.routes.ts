import { Routes } from '@angular/router';
import {HomepageComponent} from "./homepage/homepage.component";
import {RedirectComponent} from "./redirect/redirect.component";

export const routes: Routes = [
  {
    path: "",
    component: HomepageComponent
  },
  {
    path:"rd/:id",
    component: RedirectComponent
  }
];
