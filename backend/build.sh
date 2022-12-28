#!/bin/bash

appName="api-clean-twitter"

echo "👾 Start building script api ..."

go clean -cache ./...
go build -o $appName ./...

if [ ! -d "dist" ] 
then
    echo "🔨 Creation of dist folder"
    mkdir dist
fi

mv $appName ./dist

echo "🧰 Finish building script api"
