mysql:
	docker run --name=go_db -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=go_db -p 3306:3306 -d mysql

admin:
	docker run --name myadmin -d --link go_db:db -p 8081:80 phpmyadmin/phpmyadmin
run:
	go run main.go
tidy:
	go mod tidy
get:
	go get
push:
	git add . && git commit -am "updated code" && git push
portainer:
	docker run -d -p 8000:8000 -p 9443:9443 --name portainer --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer-ce:latest
build:
	go build -o bin/hello
run:
	air