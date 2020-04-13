
PROGRAM := pg_featureserv
CONTAINER := crunchydata/$(PROGRAM)
APPVERSION=CI

RM = /bin/rm
CP = /bin/cp
MKDIR = /bin/mkdir

.PHONY: all check clean docs test docker-ci release install uninstall

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

$(PROGRAM): $(GOFILES)
	go build -v

release: $(PROGRAM) Dockerfile
	docker build -f Dockerfile --build-arg VERSION=$(APPVERSION) -t $(CONTAINER):$(APPVERSION) .
	docker image prune --force

docker-ci: release
	docker tag $(CONTAINER):"$(APPVERSION)" $(CONTAINER):latest

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
