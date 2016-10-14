import {Component, OnInit} from '@angular/core';
import {Observable} from 'rxjs/Observable';

import {BackendService} from './backend.service';

@Component({
  selector: 'feature-list',
  templateUrl: './feature-list.html',
  styleUrls: ['./common.css', './feature-list.css']
})
export class FeatureListComponent implements OnInit {
  features: Observable<any>;
  error: string;

  constructor(private backend: BackendService) {}

  ngOnInit() {
    this.backend.listFeatures().subscribe(
      features => this.features = features,
      error =>  this.error = error);
  }
}
