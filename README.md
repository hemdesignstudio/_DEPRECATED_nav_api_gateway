# NAV API GATEWAY V-0.1.0

GraphQL API Gateway for Microsoft Navision Written in GO

![alt Golang](./res/go.png)
![alt Hem Design Studio](./res/hem.png)
![alt Docker](./res/docker.png)
![alt Kubernetes](./res/kubernetes.png)
![alt Google Cloud](./res/googlecloud.png)
![alt Microsoft Navision](./res/nav.png)

## Getting Started
Microsoft NAV web client
```
https://webmyway182.navmoln.se/myway182/
```

Microsoft NAV Odata API

```
https://odatamyway182.navmoln.se:18202/MyWay182Services/ODataV4
```

### Prerequisites


```
go version: 1.11.4

dep version: 0.5.0
```

### Installing and running

```
git clone https://github.com/hemdesignstudio/hem-nav-gateway.git

cd hem-nav-gateway

dep ensure

go run main.go
```

## Deployment

Deployment requires docker ...

## Built With

* [go](https://golang.org/) - Golang Programming Language
* [dep](https://github.com/golang/dep) - Dependency Management Tool
* [graphql-go](https://github.com/graphql-go/graphql) - Graphql Library for Go




## Authors

* **Ahmed Nafies** 
## License

This project is licensed under the Proprietary License - see the [LICENSE.md](LICENSE.md) file for details
