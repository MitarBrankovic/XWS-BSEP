import { formatDate } from '@angular/common';
import { Component, Input, OnInit } from '@angular/core';
import { Observable, Subscriber } from 'rxjs';
import Swal from 'sweetalert2';
import { LoggedUser } from '../model/logged-user';
import { Post } from '../model/post';
import { User } from '../model/user.model';
import { ConnectionService } from '../services/connection.service';
import { PostService } from '../services/post.service';
import { UserService } from '../services/user.service';

@Component({
  selector: 'app-user-profile-page',
  templateUrl: './user-profile-page.component.html',
  styleUrls: ['./user-profile-page.component.css']
})
export class UserProfilePageComponent implements OnInit {

  connections: any = [];
  isConnected: boolean = false;
  isApproved: boolean = false;
  isClickedOnCommentButton: Array<boolean> = [];
  commentContent: any = "";

  user: User = new User();
  posts: Array<any> = [];
  loggedUser: LoggedUser = new LoggedUser();
  newPost: Post = new Post();

  url: any = "";
  msg = "";
  postImage: any;
  postImageBase64: any;


  constructor(private userService: UserService, private postService: PostService , private connectionService: ConnectionService) { }

  ngOnInit(): void {
    this.loggedUser = this.userService.loggedUser;

    let checkUser = localStorage.getItem('currentUser')
    if (checkUser)
      this.user = JSON.parse(checkUser);
    this.getPosts()
    this.getAllConnections();
  }

  sortPostsByDate(posts:any){
    return posts.sort((a:any, b:any) => {
      return new Date(b.createdAt).getTime() - new Date(a.createdAt).getTime()
    })
  }

  getPosts() {
    this.postService.getPosts(this.user.username).subscribe(
      data => {
        this.posts = this.sortPostsByDate(data.userPosts);
      }
    )
  }



  formatDates(date: any) {
    date = formatDate(date, 'dd MMMM yyyy hh:mm', 'en_US');
    return date;
  }

  createPost() {
    if(this.newPost.content.text != "" || this.url != ""){
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
      window.location.reload();
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
        title: 'Fill status or image!'
      })
    }
  }


  checkOwnership() {
    if (this.loggedUser.username === this.user.username) {
      return true;
    }
    return false;
  }

  selectFile(event: any) { //Angular 11, for stricter type
    if (!event.target.files[0] || event.target.files[0].length == 0) {
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

  reactOnPost(postId: any, reactionType: any) {
    let data = {
      reaction: {
        id: '',
        username: this.loggedUser.username,
        type: reactionType,
        createdAt: formatDate(new Date(), 'yyyy-MM-ddThh:mm:ss', 'en_US') + 'Z'

      },
      postId: postId

    }

    this.postService.reactOnPost(data).subscribe();
    this.getPosts();
    window.location.reload();
  }

  openCommentDiv(i: number) {
    this.isClickedOnCommentButton[i] = !this.isClickedOnCommentButton[i]
  }

  sendComment(postId: any, i: number) {
    if(this.commentContent != ""){
      let data = {
        comment: {
          id: "",
          content: this.commentContent,
          username: this.loggedUser.username,
          dateCreated: formatDate(new Date(), 'yyyy-MM-ddThh:mm:ss', 'en_US') + 'Z'
        },
        postId: postId
      }
      
      this.postService.sendComment(data).subscribe(() => {
        Swal.fire({
          icon: 'success',
          title: 'Success',
          text: 'Comment is sent',
        })
      },
        () => { }
      );
      this.commentContent = "";
      window.location.reload();
      this.isClickedOnCommentButton[i] = false
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

  getNumLikes(post: Post, type:number){ 
     return post.reactions.filter((reaction:any) => reaction.type == type).length
  }

  alreadyReacted(post: Post){
    return post.reactions.some((reaction:any) => reaction.username == this.loggedUser.username)
  }

  requestConnect(){
    if(this.loggedUser.username != ""){
      this.userService.requestConnect(this.user.username).subscribe(
        (data) => {
          this.getAllConnections();
          Swal.fire({
            icon: 'success',
            title: 'Success',
            text: 'Request sent',
          })
        }
      )
    }
  }

  getAllConnections(){
    this.connectionService.getAllConnectionsByUser(this.user.username).subscribe(
      (data) => {
        this.connections = data.connections;
        this.isConnected = this.connections.some((connection:any) => connection.issuerUsername == this.loggedUser.username && connection.subjectUsername == this.user.username);
        this.isApproved = this.connections.some((connection:any) => connection.issuerUsername == this.loggedUser.username && connection.subjectUsername == this.user.username && connection.isApproved == true);
      }
    )
  }

  requestIsAccepted(){
    return this.connections.some((connection:any) => connection.issuerUsername == this.loggedUser.username && connection.subjectUsername == this.user.username && connection.isApproved == true);
  }


  //OVA METODA CE SE KORISTITI ZA DOPISIVANJE
  /*checkIfConnectionIsMutual(){
    let firstConnection = this.connections.some((connection:any) => connection.issuerUsername == this.loggedUser.username && connection.subjectUsername == this.user.username && connection.isApproved == true);
    let secondConnection = this.connections.some((connection:any) => connection.issuerUsername == this.user.username && connection.subjectUsername == this.loggedUser.username && connection.isApproved == true);
    if(firstConnection && secondConnection)
      return true
    else return false
  }*/

}
