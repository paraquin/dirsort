#!/usr/bin/bash

package_name='dirsort'
platforms=("windows/amd64" "windows/386" "linux/amd64" "linux/386" "linux/arm64" "darwin/amd64" "darwin/arm64")
save_dir="release/"

if [[ ! -d "$save_dir" ]]; then
    mkdir $save_dir
fi

for platform in "${platforms[@]}"
do
    echo "build for $platform"
    IFS='/' read -ra platform_split <<< "$platform"
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi
    GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name .
    if [ $? -ne 0 ]; then
        echo 'An error has occured! Aborting the script execution...'
        exit 1
    fi
    archieve_name=$package_name'_'$GOOS'_'$GOARCH'.zip'
    zip -r -1 $save_dir$archieve_name $output_name mapping.yaml
    rm $output_name
done