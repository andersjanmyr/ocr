#!/bin/bash

set -o errexit

echoerr() { echo "$@" 1>&2; }

if [ $# -ne 1 ]; then
  echoerr "Usage: tag.sh <version>"
  exit 1
fi

new_version=$1

if ! grep $new_version ./RELEASE_NOTES.md; then
  echoerr "RELEASE_NOTES does not contain a section for $new_version"
  exit 1
fi

if ! git diff --quiet HEAD; then
  echoerr "Cannot create release with a dirty repo."
  echoerr "Commit or stash changes and try again."
  git status -sb
  exit 1
fi

sed -i.bak -E "s/v[0-9]\.[0-9]\.[0-9]/$new_version/g" README.md
git add README.md
git commit -am "Changed version to $new_version"
git tag $new_version -am "Release $new_version"

rm README.md.bak

git push --tags

