#!/bin/bash

PROJECT=pg_featureserv
DOCKER_REPO=pramsey/$PROJECT

if [ "$TARGET" = "windows" ]; then
    BINARY=$PROJECT.exe
else
    BINARY=$PROJECT
fi

if [ "$TRAVIS_TAG" = "" ]; then
    TAG=latest
else
    TAG=$TRAVIS_TAG
fi

if [ "$TARGET" = "docker" ]; then
    VERSION=`./$PROJECT --version | cut -f2 -d' '`
    DATE=`date +%Y%m%d`
    docker build -f Dockerfile.ci --build-arg VERSION=$VERSION -t $DOCKER_REPO:$TAG .
    docker tag $DOCKER_REPO:$TAG $DOCKER_REPO:$DATE
    if [ "$TRAVIS_BRANCH" = "master" ]; then
        docker login -u "$DOCKER_USER" -p "$DOCKER_PASS"
        docker push $DOCKER_REPO
    fi
else
    mkdir upload
    zip -r upload/${PROJECT}_${TAG}_${TARGET}.zip ${BINARY} README.md LICENSE.md config/ assets/
fi

