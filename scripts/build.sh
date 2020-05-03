#!/bin/sh
srcPath="."
pkgFile="main.go"
outputPath="build"
app="go-server"
output="$outputPath/$app"
src="$srcPath/$pkgFile"

printf "\nBuilding: $app\n"
time go build -o $output $src
printf "\nBuilt: $app size:"
ls -lah $output | awk '{print $5}'
printf "\nDone building: $app\n\n"