import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Injectable } from '@angular/core';
import { Router } from '@angular/router';
import { Observable } from 'rxjs';
import { LoggedUser } from '../model/logged-user';

@Injectable({
  providedIn: 'root'
})
export class JobOffersService {

  private _url = 'https://localhost:8000';
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

  public getAllJobOffers(): Observable<any>{
    return this.http.get(this._url + '/offer', { headers: this.header });
  }

}
