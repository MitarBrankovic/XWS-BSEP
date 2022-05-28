import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { LoggedUser } from '../model/logged-user';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  token: any;

  constructor(private router:Router, private userService:UserService) { }

  loggedUser: LoggedUser = new LoggedUser(); 

  ngOnInit(): void {
    this.loggedUser = this.userService.loggedUser
  }

  register(){
    this.router.navigate(['/register']);
  }

  login(){
    this.router.navigate(['/login']);
  }

  logout(){
    localStorage.removeItem('token');
    localStorage.removeItem('currentUser');
    window.location.href = '/homePage';
  }

  homePage(){
    this.router.navigate(['/homePage']);
  }

  editProfileRedirect(){
    this.router.navigate(['/editProfile']);
  }

  changePasswordRedirect(){
    this.router.navigate(['/changePassword']);
  }

  isExpired(): boolean{
    return this.loggedUser.exp < Date.now() / 1000
  }

}
