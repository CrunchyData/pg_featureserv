.DEFAULT_GOAL := help

PROJECT_DIR := $(shell pwd)
BUILD_DIR := $(PROJECT_DIR)/build
VENDOR_DIR := $(PROJECT_DIR)/vendor
DOCKER_DIR := $(PROJECT_DIR)/docker
TOOLS_DIR := $(PROJECT_DIR)/tools

VERSION := 1.1.0


.PHONY: all
all: check clean build docker-images

.PHONY: check
check:
	@go version

.PHONY: clean
clean:
	$(info Cleaning project...)
	@rm -rf $(BUILD_DIR)
	@rm -rf $(VENDOR_DIR)

.PHONY: build
build: ## Build Go executables
build: build/pg_featureserv


.PHONY: build/pg_featureserv
build/pg_featureserv:
	$(info Building crunchy-cloud-api...)
	go build -o $@ -ldflags='-s -w' .

.PHONY: build/operator-events-agent
build/operator-events-agent:
	$(info Building operator-events-agent...)
	go build -o $@ -ldflags='-s -w' ./apps/operator-events-agent

.PHONY: build/plugin-crunchy-cloud.so
build/plugin-crunchy-cloud.so:
	$(info Building plugin-crunchy-cloud...)
	go build -o $@ -buildmode=plugin -ldflags='-s -w' ./plugins/crunchy_cloud

.PHONY: docker-build
docker-build: ## Build Go executables inside a Linux container
	docker-compose -f docker/docker-compose.build.yaml run --rm -T golang make check build

.PHONY: docker-clean
docker-clean: ## Remove any Docker volumes created during development
docker-clean: | docker-clean-volumes

.PHONY: docker-clean-volumes
docker-clean-volumes:
	docker-compose -f docker/docker-compose.build.yaml down --volumes

.PHONY: docker-images
docker-images: apiserver-image keycloak-image operator-events-agent-image

.PHONY: help
help: ALIGN=14
help: ## Print this message
	@awk -F ': ## ' -- "/^[^':]+: ## /"' { printf "'$$(tput bold)'%-$(ALIGN)s'$$(tput sgr0)' %s\n", $$1, $$2 }' $(MAKEFILE_LIST)

.PHONY: apiserver-image
apiserver-image:
	@docker build -t crunchydata/crunchy-cloud-apiserver:$(VERSION) \
			--build-arg VERSION=$(VERSION) \
			-f $(DOCKER_DIR)/Dockerfile.apiserver .

.PHONY: keycloak-image
keycloak-image: ensure-keycloak-user-storage-provider
	@echo "Building keycloak image..."
	@docker build \
			--build-arg USER_PROVIDER_VERSION=$(KEYCLOAK_USER_STORE_VERSION) \
			-t crunchydata/crunchy-cloud-keycloak:$(KEYCLOAK_IMAGE_VERSION) \
			-f $(DOCKER_DIR)/Dockerfile.keycloak .

.PHONY: operator-events-agent-image
operator-events-agent-image:
	@$(info Building operator-events-agent image...)
	@docker build -t crunchydata/operator-events-agent:$(VERSION) \
			--build-arg VERSION=$(VERSION) \
			-f $(DOCKER_DIR)/Dockerfile.operator-events-agent .

.PHONY: ensure-keycloak-user-storage-provider
ensure-keycloak-user-storage-provider:
ifeq ("$(wildcard $(BUILD_DIR)/$(KEYCLOAK_USER_STORE_FILE))","")
	$(info Artifact $(BUILD_DIR)/$(KEYCLOAK_USER_STORE_FILE) does not exist.)
	$(info Downloading: $(KEYCLOAK_USER_STORE_FILE))

	@$(TOOLS_DIR)/gh-dl-release \
		$(KEYCLOAK_USER_STORE_REPO) \
		$(KEYCLOAK_USER_STORE_FILE) \
		$(KEYCLOAK_USER_STORE_VERSION)\
		$(BUILD_DIR)/$(KEYCLOAK_USER_STORE_FILE)
endif

.PHONY: test
test:
	go test ./... -v
