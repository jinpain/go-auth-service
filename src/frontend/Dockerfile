FROM node:current-alpine3.22 AS build

WORKDIR /build

COPY package.json package-lock.json ./

RUN npm install

COPY . .

RUN npm run build

FROM nginx:alpine3.22 AS final

COPY --from=build /build/dist /usr/share/nginx/html