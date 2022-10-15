#!/usr/bin/env bash

NOW=$(date +"%F %T")

unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     PLATFORM=linux;;
    Darwin*)    PLATFORM=darwin;;
    *)          PLATFORM="UNKNOWN:${unameOut}"
esac

printf "%s >> Installing golang-migrate\n" "$NOW"
curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.$PLATFORM-amd64.tar.gz | tar xvz

printf "%s >> Installing gen\n" "$NOW"
go install github.com/smallnest/gen@latest

printf "%s >> Installing golang-migrate\n" "$NOW"

printf "%s Please install these manually: docker docker-compose\n" "$NOW"

