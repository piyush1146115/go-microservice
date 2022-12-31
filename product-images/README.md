# Product Images
A http file-server to serve file and upload files using Go standard library and gzip.

Note: need to use `--data-binary` to ensure file is not converted to text

```
curl -vv localhost:9090/1/go.mod -X PUT --data-binary @test.png
```
- To test the normal mode: `curl -v localhost:9091/images/1/holding.png -o file.png`
- To test the compression mode: `curl -v localhost:9091/images/1/holding.png --compressed -o file.png`