# k8snodeissues
This command is designed to point out potential issues in a kubernetes node.
such as pods stuck terminating, or pending

## Build
```/bin/bash
export GO111MODULE=on
go mod init
go build
```

## Usage
Login to your kubernetes cluster, then run
```/bin/bash
./k8snodeissues
```

It stays running in a loop with a 60 second sleep, control c when ready to quit.
