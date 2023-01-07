mysql:
	docker run --name mysql -p 3306  -e MYSQL_ROOT_USER=root -d mysql

createdb:
	docker exec -i sefi mysql -uroot -proot <<< "CREATE SCHEMA IF NOT EXISTS sefi;"

droppedb:
	docker exec -it sefi mysql -uroot -proot <<< "DROP DATABASE sefi-userService-dev;"
