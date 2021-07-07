Communication 通信
=============

In this section we will describe in detail the communication process between the client and the server. From establishing the connection, sending a request and receiving a response. The example is going to assume the application is running in cluster mode and that the target server is not the same as the one the client is connected to.   
>在本节中，我们将详细描述客户机和服务器之间的通信过程。从建立连接，发送请求和接收响应。该示例将假设应用程序以集群模式运行，并且目标服务器与客户机连接到的服务器不相同。

## Establishing the connection 建立连接

The overview of what happens when a client connects and makes a request is:  
>当客户端连接并发出请求时，会发生什么?

* Establish low level connection with acceptor  
  >与接收方建立低层联系
* Pass the connection to the handler service  
> 将连接传递到处理程序服务
* Handler service creates a new agent for the connection  
> 处理程序服务为连接创建一个新的代理
* Handler service reads message from the connection  
  >处理程序服务从连接读取消息
* Message is decoded with configured decoder  
> 消息用配置的解码器解码
* Decoded packet from the message is processed  
> 处理消息中已解码的数据包
* First packet must be a handshake request, to which the server returns a handshake response with the serializer, route dictionary and heartbeat timeout  
> 第一个包必须是一个握手请求，服务器返回一个握手响应，其中包含序列化程序、路由字典和心跳超时
* Client must then reply with a handshake ack, connection is then established  
> 然后客户端必须回复一个握手确认，然后连接建立
* Data messages are processed by the handler and the target server type is extracted from the message route, the message is deserialized using the specified method  
> 数据消息由处理程序处理，目标服务器类型从消息路由中提取，消息使用指定的方法反序列化
* If the target server type is different from the current server, the server makes a remote call to the right type of server, selecting one server according to the routing function logic. The remote call includes the current representation of the client's session  
>  如果目标服务器类型与当前服务器类型不同，则远程调用相应类型的服务器，根据路由功能逻辑选择服务器。远程调用包括客户机会话的当前表示
* The receiving remote server receives the request and handles it as a _Sys_ RPC call, creating a new remote agent to handle the request, this agent receives the session's representation  
> 接收请求的远程服务器接收请求并将其作为Sys RPC调用处理，创建一个新的远程代理来处理请求，该代理接收会话的表示
* The before pipeline functions are called and the handler message is deserialized  
> 调用before管道函数，反序列化处理程序消息
* The appropriate handler is then called by the remote server, which returns the response that is then serialized and the after pipeline functions are executed  
> 然后，远程服务器调用适当的处理程序，该处理程序返回被序列化的响应，并执行after管道函数
* If the backend server wants to modify the session it needs to modify and push the modifications to the frontend server explicitly  
> 如果后端服务器想要修改会话，它需要修改并显式地将修改推送到前端服务器
* Once the frontend server receives the response it forwards the message to the session specifying the request message ID  
  >一旦前端服务器接收到响应，它就将消息转发给指定请求消息ID的会话
* The agent receives the requests, encodes it and sends to the low-level connection  
> 代理接收请求，对其进行编码并将其发送到低级连接

### Acceptors 接收器

The first thing the client must do is establish a connection with the Pitaya server. And for that to happen, the server must have specified one or more acceptors.  
>客户端必须做的第一件事是与火龙果服务器建立连接。为了实现这一点，服务器必须指定一个或多个接收器。

Acceptors are the entities responsible for listening for connections, establishing them, abstracting and forwarding them to the handler service. Pitaya comes with support for TCP and websocket acceptors. Custom acceptors can be implemented and added to Pitaya applications, they just need to implement the proper interface.  
>接收器是负责侦听连接、建立连接、抽象连接并将其转发给处理程序服务的实体。火龙果支持TCP和websocket接收器。定制的接受程序可以被实现并添加到火龙果应用程序中，它们只需要实现适当的接口。

### Handler service 处理服务

After the low level connection is established it is passed to the handler service to handle. The handler service is responsible for handling the lifecycle of the clients' connections. It reads from the low-level connection, decodes the received packets and handles them properly, calling the local server's handler if the target server type is the same as the local one or forwarding the message to the remote service otherwise.  
>在建立了低级连接之后，它将被传递给处理程序服务来处理。处理程序服务负责处理客户端连接的生命周期。它从低级连接读取、解码接收到的包并正确处理它们，如果目标服务器类型与本地服务器类型相同，则调用本地服务器的处理程序，否则将消息转发给远程服务。

Pitaya has a configuration to define the number of concurrent messages being processed at the same time, both local and remote messages count for the concurrency, so if the server expects to deal with slow routes this configuration might need to be tweaked a bit. The configuration is `pitaya.concurrency.handler.dispatch`.  
>Pitaya有一个配置来定义同时处理的并发消息的数量，本地和远程消息都计算并发性，所以如果服务器希望处理较慢的路由，这个配置可能需要稍微调整一下。配置是pitaya. concurrent .handler.dispatch。

### Agent 代理

The agent entity is responsible for storing information about the client's connection, it stores the session, encoder, serializer, state, connection, among others. It is used to communicate with the client to send messages and also ensure the connection is kept alive.  
>代理实体负责存储关于客户端连接的信息，它存储会话、编码器、序列化程序、状态、连接等。它用于与客户端通信以发送消息，并确保连接保持活跃。

### Route compression 路由压缩

The application can define a dictionary of compressed routes before starting, these routes are sent to the clients on the handshake. Compressing the routes might be useful for the routes that are used a lot to reduce the communication overhead.  
>应用程序可以在开始之前定义一个压缩路由字典，这些路由在握手时被发送到客户端。压缩路由对于那些需要大量使用以减少通信开销的路由可能是有用的。

### Handshake 握手

The first operation that happens when a client connects is the handshake. The handshake is initiated by the client, who sends informations about the client, such as platform, version of the client library, and others, and can also send user data in this step. This data is stored in the client's session and can be accessed later. The server replies with heartbeat interval, name of the serializer and the dictionary of compressed routes.  
>客户端连接时发生的第一个操作是握手。握手由客户机发起，客户机发送关于客户机的信息，例如平台、客户机库的版本，以及其他信息，并且还可以在此步骤中发送用户数据。这些数据存储在客户机的会话中，可以稍后访问。服务器用心跳间隔、序列化程序名称和压缩路由字典进行响应。

### Remote service 远程服务

The remote service is responsible both for making RPCs and for receiving and handling them. In the case of a forwarded client request the RPC is of type _Sys_.  
>远程服务既负责制作rpc，也负责接收和处理它们。在转发客户端请求的情况下，RPC类型为Sys。

In the calling side the service is responsible for identifying the proper server to be called, both by server type and by routing logic.  
>在调用端，服务负责根据服务器类型和路由逻辑识别要调用的正确服务器。

In the receiving side the service identifies it is a _Sys_ RPC and creates a remote agent to handle the request. This remote agent is short-lived, living only while the request is alive, changes to the backend session do not automatically reflect in the associated frontend session, they need to be explicitly committed by pushing them. The message is then forwarded to the appropriate handler to be processed.
>在接收端，服务标识它是一个Sys RPC，并创建一个远程代理来处理请求。此远程代理是短暂的，仅在请求处于活动状态时存在，对后端会话的更改不会自动反映到相关的前端会话中，它们需要通过推送显式提交。然后将消息转发给适当的处理程序进行处理。

### Pipeline 管道

The pipeline in Pitaya is a set of functions that can be defined to be run before or after every handler request. The functions receive the context and the raw message and should return the request object and error, they are allowed to modify the context and return a modified request. If the before function returns an error the request fails and the process is aborted.   
>Pitaya中的管道是一组函数，可以定义为在每个处理程序请求之前或之后运行。函数接收上下文和原始消息，并返回请求对象和错误，它们可以修改上下文并返回修改后的请求。如果before函数返回错误，则请求失败，进程中止。

### Serializer 序列化器

The handler must first deserialize the message before processing it. So the function responsible for calling the handler method first deserializes the message, calls the method and then serializes the response returned by the method and returns it back to the remote service.  
>处理程序在处理消息之前必须首先反序列化消息。因此，负责调用处理程序方法的函数首先对消息进行反序列化，然后调用方法，然后对方法返回的响应进行序列化，并将其返回给远程服务。

### Handler 处理器

Each Pitaya server can register multiple handler structures, as long as they have different names. Each structure can have multiple methods and Pitaya will choose the right structure and methods based on the called route.  
>每个火龙果服务器可以注册多个处理器结构，只要它们有不同的名称。每个结构可以有多种方法，火龙果会根据调用的路由选择合适的结构和方法。
