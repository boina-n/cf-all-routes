# cf-all-routes Plugin
This CF CLI Plugin shows all routes for each org and space you have permission to access on your Cloud Foundry instance.
Results is returned in csv format with values separated with commas.

# Usage
`$ cf all-routes`
```
myapp,cfapps.io,ctne,prod,ctne
www,boina.fr,ctne,prod,ctne
,boina.fr,ctne,prod,ctne
test1,boina.fr,ctne,react,
test2,boina.fr,ctne,react,
```

| hostname | domain name | organization | space | app  |
|----------|-------------|--------------|-------|------|
| myapp    | cfapps.io   | ctne         | prod  | ctne |
| www      | boina.fr    | ctne         | prod  | ctne |
|          | boina.fr    | ctne         | prod  | ctne |
| test1    | boina.fr    | ctne         | react |      |
| test2    | boina.fr    | ctne         | react |      |

## Installation

##### Install from Source (need to have [Go](http://golang.org/dl/) installed)
  ```
  $ go get code.cloudfoundry.org/cli/plugin
  $ go get github.com/cloudfoundry/cli
  $ go get github.com/krujos/cfcurl
  $ go get github.com/boina-n/cf-all-routes
  $ cd $GOPATH/src/github.com/boina-n/cf-all-routes
  $ go build
  $ cf install-plugin cf-all-routes
  ```
##### Install from internet (need to have cf CLI installed)
```
$ cf install-plugin  https://github.com/boina-n/cf-all-routes/releases/download/v2.0.0/linux-amd64-cf-all-routes-release
```
## ToDo
- as I user I also want to be able to see the path of the routes
- Add header on the csv response.

