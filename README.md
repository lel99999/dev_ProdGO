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
