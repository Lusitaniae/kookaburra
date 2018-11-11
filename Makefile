.PHONY: install format vet build serve clean pack deploy ship

TAG  ?= $(shell git rev-list HEAD --max-count=1 --abbrev-commit)
REPO ?= lusotycoon/kookaburra
PKG  ?= kookaburra
pkgs  = $(shell go list ./...)

export TAG

# test: install
# 	@echo ">> running tests"
# 	go test ./...

format:
	@echo ">> formatting code"
	go fmt $(pkgs)

vet:
	@echo ">> vetting code"
	go vet $(pkgs)

build: install
	@echo ">> building binary"
	go build -ldflags "-X main.version=$(TAG)" -o $(PKG) .

serve: build
	@echo ">> serving application"
	./$(PKG)

clean:
	@echo ">> removing application"
	rm ./$(PKG)

pack:
	@echo ">> building docker image"
	GOOS=linux make build
	docker build -t $(REPO):$(TAG) .

upload:
	@echo ">> pushing docker image"
	docker push $(REPO):$(TAG)

deploy:
	@echo ">> deploying to Kubernetes"
	envsubst < k8s/deployment.yml | kubectl apply -f -

ship: pack upload deploy clean

prometheus:
	@echo ">> deploying prometheus"
	kubectl apply -f k8s/prometheus-deployment.yml
	kubectl apply -f k8s/node-exporter.yml
