import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-redirect',
  templateUrl: './redirect.component.html',
  styleUrls: ['./redirect.component.css']
})
export class RedirectComponent implements OnInit {

  constructor(private userService: UserService, private router: Router) { }

  ngOnInit(): void {
    this.userService.loginPaswrodless(this.router.url.split('/')[2]).subscribe((f) => {
      localStorage.setItem('token', JSON.stringify(f));  
      this.router.navigate(['/userHomePage']) })
  }

}
