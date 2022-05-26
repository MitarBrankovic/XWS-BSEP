import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { Observable } from "rxjs";
import { User } from "../model/user.model";

@Injectable({
  providedIn: 'root',
})
export class UserService {
  private _url = 'http://localhost:8000';
  header: any;

  constructor(private http: HttpClient, private router: Router) {
    let token = localStorage.getItem('token')
    if (token === null) {
      token = ""
    }
    this.header = new HttpHeaders().set("Authorization", JSON.parse(token).accessToken);
   }

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

  public recoveryPassword(token:string, newPassword: string){
    let body = {
      newPassword: newPassword
    }
    return this.http.put(this._url + '/recover/' + token, body);
  }

  public getAllPublicUsers(): Observable<any> {
    return this.http.get(this._url + '/publicUsers');
  }

}