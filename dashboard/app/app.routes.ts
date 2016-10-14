import {Routes} from '@angular/router';
import {FeatureListComponent} from './feature-list.component';
import {EnvironmentListComponent} from './environment-list.component';
import {FeatureCreateComponent} from './feature-create.component';
import {EnvironmentCreateComponent} from './environment-create.component';
import {FeatureDetailComponent} from './feature-detail.component';

export const routes: Routes = [
  {path: '', redirectTo: 'features', pathMatch: 'full'},
  {path: 'features', component: FeatureListComponent},
  {path: 'features/new', component: FeatureCreateComponent},
  {path: 'features/:feature-name', component: FeatureDetailComponent},
  {path: 'environments', component: EnvironmentListComponent},
  {path: 'environments/new', component: EnvironmentCreateComponent}
];
