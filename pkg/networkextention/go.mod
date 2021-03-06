module mosn.io/mosn/pkg/networkextention

go 1.12

require (
	github.com/fsnotify/fsnotify v1.4.10-0.20200417215612-7f4cf4dd2b52
	github.com/hashicorp/go-hclog v0.9.1 // indirect
	github.com/hashicorp/yamux v0.0.0-20181012175058-2f1d1f20f75d // indirect
	github.com/mitchellh/go-testing-interface v1.0.0 // indirect
	mosn.io/api v0.0.0-20210129030635-d7dc8206d7b7
	mosn.io/mosn v0.21.0
	mosn.io/pkg v0.0.0-20201228090327-daaf86502a50
)

replace (
	github.com/apache/dubbo-go-hessian2 => github.com/apache/dubbo-go-hessian2 v1.4.1-0.20200516085443-fa6429e4481d // perf: https://github.com/apache/dubbo-go-hessian2/pull/188
	github.com/golang/protobuf => github.com/golang/protobuf v1.3.5
	google.golang.org/grpc => google.golang.org/grpc v1.28.0
)
