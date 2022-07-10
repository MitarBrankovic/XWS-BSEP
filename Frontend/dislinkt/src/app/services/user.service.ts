import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { Observable, Subject } from "rxjs";
import { Connection } from "../model/connection";
import { LoggedUser } from "../model/logged-user";
import { User } from "../model/user.model";
import firebase from 'firebase/compat/app';
import "firebase/compat/database";

@Injectable({
  providedIn: 'root',
})
export class UserService {

  private _url = 'http://localhost:8000';
  header: any;
  loggedUser: LoggedUser = new LoggedUser();

  constructor(private http: HttpClient, private router: Router) {
    this.updateCredentials()
  }

  parseJwt(token: string) {
    var base64Url = token.split('.')[1];
    var base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
    var jsonPayload = decodeURIComponent(atob(base64).split('').map(function(c) {
        return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
    }).join(''));

    return JSON.parse(jsonPayload);
};

  public login(user: User) {
    return this.http.post(this._url + '/login', user);
  }

  public loginPaswordlessDemand(username: string) {
    return this.http.get(this._url + '/loginDemand/' + username);
  }

  public loginPaswrodless(token: string) {
    return this.http.get(this._url + '/login/' + token);
  }

  public register(user: User) {
    return this.http.post(this._url + '/register', user);
  }

  public loginTwoFactor(user: User) {
    return this.http.post(this._url + '/loginTwoFactor', user);
  }

  public checkTwoFactor(token: string) {
    let body = {
      token: token
    }
    return this.http.post(this._url + '/checkTwoFactor', body);
  }


  public changePassword(username: string, oldPassword: string, newPassword: string) {
    let body = {
      oldPassword: oldPassword,
      newPassword: newPassword
    }

    return this.http.put(this._url + '/changePassword/' + username, body, { headers: this.header });
  }

  public sendRecoveryMessage(email: string) {
    let body = {
      email: email
    }
    return this.http.post(this._url + '/recover', body);
  }

  public recoveryPassword(token: string, newPassword: string) {
    let body = {
      newPassword: newPassword
    }
    return this.http.put(this._url + '/recover/' + token, body);
  }

  public getAllPublicUsers(): Observable<any> {
    return this.http.get(this._url + '/publicUsers');
  }

  public getAllUsernames(): Observable<any> {
    return this.http.get(this._url + '/getAllUsernames');
  }

  public updateCredentials(){
    let token = localStorage.getItem('token')
    if (token === null) {
      token = ""
    }
    if(token != ""){
      this.loggedUser = this.parseJwt(JSON.parse(token)?.accessToken)
      this.header = new HttpHeaders().set("Authorization", JSON.parse(token).accessToken);
    }
  }

  public requestConnect(user: any): Observable<any>{
    let connection = new Connection(this.loggedUser, user);
    
    return this.http.post(this._url + '/connection', connection, { headers: this.header });

  }

  public isExpired(): boolean{
      return this.loggedUser.exp < Date.now() / 1000
  }

  public blockUser(issuerUsername: string, subjectUsername: string){
    let body = {
      issuerUsername: issuerUsername,
      subjectUsername: subjectUsername
    }
    return this.http.post(this._url + '/block', body, { headers: this.header})
  } 

  public unblockUser(issuerUsername: string, subjectUsername: string){
    let body = {
      issuerUsername: issuerUsername,
      subjectUsername: subjectUsername
    }
    return this.http.post(this._url + '/unblock', body, { headers: this.header})
  } 

  public getBlocked(){
    return this.http.get(this._url + '/block', {headers: this.header})
  }

  public createNotification(username: string, message: string, type: number){
    return this.http.post(this._url + '/notification', {username: username, message: message, type: type}, {headers: this.header})
  }
}