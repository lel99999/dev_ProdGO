# dev_ProdGO
Production Ready GO - Development Workspace

##### Install and Check Version
- MacOS
  ```
  $brew install go/golang
  $go version
  
  $mkdir -p $HOME/go/{bin,src,pkg}
  ```
  Set up Environment
  ```
  export GOPATH=$HOME/go
  export GOROOT="$(brew --prefix golang)/libexec"
  export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"
  ```

- Ubuntu (go1.16.7.linux-amd64.tar.gz = 123MB):
  ```
  $curl -OL https://golang.org/dl/go1.16.7.linux-amd64.tar.gz
  $tar xvf go1.16.7.linux-amd64.tar.gz
  $sudo chown -R root:root ./go
  $sudo mv go /usr/local
  ```

##### Run Go
```
$go run main.go
```


