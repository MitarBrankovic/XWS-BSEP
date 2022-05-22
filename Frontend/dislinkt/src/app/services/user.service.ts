import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { User } from "../model/user.model";

@Injectable({
  providedIn: 'root',
})
export class UserService {
  private _url = 'http://localhost:8000';

  constructor(private http: HttpClient, private router: Router) { }

  public login(user: User) {
    return this.http.post(this._url + '/login', user);
  }

  public register(user: User) {
    return this.http.post(this._url + '/register', user);
  }

  public passwordlessLogin(username: string) {
    return this.http.get(this._url + '/loginDemand');
  }

  public changePassword(username: string, oldPassword: string, newPassword: string) {
    let body = {
      oldPassword: oldPassword,
      newPassword: newPassword
    }

    let token = localStorage.getItem('token')
    if (token === null) {
      token = ""
    } 
    let header = new HttpHeaders().set("Authorization", JSON.parse(token).accessToken);
    return this.http.put(this._url + '/changePassword/' + username, body, { headers: header });
  }

  public getUserByUsername() {
    let username = localStorage.getItem('username')
    let token = localStorage.getItem('token')
    if (token === null) {
      token = ""
    } 
    let header = new HttpHeaders().set("Authorization", JSON.parse(token).accessToken);
    return this.http.get(this._url + '/user/' + username, { headers: header });
  }

}