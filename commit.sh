#!/usr/bin/env bash

set -eu

author=${1:-Louis Lin <lin.louis@pm.me>}
project_path=${2:-/home/llin/perso/daily-logs}

pushd "${project_path}"
/usr/bin/git add -A
/usr/bin/git commit --author "${author}" -m "$(date '+%m'): add $(date '+%d')"
popd
