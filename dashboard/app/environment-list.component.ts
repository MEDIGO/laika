import {Component, OnInit} from '@angular/core';
import {Observable} from 'rxjs/Observable';

import {BackendService} from './backend.service';

@Component({
  selector: 'environment-list',
  templateUrl: './environment-list.html',
  styleUrls: ['./common.css', './environment-list.css']
})
export class EnvironmentListComponent implements OnInit {
  envs: Observable<any>;
  error: string;

  constructor(private backend: BackendService) {}

  ngOnInit() {
    this.backend.listEnvironments().subscribe(
      envs => this.envs = envs,
      error =>  this.error = error);
  }
}
