mysql:
	docker run --name=go_db -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=go_db -p 3306:3306 -d mysql

myadmin:
	docker run --name myadmin -d --link go_db:db -p 8081:80 phpmyadmin/phpmyadmin
run:
	go run main.go
tidy:
	go mod tidy
goget:
	go get
push:
	git add . && git commit -am "updated" && git push