function request(method, endpoint, payload) {
  const opts = {
    headers: {},
    credentials: 'same-origin',
    method
  }

  if (payload) {
    opts.headers['Content-Type'] = 'application/json'
    opts.body = JSON.stringify(payload)
  }

  return fetch(endpoint, opts).then(res => {
    if (!res.ok) {
      return res.json().then(err => {
        throw new Error(err.message)
      })
    }
    return res.json()
  })
}

function get(endpoint) {
  return request('GET', endpoint)
}

function post(endpoint, payload) {
  return request('POST', endpoint, payload)
}

export function listFeatures() {
  return get('/api/features')
}

export function createFeature(feature) {
  return post('/api/events/feature_created', {
    name: feature.name
  })
}

export function getFeature(name) {
  return get(`/api/features/${window.encodeURIComponent(name)}`)
}

export function listEnvironments() {
  return get('/api/environments')
}

export function createEnvironment(environment) {
  return post('/api/events/environment_created', {
    name: environment.name
  })
}

export function toggleFeature(environment, feature, status) {
  return post('/api/events/feature_toggled', {
    environment: environment,
    feature: feature,
    status: status
  })
}

export function deleteFeature(name) {
  return post('/api/events/feature_deleted', { name })
}
