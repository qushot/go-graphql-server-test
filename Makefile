# MySQL
.PHONY: mysql/up mysql/log mysql/exec mysql/rm
mysql/up:
	docker-compose -f docker/docker-compose.yml up -d mysql
mysql/log:
	docker-compose -f docker/docker-compose.yml logs -f mysql
mysql/exec:
	docker-compose -f docker/docker-compose.yml exec mysql mysql -u user -p
mysql/rm:
	docker-compose -f docker/docker-compose.yml rm -fsv mysql

# Go
.PHONY: connectiontest
connectiontest:
	go run ./connectiontest/main.go

# All
.PHONY: down
down:
	docker-compose -f docker/docker-compose.yml down