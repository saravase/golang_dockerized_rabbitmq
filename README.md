# golang_dockerized_rabbitmq

### Golang:

The Go programming language is an open source project to make programmers more productive.

Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it easy to write programs that get the most out of multicore and networked machines, while its novel type system enables flexible and modular program construction. Go compiles quickly to machine code yet has the convenience of garbage collection and the power of run-time reflection. It's a fast, statically typed, compiled language that feels like a dynamically typed, interpreted language.

### gofiber:

Fiber is an Express inspired web framework built on top of Fasthttp, the fastest HTTP engine for Go. Designed to ease things up for fast development with zero memory allocation and performance in mind.

### RabbitMQ:

 RabbitMQ is one of the most popular open source message brokers. From T-Mobile to Runtastic, RabbitMQ is used worldwide at small startups and large enterprises.

RabbitMQ is lightweight and easy to deploy on premises and in the cloud. It supports multiple messaging protocols. RabbitMQ can be deployed in distributed and federated configurations to meet high-scale, high-availability requirements

### Environment Setup:

#### Create unique network:
    $ docker network create rabbits

#### Run rabbitmq docker image [root@rabbit-1]: 
    $ docker run -d --rm --net rabbits --hostname  rabbit-1 --name rabbit-1 rabbitmq:3.8

#### Show running docker container:
    $ docker ps

    CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                                                 NAMES
    df8220c6551e        rabbitmq:3.8        "docker-entrypoint.s…"   12 minutes ago      Up 12 minutes       4369/tcp, 5671-5672/tcp, 15691-15692/tcp, 25672/tcp   rabbit-1

#### Show rabbitmq docker container logs:
    $ docker logs rabbit-1

#### Execute rabbitmq docker container:
    $ docker exec -it rabbit-1 bash

#### Show rabbitmq CLI commands:
    root@rabbit-1:/# rabbitmqctl

#### Show rabbitmq plugin list:
    root@rabbit-1:/# rabbitmq-plugins list

### Add Management Plugin:

#### Exit from rabbitmq CLI:
    root@rabbit-1:/# exit

#### Remove running rabbitmq container:
    $ docker rm -f rabbit-1

#### Create rabbitmq docker container with management port[-p 9090:15672]:
    $ docker run -d --rm --net rabbits -p 9090:15672 --hostname  rabbit-1 --name rabbit-1 rabbitmq:3.8

#### Execute Container:
    $ docker exec -it rabbit-1 bash

#### Enable rabbitmq management plugins:
    root@rabbit-1:/# rabbitmq-plugins enable rabbitmq_management

#### Show rabbitmq plugin list:
    root@rabbit-1:/# rabbitmq-plugins list

#### Execute managment plugin:
    [guest/guest] http://localhost:9090/

### Build Pubisher Application

#### Move to docker file path:
    $ cd /app/publisher

#### Build docker publisher application:
    $ docker build . -t <docker-id>/rabbitmq-publisher:v1.0.0

#### Run docker publisher application:
    $ docker run -it --rm --net rabbits -e RABBIT_HOST=rabbit-1 -e RABBIT_PORT=5672 -e RABBIT_USER=guest  -e RABBIT_PASSWORD=guest -p 9000:9000 <docker-id>/rabbitmq-publisher:v1.0.0

#### Publish message in POST [localhost:9000/publish]:
    localhost:9000/publish
    
    {
        "msg": "data1"
    }

#### Publish message O/P:
    2020/12/05 16:38:04 Published message: data1
    {
        "message": "Message publish successfully"
    }

### Build Consumer Application

#### Move to docker file path:
    $ cd /app/consumer

#### Build docker publisher application:
    $ docker build . -t <docker-id>/rabbitmq-consumer:v1.0.0

#### Run docker publisher application:
    $ docker run -it --rm --net rabbits -e RABBIT_HOST=rabbit-1 -e RABBIT_PORT=5672 -e RABBIT_USER=guest  -e RABBIT_PASSWORD=guest <docker-id>/rabbitmq-consumer:v1.0.0

#### Consume O/P:
    2020/12/05 16:38:31  [*] Waiting for messages. To exit press CTRL+C
    2020/12/05 16:38:31 Received a message: data1

## Clustering


    $ docker network create rabbits

    $ docker run -d --rm --net rabbits --hostname rabbit-1 --name rabbit-1 rabbitmq:3.8

    $ docker exec -it rabbit-1 cat /var/lib/rabbitmq/.erlang.cookie

        BRYMAXXWBGGLNODJWFPF

#### Manual clutering:
   
    $ docker run -d --rm --net rabbits --hostname rabbit-1 --name rabbit-1 -p 9091:15672 rabbitmq:3.    8-management
    docker run -d --rm --net rabbits --hostname rabbit-2 --name rabbit-2 -p 9092:15672 rabbitmq:3.8-management
    docker run -d --rm --net rabbits --hostname rabbit-3 --name rabbit-3 -p 9093:15672 rabbitmq:3.8-management

    $ docker exec -it rabbit-1 rabbitmqctl cluster_status

    Node-2:

    docker exec -it rabbit-2 rabbitmqctl stop_app
    docker exec -it rabbit-2 rabbitmqctl reset
    docker exec -it rabbit-2 rabbitmqctl join_cluster rabbit@rabbit-1
    docker exec -it rabbit-2 rabbitmqctl start_app
    docker exec -it rabbit-2 rabbitmqctl cluster_status

    Node-3:

    docker exec -it rabbit-3 rabbitmqctl stop_app
    docker exec -it rabbit-3 rabbitmqctl reset
    docker exec -it rabbit-3 rabbitmqctl join_cluster rabbit@rabbit-1
    docker exec -it rabbit-3 rabbitmqctl start_app
    docker exec -it rabbit-3 rabbitmqctl cluster_status

    docker rm -f rabbit-1
    docker rm -f rabbit-2
    docker rm -f rabbit-3

#### Authentication Erlang cookie:

    docker run -d --rm --net rabbits --hostname rabbit-1 --name rabbit-1 -p 9091:15672 -e RABBITMQ_ERLANG_COOKIE=BRYMAXXWBGGLNODJWFPF rabbitmq:3.8-management
    docker run -d --rm --net rabbits --hostname rabbit-2 --name rabbit-2 -p 9092:15672 -e RABBITMQ_ERLANG_COOKIE=BRYMAXXWBGGLNODJWFPF rabbitmq:3.8-management
    docker run -d --rm --net rabbits --hostname rabbit-3 --name rabbit-3 -p 9093:15672 -e RABBITMQ_ERLANG_COOKIE=BRYMAXXWBGGLNODJWFPF rabbitmq:3.8-management


#### Override primary config file:

    docker run -d --rm --net rabbits --hostname rabbit-1 --name rabbit-1 -p 9091:15672 -e RABBITMQ_ERLANG_COOKIE=BRYMAXXWBGGLNODJWFPF -e RABBITMQ_CONFIG_FILE=/config/rabbitmq rabbitmq:3.8-management
    docker run -d --rm --net rabbits --hostname rabbit-2 --name rabbit-2 -p 9092:15672 -e RABBITMQ_ERLANG_COOKIE=BRYMAXXWBGGLNODJWFPF -e RABBITMQ_CONFIG_FILE=/config/rabbitmq rabbitmq:3.8-management
    docker run -d --rm --net rabbits --hostname rabbit-3 --name rabbit-3 -p 9093:15672 -e RABBITMQ_ERLANG_COOKIE=BRYMAXXWBGGLNODJWFPF -e RABBITMQ_CONFIG_FILE=/config/rabbitmq rabbitmq:3.8-management

#### Replication


#### Mirroring: