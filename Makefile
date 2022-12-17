runUser:
	cd UserService && go run main.go

runMain:
	cd MainService && go run main.go

runFe:
	cd fe && npm start

mysql:
	docker run --name mysql -p 3306  -e MYSQL_ROOT_USER=root -d mysql

createdb:
	docker exec -i sefi mysql -uroot -proot <<< "CREATE SCHEMA IF NOT EXISTS sefi;"

droppedb:
	docker exec -it sefi mysql -uroot -proot <<< "DROP DATABASE sefi-userService-dev;"

migrateup:
	cd UserService && migrate -path app/db/mysql/migration -database "mysql://root:root@tcp(localhost:3306)/test_db?parseTime=true" -verbose up

migratedown:
	cd UserService && migrate -path app/db/mysql/migration -database "mysql://root:root@tcp(localhost:3306)/test_db?parseTime=true" -verbose down

sqlc:
	cd UserService && sqlc generate