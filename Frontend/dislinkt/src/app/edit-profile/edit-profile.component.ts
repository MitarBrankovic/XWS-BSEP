import { formatDate } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { ModalDismissReasons, NgbModal } from '@ng-bootstrap/ng-bootstrap';
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
  //password: string = "";
  email: string = "";
  dateOfBirth: string = "";
  skills: any = [];
  interests: any = [];
  todayDate: string = formatDate(new Date(), 'yyyy-MM-dd', 'en_US');
  educations: any = [];
  workExperiences: any = [];
  user: any;

  isEdit: boolean = false;
  closeResult = '';


  constructor(private editProfileService: EditProfileService, private router: Router, private modalService: NgbModal) { }

  ngOnInit(): void {
    let token = localStorage.getItem('token')
    if (token === null) {
      token = ""
      this.router.navigate(['/'])
    }
    let username = this.parseJwt(JSON.parse(token)?.accessToken)?.username

    this.editProfileService.getLoggedUserFromServer(username).subscribe(f => {
      this.user = f.user;
      this.name = this.user.firstName;
      this.lastName = this.user.lastName;
      this.username = this.user.username;
      this.email = this.user.email;
      this.skills = this.user.skills;
      this.interests = this.user.interests;
      this.educations = this.user.education;
      this.workExperiences = this.user.workExperience;
      this.dateOfBirth = formatDate(this.user.dateOfBirth, 'yyyy-MM-dd', 'en_US');
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
      //password
      firstName: this.name,
      lastLame: this.lastName,
      dateOfBirth: this.dateOfBirth,
      email: this.email,
      education: this.educations,
      workExperience: this.workExperiences,
      skills: this.skills,
      interests: this.interests
    }
    this.editProfileService.editProfile(user)
  }




  //################## MODAL EDUCATION ################
  open(content: any) {
    this.modalService.open(content, {ariaLabelledBy: 'modal-basic-title'}).result.then((result) => {
      this.closeResult = `Closed with: ${result}`;
    }, (reason) => {
      this.closeResult = `Dismissed ${this.getDismissReason(reason)}`;
    });
  }

  private getDismissReason(reason: any): string {
    if (reason === ModalDismissReasons.ESC) {
      return 'by pressing ESC';
    } else if (reason === ModalDismissReasons.BACKDROP_CLICK) {
      return 'by clicking on a backdrop';
    } else {
      return `with: ${reason}`;
    }
  }
    //#################################################
}
