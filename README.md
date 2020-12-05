# golang_dockerized_rabbitmq

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
