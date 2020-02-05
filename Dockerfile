FROM node:12-buster-slim as nodebuilder
WORKDIR /home/node
COPY package.json package-lock.json ./
RUN npm ci
COPY babel.config.js .browserslistrc .eslintrc.js jest.config.js tsconfig.json vue.config.js ./
COPY public/ ./public
COPY src/ ./src
RUN npm run build

FROM golang:1.13-buster as gobuilder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
COPY --from=nodebuilder /home/node/dist/ ./dist
RUN go generate
RUN go build -ldflags="-w -s" -o /go/bin/app

FROM gcr.io/distroless/base
COPY --from=gobuilder /go/bin/app /
CMD ["/app"]
