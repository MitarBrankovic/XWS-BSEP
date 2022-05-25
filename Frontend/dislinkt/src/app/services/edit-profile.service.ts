import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { User } from "../model/user.model";

@Injectable({
  providedIn: 'root',
})
export class EditProfileService {
  private _url = 'http://localhost:8000';

  constructor(private http: HttpClient, private router: Router) { }

  public getLoggedUserFromServer(username:any): any {
    let token = localStorage.getItem('token')
    if (token === null) {
      token = ""
    }
    let header = new HttpHeaders().set("Authorization", JSON.parse(token).accessToken);
    return this.http.get(this._url + '/user/' + username, { headers: header });
  }


}