import { DatePipe, formatDate } from '@angular/common';
import { Component, Input, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Observable, Subscriber } from 'rxjs';
import Swal from 'sweetalert2';
import { LoggedUser } from '../model/logged-user';
import { Post } from '../model/post';
import { User } from '../model/user.model';
import { ConnectionService } from '../services/connection.service';
import { PostService } from '../services/post.service';
import { UserService } from '../services/user.service';
import firebase from 'firebase/compat/app';
import { TOUCH_BUFFER_MS } from '@angular/cdk/a11y/input-modality/input-modality-detector';

export const snapshotToArray = (snapshot: any) => {
  const returnArr: any[] = [];

  snapshot.forEach((childSnapshot: any) => {
      const item = childSnapshot.val();
      item.key = childSnapshot.key;
      returnArr.push(item);
  });

  return returnArr;
};

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
  roomname = ""
  isBlocked: boolean = false;
  blocked: any;
  amIBlocked: boolean = false;

  user: User = new User();
  posts: Array<any> = [];
  loggedUser: LoggedUser = new LoggedUser();
  newPost: Post = new Post();

  url: any = "";
  msg = "";
  postImage: any;
  postImageBase64: any;


  constructor(private userService: UserService, private postService: PostService , private connectionService: ConnectionService, private router: Router, public datepipe: DatePipe) { }

  ngOnInit(): void {
    this.loggedUser = this.userService.loggedUser;

    let checkUser = localStorage.getItem('currentUser')
    if (checkUser)
      this.user = JSON.parse(checkUser);
    this.getPosts()
    this.getAllConnections();
    this.getBlocked();

    firebase.database().ref('roomusers/').on('value', (resp: any) => {
      let roomname = this.user.username + this.loggedUser.username;
      let roomname2 = this.loggedUser.username + this.user.username;
      let roomuser = [];
      roomuser = snapshotToArray(resp);
      const user = roomuser.find(x => x.nickname === this.loggedUser.username && (x.roomname === roomname || x.roomname === roomname2));
      if (user !== undefined) {
        const userRef = firebase.database().ref('roomusers/' + user.key);
        this.roomname = user.roomname;
      }
    })

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
        this.commentContent = "";
        window.location.reload();
        this.isClickedOnCommentButton[i] = false
      },
        () => { }
      );
      /*this.commentContent = "";
      window.location.reload();
      this.isClickedOnCommentButton[i] = false*/
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
      this.userService.requestConnect(this.user).subscribe(
        (data) => {
          this.getAllConnections();
          this.swalUpRight('Request sent')

          //FIREBASE
          
          if(this.checkIfUserIsFollowingMe()){
            this.roomname = this.user.username + this.loggedUser.username;
          }else{
            this.roomname = this.loggedUser.username + this.user.username;
          }

          firebase.database().ref('roomusers/').orderByChild('roomname').equalTo(this.roomname).on('value', (resp: any) => {
            
            let roomuser = [];
            roomuser = snapshotToArray(resp);
            const user = roomuser.find(x => x.nickname === this.loggedUser.username);
            if (user !== undefined) {
              const userRef = firebase.database().ref('roomusers/' + user.key);
              userRef.update({status: 'online'});
            } else {
              const newroomuser = { roomname: '', nickname: '', status: '' };
              newroomuser.roomname = this.roomname;
              newroomuser.nickname = this.loggedUser.username;
              newroomuser.status = 'online';
              const newRoomUser = firebase.database().ref('roomusers/').push();
              newRoomUser.set(newroomuser);
            }
          });
        }
      )
    }
  }

  checkIfUserIsFollowingMe(){
    for(let connection of this.connections){
      if(connection.issuerUser.username == this.user.username && connection.subjectUser.username == this.loggedUser.username){
        return true;
      }
    }
    return false;
  }


  getAllConnections(){
    this.connectionService.getAllConnections().subscribe(
      (data) => {
        this.connections = data.connections;
        this.isConnected = this.connections.some((connection:any) => connection.issuerUser.username == this.loggedUser.username && connection.subjectUser.username == this.user.username);
        this.isApproved = this.connections.some((connection:any) => connection.issuerUser.username == this.loggedUser.username && connection.subjectUser.username == this.user.username && connection.isApproved == true);
      }
    )
  }

  requestIsAccepted(){
    return this.connections.some((connection:any) => connection.issuerUser.username == this.loggedUser.username && connection.subjectUser.username == this.user.username && connection.isApproved == true);
  }

  unFollow(){
    let connection = this.connections.filter((connection:any) => connection.issuerUser.username == this.loggedUser.username && connection.subjectUser.username == this.user.username);
    this.connectionService.deleteConnection(connection[0].id).subscribe(() => {
        this.getAllConnections();
        this.swalUpRight('Successfully unfollowed')
/*
                  //FIREBASE
                  let roomname = 'jikjhjhkl';
                  firebase.database().ref('roomusers/').orderByChild('roomname').equalTo(roomname).on('value', (resp: any) => {
                    let roomuser = [];
                    roomuser = snapshotToArray(resp);
                    const user = roomuser.find(x => x.nickname === this.loggedUser.username);
                    if (user !== undefined) {
                      //const userRef = firebase.database().ref('roomusers/' + user.key);
                      //userRef.remove(user)
                      firebase.database().ref('roomusers').child(user.key).remove()
                    }
                  });*/
      })
  }

  swalUpRight(title:string){
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
      icon: 'success',
      title: title
    })
  }

  messageRouter(){
    window.location.href = '/messages/'+ this.roomname;
  }

  getBlocked(){
    this.userService.getBlocked().subscribe(f => {
      this.blocked = f;
      this.isBlocked = this.blocked.blocks.some((block:any) => block.issuerUsername == this.loggedUser.username && block.subjectUsername == this.user.username);
      this.amIBlocked = this.blocked.blocks.some((block:any) => block.issuerUsername == this.user.username && block.subjectUsername == this.loggedUser.username);
    });
  }

  blockUser(){
    this.userService.blockUser(this.loggedUser.username, this.user.username).subscribe(f => { this.isBlocked = true; });
  }

  unblockUser(){
    this.userService.unblockUser(this.loggedUser.username, this.user.username).subscribe(f => { this.isBlocked = false; });
  }


  //OVA METODA CE SE KORISTITI ZA DOPISIVANJE
  checkIfConnectionIsMutual(){
    let firstConnection = this.connections.some((connection:any) => connection.issuerUser.username == this.loggedUser.username && connection.subjectUser.username == this.user.username && connection.isApproved == true);
    let secondConnection = this.connections.some((connection:any) => connection.issuerUser.username == this.user.username && connection.subjectUser.username == this.loggedUser.username && connection.isApproved == true);
    if(firstConnection && secondConnection)
      return true
    else return false
  }

}
