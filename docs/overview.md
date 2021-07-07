Overview 综述
========

Pitaya is an easy to use, fast and lightweight game server framework inspired by [starx](https://github.com/lonnng/starx) and [pomelo](https://github.com/NetEase/pomelo) and built on top of [nano](https://github.com/lonnng/nano)'s networking library.  
>火龙果是一个易于使用，快速和轻量级的游戏服务器框架，灵感来自starx和pomelo，并建立在[nano](https://github.com/lonnng/nano)'s网络库。

The goal of pitaya is to provide a basic, robust development framework for distributed multiplayer games and server-side applications.  
>pitaya的目标是为分布式多人游戏和服务器端应用程序提供一个基本的、健壮的开发框架。

## Features 特性

* **User sessions** - Pitaya has support for user sessions, allowing binding sessions to user ids, setting custom data and retrieving it in other places while the session is active  
  >**用户会话** - Pitaya支持用户会话，允许将会话绑定到用户id，设置自定义数据，并在会话激活时在其他地方检索它
* **Cluster support** - Pitaya comes with support to default service discovery and RPC modules, allowing communication between different types of servers with ease  
  >**集群支持** - Pitaya提供了对默认服务发现和RPC模块的支持，允许不同类型的服务器之间轻松地通信
* **WS and TCP listeners** - Pitaya has support for TCP and Websocket acceptors, which are abstracted from the application receiving the requests  
  >WS和TCP监听器——Pitaya支持TCP和Websocket接收器，它们是从接收请求的应用程序中抽象出来的
* **Handlers and remotes** - Pitaya allows the application to specify its handlers, which receive and process client messages, and its remotes, which receive and process RPC server messages. They can both specify custom init, afterinit and shutdown methods  
  >处理程序和远程程序——Pitaya允许应用程序指定接收和处理客户端消息的处理程序，以及接收和处理RPC服务器消息的远程程序。它们都可以指定自定义init、afterinit和shutdown方法
* **Message forwarding** - When a server receives a handler message it forwards the message to the server of the correct type  
  >消息转发——当服务器接收到handler消息时，它将消息转发给正确类型的服务器
* **Client library SDK** - [libpitaya](https://github.com/topfreegames/libpitaya) is the official client library SDK for Pitaya  
  >客户端库SDK - libpitaya是火龙果的官方客户端库SDK
* **Monitoring** - Pitaya has support for Prometheus and statsd by default and accepts other custom reporters that implement the Reporter interface  
  >监视-火龙星默认支持普罗米修斯和statsd，并接受其他定制的记者，实现记者接口
* **Open tracing compatible** - Pitaya is compatible with [open tracing](http://opentracing.io/), so using [Jaeger](https://github.com/jaegertracing/jaeger) or any other compatible tracing framework is simple  
  >开放跟踪兼容-火龙果兼容开放跟踪，所以使用Jaeger或任何其他兼容的跟踪框架是很简单的
* **Custom modules** - Pitaya already has some default modules and supports custom modules as well  
  >自定义模块- Pitaya已经有一些默认模块，并支持自定义模块
* **Custom serializers** - Pitaya natively supports JSON and Protobuf messages and it is possible to add other custom serializers as needed  
  >自定义序列化器——火龙果本身支持JSON和Protobuf消息，也可以根据需要添加其他自定义序列化器
* **Write compatible servers in other languages** - Using [libpitaya-cluster](https://github.com/topfreegames/libpitaya-cluster) its possible to write pitaya-compatible servers in other languages that are able to register in the cluster and handle RPCs, there's already a csharp library that's compatible with unity and a WIP of a python library in the repo.  
  >写在其他语言兼容的服务器——使用libpitaya-cluster可能写pitaya-compatible服务器集群中的其他语言能够注册和处理rpc,已经有一个csharp库兼容统一和在制品的python库的回购。
* **REPL Client for development/debugging** - [Pitaya-cli](https://github.com/topfreegames/pitaya-cli) is a REPL client that can be used for making development and debugging of pitaya servers easier.  
  >REPL客户端用于开发/调试- pitaya -cli是一个REPL客户端，可以用来使开发和调试火龙果服务器更容易。
* **Bots for integration/stress tests** - [Pitaya-bot](https://github.com/topfreegames/pitaya-bot) is a server test framework that can easily copy users behaviour to test corner case scenarios, which can validate the responses received, or make massive accesses into pitaya servers.   
  >用于集成/压力测试的机器人——火龙果-机器人是一个服务器测试框架，可以很容易地复制用户的行为来测试极端情况，这可以验证收到的响应，或大量访问火龙果服务器。

## Architecture 架构

Pitaya was developed considering modularity and extendability at its core, while providing solid basic functionalities to abstract client interactions to well defined interfaces. The full API documentation is available in Godoc format at [godoc](https://godoc.org/github.com/topfreegames/pitaya).  
>火龙果的开发以模块化和可扩展性为核心，同时提供了坚实的基本功能，将客户端交互抽象为定义良好的接口。完整的API文档以Godoc格式在Godoc上提供。

## Who's Using it 谁在使用它

Well, right now, only us at TFG Co, are using it, but it would be great to get a community around the project. Hope to hear from you guys soon!  
>现在，只有我们TFG公司在使用它，但如果能围绕这个项目建立一个社区就太好了。希望很快能收到你们的来信!

## How To Contribute? 如何贡献

Just the usual: Fork, Hack, Pull Request. Rinse and Repeat. Also don't forget to include tests and docs (we are very fond of both).  
>就像平常一样:Fork, Hack, Pull Request。清洗和重复的方法。另外，不要忘记包含测试和文档(我们非常喜欢这两者)。
