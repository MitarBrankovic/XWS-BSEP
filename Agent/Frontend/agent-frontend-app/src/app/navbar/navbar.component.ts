import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import Swal from 'sweetalert2';
import { AgentService } from '../services/agent.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  user: any;

  companyRegistrationRequests : any = []

  constructor(public router: Router, private agentService: AgentService) { }

  ngOnInit(): void {
    this.user = this.agentService.loggedUser
  }

  checkIfLoggedIn() {
    if(this.agentService.loggedUser == null){
      return false;
    }else{
      return true;
    }
  }

  logout() {
    this.agentService.loggedUser = null;
    localStorage.removeItem('agentUser');
    this.router.navigate(['/']);
  }


  userIsCompanyOwner(): boolean{
    let a = this.agentService.loggedUser.username
    return this.agentService.loggedUser.role == 'CompanyOwner'
  }

  userIsCommon(): boolean{
    let a = this.agentService.loggedUser.username
    return this.agentService.loggedUser.role == 'Common'
  }

  redirectToCompany() {
    if(this.agentService.loggedUser != null) {
      this.router.navigate(['/company', this.agentService.loggedUser.company.id])
    }
  }

  redirectToCompanyRegistration(){
    if(this.agentService.loggedUser != null) {
      this.router.navigate(['/companyRegistration'])
    }
  }

  async generateToken(){
    const { value: formValues } = await Swal.fire({
      title: 'Multiple inputs',
      html:
        '<input placeholder="username" id="swal-input1" [(ngModel)]="input1" class="swal2-input">' +
        '<input type="password" placeholder="password" id="swal-input2" [(ngModel)]="input2" class="swal2-input">',
      focusConfirm: false,
      preConfirm: () => {
        return [
          (<HTMLInputElement>document.getElementById("swal-input1")).value,
          (<HTMLInputElement>document.getElementById("swal-input2")).value,
        ]
      }
    })
    
    if (formValues) {
      this.agentService.generateToken(formValues[0], formValues[1]).subscribe(
        data => {
          Swal.fire({
            title: 'Token generated',
            text: data.token,
            icon: 'success'
          });
        this.agentService.saveToken(this.agentService.loggedUser.id, data.token).subscribe(
          (user) => {        this.agentService.loggedUser = user;
            localStorage.setItem('agentUser', JSON.stringify(user));}
        )
        },
          ()=>{
            Swal.fire({
              title: 'Error',
              text: 'Invalid credentials',
              icon: 'error'
            })},
          () => {}    
      )
    }
  }

  hasToken(){
    let user = localStorage.getItem('agentUser')
    if(user != null){
      return (JSON.parse(user).apiToken !== '' && JSON.parse(user).apiToken !== null)
    }
    return false
  }
  
}
