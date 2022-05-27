import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { Observable } from "rxjs";
import { LoggedUser } from "../model/logged-user";
import { Post } from "../model/post";
import { User } from "../model/user.model";

@Injectable({
  providedIn: 'root',
})
export class UserService {

  private _url = 'http://localhost:8000';
  header: any;
  loggedUser: LoggedUser = new LoggedUser();

  currentUser: User = new User()
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
    this.loggedUser = this.parseJwt(JSON.parse(token)?.accessToken)
    this.header = new HttpHeaders().set("Authorization", JSON.parse(token).accessToken);
  }
}