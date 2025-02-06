#!/usr/bin/env bash

set -eu

cd "$(dirname "$0")/.."

help=$(ghproj help-all)

echo "# Usage

<!-- This is generated by scripts/generate-usage.sh. Don't edit this file directly. -->

$help" > USAGE.md
