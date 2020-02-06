# serverless-sample

## Development

* Frontend: Use Vue's development server
* Backend:  Go server receives requests via Vue's proxy

```sh
# launch server

$ go generate
$ go build
$ ./serverless-sample

# launch frontend
$ npm install
$ npm run serve
```

## Production Build

* Frontend: JS/HTML/CSS assets are bundled and served via [brbundle](https://github.com/pyspa/brbundle)
* Backend:  Go server receives requests directly

```sh
$ npm insatall
$ npm run build
$ go generate
$ go build
```

## Make Docker image and test

```sh
$ DOCKER_BUILDKIT=1 docker build -t ${IMAGE_NAME} .
$ docker run -p 8080:8080 --rm ${IMAGE_NAME}
```

## Publish to Cloud Run

```sh
$ gcloud builds submit --tag asia.gcr.io/${GCP_PROJECT}/${IMAGE_NAME}
$ gcloud beta run deploy --image asia.gcr.io/${GCP_PROJECT}/${IMAGE_NAME}
```

## License

Apache2
