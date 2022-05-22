import { Component, OnInit } from '@angular/core';
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

  constructor(private userService: UserService) { }

  ngOnInit(): void {
    this.username = localStorage.getItem('username')
  }

  changePassword() {
    this.userService.changePassword(this.username, this.oldPassword, this.newPassword).subscribe()
  }

}
