import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { Observable } from "rxjs";
import { User } from "../model/user.model";

@Injectable({
  providedIn: 'root',
})
export class EditProfileService {
  private _url = 'http://localhost:8000';
  
  header: any;

  constructor(private http: HttpClient, private router: Router) { 
    let token = localStorage.getItem('token')
    if (token === null) {
      token = ""
    }
    this.header = new HttpHeaders().set("Authorization", JSON.parse(token).accessToken);
  }

  public getLoggedUserFromServer(username:any): Observable<any> {
    return this.http.get(this._url + '/user/findByUsername/' + username, { headers: this.header });
  }

  public editProfile(user:any): Observable<any> {
    let userId = user.id
    return this.http.put(this._url + '/user/' + userId, user, { headers: this.header });
  }


}