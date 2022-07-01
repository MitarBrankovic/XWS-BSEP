import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatCardModule } from '@angular/material/card';
import { MatIconModule } from '@angular/material/icon';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { LoginComponent } from './login/login.component';
import { RegisterComponent } from './register/register.component';
import { HttpClientModule } from '@angular/common/http';
import { FormsModule, ReactiveFormsModule  } from '@angular/forms';
import { ChangePasswordComponent } from './change-password/change-password.component';
import { RecoveryEmailComponent } from './recovery-email/recovery-email.component';
import { RecoveryPasswordComponent } from './recovery-password/recovery-password.component';
import { RedirectComponent } from './redirect/redirect.component';
import { EditProfileComponent } from './edit-profile/edit-profile.component';
import { NgbModule } from '@ng-bootstrap/ng-bootstrap';
import { HomePageComponent } from './home-page/home-page.component';
import { NavbarComponent } from './navbar/navbar.component';
import { UserProfilePageComponent } from './user-profile-page/user-profile-page.component';
import { JobOffersComponent } from './job-offers/job-offers.component';
import { MessagesComponent } from './messages/messages.component';
import { DatePipe } from '@angular/common';

import firebase from 'firebase/compat/app';
export const config = {
  apiKey: "AIzaSyDLcNmPiTMMzY-2MbSP_v5-hJUgSUDxxdA",
  authDomain: "xml-bsep.firebaseapp.com",
  databaseURL: "https://xml-bsep-default-rtdb.europe-west1.firebasedatabase.app",
  projectId: "xml-bsep",
  storageBucket: "xml-bsep.appspot.com",
  messagingSenderId: "350946272842",
  appId: "1:350946272842:web:f567b5337dbbee29f6a8b3"
};
firebase.initializeApp(config);

@NgModule({
  declarations: [
    AppComponent,
    LoginComponent,
    RegisterComponent,
    ChangePasswordComponent,
    RecoveryEmailComponent,
    RecoveryPasswordComponent,
    RedirectComponent,
    EditProfileComponent,
    HomePageComponent,
    NavbarComponent,
    UserProfilePageComponent,
    JobOffersComponent,
    MessagesComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    HttpClientModule,
    FormsModule,
    ReactiveFormsModule,
    NgbModule,
    MatSidenavModule,
    MatCardModule,
    MatFormFieldModule,
    MatIconModule,
    BrowserAnimationsModule,
    MatInputModule,
  ],
  providers: [DatePipe],
  bootstrap: [AppComponent]
})
export class AppModule { }
