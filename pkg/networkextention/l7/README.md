# TODO

* request header 需要封装为 stream 做内存复用，并使用 stream id 等相关做请求和响应的关联
* 封装 stream 的运行框架
* 如上解决后，需要按照其架构目录进行拆分 package，如：cgo、shim、filter、stream manager、dynamic cluster

# build

* 编译 cgo 适配层

```
go build -mod=vendor -o  golang_extention.so -buildmode=c-shared golang_extention.go mosn.go shim.go dynamic_cluster.go
```

* 编译 mosn on envoy

```
# enable wasm and mosn on envoy
bazel build -c opt  //source/exe:envoy-static  --define=wasm=enabled   --define=mosn_on_envoy=enabled  --remote_http_cache=http://bazel-test.cache.alipay.net  --verbose_failures
# disable wasm and mosn on envoy
bazel build -c opt  //source/exe:envoy-static  --define=wasm=disabled   --define=mosn_on_envoy=disabled  --remote_http_cache=http://bazel-test.cache.alipay.net  --verbose_failures
```
