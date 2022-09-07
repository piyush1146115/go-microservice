# go-microservice
An attempt to learn microservices using Golang. 

### Refs:
- [Nic Jackson's Youtube Playlist](https://www.youtube.com/playlist?list=PLmD8u-IFdreyh6EUfevBcbiuCKzFk0EW_)


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
