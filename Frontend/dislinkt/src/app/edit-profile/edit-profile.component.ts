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
  
  name: string = "";
  lastName: string = "";
  username: string = "";
  password: string = "";
  email: string = "";
  dateOfBirth: string = "";
  skills: string = "";
  interests: string = "";
  todayDate: string = formatDate(new Date(), 'yyyy-MM-dd', 'en_US');
  educations: any = [];
  workExperiences: any = [];
  user:any;
  user2: User = new User;

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
      this.user2 = f.user;
      console.log(this.user2)
      this.user = f.user;
      this.name = this.user.firstName;
      this.lastName = this.user.lastName;
      this.username = this.user.username;
      this.password = '';
      this.email = this.user.email;
      this.skills = this.user.skills;
      this.interests = this.user.interests;
      this.educations = this.user.education;
      this.workExperiences = this.user.workExperience;
      this.dateOfBirth = formatDate(this.user.dateOfBirth, 'yyyy-MM-dd', 'en_US') + "T00:00:00Z";
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
    const index = this.educations.indexOf(e, 0);
    if (index > -1) {
      this.educations.splice(index, 1);
    }
  }
  removeWork(w:any): void{
    const index = this.workExperiences.indexOf(w, 0);
    if (index > -1) {
      this.workExperiences.splice(index, 1);
    }
  }

  saveUser(): void{
    const user= {
      id: this.user.id,
      username: this.username,
      password: this.password,
      firstName: this.name,
      lastName: this.lastName,
      dateOfBirth: this.dateOfBirth,
      email: this.email,
      education: this.educations,
      workExperience: this.workExperiences,
      skills: this.skills.toString().split(','),	
      interests: this.interests.toString().split(',')
    }
    this.editProfileService.editProfile(user).subscribe(f => this.user = f)
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
      this.educations.push(education)
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
      this.workExperiences.push(workExperience)
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
