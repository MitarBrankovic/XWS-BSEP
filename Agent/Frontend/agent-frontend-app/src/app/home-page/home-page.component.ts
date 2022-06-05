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

  constructor(private agentService: AgentService, public router: Router) { }

  ngOnInit(): void {
    this.getLoggedUser()
    this.findAllCompanyRegistrationRequests()
    this.getAllCompanies()
  }

  userIsCommon(): boolean{
    return this.agentService.loggedUser.role == 'Common'
  }

  userIsAdmin(): boolean{
    return this.agentService.loggedUser.role == 'Admin'
  }

  userIsCompanyOwner(): boolean{
    return this.agentService.loggedUser.role == 'CompanyOwner'
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
    //this.user = this.agentService.loggedUser
    this.user = localStorage.getItem('agentUser')
  }

  getAllCompanies() {
    this.agentService.getAllCompanies().subscribe(companies => {
      this.companies = companies
    })
  }
}
