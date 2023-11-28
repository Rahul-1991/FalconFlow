# FalconFlow 
## A publish-subscribe message distribution model for 1-to-many communication

FalconFlow is an attempt to create a lite version of Message Oriented Middleware NATS using golang.
NATS is an open-source messaging broker (also sometimes referred to as message-oriented middleware). It’s written in Go and you can find it’s source code in the [nats.io](https://github.com/nats-io/nats-server/) GitHub repository.

## Features

- Connect via telnet
- Basic operations like PUB, SUB, CONNECT
- Handle concurrent connections
- Subscription based on topics as well as subscription id
- A lighter version of NATS protocol parser

The project helps to understand how to manage subscription to topics and route published message in a topic to its subscribers. It is really interesting if you want to understand how message queues work under the hood.

## Setting it up

To run falconflow locally follow the below steps
```
git clone https://github.com/Rahul-1991/FalconFlow.git
cd falconflow
go run main.go
```

## Sample Commands
After the server is running and accepting new tcp connections you can connect to it using telnet and test the commands
```
telnet localhost 4222
SUB AA 1
PUB AA 12\r\nHello World!\r\n
```
The above commands allow a client to create a tcp connection with the server. Then the client subscribes to a topic AA with subscriber id as 1. Next it uses the PUB command to publish a message of 12 bytes to the topic AA. All clients which have subscribed to the topic will receive this message.
For more information you can refer the [NATS documentation](https://docs.nats.io)

## Development

Want to contribute? Great!

Would love to have people go through my code and help in adding more features to falconflow. You can always create a pull request by forking this repository and adding any enhancement in your branch.

