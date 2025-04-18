generate-mocks:
	mockery --dir src/core/_interfaces/ --all --output src/core/_mocks

terraform:
	terraform init
	terraform plan
	terraform apply -auto-approve

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bootstrap src/main.go

zip:
	zip -r ./build/video-users.zip bootstrap ./src go.mod go.sum
