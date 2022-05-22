import { Component, OnInit } from '@angular/core';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-user-home-page',
  templateUrl: './user-home-page.component.html',
  styleUrls: ['./user-home-page.component.css']
})
export class UserHomePageComponent implements OnInit {

  oldPassword: string = ''
  newPassword: string = ''
  loggedUser: any

  constructor(private userService: UserService) { }

  ngOnInit(): void {
    this.loggedUser = localStorage.getItem('username')
  }

  changePassword() {
    this.userService.changePassword(this.loggedUser, this.oldPassword, this.newPassword).subscribe()
  }

}
