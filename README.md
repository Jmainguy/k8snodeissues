# k8sNodeIssues
This command is designed to point out potential issues in a kubernetes node.
such as pods stuck terminating, or pending

## Build
```/bin/bash
export GO111MODULE=on
go mod init
go get k8s.io/client-go@v12.0.0
go build
```

## Usage
Login to your kubernetes cluster, then run
```/bin/bash
./k8sNodeIssues
```
