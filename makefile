start:
	go run main.go

# DB
sql-migrate:
	which $@ || go get github.com/rubenv/sql-migrate/...

migrate/first:
	make migrate/init
	make migrate/up

migrate/init:
	mysql -u iosk -h localhost --protocol tcp -piosk \
	-e "CREATE DATABASE IF NOT EXISTS halCinema /*!40100 DEFAULT CHARACTER SET utf8 */"

migrate/up:
	sql-migrate $(@F) -env="development" -config="dbconfig.yml"

migrate/drop:
	mysql -u iosk -h localhost --protocol tcp -piosk \
	-e "DROP DATABASE IF EXISTS halCinema"

migrate/clean:
	make migrate/drop
	make migrate/init
	make migrate/up
