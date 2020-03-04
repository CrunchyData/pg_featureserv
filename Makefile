
PROGRAM := pg_featureserv
CONTAINER := pramsey/$(PROGRAM)

.PHONY: all check clean test docker install uninstall

GOFILES := $(shell find . -type f -name '*.go')

all: $(PROGRAM)

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

docker: $(PROGRAM) Dockerfile
	docker build -f Dockerfile.ci --build-arg VERSION=`./$(PROGRAM) --version | cut -f2 -d' '` -t $(CONTAINER):latest .
	docker image prune --force

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
