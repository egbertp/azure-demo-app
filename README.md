### Build and Run

Install dependencies using [Glide Package Management for Go](https://glide.sh/)

```
$ glide install
```

Build the binary
```
$ make build
```

Run the app
```
$ ./azure-demo-app
```

Release the app
``` 
$ make release
```


### Use tha app

Point your browser to `http://localhost:7000/version`

### Upload the docker container

docker push poteb/azure-demo-app