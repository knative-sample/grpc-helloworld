# grpc-helloworld
An example of gRPC helloworld

## usage
- start server

```
docker run --rm -it -p 8080:8080 registry.cn-hangzhou.aliyuncs.com/knative-sample/helloworld-server:2020-03-24_173752
```

- start client

```
docker run --rm -it -e GRPC_CONCURRENT="10000" -e GRPC_SERVER_ADDR=${server_addr}:8080 registry.cn-hangzhou.aliyuncs.com/knative-sample/helloworld-client:2020-03-24_212051
```
  - GRPC_CONCURRENT 指定并发数
  - GRPC_SERVER_ADDR 指定 server 地址

```
└─# docker run --rm -it -e GRPC_SERVER_ADDR=30.5.123.238:8080 client
020/04/21 15:03:18 resp: key: Client Hello for [322]  --> resp index:322, latency:10
2020/04/21 15:03:18 resp: key: Client Hello for [625]  --> resp index:625, latency:10
2020/04/21 15:03:18 resp: key: Client Hello for [517]  --> resp index:517, latency:10
2020/04/21 15:03:18 resp: key: Client Hello for [384]  --> resp index:384, latency:10
2020/04/21 15:03:18 resp: key: Client Hello for [377]  --> resp index:377, latency:10
2020/04/21 15:03:18 resp: key: Client Hello for [786]  --> resp index:786, latency:10
2020/04/21 15:03:18 resp: key: Client Hello for [874]  --> resp index:874, latency:10
2020/04/21 15:03:18 resp: key: Client Hello for [426]  --> resp index:426, latency:10
2020/04/21 15:03:18 resp: key: Client Hello for [979]  --> resp index:979, latency:10
... ... 
```
