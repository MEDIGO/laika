function request(method, endpoint, payload) {
  const opts = {
    headers: {},
    credentials: 'same-origin',
    method,
  };

  if (payload) {
    opts.headers['Content-Type'] = 'application/json';
    opts.body = JSON.stringify(payload);
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
