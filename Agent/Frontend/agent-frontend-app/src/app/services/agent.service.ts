import { Injectable } from '@angular/core';
import { HttpClient } from "@angular/common/http";
import { User } from '../model/user';
import { Observable } from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class AgentService {
  private _url = 'http://localhost:8081/api/agent/';

  private helper :any = localStorage.getItem('agentUser')
  public loggedUser: any = JSON.parse(this.helper)

  constructor(private http: HttpClient) { }

  public registerUser(user: User) {
    return this.http.post(this._url + 'saveUser', user);
  }

  public login(username: string, password: string): Observable<any> {
    return this.http.get<any>(this._url + 'findUser?username=' + username + '&password=' + password);
  }

  public findAllCompanies(): Observable<any> {
    return this.http.get<any>(this._url + 'findAllCompanies');
  }

  public sendRegistrationRequest(request: any) {
    return this.http.post(this._url + 'saveCompanyRegistrationRequest', request);
  }

  public findAllCompanyRegistrationRequests(): Observable<any> {
    return this.http.get<any>(this._url + 'findAllCompanyRegistrationRequests');
  }

  public registerCompany(companyRegistrationRequest: any) {
    return this.http.post(this._url + 'saveCompany', companyRegistrationRequest);
  }

  public findOneCompanyById(id: number): Observable<any> {
    return this.http.get<any>(this._url + 'findOneCompanyById?companyId=' + id);
  }

  editCompanyInfo(data: any) {
    return this.http.post(this._url + 'editCompanyInfo', data);
  }

  public findAllCommentsByCompanyId(companyId: number): Observable<any> {
    return this.http.get<any>(this._url + 'findAllCommentsByCompanyId/' + companyId);
  }

  public getAllCompanies(){
    return this.http.get<any>(this._url + 'findAllCompanies');
  }

  public saveComment(dto:any){
    return this.http.post(this._url + 'saveComment', dto);
  }

}
