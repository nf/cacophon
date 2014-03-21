#!/bin/bash

set -ex

NODE_VERSION=v0.10.26
NODE_PLATFORM=linux-x64
NODE_DIST=node-${NODE_VERSION}-${NODE_PLATFORM}

apt-get update && apt-get install --no-install-recommends -y -q --force-yes curl git build-essential
curl http://nodejs.org/dist/${NODE_VERSION}/${NODE_DIST}.tar.gz | tar xvzf -
ln -s /${NODE_DIST}/bin/node /usr/bin/node
ln -s /${NODE_DIST}/bin/npm /usr/bin/npm

/node-v0.10.26-linux-x64/bin/npm -g install bower
ln -s /${NODE_DIST}/bin/bower /usr/bin/bower
