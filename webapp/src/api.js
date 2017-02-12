const API_URL='http:localhost:8080/api/'

// Nav

export const fetchProjectNavList = () => {
  return fetch('/api/project/list')
    .then(response => {
      return response.json() // TODO: Error Handling of this all
    })
}

// ABOUT ME

export const fetchAboutMeData = () => {
  return fetch('/api/aboutme')
    .then(response => {
      return response.json() // TODO: Error Handling of this all
    })
}

// HOME

export const fetchTagLine = () => {
  return fetch('/api/tagline')
    .then(response => {
      return response.text() // TODO: Error Handling of this all
    })
}

export const fetchGlanceItems = () => {
  return fetch('/api/glance')
    .then(response => {
      return response.json() // TODO: Error Handling of this all
    })
}

// Projects

export const fetchProjects = () => {
  return fetch('/api/project')
    .then(response => {
      return response.json() // TODO: Error handling
    })
}
