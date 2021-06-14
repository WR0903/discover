## go-kit 学习

用go-kit简单搭建一个http微服务，实现向consul进行服务注册，以及获取consul注册中有多少服务器实例的接口

```http://127.0.0.1:10086/discovery?serviceName=SayHello```

### 使用方法
首先启动consul
```consul agent -dev```

然后运行本程序
