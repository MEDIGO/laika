import {Component} from '@angular/core';
import {Router} from '@angular/router';

import {BackendService} from './backend.service';

@Component({
  selector: 'environment-create',
  templateUrl: './environment-create.html',
  styleUrls: ['./common.css']
})
export class EnvironmentCreateComponent {
  env = {}
  error: string

  constructor(
    private router: Router,
    private backend: BackendService) {}

  handleCreate() {
    this.backend.createEnvironment(this.env).subscribe(
      env => this.router.navigate(['/environments']),
      error =>  this.error = error
    );
  }
}
