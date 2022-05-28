import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User } from '../model/user.model';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {

  publicUsers: any;
  filteredUsers: any;
  searchValue: string = "";

  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {
    this.getAllPublicUsers()
  }

  getAllPublicUsers() {
    this.userService.getAllPublicUsers().subscribe(
      f => {
        this.publicUsers = f.users;
        this.filteredUsers = this.publicUsers;
      })

  }

  searchUsers(username: string) {
    this.filteredUsers = this.publicUsers.filter(
      (user: any) => user.username.toLowerCase() === username.toLowerCase());

    if (username === "")
      this.filteredUsers = this.publicUsers;
  }

  redirectToUserProfile(user: User) {
    localStorage.setItem('currentUser', JSON.stringify(user))
    this.router.navigate(['/profile'])
  }

}
