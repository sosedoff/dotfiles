#!/bin/bash

if [ $# -ne 3 ]; then
  echo "Usage: git merge-repo <remote> <new branch> <new dir>"
  exit 1
fi

git remote add new-repo $1
git fetch new-repo
git branch $2 new-repo/master
git checkout $2

mkdir -vp $3

git ls-tree -z --name-only HEAD | xargs -0 -I {} git mv {} "${GIT_PREFIX}$3"
git commit -m "Moved files to '${GIT_PREFIX}$3'"
git checkout master
git merge --allow-unrelated-histories --no-edit -s recursive -X no-renames $2
git branch -D $2
git remote remove new-repo