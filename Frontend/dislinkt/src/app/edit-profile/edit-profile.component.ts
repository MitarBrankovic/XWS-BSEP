import { formatDate } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
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
  todayDate: string = formatDate(new Date(), 'yyyy-MM-dd', 'en_US');

  user: any;


  constructor(private editProfileService: EditProfileService, private router: Router) { }

  ngOnInit(): void {
    let token = localStorage.getItem('token')
    if (token === null) {
      token = ""
      this.router.navigate(['/'])
    }
    let username = this.parseJwt(JSON.parse(token)?.accessToken)?.username

    this.editProfileService.getLoggedUserFromServer(username).subscribe((f: any)=> {
      alert(f.username)
      this.user = f;
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

}
