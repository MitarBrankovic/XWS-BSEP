import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ChangePasswordComponent } from './change-password/change-password.component';
import { LoginComponent } from './login/login.component';
import { RecoveryEmailComponent } from './recovery-email/recovery-email.component';
import { RecoveryPasswordComponent } from './recovery-password/recovery-password.component';
import { RedirectComponent } from './redirect/redirect.component';
import { RegisterComponent } from './register/register.component';
import { EditProfileComponent } from './edit-profile/edit-profile.component';
import { HomePageComponent } from './home-page/home-page.component';
import { JobOffersComponent } from './job-offers/job-offers.component';
import { UserProfilePageComponent } from './user-profile-page/user-profile-page.component';
import { User } from './model/user.model';

const routes: Routes = [
  { path: "", redirectTo: '/homePage', pathMatch: 'full' },
  { path: "login", component: LoginComponent },
  { path: "register", component: RegisterComponent },
  { path: "changePassword", component: ChangePasswordComponent },
  { path: "recovery", component: RecoveryEmailComponent },
  { path: "recover/:token", component: RecoveryPasswordComponent },
  { path: "redirect/:token", component: RedirectComponent },
  { path: "editProfile", component: EditProfileComponent },
  { path: "homePage", component: HomePageComponent },
  { path: "jobOffers", component: JobOffersComponent },
  { path: "profile", component: UserProfilePageComponent, data: { user: User } },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
