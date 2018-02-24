const request = (method, endpoint, payload) => {
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

const get = (endpoint) =>
  request('GET', endpoint)

const post = (endpoint, payload) =>
  request('POST', endpoint, payload)

const listFeatures = () =>
  get('/api/features')

const createFeature = (feature) =>
  post('/api/events/feature_created', {
    name: feature.name
  })

const getFeature = (name) =>
  get(`/api/features/${window.encodeURIComponent(name)}`)

const listEnvironments = () =>
  get('/api/environments')

const createEnvironment = (environment) =>
  post('/api/events/environment_created', {
    name: environment.name
  })

const toggleFeature = (environment, feature, status) =>
  post('/api/events/feature_toggled', {
    environment: environment,
    feature: feature,
    status: status
  })

const deleteFeature = (name) =>
  post('/api/events/feature_deleted', { name })

export { listFeatures, createFeature, getFeature, listEnvironments, createEnvironment, toggleFeature, deleteFeature }
