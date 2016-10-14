import {NgModule} from '@angular/core';
import {RouterModule} from "@angular/router";
import {BrowserModule} from '@angular/platform-browser';
import {HttpModule} from "@angular/http";
import {FormsModule} from '@angular/forms';
import {NgbModule} from '@ng-bootstrap/ng-bootstrap';
import {MomentModule} from 'angular2-moment';

import {routes} from "./app.routes";
import {AppComponent} from './app.component';
import {FeatureListComponent} from './feature-list.component';
import {FeatureDetailComponent} from './feature-detail.component';
import {EnvironmentListComponent} from './environment-list.component';
import {FeatureCreateComponent} from './feature-create.component';
import {EnvironmentCreateComponent} from './environment-create.component';
import {BackendService} from './backend.service';
import {ItermapPipe} from './itermap.pipe';

@NgModule({
  imports: [
    NgbModule.forRoot(),
    BrowserModule,
    HttpModule,
    RouterModule.forRoot(routes),
    FormsModule,
    MomentModule
  ],
  providers: [BackendService],
  declarations: [
    AppComponent,
    FeatureListComponent,
    FeatureCreateComponent,
    EnvironmentListComponent,
    EnvironmentCreateComponent,
    FeatureDetailComponent,
    ItermapPipe
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
