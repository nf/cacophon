#!/bin/bash
APP_YAML=$1
DEPS=$(sed -n -e 's/\  apt_get_install: \(.*\)/\1/p' ${APP_YAML})
DEBIAN_FRONTEND=noninteractive apt-get install -y --force-yes -q --no-install-recommends ${DEPS}
