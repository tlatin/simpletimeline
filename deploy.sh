#!/usr/bin/env bash

BRANCH=`git rev-parse --abbrev-ref HEAD`

while getopts ":f" opt; do
  case $opt in
    f)
      f=1
      ;;
    *)
      echo "usage: $0 [-f]" >&2
      exit 1
      ;;
  esac
done

# "Colorizing" Scripts: http://tldp.org/LDP/abs/html/colorizing.html
GREEN='\x1B[32;40m' 
RED='\x1B[31;40m'
alias Reset="tput sgr0" # Reset text attributes to normal

deploy () {
  local version=$1
  local color=${2:-$GREEN}

  echo -n "Deploying version "
  echo -e $color$version; Reset
  goapp deploy -oauth --version=$version app
}

# predeploy () {
#   if [ $f ]; then
#     echo "Skipping predeploy."
#     return
#   fi

#   local version=$1
#   local color=${2:-$GREEN}
#   echo -n "Pre-deploy checks for "
#   echo -e $color$version; Reset

#   if ! ./run_tests.py ; then
#     echo "Tests failed. Abandoning deploy."
#     exit
#   else
#     echo "Tests passed."
#   fi
# }

if [ $BRANCH == 'master' ]; then
  date_commit=`git log -1 --pretty=format:%cd-%h --date=short`
  previous_release=`git describe --abbrev=0 --tags`
  release="rel-$date_commit"

  # predeploy $release $RED

  # add a lightweight tag and push it to origin
  git tag $release
  git push origin $release 
  
  # deploy the tagged version
  deploy $release $RED

  # generate the URL for the Github tag comparison
  url="https://github.com/tlatin/simpletimeline/compare/$previous_release...$release"
  echo "Deploy comparison: $url"
else
  version=$BRANCH  
  # deploy a version named as the branch
  predeploy $version
  deploy $version
fi
