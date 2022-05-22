import { Component, OnInit } from '@angular/core';
import { User } from '../model/user.model';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

  newUser: User = new User;

  constructor(private userService: UserService) { }

  ngOnInit(): void {
  }

  createAccount() {
    this.userService.register(this.newUser).subscribe();
  }

}
