FROM node:8 as react-build
WORKDIR /app
COPY package.json package.json
COPY yarn.lock yarn.lock
RUN yarn
COPY . ./
