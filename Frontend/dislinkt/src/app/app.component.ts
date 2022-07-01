import { Component } from '@angular/core';
import { Router } from '@angular/router';
import * as firebase from 'firebase/app';

const config = {
  apiKey: "AIzaSyDLcNmPiTMMzY-2MbSP_v5-hJUgSUDxxdA",
  authDomain: "xml-bsep.firebaseapp.com",
  databaseURL: "https://xml-bsep-default-rtdb.europe-west1.firebasedatabase.app",
  projectId: "xml-bsep",
  storageBucket: "xml-bsep.appspot.com",
  messagingSenderId: "350946272842",
  appId: "1:350946272842:web:f567b5337dbbee29f6a8b3"
};


@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'dislinkt';

  //constructor() {
    //firebase.initializeApp(config);
    //var db = firebase.firestore();
  //}
}
