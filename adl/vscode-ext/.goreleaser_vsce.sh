#!/bin/bash

rm *.vsix 2> /dev/null || true

TAG=`git tag -l --points-at HEAD`
if [ "$TAG" != "" ]; then
    T2=$(echo $TAG | sed 's/v\(.*\)/"\1"/')
    sed -i 's!"version":.*,!"version": '"$T2"',!' package.json
    sed -i 's!"version":.*,!"version": '"$T2"',!' client/package.json
fi

vsce package