import { formatDate } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import Swal from 'sweetalert2';
import { User } from '../model/user.model';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-register',
  templateUrl: './register.component.html',
  styleUrls: ['./register.component.css']
})
export class RegisterComponent implements OnInit {

  newUser: User = new User;

  //modals
  schoolModal:string = "";
  degreeModal:string = "";
  fieldOfStudyModal:string = "";
  startDateModal:any;
  endDateModal:any;
  titleModal:string = "";
  companyModal:string = "";
  employmentTypeModal:string = "";
  locationModal:string = "";
  startDateWorkModal:any;
  endDateWorkModal:any;
  todayDate: string = formatDate(new Date(), 'yyyy-MM-dd', 'en_US');

  allUsernames:any;

  constructor(private userService: UserService, private modalService: NgbModal, private router:Router) { }

  ngOnInit(): void {
    this.getAllUsernames();
  }

  formatTime(){ 
    this.newUser.dateOfBirth = `${this.newUser.dateOfBirth}T00:00:00.000Z` 
  }

  createAccount() {
    this.formatTime();
    this.newUser.skills = this.newUser.skills.toString().split(','),	
    this.newUser.interests =this.newUser.interests.toString().split(',')
    this.userService.register(this.newUser).subscribe(
    ()=>{
      this.router.navigate(['/']);
      Swal.fire(
        'Good job!',
        'Successfully created account!',
        'success'
      )},
    ()=>{
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: 'Wrong inputs!',
      })
    },
    ()=>{});
  }
  
  getAllUsernames(){
    this.userService.getAllUsernames().subscribe(
      (data:any)=>{
        this.allUsernames = data;
      },
      ()=>{},
      ()=>{}
    )
  }

  removeEducation(e:any): void{
    const index = this.newUser.education.indexOf(e, 0);
    if (index > -1) {
      this.newUser.education.splice(index, 1);
    }
  }
  removeWork(w:any): void{
    const index = this.newUser.workExperience.indexOf(w, 0);
    if (index > -1) {
      this.newUser.workExperience.splice(index, 1);
    }
  }

  addEducation(){
    const education= {
      school: this.schoolModal,
      degree: this.degreeModal,
      fieldOfStudy: this.fieldOfStudyModal,
      startDate: `${this.startDateModal}T00:00:00.000Z`,
      endDate: `${this.endDateModal}T00:00:00.000Z`
    }
    if(education.school == "" || education.degree == "" || education.fieldOfStudy == "" || education.startDate == "T00:00:00.000Z" || education.endDate == "T00:00:00.000Z"){
      Swal.fire({
        icon: 'error',
        title: 'Fill all inputs',
        text: 'Something went wrong!',
      })
    }else{
      this.newUser.education.push(education)
      this.modalService.dismissAll()
    }

  }

  addWorkExperience(){
    const workExperience= {
      title: this.titleModal,
      company: this.companyModal,
      employmentType: this.employmentTypeModal,
      location: this.locationModal,
      startDate: `${this.startDateWorkModal}T00:00:00.000Z`,
      endDate: `${this.endDateWorkModal}T00:00:00.000Z`
    }
    if(workExperience.title == "" || workExperience.company == "" || workExperience.employmentType == "" || workExperience.startDate == "T00:00:00.000Z" || workExperience.endDate == "T00:00:00.000Z" || workExperience.location == ""){
      Swal.fire({
        icon: 'error',
        title: 'Fill all inputs',
        text: 'Something went wrong!',
      })
    }else{
      this.newUser.workExperience.push(workExperience)
      this.modalService.dismissAll()
    }
  }

  open(content: any) {
    this.schoolModal = ""
    this.degreeModal = ""
    this.fieldOfStudyModal = ""
    this.startDateModal = ""
    this.endDateModal = ""
    this.titleModal = ""
    this.companyModal = ""
    this.employmentTypeModal = ""
    this.locationModal = ""
    this.startDateWorkModal = ""
    this.endDateWorkModal = ""
    this.modalService.open(content, {ariaLabelledBy: 'modal-basic-title'}).result.then((result) => {}, (reason) => {});
  }

}
