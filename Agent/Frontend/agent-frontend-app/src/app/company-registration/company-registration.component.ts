import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import Swal from 'sweetalert2';
import { AgentService } from '../services/agent.service';

@Component({
  selector: 'app-company-registration',
  templateUrl: './company-registration.component.html',
  styleUrls: ['./company-registration.component.css']
})
export class CompanyRegistrationComponent implements OnInit {

  companyName: string = ''
  companyContactInfo: string = ''
  companyDescription : string = ''

  constructor(private agentService: AgentService, private router: Router) { }

  ngOnInit(): void {
  }

  sendRegistrationRequest(){
    let request = {
      companyOwnerUsername: this.agentService.loggedUser.username,
      companyOwnerName: this.agentService.loggedUser.firstName + ' ' + this.agentService.loggedUser.lastName,
      companyName: this.companyName,
      companyContactInfo: this.companyContactInfo,
      companyDescription: this.companyDescription,
      username:  this.agentService.loggedUser.username
    }
    this.agentService.sendRegistrationRequest(request).subscribe(
      data => {
        const Toast = Swal.mixin({
          toast: true,
          position: 'top-end',
          showConfirmButton: false,
          timer: 1500,
          timerProgressBar: true,
          didOpen: (toast) => {
            toast.addEventListener('mouseenter', Swal.stopTimer)
            toast.addEventListener('mouseleave', Swal.resumeTimer)
          }
        })
        Toast.fire({
          icon: 'success',
          title: 'Company registration request sent successfully'
        })
        this.router.navigate(['/homePage'])
      },
      error => {
        const Toast = Swal.mixin({
          toast: true,
          position: 'top-end',
          showConfirmButton: false,
          timer: 1500,
          timerProgressBar: true,
          didOpen: (toast) => {
            toast.addEventListener('mouseenter', Swal.stopTimer)
            toast.addEventListener('mouseleave', Swal.resumeTimer)
          }
        })
        
        Toast.fire({
          icon: 'error',
          title: 'You need to insert valid inputs'
        })
      }
    )
  }

}
