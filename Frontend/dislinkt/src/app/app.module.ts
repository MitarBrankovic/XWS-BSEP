import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule } from '@angular/forms';
import { UserHomePageComponent } from './user-home-page/user-home-page.component';
import { ChangePasswordComponent } from './change-password/change-password.component';
import { RecoveryEmailComponent } from './recovery-email/recovery-email.component';
import { RecoveryPasswordComponent } from './recovery-password/recovery-password.component';
import { RedirectComponent } from './redirect/redirect.component';

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    RegisterComponent,
    UserHomePageComponent,
    ChangePasswordComponent,
    RecoveryEmailComponent,
    RecoveryPasswordComponent,
    RedirectComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
