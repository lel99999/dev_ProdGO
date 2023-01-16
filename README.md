# dev_ProdGO
Production Ready GO - Development Workspace

##### Install and Check Version
- MacOS
  ```
  $brew install go/golang
  $go version
  
  $mkdir -p $HOME/go/{bin,src,pkg}
  ```
  Set up Environment/GO Paths
  ```
  export GOPATH=$HOME/go
  export GOROOT="$(brew --prefix golang)/libexec"
  export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
  ```

- Ubuntu (go1.16.7.linux-amd64.tar.gz = 123MB):
  ```
  $curl -OL https://golang.org/dl/go1.16.7.linux-amd64.tar.gz
  $sudo tar -C /usr/local -xvf go1.16.7.linux-amd64.tar.gz
  $sudo chown -R root:root ./go
  $sudo mv go /usr/local
  ```
  Set up Environment/GO Paths
  ```
  $vi ~/.profile
  
  ADD the following: 
  *** GOROOT is location where Go package is installed on your system while GOPATH is location of working directory ***

  export GOROOT=/usr/local/go
  export GOPATH=$HOME/go_projects
  export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
  ```

##### Run Go
```
$go run main.go
```

##### Turn Go Code into a binary executable
```
$go build main.go
$./main
```
##### In order to run go install successfully, you must pass it the install path of the binary you created with go build. To find the binary’s install path, run the following go list command: 
```
$go list -f '{{.Target}}'
```

##### Add GoLang Linters
- [https://golangci-lint.run/usage/install/#local-installation](https://golangci-lint.run/usage/install/#local-installation) <br/>
  - MacOS 
  ```
  $brew install golangci-lint
  $brew upgrade golangci-lint
  ```

  - Linux
  ```
  $curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.43.0
  $golangci-lint --version
  ```

- To Run default set of linters:
  ```
  $golanci-lint run
  ```

- Go vet is a tool concerned with code correctness
  - Use:
  ```
  $go vetn main.go
  ```
  
##### Add Go Debugger
- Add go-devel
  [https://github.com/go-delve/delve](https://github.com/go-delve/delve) <br/>

##### Hexagonal Architecture (introduced 2015)
- Original Intent:
  `Allow an application to equally be driven by users, programs, automated test or batch scripts, and to be developed and tested in isolation from its eventual run-time devices and databases.` <br/>

- 3 Principles of Hexagonal Architecture:
  - Explicitly separate User-Side, Business Logic, and Server-Side
  - Dependencies are going from User-Side and Server-Side to the Business Logic
  - We isolate the boundaries by using Ports and Adapters
