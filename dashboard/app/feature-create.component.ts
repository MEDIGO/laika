import {Component} from '@angular/core';
import {Router} from '@angular/router';

import {BackendService} from './backend.service';

@Component({
  selector: 'feature-create',
  templateUrl: './feature-create.html',
  styleUrls: ['./common.css', './feature-create.css']
})
export class FeatureCreateComponent {
  feature = {}
  error: string

  constructor(
    private router: Router,
    private backend: BackendService) {}

  handleCreate() {
    this.backend.createFeature(this.feature).subscribe(
      feature => this.router.navigate(['/features']),
      error =>  this.error = error);
  }
}
