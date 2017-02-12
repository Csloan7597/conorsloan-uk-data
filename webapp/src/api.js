const API_URL='http:localhost:8080/api/'

// ABOUT ME

export const fetchAboutMeData = () => {
  return fetch('api/aboutme')
    .then(response => {
      return response.json() // TODO: Error Handling of this all
    })
}

// HOME

export const fetchTagLine = () => {
  return fetch('api/tagline')
    .then(response => {
      return response.text() // TODO: Error Handling of this all
    })
}

export const fetchGlanceItems = () => {
  return fetch('api/glance')
    .then(response => {
      return response.json() // TODO: Error Handling of this all
    })
}
