import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Company } from '../model/company';
import { AgentService } from '../services/agent.service';

@Component({
  selector: 'app-company-details',
  templateUrl: './company-details.component.html',
  styleUrls: ['./company-details.component.css']
})
export class CompanyDetailsComponent implements OnInit {

  company: Company = new Company

  constructor(private agentService: AgentService, private route: ActivatedRoute) { }

  ngOnInit(): void {
    const id = Number(this.route.snapshot.paramMap.get('id'));
    if(id)
      this.agentService.findOneCompanyById(id).subscribe(company => this.company = company)
  }

  editCompanyInfo(){
    let data = {
      id : this.company.id,
      contactInfo: this.company.contactInfo,
      description: this.company.description,
    }
    this.agentService.editCompanyInfo(data).subscribe(() => alert("Company info edited successfully"))
  }

}
