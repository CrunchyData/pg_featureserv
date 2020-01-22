FROM centos:7

ARG VERSION=0.1

LABEL vendor="Crunchy Data" \
	url="https://crunchydata.com" \
	release="${VERSION}" \
	org.opencontainers.image.vendor="Crunchy Data" \
	os.version="7.7"

ADD ./pg_featureserv .
ADD ./assets ./assets

VOLUME ["/config"]
VOLUME [ "/assets"]

USER 1001
EXPOSE 9000

ENTRYPOINT ["./pg_featureserv"]
CMD []
