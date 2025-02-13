#!/bin/bash
set -e

microservice_package="userapiserver"

pushd swagger
docker build -t swagger:2.4.37 .
popd

echo "Generating models"
mkdir -p models
docker run --rm -v ${PWD}:/local --user $(id -u):$(id -g) swagger:2.4.37 generate -i /local/api/restAPI.yaml -l go-server -o /local -t /templates --additional-properties hideGenerationTimestamp=true -Dservice -Dmodels -DpackageName=models
mv go/* models
rm -rf go

docker rmi swagger:2.4.37