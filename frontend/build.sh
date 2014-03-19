#!/bin/bash

set -ex

cd /app
find .
npm install
bower install --allow-root --config.interactive=false
