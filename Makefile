
APPVERSION := CI
GOVERSION := 1.13
PROGRAM := pg_featureserv
CONTAINER := crunchydata/$(PROGRAM)

RM = /bin/rm
CP = /bin/cp
MKDIR = /bin/mkdir

.PHONY: all build-local check clean docker-build docs install release test uninstall

GOFILES := $(shell find . -type f -name '*.go')

all: docs release

check:
	go version

clean:
	$(info Cleaning project...)
	@rm -f $(PROGRAM)
	@rm -rf docs/*

docs:
	@rm -rf docs/* && cd hugo && hugo && cd ..

build-local: $(GOFILES)
	go build -v

docker-build: Dockerfile clean
	sudo docker run --rm -v "$(PWD)":/usr/src/myapp:z -w /usr/src/myapp golang:$(GOVERSION) make APPVERSION=$(APPVERSION) build-local
	docker build -f Dockerfile --build-arg VERSION=$(APPVERSION) -t $(CONTAINER):$(APPVERSION) .

release: docker-build clean build-local

test:
	go test -v

install: $(PROGRAM) docs
	$(MKDIR) -p $(DESTDIR)/usr/bin
	$(MKDIR) -p $(DESTDIR)/usr/share/$(PROGRAM)
	$(MKDIR) -p $(DESTDIR)/etc
	$(CP) $(PROGRAM) $(DESTDIR)/usr/bin/$(PROGRAM)
	$(CP) config/$(PROGRAM).toml.example $(DESTDIR)/etc/$(PROGRAM).toml
	$(CP) -r assets $(DESTDIR)/usr/share/$(PROGRAM)/assets
	$(CP) -r docs $(DESTDIR)/usr/share/$(PROGRAM)/docs

uninstall:
	$(RM) $(DESTDIR)/usr/bin/$(PROGRAM)
	$(RM) $(DESTDIR)/etc/$(PROGRAM).toml
	$(RM) -r $(DESTDIR)/usr/share/$(PROGRAM)
