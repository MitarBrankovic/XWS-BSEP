import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AgentService } from '../services/agent.service';

@Component({
  selector: 'app-company-registration',
  templateUrl: './company-registration.component.html',
  styleUrls: ['./company-registration.component.css']
})
export class CompanyRegistrationComponent implements OnInit {

  companyContactInfo: string = ''
  companyDescription : string = ''

  constructor(private agentService: AgentService, private router: Router) { }

  ngOnInit(): void {
  }

  sendRegistrationRequest(){
    let request = {
      companyOwnerUsername: this.agentService.loggedUser.username,
      companyOwnerName: this.agentService.loggedUser.firstName + ' ' + this.agentService.loggedUser.lastName,
      companyContactInfo: this.companyContactInfo,
      companyDescription: this.companyDescription
    }
    this.agentService.sendRegistrationRequest(request).subscribe(
      data => {
        alert('Request is sent')
        this.router.navigate(['/homePage'])
      }
    )
  }

}
