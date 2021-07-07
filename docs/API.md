Pitaya API
==========

## Handlers 处理器

Handlers are one of the core features of Pitaya, they are the entities responsible for receiving the requests from the clients and handling them, returning the response if the method is a request handler, or nothing, if the method is a notify handler.  
>处理程序是Pitaya的核心特性之一，它们是负责从客户端接收请求并处理请求的实体，如果方法是请求处理程序则返回响应，如果方法是通知处理程序则不返回响应。

### Signature 签名

Handlers must be public methods of the struct and have a signature following:  
>处理程序必须是结构体的公共方法，并具有以下签名:

Arguments 参数
* `context.Context`: the context of the request, which contains the client's session.
  >`context.Context`:请求的上下文，其中包含客户机的会话.
* `pointer or []byte`: the payload of the request (_optional_).
> `pointer or []byte`:请求的有效负载(_可选_)。

Notify handlers return nothing, while request handlers must return:  
>通知处理程序不返回任何内容，而请求处理程序必须返回:
* `pointer or []byte`: the response payload
  >`pointer or []byte`:响应有效负载
* `error`: an error variable  
  >`error`: 错误变量


### Registering handlers  注册处理程序

Handlers must be explicitly registered by the application by calling `pitaya.Register` with a instance of the handler component. The handler's name can be defined by calling `pitaya/component`.WithName(`"handlerName"`) and the methods can be renamed by using `pitaya/component`.WithNameFunc(`func(string) string`).  
>应用程序必须通过调用pitaya来显式地注册处理程序。注册handler组件的实例。处理程序的名称可以通过调用`pitaya/component`.WithName(`"handlerName"`)来定义，方法可以通过使用`pitaya/component`.WithNameFunc(`func(string) string`)来重命名。

The clients can call the handler by calling `serverType.handlerName.methodName`.
>客户端可以通过调用`serverType.handlerName.methodName`来调用处理程序。


### Routing messages 路由消息

Messages are forwarded by pitaya to the appropriate server type, and custom routers can be added to the application by calling `pitaya.AddRoute`, it expects two arguments:  
>消息由pitaya转发到适当的服务器类型，通过调用`pitaya.AddRoute`可以将自定义路由器添加到应用程序中 ，它需要两个参数:

* `serverType`: the server type of the target requests to be routed  
  >`serverType`: 要路由的目标请求的服务器类型
* `routingFunction`: the routing function with the signature `func(*session.Session, *route.Route, []byte, map[string]*cluster.Server) (*cluster.Server, error)`, it receives the user's session, the route being requested, the message and the map of valid servers of the given type, the key being the servers' ids  
  >routingFunction:带有签名func(*session)的路由函数。会话,*路线。(*cluster. server) (*cluster. server) (*cluster. server) (*cluster. server)它接收用户的会话、被请求的路由、消息和给定类型的有效服务器的映射，密钥是服务器的id

The server will then use the routing function when routing requests to the given server type.  
>然后，当将请求路由到给定的服务器类型时，服务器将使用路由函数。


### Lifecycle Methods 生命周期方法

Handlers can optionally implement the following lifecycle methods:  
>处理程序可以选择性地实现以下生命周期方法:

* `Init()` - Called by Pitaya when initializing the application  
> ' Init() ' -在初始化应用程序时由Pitaya调用
* `AfterInit()` - Called by Pitaya after initializing the application  
> ' AfterInit() ' -由Pitaya在初始化应用程序后调用
* `BeforeShutdown()` - Called by Pitaya when shutting down components, but before calling shutdown  
> ' BeforeShutdown() ' - Pitaya在关闭组件时调用，但在调用shutdown之前
* `Shutdown()` - Called by Pitaya after the start of shutdown  
> ' Shutdown() ' -在关闭开始后由Pitaya调用


### Handler example  处理程序的例子

Below is a very barebones example of a handler definition, for a complete working example, check the [cluster demo](https://github.com/topfreegames/pitaya/tree/master/examples/demo/cluster).  
>下面是一个非常简单的处理程序定义示例，要获得一个完整的工作示例，请查看cluster演示。

```go
import (
  "github.com/topfreegames/pitaya"
  "github.com/topfreegames/pitaya/component"
)

type Handler struct {
  component.Base
}

type UserRequestMessage struct {
  Name    string `json:"name"`
  Content string `json:"content"`
}

type UserResponseMessage {
}

type UserPushMessage{
  Command string `json:"cmd"`
}

// Init runs on service initialization (not required to be defined)
func (h *Handler) Init() {}

// AfterInit runs after initialization (not required to be defined)
func (h *Handler) AfterInit() {}

// TestRequest can be called by the client by calling <servertype>.testhandler.testrequest
func (h *Handler) TestRequest(ctx context.Context, msg *UserRequestMessage) (*UserResponseMessage, error) {
  return &UserResponseMessage{}, nil
}

func (h *Handler) TestPush(ctx context.Context, msg *UserPushMessage) {
}

func main() {
  pitaya.Register(
    &Handler{}, // struct to register as handler
    component.WithName("testhandler"), // name of the handler, used by the clients
    component.WithNameFunc(strings.ToLower), // naming conversion scheme to be used by the clients
  )

  ...
}

```

## Remotes 远程

Remotes are one of the core features of Pitaya, they are the entities responsible for receiving the RPCs from other Pitaya servers.  
>远程是火龙果的核心功能之一，它们是负责接收来自其他火龙果服务器的rpc的实体。

### Signature 签名

Remotes must be public methods of the struct and have a signature following:  
>Remotes必须是该结构的公共方法，并具有以下签名:

Arguments 参数
* `context.Context`: the context of the request.  
> `context.Context`: 请求的上下文。
* `proto.Message`: the payload of the request (_optional_).  
> `proto.Message`: 请求的有效载荷(_可选_)。

Remote methods must return:
>远程方法必须返回:
* `proto.Message`: the response payload in protobuf format  
> `proto.Message`: protobuf格式的响应负载
* `error`: an error variable  
> `error`: 错误变量


### Registering remotes 注册远端

Remotes must be explicitly registered by the application by calling `pitaya.RegisterRemote` with a instance of the remote component. The remote's name can be defined by calling `pitaya/component`.WithName(`"remoteName"`) and the methods can be renamed by using `pitaya/component`.WithNameFunc(`func(string) string`).  
>remote必须由应用程序通过调用pitaya显式注册。用远程组件的实例RegisterRemote。远程的名称可以通过调用pitaya/component.WithName("remoteName")来定义，方法可以通过使用pitaya/component.WithNameFunc(func(string) string)来重命名。

The servers can call the remote by calling `serverType.remoteName.methodName`.  
>服务器可以通过调用' serverType.remoteName.methodName '来调用远程服务器。


### RPC calls  远程过程调用

There are two options when sending RPCs between servers:  
>当在服务器之间发送rpc时有两个选项:
* **Specify only server type**: In this case Pitaya will select one of the available servers at random  
> **只指定服务器类型**:在这种情况下，火龙果将随机选择一个可用的服务器
* **Specify server type and ID**: In this scenario Pitaya will send the RPC to the specified server  
> **指定服务器类型和ID**:在这个场景中，火龙果将发送RPC到指定的服务器


### Lifecycle Methods 生命周期方法

Remotes can optionally implement the following lifecycle methods:  
>Remotes可以选择性地实现以下生命周期方法:

* `Init()` - Called by Pitaya when initializing the application  
> ' Init() ' -在初始化应用程序时由Pitaya调用
* `AfterInit()` - Called by Pitaya after initializing the application  
> ' AfterInit() ' -由Pitaya在初始化应用程序后调用
* `BeforeShutdown()` - Called by Pitaya when shutting down components, but before calling shutdown  
> ' BeforeShutdown() ' - Pitaya在关闭组件时调用，但在调用shutdown之前
* `Shutdown()` - Called by Pitaya after the start of shutdown  
> ' Shutdown() ' -在关闭开始后由Pitaya调用

### Remote example  远程的例子

For a complete working example, check the [cluster demo](https://github.com/topfreegames/pitaya/tree/master/examples/demo/cluster).  
>要获得一个完整的工作示例，请查看[cluster demo](https://github.com/topfreegames/pitaya/tree/master/examples/demo/cluster)。
