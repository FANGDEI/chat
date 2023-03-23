all:
	make mysql
	make redis

mysql:
	docker pull mysql:5.7
	docker run --name chat-mysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=chat -d mysql:5.7
	docker cp ./chat.sql chat:/

redis:
	docker pull redis:latest
	docker run -it -d -p 6379:6379 --name chat-redis redis --requirepass "password"

clean:
	docker stop chat-mysql
	docker rm chat-mysql
	docker rmi mysql:5.7
	docker stop chat-redis
	docker rm chat-redis
	docker rmi redis