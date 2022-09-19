# go-microservice
An attempt to learn microservices using Golang. 

### Refs:
- [Nic Jackson's Youtube Playlist](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_)
- [Nic Jackson's Github Repo](https://github.com/nicholasjackson/building-microservices-youtube)

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
