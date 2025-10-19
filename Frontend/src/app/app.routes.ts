import { Routes } from '@angular/router';
import {HomepageComponent} from "./homepage/homepage.component";
import {RedirectComponent} from "./redirect/redirect.component";
import {LoginComponent} from "./login/login.component";
import {RegisterComponent} from "./register/register.component";
import {ProfileComponent} from "./profile/profile.component";
import {authGuard} from "./guards/auth.guard";

export const routes: Routes = [
  {
    path: "",
    component: HomepageComponent
  },
  {
    path:"rd/:id",
    component: RedirectComponent
  },
  {
    path:"login",
    component: LoginComponent
  },
  {
    path:"register",
    component: RegisterComponent
  },
  {
    path:"profile",
    component: ProfileComponent
    //canActivate: [authGuard]
  }
];
