function basic(username, password) {
  return `Basic ${btoa(`${username}:${password}`)}`;
}

function request(method, endpoint, payload) {
  const opts = {
    headers: {},
    method,
  };

  if (payload) {
    opts.headers['Content-Type'] = 'application/json';
    opts.body = JSON.stringify(payload);
  }

  const username = localStorage.getItem('credentials.username');
  const password = localStorage.getItem('credentials.password');

  if (username && password) {
    opts.headers.Authorization = basic(username, password);
  }

  return fetch(endpoint, opts)
    .then((res) => {
      if (!res.ok) throw new Error(`${res.status}`);
      return res.json();
    });
}

function get(endpoint) {
  return request('GET', endpoint);
}

function post(endpoint, payload) {
  return request('POST', endpoint, payload);
}

function patch(endpoint, payload) {
  return request('PATCH', endpoint, payload);
}

export function listFeatures() {
  return get('/api/features');
}

export function createFeature(feature) {
  return post('/api/features', feature);
}

export function getFeature(name) {
  return get(`/api/features/${name}`);
}

export function listEnvironments() {
  return get('/api/environments');
}

export function createEnvironment(environment) {
  return post('/api/environments', environment);
}

export function getUser(username) {
  return get(`/api/users/${username}`);
}

export function updateFeature(name, feature) {
  return patch(`/api/features/${name}`, feature);
}
