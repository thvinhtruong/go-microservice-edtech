# Go Microservice Edtech Template

This is a template for building a microservice with Go, Gin, Gorm, and RabbitMQ.
## How to use
1. Clone this repository
2. Run `make init` to install dependencies
3. Run `make run` to start the server

## How to build
1. Run `make build` to build the binary
2. Run `make docker` to build the docker image
3. Run `make docker-run` to run the docker image
4. Run `make docker-push` to push the docker image to Docker Hub
5. Run `make docker-clean` to remove the docker image
6. Run `make clean` to remove the binary

## Our Architecture
### Stack
Building blocks: 
grpc-ecosystem/grpc-gateway/v2, golang-migrate/migrate/v4, ... (tbu)

Utils: 
ilyakaznacheev/cleanenv, samber/lo, sirupsen/logrus, google.golang.org/genproto, google.golang.org/grpc, google.golang.org/protobuf, ... (tbu)

Frontend:
react, redux, redux-saga, ... (tbu)


## How to contribute
1. Fork this repository
2. Create a new branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request
6. Wait for review
7. If approved, merge the PR

## License
[MIT](LICENSE)

## Credits
- Minh Bao Co Tam @Lang0808
- Thanh Vinh Truong @thvinhtruong

