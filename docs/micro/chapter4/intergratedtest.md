## 集成测试
### 概念
集成测试的测试对象是整个系统或者某个功能模块，比如测试用户注册、登录功能是否正常，是一种端到端的测试。如果测试用例使用到真实的系统时间、真实的文件系统、真实的数据库，亦或者是其他真实的外部依赖，那么该测试已经进入到了集成测试的领域。
![img.png](img.png)

在集群测试中，代码对外部资源存在依赖关系，尽管代码本身的逻辑是完全正确的，但如果依赖不完善，则可能会导致测试的失败，我们需要认识到这个问题。

### 引用
*[Golang单元测试实践指南](https://mp.weixin.qq.com/s?src=11&timestamp=1630576331&ver=3290&signature=BcOBn5vIjb9uNyOe0162iqGckbT6CfTNSI2KR5EKTKRYmFPyajDLlGMBIdaYo91FcVraCTTH71BcgHAjL2nVF76p4UJYXvto5uWAqnTWLwFW7d6F15C-HkpFjGH90myW&new=1)