import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { User } from '../model/user';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AgentService {

  private _url = 'http://localhost:8081/api/agent/';

  public loggedUser: any

  constructor(private http: HttpClient) { }

  public registerUser(user: User) {
    return this.http.post(this._url + 'saveUser', user);
  }

  public login(username: string, password: string): Observable<any> {
    this.loggedUser = this.http.get<any>(this._url + 'findUser?username=' + username + '&password=' + password);

    return this.loggedUser;
  }

  public findAllCompanies(): Observable<any> {
    return this.http.get<any>(this._url + 'findAllCompanies');
  }

}
