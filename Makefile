generate-mocks:
	mockery --dir src/core/useCases/ --all --output src/mocks
	mockery --dir src/repositories/ --all --output src/mocks

terraform:
	terraform init
	terraform plan
	terraform apply -auto-approve

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bootstrap src/main.go

zip:
	zip -r ./build/video-users.zip bootstrap ./src go.mod go.sum
