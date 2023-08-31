run:
	chmod +x cmd/segmentation/wait-for-it.sh ; docker-compose up  --remove-orphans --build postgres segmentation

lint:
	gofmt -l -w .
	golangci-lint run --fix