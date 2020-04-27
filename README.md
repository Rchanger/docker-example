# docker-example
## Small Example of connecting server and client by creating docker image and using docker-compose.
Just execute following command to make it work
  ~~~
  docker-compose up
  ~~~

## The other way of linking two container is using the ___--link___ option.
Perform following step to achieve the same.
1. First inside __server__ folder update the copy command in  dockerfile ___from /server . to . .___ and then execute following command to build server image.
   ~~~ 
   sudo docker build -t server . 
   ~~~
2. Run the server container by executing following command.   
   ~~~ 
   sudo docker run -d -p 8090:8090 -e Port=3030 --name server -v test-volume:/serverdata -t server 
   ~~~
3. Modify the copy command of dockerfile inside __client__ folder ___from /client . to . .___ and excute following command to build the  client image.
   ~~~
   sudo docker build -t client .
   ~~~ 
4. Run the client container by executing following command.
   ~~~
   sudo docker run -d -p 9090:9090 -e ServerHost=server -e ServerPort=3030 --name client --link server -v test-volume:/clientdata -t  client
   ~~~
