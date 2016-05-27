#!/usr/bin/env bash

# try to adjust base dependence
if [ "$APPROOT" == "" ];
then
    parent=$(cd "$(pwd)/../"; pwd)/app
    if [ -d $parent ];
    then
        export APPROOT=$(cd "$parent"; pwd)

        source "$APPROOT/env.sh"
    fi

    grand=$(cd "$(pwd)/../"; pwd)/go
    if [ -d $grand ];
    then
        export APPROOT=$(cd "$grand"; pwd)

        source "$APPROOT/env.sh"
    fi
fi

# adjust GOPATH
case ":$GOPATH:" in
    *":$(pwd):"*) :;;
    *) GOPATH=$(pwd):$GOPATH;;
esac
export GOPATH


# adjust PATH
if [ -n "$ZSH_VERSION" ]; then
    readopts="Ra"
else
    readopts="ra"
fi
while IFS=':' read -$readopts ARR; do
    for i in "${ARR[@]}"; do
        case ":$PATH:" in
            *":$i/bin:"*) :;;
            *) PATH=$i/bin:$PATH
        esac
    done
done <<< "$GOPATH"
export PATH
