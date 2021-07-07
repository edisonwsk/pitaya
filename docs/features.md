Features 特性
========

Pitaya has a modular and configurable architecture which helps to hide the complexity of scaling the application and managing clients' sessions and communications.  
>Pitaya有一个模块化和可配置的架构，这有助于隐藏扩展应用程序和管理客户端会话和通信的复杂性。

Some of its core features are described below.  
>它的一些核心特性如下所述。

## Frontend and backend servers 前端和后端服务器

In cluster mode servers can either be a frontend or backend server.  
>在集群模式下，服务器既可以是前端服务器，也可以是后端服务器。

Frontend servers must specify listeners for receiving incoming client connections. They are capable of forwarding received messages to the appropriate servers according to the routing logic.  
>前端服务器必须指定侦听器来接收传入的客户端连接。它们能够根据路由逻辑将接收到的消息转发到适当的服务器。

Backend servers don't listen for connections, they only receive RPCs, either forwarded client messages (sys rpc) or RPCs from other servers (user rpc).  
>后端服务器不监听连接，它们只接收rpc，或者转发的客户端消息(sys rpc)，或者来自其他服务器(用户rpc)的rpc。

## Groups 组

Groups are structures which store information about target users and allows sending broadcast messages to all users in the group and also multicast messages to a subset of the users according to some criteria.  
>组是一种结构，它存储关于目标用户的信息，并允许向组中的所有用户发送广播消息，也可以根据某些标准向用户子集发送多播消息。

They are useful for creating game rooms for example, you just put all the players from a game room into the same group and then you'll be able to broadcast the room's state to all of them.  
>它们对于创建游戏房间非常有用。例如，你只需将一个游戏房间的所有玩家放到同一个组中，然后你就可以向所有人广播房间的状态。

## Listeners 监听器

Frontend servers must specify one or more acceptors to handle incoming client connections, Pitaya comes with TCP and Websocket acceptors already implemented, and other acceptors can be added to the application by implementing the acceptor interface.  
>前端服务器必须指定一个或多个接收端来处理传入的客户端连接，Pitaya附带已经实现的TCP和Websocket接收端，其他接收端可以通过实现acceptor接口添加到应用程序中。

## Acceptor Wrappers 接收器包装器

Wrappers can be used on acceptors, like TCP and Websocket, to read and change incoming data before performing the message forwarding. To create a new wrapper just implement the Wrapper interface (or inherit the struct from BaseWrapper) and add it into your acceptor by using the WithWrappers method. Next there are some examples of acceptor wrappers.   
>包装器可以在接收者(如TCP和Websocket)上使用，在执行消息转发之前读取和更改传入数据。要创建一个新的包装器，只需实现包装器接口(或从BaseWrapper继承该结构)，并使用WithWrappers方法将其添加到接受程序中。下面是一些接受器包装的示例。

### Rate limiting 频率限制
Read the incoming data on each player's connection to limit requests troughput. After the limit is exceeded, requests are dropped until slots are available again. The requests count and management is done on player's connection, therefore it happens even before session bind. The used algorithm is the [Leaky Bucket](https://en.wikipedia.org/wiki/Leaky_bucket#Comparison_with_the_token_bucket_algorithm). This algorithm represents a leaky bucket that has its output flow slower than its input flow. It saves each request timestamp in a `slot` (of a total of `limit` slots) and this slot is freed again after `interval`. For example: if `limit` of 1 request in an `interval` of 1 second, when a request happens at 0.2s the next request will only be handled by pitaya after 1s (at 1.2s).  
>读取每个玩家连接上的传入数据，以限制请求吞吐量。超过限制后，请求将被丢弃，直到插槽再次可用。请求计数和管理是在玩家的连接上完成的，因此它甚至发生在会话绑定之前。使用的算法是漏桶。该算法表示一个泄漏的桶，其输出流比输入流慢。它将每个请求时间戳保存在一个插槽(总共的限制插槽)中，这个插槽在间隔之后再次释放。例如:如果在1秒的间隔内限制一个请求，当一个请求在0.2秒发生时，下一个请求将在1秒(1.2秒)后被pitaya处理。

```
0     request 请求
|--------|
   0.2s
0                 available again 再次可用
|------------------------|
|- 0.2s -|----- 1s ------|
```

## Message forwarding 消息转发

When a server instance receives a client message, it checks the target server type by looking at the route. If the target server type is different from the receiving server type, the instance forwards the message to an appropriate server instance of the correct type. The client doesn't need to take any action to forward the message, this process is done automatically by Pitaya.  
>当服务器实例接收到客户端消息时，它通过查看路由来检查目标服务器类型。如果目标服务器类型与接收服务器类型不同，则实例将消息转发到正确类型的适当服务器实例。客户端不需要采取任何行动转发消息，这个过程是由Pitaya自动完成的。

By default the routing function chooses one instance of the target server type at random. Custom functions can be defined to change this behavior.  
>缺省情况下，路由功能随机选择一个目标服务器类型的实例。可以定义自定义函数来更改此行为。

## Message push 消息推送

Messages can be pushed to users without previous information about either session or connection status. These push messages have a route (so that the client can identify the source and treat properly), the message, the target ids and the server type the client is expected to be connected to.  
>消息可以推送给用户，而不需要先前关于会话或连接状态的信息。这些推送消息有一个路由(以便客户机能够识别源并正确处理)、消息、目标id和客户机希望连接到的服务器类型。

## Modules 模块

Modules are entities that can be registered to the Pitaya application and must implement the defined [interface](https://github.com/topfreegames/pitaya/tree/master/interfaces/interfaces.go#L24). Pitaya is responsible for calling the appropriate lifecycle methods as needed, the registered modules can be retrieved by name.  
>模块是可以注册到火龙果应用程序的实体，并且必须实现定义的接口。Pitaya负责根据需要调用适当的生命周期方法，注册的模块可以按名称检索。

Pitaya comes with a few already implemented modules, and more modules can be implemented as needed. The modules Pitaya has currently are:  
>Pitaya带有一些已经实现的模块，更多的模块可以根据需要实现。火龙果目前的模块有:

### Binary 二进制

This module starts a binary as a child process and pipes its stdout and stderr to info and error log messages, respectively.  
>该模块将二进制文件作为子进程启动，并将其stdout和stderr分别输送到info和error日志消息。

### Unique session 唯一会话

This module adds a callback for `OnSessionBind` that checks if the id being bound has already been bound in one of the other frontend servers.  
>该模块为OnSessionBind添加了一个回调，用于检查被绑定的id是否已经绑定到其他前端服务器之一。

### Binding storage 绑定存储

This module implements functionality needed by the gRPC RPC implementation to enable the functionality of broadcasting session binds and pushes to users without knowledge of the servers the users are connected to.  
>这个模块实现了gRPC RPC实现所需要的功能，使广播会话绑定和推送功能能够在不知道用户连接的服务器的情况下发送给用户。

## Monitoring 监控

Pitaya has support for metrics reporting, it comes with Prometheus and Statsd support already implemented and has support for custom reporters that implement the `Reporter` interface. Pitaya also comes with support for open tracing compatible frameworks, allowing the easy integration of Jaeger and others.  
>Pitaya支持指标报告，它已经实现了Prometheus和Statsd支持，并支持实现Reporter接口的定制Reporter。火龙果还支持开放跟踪兼容框架，允许Jaeger和其他容易集成。

The list of metrics reported by the `Reporter` is:
>由`Reporter`报告的指标列表是:

- Response time: the time to process a message, in nanoseconds. It is segmented
  by route, status, server type and response code;  
> 响应时间:处理消息的时间，以纳秒为单位。这是分段
  按路由、状态、服务器类型和响应代码;
- Process delay time: the delay to start processing a message, in nanoseconds;
  It is segmented by route and server type;  
> 处理延迟时间:开始处理消息的延迟，单位为纳秒;
  它根据路由和服务器类型进行分段;
- Exceeded Rate Limit: the number of blocked requests by exceeded rate limiting;  
> 超过速率限制:由于超过速率限制而阻塞的请求的数量;
- Connected clients: number of clients connected at the moment;  
> 已连接客户端:当前已连接的客户端数量;
- Server count: the number of discovered servers by service discovery. It is
  segmented by server type;  
> 服务器数量:按服务发现的服务器数量。它是
  按服务器类型划分;
- Channel capacity: the available capacity of the channel;  
> 信道容量:信道的可用容量;
- Dropped messages: the number of rpc server dropped messages, that is, messages that are not handled;  
> 删除消息:rpc服务器删除消息的数量，即未处理的消息;
- Goroutines count: the current number Goroutines;  
> Goroutines count:当前Goroutines数量;
- Heap size: the current heap size;  
> 堆大小:当前堆大小;
- Heap objects count: the current number of objects at the heap;  
> 堆对象计数:堆上对象的当前数量;
- Worker jobs retry: the current amount of RPC reliability worker job retries;   
> Worker job retry:当前RPC可靠性Worker job的重试次数;
- Worker jobs total: the current amount of RPC reliability worker jobs. It is
  segmented by job status;  
> Worker job total:当前RPC可靠性Worker的数量。它是
  按工作状态分类;
- Worker queue size: the current size of RPC reliability worker job queues. It
  is segmented by each available queue.  
> Worker queue size: RPC reliability Worker作业队列的当前大小。它
  由每个可用的队列分段。

### Custom Metrics 自定义指标

Besides pitaya default monitoring, it is possible to create new metrics. If using only Statsd reporter, no configuration is needed. If using Prometheus, it is necessary do add a configuration specifying the metrics parameters. More details on [doc](configuration.html#metrics-reporting) and this [example](https://github.com/topfreegames/pitaya/tree/master/examples/demo/custom_metrics).  
>除了火龙果默认监控，我们还可以创造新的指标。如果只使用Statsd reporter，则不需要配置。如果使用Prometheus，那么有必要添加一个配置来指定度量参数。关于doc和这个例子的更多细节。

## Pipelines 管道

Pipelines are middlewares which allow methods to be executed before and after handler requests, they receive the request's context and request data and return the request data, which is passed to the next method in the pipeline.  
>管道是允许方法在处理程序请求之前和之后执行的中间件，它们接收请求的上下文和请求数据，并返回请求数据，这些数据被传递给管道中的下一个方法。

## RPCs

Pitaya has support for RPC calls when in cluster mode, there are two components to enable this, RPC client and RPC server. There are currently two options for using RPCs implemented for Pitaya, NATS and gRPC, the default is NATS.  
>火龙果支持在集群模式下的RPC调用，有两个组件来启用它，RPC客户端和RPC服务器。目前在火龙果上使用rpc有两个选项，NATS和gRPC，默认是NATS。

There are two types of RPCs, _Sys_ and _User_.  
>有两种类型的rpc, _Sys_ 和 _User_ 。

### Sys RPCs 系统 RPCs

These are the RPCs done by the servers when forwarding handler messages to the appropriate server type.  
>这些是服务器在将处理程序消息转发到适当的服务器类型时完成的rpc。

### User RPCs 用户 RPCs

User RPCs are done when the application actively calls a remote method in another server. The call can specify the ID of the target server or let Pitaya choose one according to the routing logic.  
>当应用程序主动调用另一个服务器中的远程方法时，用户rpc就完成了。调用可以指定目标服务器的ID，或者让Pitaya根据路由逻辑选择一个。

### User Reliable RPCs  用户可靠的  RPCs

These are done when the application calls a remote using workers, that is, Pitaya retries the RPC if any error occurrs.  
>这些都是在应用程序使用workers调用远程时完成的，也就是说，如果发生任何错误，Pitaya将重新尝试RPC。

**Important**: the remote that is being called must be idempotent; also the ReliableRPC will not return the remote's reply since it is asynchronous, it only returns the job id (jid) if success.  
>重要提示:被调用的远程必须是幂等的;此外，reliablepc不会返回远程的应答，因为它是异步的，如果成功，它只返回作业id (jid)。

## Server operation mode 服务器运行方式

Pitaya has two types of operation: standalone and cluster mode.  
>火龙果有两种操作模式:单机模式和集群模式。

### Standalone mode 独立模式

In standalone mode the servers don't interact with one another, don't use service discovery and don't have support to RPCs. This is a limited version of the framework which can be used when the application doesn't need to have different types of servers or communicate among them.  
>在独立模式下，服务器不相互交互，不使用服务发现，不支持rpc。这是框架的一个有限版本，当应用程序不需要使用不同类型的服务器或在它们之间通信时，可以使用该框架。

### Cluster mode  集群模式

Cluster mode is a more complete mode, using service discovery, RPC client and server and remote communication among servers of the application. This mode is useful for more complex applications, which might benefit from splitting the responsabilities among different specialized types of servers. This mode already comes with default services for RPC calls and service discovery.  
>集群模式是一种较为完整的模式，利用服务发现、RPC客户机和服务器以及应用程序的服务器之间的远程通信。这种模式对于更复杂的应用程序很有用，它们可以通过在不同的特定类型的服务器之间划分职责而获益。这种模式已经附带了用于RPC调用和服务发现的默认服务。

## Serializers  序列化器

Pitaya has support for different types of message serializers for the messages sent to and from the client, the default serializer is the JSON serializer and Pitaya comes with native support for the Protobuf serializer as well. New serializers can be implemented by implementing the `serialize.Serializer` interface.  
>对于发送到客户端和从客户端发送的消息，火龙果支持不同类型的消息序列化器，默认的序列化器是JSON序列化器，火龙果也支持原生的Protobuf序列化器。新的序列化器可以通过实现serialize来实现。序列化器接口。

The desired serializer can be set by the application by calling the `SetSerializer` method from the `pitaya` package.  
>应用程序可以通过从火龙果包中调用SetSerializer方法来设置所需的序列化器。

## Service discovery 服务发现

Servers operating in cluster mode must have a service discovery client to be able to work. Pitaya comes with a default client using etcd, which is used if no other client is defined. The service discovery client is responsible for registering the server and keeping the list of valid servers updated, as well as providing information about requested servers as needed.  
>在集群模式下运行的服务器必须有一个服务发现客户端才能工作。Pitaya附带了一个使用etcd的默认客户端，如果没有定义其他客户端，就使用etcd。服务发现客户端负责注册服务器并保持有效服务器列表的更新，并根据需要提供有关所请求服务器的信息。

## Sessions 会话

Every connection established by the clients has an associated session instance, which is ephemeral and destroyed when the connection closes. Sessions are part of the core functionality of Pitaya, because they allow asynchronous communication with the clients and storage of data between requests. The main features of sessions are:  
>客户端建立的每个连接都有一个相关的会话实例，它是短暂的，当连接关闭时将被销毁。会话是Pitaya核心功能的一部分，因为它们允许与客户端进行异步通信，并在请求之间存储数据。会议的主要特点是:

* **ID binding** - Sessions can be bound to an user ID, allowing other parts of the application to send messages to the user without needing to know which server or connection the user is connected to  
> **ID binding** - 会话可以绑定到用户ID，允许应用程序的其他部分向用户发送消息，而不需要知道用户连接到哪个服务器或连接
* **Data storage** - Sessions can be used for data storage, storing and retrieving data between requests  
> **Data storage** - 会话可用于数据存储，存储和检索请求之间的数据
* **Message passing** - Messages can be sent to connected users through their sessions, without needing to have knowledge about the underlying connection protocol  
> **Message passing** - 消息可以通过会话发送给已连接的用户，而不需要了解底层连接协议
* **Accessible on requests** - Sessions are accessible on handler requests in the context instance  
> **Accessible on requests** - 会话可以在上下文实例中的处理程序请求上访问
* **Kick** - Users can be kicked from the server through the session's `Kick` method  
> **Kick** - 用户可以通过会话的Kick方法被踢出服务器

Even though sessions are accessible on handler requests both on frontend and backend servers, their behavior is a bit different if they are a frontend or backend session. This is mostly due to the fact that the session actually lives in the frontend servers, and just a representation of its state is sent to the backend server.  
>尽管会话可以通过前端和后端服务器上的处理程序请求访问，但如果它们是前端或后端会话，它们的行为会有所不同。这主要是由于会话实际上存在于前端服务器中，只是将其状态的表示发送到后端服务器。

A session is considered a frontend session if it is being accessed from a frontend server, and a backend session is accessed from a backend server. Each kind of session is better described below.
>如果会话是从前端服务器访问的，而后端会话是从后端服务器访问的，则会话被认为是前端会话。下面将更好地描述每种会话。

### Frontend sessions 前端会话

Sessions are associated to a connection in the frontend server, and can be retrieved by session ID or bound user ID in the server the connection was established, but cannot be retrieved from a different server.  
>会话与前端服务器中的连接相关联，可以通过建立连接的服务器中的会话ID或绑定用户ID检索会话，但不能从其他服务器检索会话。  

Callbacks can be added to some session lifecycle changes, such as closing and binding. The callbacks can be on a per-session basis (with `s.OnClose`) or for every session (with `OnSessionClose`, `OnSessionBind` and `OnAfterSessionBind`).  
>回调函数可以添加到某些会话生命周期更改中，例如关闭和绑定。回调函数可以基于每个会话(s.OnClose)，也可以针对每个会话(OnSessionClose, OnSessionBind和OnAfterSessionBind)。

### Backend sessions 后端会话

Backend sessions have access to the sessions through the handler's methods, but they have some limitations and special characteristics. Changes to session variables must be pushed to the frontend server by calling `s.PushToFront` (this is not needed for `s.Bind` operations), setting callbacks to session lifecycle operations is also not allowed. One can also not retrieve a session by user ID from a backend server.  
>后端会话可以通过处理程序的方法访问会话，但是它们有一些限制和特殊的特征。对会话变量的更改必须通过调用s.PushToFront来推送到前端服务器(这对于s.Bind操作是不需要的)，设置回调到会话生命周期操作也是不允许的。也不能通过用户ID从后端服务器检索会话。
