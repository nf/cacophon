#!/bin/bash

set -ex

cd /app
npm install
bower install --allow-root --config.interactive=false
