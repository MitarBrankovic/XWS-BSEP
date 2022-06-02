import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AgentService } from '../services/agent.service';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {

  user: any

  companyRegistrationRequests : any = []

  constructor(private agentService: AgentService, public router: Router) { }

  ngOnInit(): void {
    this.user = this.agentService.loggedUser
    this.findAllCompanyRegistrationRequests()
  }

  userIsCommon(): boolean{
    return this.user?.role == 'Common'
  }

  userIsAdmin(): boolean{
    return this.user?.role == 'Admin'
  }

  findAllCompanyRegistrationRequests(){
    this.agentService.findAllCompanyRegistrationRequests().subscribe(data => this.companyRegistrationRequests = data)
  }

  registerCompany(companyRegistrationRequest: any){
    this.agentService.registerCompany(companyRegistrationRequest).subscribe(() => {
      alert('Company registered successfully')
      this.findAllCompanyRegistrationRequests()
    })
  }

}
