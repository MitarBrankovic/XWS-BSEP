import { formatDate } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ModalDismissReasons, NgbModal } from '@ng-bootstrap/ng-bootstrap';
import Swal from 'sweetalert2';
import { User } from '../model/user.model';
import { EditProfileService } from '../services/edit-profile.service';

@Component({
  selector: 'app-edit-profile',
  templateUrl: './edit-profile.component.html',
  styleUrls: ['./edit-profile.component.css']
})
export class EditProfileComponent implements OnInit {
  
  todayDate: string = formatDate(new Date(), 'yyyy-MM-dd', 'en_US');
  user: User = new User;

  isEdit: boolean = false;
  closeResult = '';

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


  constructor(private editProfileService: EditProfileService, private router: Router, private modalService: NgbModal) { }

  ngOnInit(): void {
    let token = localStorage.getItem('token')
    if (token === null) {
      token = ""
      this.router.navigate(['/'])
    }

    let oldUsername = this.parseJwt(JSON.parse(token)?.accessToken)?.username
    
    this.editProfileService.getLoggedUserFromServer(oldUsername).subscribe(f => {
      this.user = f.user;
      this.user.dateOfBirth = formatDate(this.user.dateOfBirth, 'yyyy-MM-dd', 'en_US') + "T00:00:00Z";
    });

  }

  parseJwt(token: string) {
    var base64Url = token.split('.')[1];
    var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    var jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
};

  onSubmit(){}

  removeEducation(e:any): void{
    const index = this.user.education.indexOf(e, 0);
    if (index > -1) {
      this.user.education.splice(index, 1);
    }
  }
  removeWork(w:any): void{
    const index = this.user.workExperience.indexOf(w, 0);
    if (index > -1) {
      this.user.workExperience.splice(index, 1);
    }
  }

  saveUser(): void{
    this.user.skills = this.user.skills.toString().split(','),	
    this.user.interests =this.user.interests.toString().split(',')
    
    this.editProfileService.editProfile(this.user).subscribe(f => this.user = f)
  }

  addEducation(){
    const education= {
      school: this.schoolModal,
      degree: this.degreeModal,
      fieldOfStudy: this.fieldOfStudyModal,
      startDate: this.startDateModal,
      endDate: this.endDateModal
    }
    if(education.school == "" || education.degree == "" || education.fieldOfStudy == "" || education.startDate == "" || education.endDate == ""){
      Swal.fire({
        icon: 'error',
        title: 'Fill all inputs',
        text: 'Something went wrong!',
      })
    }else{
      this.user.education.push(education)
      this.modalService.dismissAll()
    }

  }

  addWorkExperience(){
    const workExperience= {
      title: this.titleModal,
      company: this.companyModal,
      employmentType: this.employmentTypeModal,
      location: this.locationModal,
      startDate: this.startDateWorkModal,
      endDate: this.endDateWorkModal
    }
    if(workExperience.title == "" || workExperience.company == "" || workExperience.employmentType == "" || workExperience.startDate == "" || workExperience.endDate == "" || workExperience.location == ""){
      Swal.fire({
        icon: 'error',
        title: 'Fill all inputs',
        text: 'Something went wrong!',
      })
    }else{
      this.user.workExperience.push(workExperience)
      this.modalService.dismissAll()
    }
  }




  //################## MODALS #########################
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

    //#################################################
}
