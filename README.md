# apartment-site

## Technologies
- [Nuxt.js](https://nuxtjs.org/)
- [Vue.js](https://vuejs.org/) + [Vuetify](https://next.vuetifyjs.com/en/)
- [Go](https://golang.org/)
- [MongoDB](https://www.mongodb.com/)

## Libraries
- [axios](https://axios.nuxtjs.org/)
- [mongo-go-driver](https://github.com/mongodb/mongo-go-driver)

## Usage
### Build Site
```bash
# navigate to site sub-directory
$ cd site/

# install dependencies
$ yarn install

# serve with hot reload at localhost:3000
$ yarn dev
```

### Docker Compose
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
