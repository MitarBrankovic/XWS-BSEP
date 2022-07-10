import { Component, OnInit, ElementRef, ViewChild } from '@angular/core';
import { Router, ActivatedRoute } from '@angular/router';
import { FormControl, FormGroupDirective, FormBuilder, FormGroup, NgForm, Validators } from '@angular/forms';
import firebase from 'firebase/compat/app';
import { DatePipe } from '@angular/common';
import { UserService } from '../services/user.service';
import { ErrorStateMatcher } from '@angular/material/core';
import { LoggedUser } from '../model/logged-user';


/*export class MyErrorStateMatcher implements ErrorStateMatcher {
  isErrorState(control: FormControl | null, form: FormGroupDirective | NgForm | null): boolean {
    const isSubmitted = form && form.submitted;
    return !!(control && control.invalid && (control.dirty || control.touched || isSubmitted));
  }
}*/


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
  selector: 'app-messages',
  templateUrl: './messages.component.html',
  styleUrls: ['./messages.component.css']
})

export class MessagesComponent implements OnInit {

  @ViewChild('chatcontent') chatcontent:ElementRef = new ElementRef({});
  scrolltop: number = 0;

  chatForm: FormGroup = new FormGroup({});
  nickname:any = '';
  roomname:any = '';
  message:any = '';
  users:any = [];
  chats:any = [];


  constructor(private router: Router, private formBuilder: FormBuilder, public datepipe: DatePipe, private userService: UserService) {
      this.nickname = this.userService.loggedUser.username;
      this.roomname = this.router.url.split('/')[2];  //ako je link/21312
      firebase.database().ref('chats/').on('value', (resp: any) => {
        this.chats = [];
        this.chats = snapshotToArray(resp).filter(x => x.roomname === this.roomname);;
        setTimeout(() => this.scrolltop = this.chatcontent.nativeElement.scrollHeight, 500);
      });
      firebase.database().ref('roomusers/').orderByChild('roomname').equalTo(this.roomname).on('value', (resp2: any) => {
        const roomusers = snapshotToArray(resp2);
        this.users = roomusers;
      });
    }

  ngOnInit(): void {
    this.chatForm = this.formBuilder.group({
      'message' : [null, Validators.required]
    });
  }

  onFormSubmit(form: any) {
    if(form.message != null && form.message != ""){
      const chat = form;
      chat.roomname = this.roomname;
      chat.nickname = this.nickname;
      chat.date = this.datepipe.transform(new Date(), 'dd/MM/yyyy HH:mm:ss');
      chat.type = 'message';
      const newMessage = firebase.database().ref('chats/').push();
      newMessage.set(chat);
      this.chatForm = this.formBuilder.group({
        'message' : [null, Validators.required]
      });

      this.createNotification(`${this.nickname}: ${form.message}`, 2);
    }
  }

  createNotification(message: string, type: number){
    
    let tempUser = "";
    if(this.users[0].nickname == this.userService.loggedUser.username)
      tempUser = this.users[1].nickname;
    else
      tempUser = this.users[0].nickname;

    this.userService.getByUsername(tempUser).subscribe((user:any) => {
      if ((type == 0 && user.user.followNotification) || (type==1 && user.user.postNotification) || (type == 2 && user.user.messageNotification)) 
      this.userService.createNotification(user.user.username, message, type).subscribe();
  })
  }
}
