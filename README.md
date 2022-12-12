# go-microservice
A hands-on attempt to learn microservices using Golang. 

### Refs:
- [Nic Jackson's Youtube Playlist](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_)
- [Nic Jackson's Github Repo](https://github.com/nicholasjackson/building-microservices-youtube)

## Services

### Product API [./product-api](./product-api)
Simple Go based JSON API built using the Gorilla framework. The API allows CRUD based operations on a product list.

### Frontend website [./frontend](./frontend)
ReactJS website for presenting the Product API information

## Series Content

Over the weeks we will look at the following topics, teaching you everything you need to know regarding building microservices with the go programming language:

- Introduction to microservices
- RESTFul microservices
- gRPC microservices
- Packaging applications with Docker
- Testing microservice
- Continuous Delivery
- Observability
- Using Kubernetes
- Debugging
- Security
- Asynchronous microservices
- Caching
- Microservice reliability using a Service Mesh

### Part-1:
- Build a very basic http server using go
- Add different handlefuncs for different paths
- Add some logging and write to the response
- Handle error and return appropriate code for that 
- Run the server using `go run main.go`
- Use curl for query like `curl -v -d 'Request' localhost:9090`, `curl -v localhost:9090/goodbye`, `curl -v localhost:9090`

### Part-2:
- Add handler package
- Build custom handlers
- Use logger wih specific info
- Define custom http server
- Implement graceful shutdown
- Add ReadTimeout, IdleTimeout, WriteTimeout

### Part-3:
- Add product package
- Introduce RESTful services
- Add JSON encoding/ JSON serializing
- Add filtering to HTTP requests
- Add http get request handler
- Add Graceful shutdown

### Part-4:
- Add handlers package
- Add http PUT and POST request handlers
- Tested with commands like: `curl -v localhost:9090/1 -XPUT -d '{"id":1, "name":"tea", "description":"a nice cup of tea"}'`, `curl -v localhost:9090 -X POST -d '{"name": "Water"}'`

### Part-5:
- Add Gorilla MUX router
- Refactor previous code and replace default HTTP router with Gorilla MUX
- Add [Middleware](https://github.com/gorilla/mux#middleware) to PUT and POST methods
- Handle JSON deserializing from Middleware

### Part-6:
- Use [Go Validators](https://github.com/go-playground/validator)
- Add Json validation for product fields in Middleware
- Test validation with simple unit test

### Part-7:
- Use Swagger for documentation
- Use Runtime Middleware
- Add Redoc for visualizing the documentation
- You can now see well organized documentation on http://localhost:9090/docs by running the program
- Refactor  the code

### Part-8:
- Create auto-generated client through Swagger 
- To generate client: `swagger generate client -f ../swagger.yaml -A product-api`

### Part-9:
- Adding frontend
- CORS

### Part-10
- How to upload and serve files using the Go standard library

### Part-11
- Handling multi-part uploads 

### Part-12
- Used Gzip compression for handling http requests
- To test the compression mode: `curl -v localhost:9091/images/1/holding.png --compressed -o file.png`

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
