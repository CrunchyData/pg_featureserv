##VARIABLES -
##      APPVERSION - Variable to set the version label
##      GOVERSION - Defaults to 1.21.6 but can be overriden, uses alpine go container as base
##      PROGRAM - Name of binary, pg_featureserv
##      CONTAINER - prefix and name of the generated container
##      DATE - Date String used as alternate tag for generated containers
##      BASE_REGISTRY - This is the registry to pull the base image from
##      BASE_IMAGE - The base image to use for the final container
##      TARGETARCH - The architecture the resulting image is based on and the binary is compiled for
##      IMAGE_TAG - The container and tag to be applied to the container

APPVERSION := latest
GOVERSION := 1.21.6
PROGRAM := pg_featureserv
CONTAINER := pramsey/$(PROGRAM)
DATE := $(shell date +%Y%m%d)
BASE_REGISTRY := registry.access.redhat.com
BASE_IMAGE := ubi8-micro
SYSTEMARCH = $(shell uname -i)

ifeq ($(SYSTEMARCH), x86_64)
TARGETARCH := amd64
PLATFORM=amd64
else
TARGETARCH := arm64
PLATFORM=arm64
endif

IMAGE_TAG = $(CONTAINER):$(APPVERSION)-$(TARGETARCH)

RM = /bin/rm
CP = /bin/cp
MKDIR = /bin/mkdir
SED = /usr/bin/sed

.PHONY: build bin-for-docker check clean build-in-docker docker docker-build docs install local-docker release test uninstall

.DEFAULT_GOAL := help

GOFILES := $(shell find . -type f -name '*.go')

check:  ##              This checks the current version of Go installed locally
	go version

clean:  ##              This will clean all local build artifacts
	$(info Cleaning project...)
	@rm -f $(PROGRAM)
	@rm -rf docs/*
	@docker image inspect $(CONTAINER):$(APPVERSION) >/dev/null 2>&1 && docker rmi -f $(CONTAINER):$(APPVERSION) $(CONTAINER):$(DATE) || echo -n ""

docs:   ##               Generate docs
	@rm -rf docs/* && cd hugo && hugo && cd ..

build: $(PROGRAM)  ##              Build a local binary using APPVERSION parameter or CI as default

$(PROGRAM): $(GOFILES) 
	go build -v -ldflags "-s -w -X github.com/CrunchyData/pg_featureserv/conf.setVersion=$(APPVERSION)"

bin-for-docker: $(GOFILES)  ##     Build a local binary using APPVERSION parameter or CI as default (to be used in docker image)
# to be used in docker the built binary needs the CGO_ENABLED=0 option
	CGO_ENABLED=0 go build -v -ldflags "-s -w -X github.com/CrunchyData/pg_featureserv/conf.setVersion=$(APPVERSION)"

build-in-docker: $(GOFILES)   ##    Build a local binary based of a golang base docker image without the need of a local go environment
	docker run --rm -v "$(PWD)":/usr/src/myapp:z -w /usr/src/myapp golang:$(GOVERSION) make APPVERSION=$(APPVERSION) $(PROGRAM)

docker-build: Dockerfile 
	docker build -f Dockerfile \
		--target $(BUILDTYPE) \
    --build-arg VERSION=$(APPVERSION) \
    --build-arg GOLANG_VERSION=$(GOVERSION) \
    --build-arg TARGETARCH=$(TARGETARCH) \
		--build-arg PLATFORM=$(PLATFORM) \
    --build-arg BASE_REGISTRY=$(BASE_REGISTRY) \
    --build-arg BASE_IMAGE=$(BASE_IMAGE) \
    --label vendor="Crunchy Data" \
    --label url="https://crunchydata.com" \
    --label release="${APPVERSION}" \
    --label org.opencontainers.image.vendor="Crunchy Data" \
    --label os.version="7.7" \
    -t $(IMAGE_TAG) -t $(CONTAINER):$(DATE) .
	docker image prune --filter label=stage=featureservbuilder -f

set-local:
	$(eval BUILDTYPE = local)
	$(eval IMAGE_TAG = $(CONTAINER):$(APPVERSION))

set-inherited:
	$(eval BUILDTYPE = inherited)
	$(eval IMAGE_TAG = $(CONTAINER):$(APPVERSION)-$(TARGETARCH))

local-docker: bin-for-docker Dockerfile set-local docker-build ##       Generate a BASE_IMAGE container with APPVERSION tag, using a locally built binary

docker: Dockerfile set-inherited docker-build ##             Generate a BASE_IMAGE container with APPVERSION tag, using a binary built in an alpine golang build container

release: clean docs $(PROGRAM) local-docker  ##            Generate the docs, a local build, and then uses the local build to generate BASE_IMAGE container

test:  ##               Run the tests locally
	go test -v ./internal/cql ./internal/service

install: $(PROGRAM) docs  ##            This will install the program locally
	$(MKDIR) -p $(DESTDIR)/usr/bin
	$(MKDIR) -p $(DESTDIR)/usr/share/$(PROGRAM)
	$(MKDIR) -p $(DESTDIR)/etc
	$(CP) $(PROGRAM) $(DESTDIR)/usr/bin/$(PROGRAM)
	$(SED) 's,./assets,/user/share/$(PROGRAM)/assets,' config/$(PROGRAM).toml.example > $(DESTDIR)/etc/$(PROGRAM).toml
	$(CP) -r assets $(DESTDIR)/usr/share/$(PROGRAM)/assets
	$(CP) -r docs $(DESTDIR)/usr/share/$(PROGRAM)/docs

uninstall:  ##          This will uninstall the program from your local system
	$(RM) $(DESTDIR)/usr/bin/$(PROGRAM)
	$(RM) $(DESTDIR)/etc/$(PROGRAM).toml
	$(RM) -r $(DESTDIR)/usr/share/$(PROGRAM)

help:   ##               Prints this help message
	@echo ""
	@echo ""
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | fgrep -v : | sed -e 's/\\$$//' | sed -e 's/.*##//'
	@echo ""
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | fgrep : | sed -e 's/\\$$//' | sed -e 's/:.*##/:/'
	@echo ""
	@echo ""
