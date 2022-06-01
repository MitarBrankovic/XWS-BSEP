import { formatDate } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User } from '../model/user';
import { AgentService } from '../services/agent.service';

@Component({
  selector: 'app-user-registration',
  templateUrl: './user-registration.component.html',
  styleUrls: ['./user-registration.component.css']
})
export class UserRegistrationComponent implements OnInit {

  newUser: User = new User;

  todayDate: string = formatDate(new Date(), 'yyyy-MM-dd', 'en_US');

  constructor(private agentService: AgentService,  private router: Router) { }

  ngOnInit(): void {
  }

  registerUser() {

    this.newUser.dateOfBirth = `${this.newUser.dateOfBirth}T00:00:00.000Z`

    this.agentService.registerUser(this.newUser).subscribe(() => {
      alert('User registered successfully');
      this.router.navigate(['/']);

      this.newUser = new User;
    });
  }

}
