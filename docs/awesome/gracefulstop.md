## 服务关闭

服务关闭需要注意的点
* 服务关闭错误抛出
* 多服务并发关闭，不能因为并发报错误，停止其他服务的关闭
* 服务优雅关闭，超时后需要转成强制关闭
* 服务超时仍然无法关闭，需要转成进一步的强制关闭
* HTTP的shutdown会导致，listen and server的立刻返回，需要支持register on shutdown
* errgroup的context超时传递，需要每个服务里实现ctx.done和ctx.err，否则会导致服务无法正常cancel

### 服务关闭常规写法·