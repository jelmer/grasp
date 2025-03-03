FROM docker.io/node:alpine AS assetbuilder
WORKDIR /app
COPY package*.json ./
COPY gulpfile.js ./
COPY assets/ ./assets/
RUN npm install && NODE_ENV=production ./node_modules/gulp/bin/gulp.js

FROM docker.io/golang:latest AS binarybuilder
RUN go install github.com/gobuffalo/packr/packr@latest
WORKDIR /go/src/github.com/jelmer/grasp
COPY . /go/src/github.com/jelmer/grasp
COPY --from=assetbuilder /app/assets/build ./assets/build
ARG GOARCH=amd64
ARG GOOS=linux
RUN make ARCH=${GOARCH} OS=${GOOS} docker

FROM docker.io/alpine:latest
EXPOSE 8080
RUN apk add --update --no-cache bash ca-certificates
WORKDIR /app
COPY --from=binarybuilder /go/src/github.com/jelmer/grasp/grasp .
CMD ["./grasp", "server"]
