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
    //this.getLoggedUser()
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
    this.router.navigate(['/']);
  }

  getLoggedUser(){
    this.user = this.agentService.loggedUser
    alert()
  }

  userIsCompanyOwner(): boolean{
    return this.user?.role == 'CompanyOwner'
  }

}
