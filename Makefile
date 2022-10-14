
APPVERSION := latest
GOVERSION := 1.15
PROGRAM := pg_featureserv
CONTAINER := pramsey/$(PROGRAM)
DATE := $(shell date +%Y%m%d)

RM = /bin/rm
CP = /bin/cp
MKDIR = /bin/mkdir
SED = /usr/bin/sed

.PHONY: build bin-for-docker check clean build-in-docker docs install release test uninstall

.DEFAULT_GOAL := help

GOFILES := $(shell find . -type f -name '*.go')

check:  ##         This checks the current version of Go installed locally
	go version

clean:  ##         This will clean all local build artifacts
	$(info Cleaning project...)
	@rm -f $(PROGRAM)
	@rm -rf docs/*
	@docker image inspect $(CONTAINER):$(APPVERSION) >/dev/null 2>&1 && docker rmi $(CONTAINER):$(APPVERSION) $(CONTAINER):$(DATE) || echo -n ""

docs:   ##          Generate docs
	@rm -rf docs/* && cd hugo && hugo && cd ..

build: $(PROGRAM)  ##         just an alias

$(PROGRAM): $(GOFILES)  ##         Build a local binary using APPVERSION parameter or CI as default
	go build -v -ldflags "-s -w -X github.com/CrunchyData/pg_featureserv/conf.setVersion=$(APPVERSION)"

bin-for-docker: $(GOFILES)  ##         Build a local binary using APPVERSION parameter or CI as default (to be used in docker image)
# to be used in docker the built binary needs the CGO_ENABLED=0 option
	CGO_ENABLED=0 go build -v -ldflags "-s -w -X github.com/CrunchyData/pg_featureserv/conf.setVersion=$(APPVERSION)"

build-in-docker: $(GOFILES)   ##    Build a local binary based of a golang base docker image without the need of a local go environment
	docker run --rm -v "$(PWD)":/usr/src/myapp:z -w /usr/src/myapp golang:$(GOVERSION) make APPVERSION=$(APPVERSION) $(PROGRAM)

docker: bin-for-docker Dockerfile  ##  Generate a CentOS 7 container with APPVERSION tag, with the dedicated binary installed into
	docker build -f Dockerfile --build-arg VERSION=$(APPVERSION) -t $(CONTAINER):$(APPVERSION) -t $(CONTAINER):$(DATE) .

release: clean docs $(PROGRAM) docker  ##       Generate the docs, a local build, and then uses the local build to generate a CentOS 7 container

test:  ##          Run the tests locally
	go test -v $(shell go list ./... | grep -vw db_test)

install: $(PROGRAM) docs  ##       This will install the program locally
	$(MKDIR) -p $(DESTDIR)/usr/bin
	$(MKDIR) -p $(DESTDIR)/usr/share/$(PROGRAM)
	$(MKDIR) -p $(DESTDIR)/etc
	$(CP) $(PROGRAM) $(DESTDIR)/usr/bin/$(PROGRAM)
	$(SED) 's,./assets,/user/share/$(PROGRAM)/assets,' config/$(PROGRAM).toml.example > $(DESTDIR)/etc/$(PROGRAM).toml
	$(CP) -r assets $(DESTDIR)/usr/share/$(PROGRAM)/assets
	$(CP) -r docs $(DESTDIR)/usr/share/$(PROGRAM)/docs

uninstall:  ##     This will uninstall the program from your local system
	$(RM) $(DESTDIR)/usr/bin/$(PROGRAM)
	$(RM) $(DESTDIR)/etc/$(PROGRAM).toml
	$(RM) -r $(DESTDIR)/usr/share/$(PROGRAM)

help:   ##          Prints this help message
	@echo ""
	@echo ""
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/:.*##/:/'
	@echo ""
	@echo ""
