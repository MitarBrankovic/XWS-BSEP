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

  public findAllInterviewsByCompanyId(companyId: number): Observable<any> {
    return this.http.get<any>(this._url + 'findAllInterviewsByCompanyId/' + companyId);
  }

  public findAllPositionsByCompanyId(companyId: number): Observable<any> {
    return this.http.get<any>(this._url + 'findAllPositionsByCompanyId/' + companyId);
  }

  public getAllCompanies(){
    return this.http.get<any>(this._url + 'findAllCompanies');
  }

  public saveComment(dto:any){
    return this.http.post(this._url + 'saveComment', dto);
  }

  public saveInterview(dto:any){
    return this.http.post(this._url + 'addInterviewProcess', dto);
  }

  public savePosition(companyId:any, positionName:string, description:string, criteria:string){
    return this.http.post(this._url + 'addOpenPosition/' + companyId + '/' + positionName + '/' + description + '/' + criteria, null);
  }

  public saveSalary(dto:any){
    return this.http.post(this._url + 'addSallary', dto);
  }

  public promoteCompany(dto:any, token:string){
    return this.http.post('http://localhost:8000/offer/mono/' + token, dto);
  }

  public generateToken(username:string, password: string): Observable<any> {
    let body = {
      username: username,
      password: password
    }

    return this.http.post('http://localhost:8000/generateApiToken', body);
  }

  public saveToken(userId:number, token:string){
    return this.http.post(this._url + 'saveToken/' + userId + '/' + token, null);
  }

}
