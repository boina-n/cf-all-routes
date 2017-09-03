# cf-all-routes Plugin
This CF CLI Plugin to shows all routes for each org and space you have permission to access.

#Usage

`cf all-routes`
```
myapp.boina.fr
myapp1.boina.fr
```

##Installation
#####Install from Source (need to have [Go](http://golang.org/dl/) installed)
  ```
  $ go get github.com/cloudfoundry/cli
  $ go get github.com/krujos/cfcurl
  $ go get github.com/boina-n/cf-all-routes
  $ cd $GOPATH/src/github.com/boina-n/cf-all-routes
  $ go build
  $ cf install-plugin cf-all-routes
  ```
