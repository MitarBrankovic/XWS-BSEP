import { getMissingNgModuleMetadataErrorData } from '@angular/compiler';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User } from '../model/user.model';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-user-home-page',
  templateUrl: './user-home-page.component.html',
  styleUrls: ['./user-home-page.component.css']
})
export class UserHomePageComponent implements OnInit {

  user: any
  jwtparser: any;

  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {
    //this.getUserbyUsername() 
    let token = localStorage.getItem('token')
    if (token === null) {
      token = ""
      this.router.navigate(['/'])
    }
    if(this.parseJwt(JSON.parse(token)?.accessToken)?.role !== "user"){
      this.router.navigate(['/'])
    }
  }

  changePasswordRedirect() {
    this.router.navigate(['/changePassword'])
  }

  getUserbyUsername() {
    /*this.userService.getUserByUsername().subscribe(user => {
      this.user = user
    })*/
  }

  parseJwt(token: string) {
    var base64Url = token.split('.')[1];
    var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    var jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
};

}
