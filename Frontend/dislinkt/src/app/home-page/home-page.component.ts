import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {

  publicUsers:any;

  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {
    this.getAllPublicUsers()
  }

  getAllPublicUsers() {
    this.userService.getAllPublicUsers().subscribe(
      f=> this.publicUsers = f.users)
  }

}
