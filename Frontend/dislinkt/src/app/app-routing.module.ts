import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { ChangePasswordComponent } from './change-password/change-password.component';
import { LoginComponent } from './login/login.component';
import { RecoveryEmailComponent } from './recovery-email/recovery-email.component';
import { RecoveryPasswordComponent } from './recovery-password/recovery-password.component';
import { RedirectComponent } from './redirect/redirect.component';
import { RegisterComponent } from './register/register.component';
import { UserHomePageComponent } from './user-home-page/user-home-page.component';

const routes: Routes = [
  { path: "", component: LoginComponent },
  { path: "register", component: RegisterComponent },
  { path: "userHomePage", component: UserHomePageComponent },
  { path: "changePassword", component: ChangePasswordComponent },
  { path: "recovery", component: RecoveryEmailComponent },
  { path: "recover/:token", component: RecoveryPasswordComponent},
  { path: "redirect/:token", component: RedirectComponent}
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
