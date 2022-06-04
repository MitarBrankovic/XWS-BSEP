import { Component, OnInit } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import Swal from 'sweetalert2';
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
  content:string = ''
  user: any
  id: number = 0

  commentsChecked: boolean = true
  interviewsChecked: boolean = false
  openPositionsChecked: boolean = false

  comments:any

  constructor(private agentService: AgentService, private route: ActivatedRoute) { }

  ngOnInit(): void {
    this.id = Number(this.route.snapshot.paramMap.get('id'));
    if (this.id)
      this.agentService.findOneCompanyById(this.id).subscribe(company => {
        this.company = company
        this.averageMark = this.calculateAverageMark(company.marks)
      }
      )
    this.findAllCommentsByCompanyId(this.id)
  }

  findAllCommentsByCompanyId(companyId: any) {
    this.agentService.findAllCommentsByCompanyId(companyId).subscribe(comments => {
      this.comments = comments
    })
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
    this.agentService.editCompanyInfo(data).subscribe(() => this.swalSuccess('Company info is edited!'))
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

  checkOwnership(){
    if (this.company.username == this.agentService.loggedUser.username)
      return true
    return false
  }

  createComment(){
    if(this.content != ''){
      let commentDto = {
        content : this.content,
        userSignature : 'Software developer (Medior)',
        companyId : this.company.id,
        userId : this.agentService.loggedUser.id,
        username : this.agentService.loggedUser.username
      }
      this.agentService.saveComment(commentDto).subscribe(() => {
        this.findAllCommentsByCompanyId(this.id);
        this.content = ''
      })

    }else{
      this.swalError('Write comment first!')
    }

  }

  swalError(title: string) {
    const Toast = Swal.mixin({
      toast: true,
      position: 'top-end',
      showConfirmButton: false,
      timer: 1100,
      timerProgressBar: true,
      didOpen: (toast) => {
        toast.addEventListener('mouseenter', Swal.stopTimer)
        toast.addEventListener('mouseleave', Swal.resumeTimer)
      }
    })
    
    Toast.fire({
      icon: 'error',
      title: title
    })
  }

  swalSuccess(title: string) {
    const Toast = Swal.mixin({
      toast: true,
      position: 'top-end',
      showConfirmButton: false,
      timer: 1100,
      timerProgressBar: true,
      didOpen: (toast) => {
        toast.addEventListener('mouseenter', Swal.stopTimer)
        toast.addEventListener('mouseleave', Swal.resumeTimer)
      }
    })
    
    Toast.fire({
      icon: 'success',
      title: title
    })
  }

}
