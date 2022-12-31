# go-microservice
A hands-on attempt to learn microservices using Golang. 

### Acknowledgement:
This project followed Nic Jackson's `Building Microservice with Go` tutorial series.

- [Nic Jackson's Youtube Playlist](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_)
- [Nic Jackson's Github Repo](https://github.com/nicholasjackson/building-microservices-youtube)

## Services
This project consists of the following different services:

### Product API [./product-api](./product-api)
Simple Go based JSON API built using the Gorilla framework. The API allows CRUD based operations on a product list.

### Frontend website [./frontend](./frontend)
ReactJS website for presenting the Product API information

### Product-images [./product-images](./product-images)
Simple Go based file service that can download and upload data and use `gzip` for file compression

### Currency [./currency](./currency)
gRPC and Go based RPC service to get updated rates of different currencies from European central bank. 

## Resources:
- [Rest API Best Practices](https://docs.microsoft.com/en-us/azure/architecture/best-practices/api-design)
- [Go JSON encoding](https://pkg.go.dev/encoding/json)
- [Go regexp package](https://pkg.go.dev/regexp)
- [Testify framework](https://github.com/stretchr/testify)
- [Gorilla framework](https://www.gorillatoolkit.org/)
- [OWASP Top 10](https://www.cloudflare.com/learning/security/threats/owasp-top-10/)
- [Redoc](https://github.com/Redocly/redoc)
- [Runtime Middleware](https://github.com/go-openapi/runtime/tree/master/middleware)
- [Go-swagger](https://github.com/go-swagger/go-swagger)
- [Accept-encoding](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Accept-Encoding)
- [Go http responsewriter](https://pkg.go.dev/net/http#ResponseWriter)
- [protobuf](https://developers.google.com/protocol-buffers/docs/proto3)
- [protobuf-encoding](https://developers.google.com/protocol-buffers/docs/encoding)
- [grpc](https://grpc.io/)
- [grpcurl](https://github.com/fullstorydev/grpcurl)
- [grpc server-side streaming](https://grpc.io/docs/languages/go/basics/#server-side-streaming-rpc)
- [grpc client-side streaming](https://grpc.io/docs/languages/go/basics/#client-side-streaming-rpc)
- [grpc error handling](https://grpc.io/docs/guides/error/)