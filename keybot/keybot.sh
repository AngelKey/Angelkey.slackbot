#!/bin/bash

set -e -u -o pipefail # Fail on error

dir=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )
cd $dir

git pull --ff-only
GO15VENDOREXPERIMENT=1 go install github.com/keybase/slackbot/keybot
"$GOPATH/bin/keybot"
