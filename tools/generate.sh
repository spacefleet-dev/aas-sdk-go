#!/bin/sh

docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v6.2.0 generate \
    -i /local/api/merged.yaml \
    -g go \
    -o /local/api \
    --additional-properties=packageName=api,enumClassPrefix=true,useOneOfDiscriminatorLookup=true,isGoSubmodule=true
