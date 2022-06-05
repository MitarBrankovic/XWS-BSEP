import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Offer } from '../model/offer';
import { JobOffersService } from '../services/job-offers.service';

@Component({
  selector: 'app-job-offers',
  templateUrl: './job-offers.component.html',
  styleUrls: ['./job-offers.component.css']
})
export class JobOffersComponent implements OnInit {

  jobOffers:any;
  //offer: Offer = new Offer();

  constructor( private jobOffersService: JobOffersService, private router: Router) { }

  ngOnInit(): void {
    this.getAllJobOffers()
  }

  getAllJobOffers(){
    this.jobOffersService.getAllJobOffers().subscribe((data) => {
      this.jobOffers = data.offers;
    })
  }
}
