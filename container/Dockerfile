FROM centos:7

ARG VERSION

LABEL vendor="Crunchy Data" \
	url="https://crunchydata.com" \
	release="${VERSION}" \
	org.opencontainers.image.vendor="Crunchy Data" \
	os.version="7.7"

ADD ./build/crunchy-cloud-api /usr/local/bin
ADD ./build/plugin-crunchy-cloud.so /opt/crunchy-cloud/

VOLUME ["/config"]

USER daemon

ENTRYPOINT ["/usr/local/bin/crunchy-cloud-api"]
