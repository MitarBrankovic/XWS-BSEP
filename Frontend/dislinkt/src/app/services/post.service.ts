import { HttpClient, HttpHeaders } from "@angular/common/http";
import { Injectable } from "@angular/core";
import { Router } from "@angular/router";
import { Observable } from "rxjs";
import { Post } from "../model/post";

@Injectable({
    providedIn: 'root',
})
export class PostService {

    private _url = 'http://localhost:8000';
    header: any;

    constructor(private http: HttpClient, private router: Router) {
        let token = localStorage.getItem('token')
        if (token === null) {
            token = ""
        }
        if (token != "")
            this.header = new HttpHeaders().set("Authorization", JSON.parse(token).accessToken);
    }

    public getPosts(username: string): Observable<any> {
        return this.http.get(this._url + '/post/user/' + username);
    }

    public createPost(post: Post) {
        return this.http.post(this._url + '/post', post, { headers: this.header });
    }

    public reactOnPost(reaction: any) {
        return this.http.post(this._url + '/reaction', reaction, { headers: this.header });
    }
}