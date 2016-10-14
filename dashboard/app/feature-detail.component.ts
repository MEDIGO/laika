import {Component, OnInit} from '@angular/core';
import {Observable} from 'rxjs/Observable';
import {ActivatedRoute} from '@angular/router';

import {BackendService} from './backend.service';

@Component({
  selector: 'feature-detail',
  templateUrl: './feature-detail.html',
  styleUrls: ['./common.css', './feature-detail.css']
})
export class FeatureDetailComponent implements OnInit {
  feature: any = {};
  error: string;

  constructor(
    private route: ActivatedRoute,
    private backend: BackendService
  ) {}

  ngOnInit() {
    this.route.params.subscribe(params => {
      this.backend.getFeature(params['feature-name']).subscribe(
        feature => this.feature = feature,
        error =>  this.error = error
      );
    });
  }

  handleToggle(status: boolean, name: string) {
    this.feature.status[name] = status;

    this.backend.updateFeature(this.feature.name, {status: this.feature.status}).subscribe(
      feature => this.feature = feature,
      error =>  this.error = error
    );
  }
}
