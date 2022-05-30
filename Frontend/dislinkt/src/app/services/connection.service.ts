import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { Observable } from "rxjs";
import { LoggedUser } from "../model/logged-user";
import { Post } from "../model/post";

@Injectable({
    providedIn: 'root',
})
export class ConnectionService {

    private _url = 'http://localhost:8000';
    header: any;
    loggedUser: LoggedUser = new LoggedUser();

    constructor(private http: HttpClient, private router: Router) {
        let token = localStorage.getItem('token')
        if (token === null) {
            token = ""
        }
        if (token != "")
            this.header = new HttpHeaders().set("Authorization", JSON.parse(token).accessToken);
    }

    public getAllConnections(): Observable<any>{
        return this.http.get(this._url + '/connection', { headers: this.header });
    }

    public acceptRequest(id: number): Observable<any>{
        return this.http.put(this._url + '/connection/' + id, { headers: this.header });
    }

    public declineRequest(id: number): Observable<any>{
        return this.http.delete(this._url + '/connection/' + id, { headers: this.header });
    }

}   