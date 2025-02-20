# start app
.PHONY : run
run:
	go mod download && go run cmd/main.go --config=/app/config/application.yaml