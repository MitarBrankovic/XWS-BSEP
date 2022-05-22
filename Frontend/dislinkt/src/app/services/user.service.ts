import { HttpClient } from "@angular/common/http";
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

}