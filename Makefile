
# Script for running dev server
run-setup:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
run-dev:
	air
run-seeds:
	go run cmd/database/seeds.go
