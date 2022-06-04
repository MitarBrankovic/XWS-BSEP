import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AgentService } from '../services/agent.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  user: any;

  companyRegistrationRequests : any = []

  constructor(public router: Router, private agentService: AgentService) { }

  ngOnInit(): void {
    this.user = this.agentService.loggedUser
  }

  checkIfLoggedIn() {
    if(this.agentService.loggedUser == null){
      return false;
    }else{
      return true;
    }
  }

  logout() {
    this.agentService.loggedUser = null;
    localStorage.removeItem('agentUser');
    this.router.navigate(['/']);
  }


  userIsCompanyOwner(): boolean{
    let a = this.agentService.loggedUser.username
    return this.agentService.loggedUser.role == 'CompanyOwner'
  }

}
