import { Component, OnInit } from '@angular/core';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-recovery-email',
  templateUrl: './recovery-email.component.html',
  styleUrls: ['./recovery-email.component.css']
})
export class RecoveryEmailComponent implements OnInit {

  username: string = ''

  constructor(private userService: UserService) { }

  ngOnInit(): void {
  }

  sendEmail() {
    this.userService.sendRecoveryMessage(this.username).subscribe(() => {
      alert('Request for new password is sent on your email address.')
    })
  }

}
