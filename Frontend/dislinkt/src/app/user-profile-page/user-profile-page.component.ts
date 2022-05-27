import { formatDate } from '@angular/common';
import { Component, Input, OnInit } from '@angular/core';
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
  constructor(private userService: UserService, private postService: PostService) { }

  ngOnInit(): void {
    this.loggedUser = this.userService.loggedUser;
    this.user = this.userService.currentUser
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
}
