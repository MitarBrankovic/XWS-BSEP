import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { User } from '../model/user';

@Injectable({
  providedIn: 'root'
})
export class AgentService {

  private _url = 'http://localhost:8081/api/agent/';

  constructor(private http: HttpClient) { }

  public registerUser(user: User) {
    return this.http.post(this._url + 'saveUser', user);
  }
}
