import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import Swal from 'sweetalert2';
import { AgentService } from '../services/agent.service';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  username: string = ''
  password: string = ''

  constructor(private agentService: AgentService, private router: Router) { }

  ngOnInit(): void {
  }

  login(): void {
    this.agentService.login(this.username, this.password).subscribe(data =>
    {
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
        title: 'Signed in successfully'
      })
      this.agentService.loggedUser = data.user
      localStorage.setItem('agentUser', JSON.stringify(data.user))
      localStorage.setItem('jwtToken', JSON.stringify(data.token))
      this.router.navigate(['/homePage'])
    },
    error => {
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
        title: 'Invalid username or password'
      })
    })

  }

}
