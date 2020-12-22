FROM node:15.4 AS builder

WORKDIR /app

COPY ./package.json .

RUN yarn install

COPY . .

RUN yarn generate

FROM nginx:1.19

COPY --from=builder /app/dist /usr/share/nginx/html
