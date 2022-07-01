import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import Swal from 'sweetalert2';
import { User } from '../model/user.model';
import { UserService } from '../services/user.service';
import firebase from 'firebase/compat/app';
import "firebase/compat/database";

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  user: User = new User;
  passwordless: boolean = false;
  twoFactor: boolean = false;


  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {
  }


  login() {
    this.userService.login(this.user).subscribe((token) => {
      localStorage.setItem('token', JSON.stringify(token))
      localStorage.setItem('username', this.user.username)
      this.userService.updateCredentials();

      firebase.database().ref('users/').orderByChild('nickname').equalTo(this.user.username).once('value', snapshot => {
        if (snapshot.exists()) {
          window.location.href = '/homePage';
        } else {
          const newUser = firebase.database().ref('users/').push();
          newUser.set(this.userService.loggedUser);
          window.location.href = '/homePage';
        }
      });

      //window.location.href = '/homePage';
    },
    ()=>{
      this.userService.loginTwoFactor(this.user).subscribe(() => {
        this.checkTwoFactor();
      },
      ()=>{
        this.swalError('Username/Password incorect')},
      ()=>{})},
    ()=>{}
    );

  }

  loginPasswordlessDemand() {
    this.userService.loginPaswordlessDemand(this.user.username).subscribe()
  }

  register() {
    this.router.navigate(['/register']);
  }

  recoverPassword() {
    this.router.navigate(['/recovery']);
  }

  swalError(text: string){
    Swal.fire({
      icon: 'error',
      title: 'Error',
      text: text,
    })
  }




  async checkTwoFactor(){
    const { value: formValues } = await Swal.fire({
      title: 'Two factor authentication token',
      html:
        '<input placeholder="token" id="swal-input1" [(ngModel)]="input1" class="swal2-input">',
      focusConfirm: false,
      preConfirm: () => {
        return [
          (<HTMLInputElement>document.getElementById("swal-input1")).value,
        ]
      }
    })
    
    if (formValues) {
      this.userService.checkTwoFactor(formValues[0]).subscribe(
        data => {
          Swal.fire({
            title: 'Success',
            text: 'You are logged in',
            icon: 'success'
          });
          localStorage.setItem('token', JSON.stringify(data))
          localStorage.setItem('username', this.user.username)
          this.userService.updateCredentials();
          window.location.href = '/homePage';
        },
          ()=>{
            this.swalError('Invalid credentials')},
          () => {}    
      )
    }
  }

}
