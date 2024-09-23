local-build:
	go build -o ./bin/server cmd/server/main.go

run-server:
	./bin/server

clean:
	rm -f ./bin/server

local-run: clean local-build
	APP_SERVER_ENV=local $(MAKE) run-server
