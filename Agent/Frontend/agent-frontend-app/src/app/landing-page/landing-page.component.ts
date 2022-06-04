import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AgentService } from '../services/agent.service';

@Component({
  selector: 'app-landing-page',
  templateUrl: './landing-page.component.html',
  styleUrls: ['./landing-page.component.css']
})
export class LandingPageComponent implements OnInit {

  companies:any

  constructor(public router: Router, public agentService: AgentService) { }

  ngOnInit(): void {
    this.getAllCompanies()
  }

  getAllCompanies() {
    this.agentService.getAllCompanies().subscribe(companies => {
      this.companies = companies
    })
  }

  redirectToCompany(companyId: number) {
    if(this.agentService.loggedUser != null) {
      this.router.navigate(['/company', companyId])
    }
  }
}
