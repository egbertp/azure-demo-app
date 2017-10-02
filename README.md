### Build and Run locally

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

### Build the Docker container
```
./build.sh
```

### Test the docker container
```
$ docker run -d --rm -p 127.0.0.1:8000:8000 poteb/azure-demo-app:v1.0.0
```

Now point your browser to: `http://localhost:8000/version`


### Upload the docker container
```
$ docker push poteb/azure-demo-app:v1.0.0
```

### Host this app as Azure Web App

[Read these steps](azure-cli-deployment-scripts/README.md)
