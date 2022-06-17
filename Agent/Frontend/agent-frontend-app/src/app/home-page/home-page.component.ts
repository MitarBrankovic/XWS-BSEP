import { HttpHeaders } from '@angular/common/http';
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
  companies:any
  companyRegistrationRequests : any = []

  constructor(private agentService: AgentService, public router: Router) { 
  }

  ngOnInit(): void {
    this.getLoggedUser()
    this.findAllCompanyRegistrationRequests()
    this.getAllCompanies()
  }

  userIsCommon(): boolean{
    return this.agentService.loggedUser.role.name == 'ROLE_COMMON'
  }

  userIsAdmin(): boolean{
    return this.agentService.loggedUser.role.name == 'ROLE_ADMIN'
  }

  userIsCompanyOwner(): boolean{
    return this.agentService.loggedUser.role.name == 'ROLE_COMPANY_OWNER'
  }

  findAllCompanyRegistrationRequests(){
    this.agentService.findAllCompanyRegistrationRequests().subscribe(data => this.companyRegistrationRequests = data)
  }

  registerCompany(companyRegistrationRequest: any){
    this.agentService.registerCompany(companyRegistrationRequest).subscribe(() => {
      alert('Company registered successfully')
      this.findAllCompanyRegistrationRequests()
      this.getLoggedUser()
    })
  }

  getLoggedUser(){
    let localStorageAgentUser = localStorage.getItem('agentUser')
    if(localStorageAgentUser != null)
      this.user = JSON.parse(localStorageAgentUser)
  }

  getAllCompanies() {
    this.agentService.getAllCompanies().subscribe(companies => {
      this.companies = companies
    })
  }

  navigateToUserCompany() {
    this.router.navigate(['/company', this.user?.company.id])
  }
}
