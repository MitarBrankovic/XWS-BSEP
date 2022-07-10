import { formatDate } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { NgbModal } from '@ng-bootstrap/ng-bootstrap';
import { LoggedUser } from '../model/logged-user';
import { User } from '../model/user.model';
import { ConnectionService } from '../services/connection.service';
import { EditProfileService } from '../services/edit-profile.service';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-navbar',
  templateUrl: './navbar.component.html',
  styleUrls: ['./navbar.component.css']
})
export class NavbarComponent implements OnInit {

  token: any;

  connections: any = [];
  user: User = new User();
  notifications: any = [];

  constructor(private router:Router, private userService:UserService, private editProfileService:EditProfileService, private modalService: NgbModal, private connectionService: ConnectionService) { }

  loggedUser: LoggedUser = new LoggedUser(); 

  ngOnInit(): void {
    this.loggedUser = this.userService.loggedUser
  }

  register(){
    this.router.navigate(['/register']);
  }

  login(){
    this.router.navigate(['/login']);
  }

  logout(){
    localStorage.removeItem('token');
    localStorage.removeItem('currentUser');
    window.location.href = '/homePage';
  }

  homePage(){
    this.router.navigate(['/homePage']);
  }

  profileRedirect(){
    let token = localStorage.getItem('token')
    if (token === null) {
      token = ""
      this.router.navigate(['/'])
    }

    let user

    let oldUsername = this.parseJwt(JSON.parse(token)?.accessToken)?.username
    
    this.editProfileService.getLoggedUserFromServer(oldUsername).subscribe(f => {
      user = f.user;
      localStorage.setItem('currentUser', JSON.stringify(user))
      window.location.href="/profile"
    });
  }

  parseJwt(token: string) {
    var base64Url = token.split('.')[1];
    var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    var jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
  }


  editProfileRedirect(){
    this.router.navigate(['/editProfile']);
  }

  changePasswordRedirect(){
    this.router.navigate(['/changePassword']);
  }

  isExpired(): boolean{
    return this.loggedUser.exp < Date.now() / 1000
  }

  friendRequests(content: any) {
    this.getAllConnections()
    this.modalService.open(content, {ariaLabelledBy: 'modal-basic-title'}).result.then((result) => {}, (reason) => {});
  }

  notification(content: any){
    this.getAllNotifications();
    this.modalService.open(content, {ariaLabelledBy: 'modal-basic-title'}).result.then((result) => {}, (reason) => {});
  }

  getAllNotifications(){
    this.userService.getNotifications(this.loggedUser.username).subscribe((data:any) => this.notifications = data.notification)
  }

  //OPTIMIZOVATI - NA BACKU NAPRAVITI METODU
  getAllConnections(){
    this.connectionService.getAllConnections().subscribe(
      (data) => {
        this.connections = this.getUnapprovedConnectionsByUser(data.connections);
      }
    )
  }

  dismissNotification(notification: any){
    this.userService.removeNotification(notification).subscribe();
    this.notifications = this.notifications.filter((n:any) => n.id != notification.id)
  }

  getUnapprovedConnectionsByUser(connections:any){
    let unapprovedConnections = [];
    for(let connection of connections){
      if(this.loggedUser.username == connection.subjectUser.username && connection.isApproved == false){
        unapprovedConnections.push(connection)
      }
    }
    return unapprovedConnections
  }

  acceptRequest(connection: any){
    this.connectionService.acceptRequest(connection.id).subscribe(() => {
      this.getAllConnections();
    })
  }

  declineRequest(connection: any){
    this.connectionService.deleteConnection(connection.id).subscribe(() => {
      this.getAllConnections();
    })
  }

  jobOffersRedirect(){
    this.router.navigate(['/jobOffers']);
  }

}
