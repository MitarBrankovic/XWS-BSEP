import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { AgentService } from '../services/agent.service';

@Component({
  selector: 'app-home-page',
  templateUrl: './home-page.component.html',
  styleUrls: ['./home-page.component.css']
})
export class HomePageComponent implements OnInit {

  private user: any

  constructor(private agentService: AgentService, public router: Router) { }

  ngOnInit(): void {
    this.user = this.agentService.loggedUser
  }

  userIsCommon(): boolean{
    return this.user?.role == 'Common'
  }

}
