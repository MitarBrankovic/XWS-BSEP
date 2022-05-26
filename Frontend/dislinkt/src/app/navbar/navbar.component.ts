import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  token: any;

  constructor(private router:Router) { }

  ngOnInit(): void {
    let token = localStorage.getItem('token')
    if (token === null) {
      token = ""
      this.router.navigate(['/'])
    }
    this.token = token;
  }

  register(){
    this.router.navigate(['/register']);
  }

  login(){
    this.router.navigate(['/login']);
  }

  homePage(){
    this.router.navigate(['/homePage']);
  }

}
