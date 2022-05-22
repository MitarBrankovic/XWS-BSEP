import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-recovery-password',
  templateUrl: './recovery-password.component.html',
  styleUrls: ['./recovery-password.component.css']
})
export class RecoveryPasswordComponent implements OnInit {

  newPassword: string = '';
  confirmedPassword: string = '';
  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {
  }

  recoveryPassword(){
    if(this.newPassword === this.confirmedPassword){
      this.userService.recoveryPassword(this.router.url.split('/')[2], this.newPassword).subscribe()
    }
    else{
      alert("Wrong password confirmation");
    }
  }

}
