
CURDIR=`pwd`
OLDGOPATH="$GOPATH"
export GOPATH="$CURDIR"

rm ./bin/mytest
gofmt -w src

go install mytest

export GOPATH="$OLDGOPATH"

echo 'finished'
