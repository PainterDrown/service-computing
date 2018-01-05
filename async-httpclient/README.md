# asycn-httpclient

> 服务计算课程的最后一次作业，通过 Go 的并发思想来理解消息（事件）驱动。

参考：
  + [理解 goroutine 的并发](http://blog.csdn.net/pmlpml/article/details/78850661)
  + [Reactive JAX-RS Client API](https://jersey.github.io/documentation/latest/rx-client.html)


## Usage

```sh
# mode 1
go run main.go 1

# mode 2
go run main.go 2
```

根据参考博客，mode 1 的运行时间大约为 5400ms，mode 2 的运行时间大约为 730ms。

## 作业要求

  + 依据文档图6-1，用中文描述 Reactive 动机
  + 使用 go HTTPClient 实现图 6-2 的 Naive Approach
  + 为每个 HTTP 请求设计一个 goroutine ，利用 Channel 搭建基于消息的异步机制，实现图 6-3
  + 对比两种实现，用数据说明 go 异步 REST 服务协作的优势
  + 思考： 是否存在一般性的解决方案？

## 一般性的解决方案

消息驱动的并发模型中，我觉得可以通过无数据依赖关系的并行计算来提升服务速度，思想类比于流水线 CPU...
