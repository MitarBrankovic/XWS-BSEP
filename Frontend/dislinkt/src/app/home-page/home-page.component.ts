import { formatDate } from '@angular/common';
import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import Swal from 'sweetalert2';
import { LoggedUser } from '../model/logged-user';
import { Post } from '../model/post';
import { User } from '../model/user.model';
import { PostService } from '../services/post.service';
import { UserService } from '../services/user.service';
import { ConnectionService } from '../services/connection.service';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {

  publicUsers: any;
  filteredUsers: any;
  searchValue: string = "";
  posts: Array<any> = [];
  homePagePosts: Array<any> = [];
  loggedUser: LoggedUser = new LoggedUser();
  isClickedOnCommentButton: Array<boolean> = [];
  commentContent: any = "";
  connections: any;
  recommendedUsers: Array<any> = [];

  constructor(private userService: UserService, private postService: PostService, private connectionService: ConnectionService, private router: Router) { }

  ngOnInit(): void {
    this.loggedUser = this.userService.loggedUser;
    this.getAllPublicUsers()
    this.getConnections();
    this.getRecommendedUsers();
  }

  getAllPublicUsers() {
    this.userService.getAllPublicUsers().subscribe(
      f => {
        this.publicUsers = f.users;
        this.filteredUsers = this.publicUsers;
      })

  }

  searchUsers(username: string) {
    this.filteredUsers = this.publicUsers.filter(
      (user: any) => user.username.toLowerCase().includes(username.toLowerCase()));

    if (username === "")
      this.filteredUsers = this.publicUsers;
  }

  redirectToUserProfile(user: User) {
    localStorage.setItem('currentUser', JSON.stringify(user))
    this.router.navigate(['/profile'])
  }

  privateOrPublic(user:any){
    if(user.private)
      return "Private"
    else
      return "Public"
  }

  formatDates(date: any) {
    date = formatDate(date, 'dd MMMM yyyy hh:mm', 'en_US');
    return date;
  }

  reactOnPost(post: any, reactionType: any) {
    let data = {
      reaction: {
        id: '',
        username: this.loggedUser.username,
        type: reactionType,
        createdAt: formatDate(new Date(), 'yyyy-MM-ddThh:mm:ss', 'en_US') + 'Z'

      },
      postId: post.id

    }

    this.postService.reactOnPost(data).subscribe();
    this.createNotification(post.user.username, `${this.loggedUser.username} reacted on your post`, 1)
    window.location.reload();
  }

  getNumLikes(post: Post, type:number){ 
    return post.reactions.filter((reaction:any) => reaction.type == type).length
  }

  alreadyReacted(post: Post){
    return post.reactions.some((reaction:any) => reaction.username == this.loggedUser.username)
  }

  openCommentDiv(i: number) {
    this.isClickedOnCommentButton[i] = !this.isClickedOnCommentButton[i]
  }

  sendComment(post: any, i: number) {
    if(this.commentContent != ""){
      let data = {
        comment: {
          id: "",
          content: this.commentContent,
          username: this.loggedUser.username,
          dateCreated: formatDate(new Date(), 'yyyy-MM-ddThh:mm:ss', 'en_US') + 'Z'
        },
        postId: post.id
      }
      
      this.postService.sendComment(data).subscribe(() => {
        this.commentContent = "";
        window.location.reload();
        this.isClickedOnCommentButton[i] = false
      },
        () => { }
      );
      /*this.commentContent = "";
      window.location.reload();
      this.isClickedOnCommentButton[i] = false*/
      this.createNotification(post.user.username,`${this.loggedUser.username} commented on your post`, 1)
    }else{
      const Toast = Swal.mixin({
        toast: true,
        position: 'top-end',
        showConfirmButton: false,
        timer: 1100,
        timerProgressBar: true,
        didOpen: (toast) => {
          toast.addEventListener('mouseenter', Swal.stopTimer)
          toast.addEventListener('mouseleave', Swal.resumeTimer)
        }
      })
      
      Toast.fire({
        icon: 'error',
        title: 'Fill the comment field!'
      })
    }
  }

  getConnections(){
    this.connectionService.getAllConnections().subscribe(f => {
      this.connections = f.connections;
      this.connections.forEach((connection:any) => {
        if(connection.issuerUser.username == this.loggedUser.username && connection.isApproved == true){
          this.postService.getLatestPosts(connection.subjectUser.username).subscribe((f:any) => {
            f.posts.forEach((post:any) => {
              //post.reactions.some((reaction:any) => reaction.username == this.loggedUser.username)
              this.userService.getBlocked().subscribe((blocked:any) => {
                if(blocked.blocks.some((blockedUser:any) => blockedUser.issuerUsername == post.user.username && blockedUser.subjectUsername == this.loggedUser.username)){
                  
                }else{
                  this.homePagePosts.push(post)
                }
              })
              //this.homePagePosts.push(post);
            })
          })
        }
      },
    )})
  }

  createNotification(username:string, message: string, type: number){
    this.userService.getByUsername(username).subscribe((user:any) => {
      if ((type == 0 && user.user.followNotification) || (type==1 && user.user.postNotification) || (type == 2 && user.user.messageNotification)) 
      this.userService.createNotification(user.user.username, message, type).subscribe();
  })
    
}

getRecommendedUsers(){
  this.userService.getRecommendedUsers().subscribe((data:any) => {
    data.username.forEach((username:any) => {
      this.userService.getByUsername(username).subscribe((u:any) => {
      this.recommendedUsers.push(u.user)
      })
    }
  )})
}

}
