import { formatDate } from '@angular/common';
import { Component, Input, OnInit } from '@angular/core';
import { Observable, Subscriber } from 'rxjs';
import { LoggedUser } from '../model/logged-user';
import { Post } from '../model/post';
import { User } from '../model/user.model';
import { PostService } from '../services/post.service';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-user-profile-page',
  templateUrl: './user-profile-page.component.html',
  styleUrls: ['./user-profile-page.component.css']
})
export class UserProfilePageComponent implements OnInit {

  user: User = new User();
  posts: Array<any> = [];
  loggedUser: LoggedUser = new LoggedUser();
  newPost: Post = new Post();

  url: any;
  msg = "";
  postImage: any;
  postImageBase64: any;

  constructor(private userService: UserService, private postService: PostService) { }

  ngOnInit(): void {
    this.loggedUser = this.userService.loggedUser;
   
    let checkUser = localStorage.getItem('currentUser') 
    if(checkUser)
      this.user = JSON.parse(checkUser);
    this.getPosts()
  }

  getPosts(){
    this.postService.getPosts(this.user.username).subscribe(
      data => {
        this.posts = data.userPosts;
      }
    )
  }

  formatDates(date: any){
    date = formatDate(date, 'dd MMMM yyyy hh:mm', 'en_US');
    return date;
  }

  createPost(){
    this.newPost.content.text.match(/#\w+/g)?.forEach(element => {
      this.newPost.content.links.push(element.substring(1))
    })
    this.newPost.content.text = this.newPost.content.text.replace(/ #\S+/g, '');
    this.newPost.content.image = this.url;
    this.newPost.createdAt = new Date();
    this.newPost.user.firstName = this.user.firstName;
    this.newPost.user.lastName = this.user.lastName;
    this.newPost.user.username = this.user.username;
    console.log(this.newPost)
    this.postService.createPost(this.newPost).subscribe()
  }


  checkOwnership(){
    if(this.loggedUser.username === this.user.username){
      return true;
    }
    return false;
  }

	selectFile(event: any) { //Angular 11, for stricter type
		if(!event.target.files[0] || event.target.files[0].length == 0) {
			this.msg = 'You must select an image';
			return;
		}
		
		var mimeType = event.target.files[0].type;
		
		if (mimeType.match(/image\/*/) == null) {
			this.msg = "Only images are supported";
			return;
		}
		
		var reader = new FileReader();
		reader.readAsDataURL(event.target.files[0]);
		reader.onload = (_event) => {
			this.msg = "";
			this.url = reader.result;
		}
	}

}
