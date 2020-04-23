FROM centos:7

ARG VERSION

LABEL vendor="Crunchy Data" \
	url="https://crunchydata.com" \
	release="${VERSION}" \
	org.opencontainers.image.vendor="Crunchy Data" \
	os.version="7.7"

ADD ./pg_featureserv .
ADD ./assets ./assets

VOLUME ["/config"]
VOLUME ["/assets"]

USER 1001
EXPOSE 9000

ENTRYPOINT ["./pg_featureserv"]
CMD []

# To build
# make release

# To run
# docker run -dt -e DATABASE_URL=postgres://user:pass@host/dbname -p 9000:9000 crunchydata/pg_featureserv
