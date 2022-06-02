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

  averageMark = 0

  commentsChecked: boolean = true
  interviewsChecked: boolean = false
  openPositionsChecked: boolean = false

  constructor(private agentService: AgentService, private route: ActivatedRoute) { }

  ngOnInit(): void {
    const id = Number(this.route.snapshot.paramMap.get('id'));
    if (id)
      this.agentService.findOneCompanyById(id).subscribe(company => {
        this.company = company
        this.averageMark = this.calculateAverageMark(company.marks)
      }
      )
  }
  calculateAverageMark(marks: any): number {
    let sum = 0
    for (let mark of marks)
      sum += mark.mark
    if (sum == 0)
      return 0
    return sum / marks.length
  }

  editCompanyInfo() {
    let data = {
      id: this.company.id,
      contactInfo: this.company.contactInfo,
      description: this.company.description,
    }
    this.agentService.editCompanyInfo(data).subscribe(() => alert("Company info edited successfully"))
  }

  openCommentsDiv() {
    this.commentsChecked = true
    this.interviewsChecked = false
    this.openPositionsChecked = false
  }

  openInterviewsDiv() {
    this.commentsChecked = false
    this.interviewsChecked = true
    this.openPositionsChecked = false
  }

  openOpenPositonsDiv() {
    this.commentsChecked = false
    this.interviewsChecked = false
    this.openPositionsChecked = true
  }

}
