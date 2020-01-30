FROM node:12.9.1-alpine

RUN mkdir -p /usr/src/app

WORKDIR /usr/src/app/

COPY package.json /usr/src/app/
COPY yarn.lock /usr/src/app/

RUN yarn install

COPY . /usr/src/app/
