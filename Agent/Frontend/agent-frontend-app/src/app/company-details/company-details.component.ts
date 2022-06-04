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
  contentInterview:string = ''
  contentPosition:string = ''
  user: any
  id: number = 0

  commentsChecked: boolean = true
  interviewsChecked: boolean = false
  openPositionsChecked: boolean = false

  comments:any
  interviews:any
  positions:any

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
    this.findAllInterviewsByCompanyId(this.id)
    this.findAllPositionsByCompanyId(this.id)
  }

  findAllCommentsByCompanyId(companyId: any) {
    this.agentService.findAllCommentsByCompanyId(companyId).subscribe(comments => {
      this.comments = comments
    })
  }

  findAllInterviewsByCompanyId(companyId: any) {
    this.agentService.findAllInterviewsByCompanyId(companyId).subscribe(interviews => {
      this.interviews = interviews
    })
  }

  findAllPositionsByCompanyId(companyId: any) {
    this.agentService.findAllPositionsByCompanyId(companyId).subscribe(positions => {
      this.positions = positions
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

  createInterview(){
    if(this.contentInterview != ''){
      let interviewDto = {
        userId : this.agentService.loggedUser.id,
        companyId : this.company.id,
        interviewDescription : this.contentInterview,
        userSignature : 'Software developer (Medior)',
        username : this.agentService.loggedUser.username
      }
      this.agentService.saveInterview(interviewDto).subscribe(() => {
        this.findAllInterviewsByCompanyId(this.id);
        this.contentInterview = ''
      })

    }else{
      this.swalError('Write interview description first!')
    }
  }

  createPosition(){
    if(this.contentPosition != ''){
      this.agentService.savePosition(this.company.id, this.contentPosition).subscribe(() => {
        this.findAllPositionsByCompanyId(this.id);
        this.contentPosition = ''
      })
    }
    else{
      this.swalError('Write position name first!')
    }
  }

  calculateAverageSalary(position:any){
    let sum = 0
    for (let mark of position.sallarys)
      sum += mark.sallaryValue
    if (sum == 0)
      return 0
    return (sum / position.sallarys.length).toFixed(2)
  }

  calculateMinimumSalary(position:any){
    if(position.sallarys.length != 0){
      let min = position.sallarys[0].sallaryValue
      for (let mark of position.sallarys)
        if(mark.sallaryValue < min)
          min = mark.sallaryValue
      return min
    }else{return 0}
  }

  calculateMaximumSalary(position:any){
    if(position.sallarys.length != 0){
      let max = position.sallarys[0].sallaryValue
      for (let mark of position.sallarys)
        if(mark.sallaryValue > max)
          max = mark.sallaryValue
      return max
    }else{return 0}
  }

  checkIfAlreadySentSalary(position:any){
    for (let mark of position.sallarys)
      if(mark.userId == this.agentService.loggedUser.id)
        return true
    return false
  }

  

  async addSalary(position:any){
    
    const { value: salaryValue } = await Swal.fire({
      title: 'How much do you earn?',
      icon: 'question',
      input: 'range',
      inputLabel: 'Your salary',
      inputAttributes: {
        min: '0',
        max: '3000',
        step: '5'
      },
      inputValue: 800
    })

    if (salaryValue) {
      let dto = {
        userId: this.agentService.loggedUser.id,
        positionId: position.id,
        sallary: salaryValue
      }

      this.agentService.saveSalary(dto).subscribe(() => {
        this.findAllPositionsByCompanyId(this.id);
      })
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
