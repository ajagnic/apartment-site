# apartment-site
##### Project Tracking
[ [Trello](https://trello.com/b/qh68JtT0/apartment-website) ]

## Technologies
- [Nuxt.js](https://nuxtjs.org/)
- [Vue.js](https://vuejs.org/) + [Vuetify](https://next.vuetifyjs.com/en/)
- [Go](https://golang.org/)
- [MongoDB](https://www.mongodb.com/)

## Libraries
- [axios](https://axios.nuxtjs.org/)
- [mongo-go-driver](https://github.com/mongodb/mongo-go-driver)

## Usage
### Front-End Only
```bash
# navigate to site sub-directory
$ cd site/

# install dependencies
$ yarn install

# serve with hot reload at localhost:3000
$ yarn dev
```

### Run Whole Project with Docker
```bash
# create env variables from example file
$ cp .env.example .env

# start nuxt dev server, api, mongodb and mongo-express
$ docker-compose -f docker-compose.dev.yml up
```

## Architecture and Data Model
![](arch.png)

# Authors
Adrian Agnic [ [Github](https://github.com/ajagnic) ]
