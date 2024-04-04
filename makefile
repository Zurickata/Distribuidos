docker:
	docker-compose build
	docker-compose up

capitanes:
	cd cliente && go run cliente.go
