import {Injectable} from '@angular/core';
import {Http, Response, Headers, RequestOptions} from '@angular/http';
import {Observable} from 'rxjs/Observable';

@Injectable()
export class BackendService {
  constructor(private http: Http) {}

  listFeatures() {
    return this.request('get', '/api/features');
  }

  createFeature(feature) {
    return this.request('post', '/api/events/feature_created', {
      'name': feature.name
    })
  }

  getFeature(name) {
    return this.request('get', '/api/features/' + (window as any).encodeURIComponent(name));
  }

  toggleFeature(env: string, feature: string, status: boolean) {
    return this.request('post', '/api/events/feature_toggled', {
      'environment': env,
      'feature': feature,
      'status': status,
    })
  }

  listEnvironments() {
    return this.request('get', '/api/environments');
  }

  createEnvironment(env) {
    return this.request('post', '/api/events/environment_created', {
      'name': env.name
    })
  }

  private get(path: string) {
    return this.http.get(path).map((res) => res.json()).catch(this.handleError);
  }

  private request(method: string, path: string, payload?) {
    let body = null;
    let headers = null
    let options = null;

    if (payload) {
      body = JSON.stringify(payload);
      headers = new Headers({ 'Content-Type': 'application/json' });
      options = new RequestOptions({ headers: headers });
    }

    return this.http[method](path, body, options).map((res) => res.json()).catch(this.handleError);
  }

  private handleError(res: any) {
    let error = res.json();
    return Observable.throw(error.message);
  }
}
