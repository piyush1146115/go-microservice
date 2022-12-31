# Product API

Go based Product API built using the Gorilla Toolkit [https://www.gorillatoolkit.org/](https://www.gorillatoolkit.org/)


## Documentation

OpenAPI documentation can be found in the [swagger.yaml](./swagger.yaml) file
You can see well organized documentation on `http://localhost:9090/docs` after running this product-api.

## Running

The applicaiton can be run with `go run`

```
➜ go run main.go0
products-api 2020/02/16 16:15:11 Starting server on port 9090

curl localhost:9090/products

curl -v localhost:9090/1 -XPUT -d '{"id":1, "name":"tea", "description":"a nice cup of tea"}'`, `curl -v localhost:9090 -X POST -d '{"name": "Water"}'

```
