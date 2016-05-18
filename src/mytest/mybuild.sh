#!/usr/bin/env bash

if [ ! -f install ]; then
    echo 'install must be run within its container folder' 1>&2
    exit 1
fi

CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

rm ./bin/mytest
#rm ./output/bin/accesslayer
gofmt -w src

/home/users/lanyangyang01/go/go/bin/go install mytest
#mkdir -p ./output/bin
#cp bin/accesslayer output/bin
#cp -r conf output/
#cp -r test/noahdes output/



export GOPATH="$OLDGOPATH"

echo 'finished'
