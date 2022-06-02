import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { CompanyDetailsComponent } from './company-details/company-details.component';
import { CompanyRegistrationComponent } from './company-registration/company-registration.component';
import { HomePageComponent } from './home-page/home-page.component';
import { LandingPageComponent } from './landing-page/landing-page.component';
import { LoginComponent } from './login/login.component';
import { UserRegistrationComponent } from './user-registration/user-registration.component';

const routes: Routes = [
  { path: "", redirectTo: '/landingPage', pathMatch: 'full' },

  { path: "landingPage", component: LandingPageComponent },
  { path: "userRegistration", component: UserRegistrationComponent },
  { path: "login", component: LoginComponent },
  { path: "homePage", component: HomePageComponent },
  { path: "companyRegistration", component: CompanyRegistrationComponent },
  { path: "company/:id", component: CompanyDetailsComponent },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
