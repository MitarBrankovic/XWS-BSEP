import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AgentService } from '../services/agent.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  username: string = ''
  password: string = ''

  constructor(private agentService: AgentService, private router: Router) { }

  ngOnInit(): void {
  }

  login(): void {
    this.agentService.login(this.username, this.password).subscribe(user =>
    {
      if(user == undefined)
        alert("Invalid username or password!")
      else{
        alert('User logged in successfully. ')
        this.router.navigate(['/homePage'])
      }
    })

  }

}
