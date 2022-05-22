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

  constructor(private userService: UserService, private Router: Router) { }

  ngOnInit(): void {
    this.getUserbyUsername() 
  }

  changePasswordRedirect() {
    this.Router.navigate(['/changePassword'])
  }

  getUserbyUsername(){
    this.userService.getUserByUsername().subscribe(user => {
      this.user = user
    })
  }

}
