import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User } from '../model/user.model';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-change-password',
  templateUrl: './change-password.component.html',
  styleUrls: ['./change-password.component.css']
})
export class ChangePasswordComponent implements OnInit {

  oldPassword: string = ''
  newPassword: string = ''
  username: any

  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {
    this.username = localStorage.getItem('username')
    let token = localStorage.getItem('token')
    if (token === null) {
      token = ""
      this.router.navigate(['/'])
    }
    if(this.parseJwt(JSON.parse(token)?.accessToken)?.role !== "user"){
      this.router.navigate(['/'])
    }
  }

  changePassword() {
    this.userService.changePassword(this.username, this.oldPassword, this.newPassword).subscribe()
  }

  parseJwt(token: string) {
    var base64Url = token.split('.')[1];
    var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    var jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
  }

}
