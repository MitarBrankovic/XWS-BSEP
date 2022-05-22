import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User } from '../model/user.model';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  user: User = new User;
  passwordless: boolean = false;


  constructor(private userService: UserService,
    private router: Router) { }

  ngOnInit(): void {
  }

  login() {
    this.userService.login(this.user).subscribe((token) => {

      localStorage.setItem('token', JSON.stringify(token))

      localStorage.setItem('username', this.user.username)

      this.router.navigate(['/userHomePage']);
    });
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


}
