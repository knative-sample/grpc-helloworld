#!/bin/bash
#****************************************************************#
# Create Date: 2020-02-04 10:25
#********************************* ******************************#
set -e 

BASEDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

function parse_git_dirty {
  local STATUS="$(git status 2> /dev/null)"
  local info=""
  if [[ $? -ne 0 ]]; then printf ${info}; return; fi
  if echo ${STATUS} | grep -c "renamed:"         &> /dev/null; then info="${info}>"; fi
  if echo ${STATUS} | grep -c "branch is ahead:" &> /dev/null; then info="${info}!"; fi
  if echo ${STATUS} | grep -c "new file::"       &> /dev/null; then info="${info}+"; fi
  if echo ${STATUS} | grep -c "modified:"        &> /dev/null; then info="${info}*"; fi
  if echo ${STATUS} | grep -c "deleted:"         &> /dev/null; then info="${info}-"; fi
  if [ ! -z "${info}" ]; then
      printf "dirty"
  fi
}

COMPONENT="${1}"
RELEASE="${2}"
NS="knative-sample"
IMAGE_NAME="helloworld-${1}"
TAG=$(date +"%Y%m%d-%H%M%S")

if [[ "$(parse_git_dirty)" = "dirty" && "$RELEASE" = "true" ]]; then
    echo "git status is dirty, can not release"
    exit 1
fi

if [ "$RELEASE" = "true" ]; then
    NS="knative-release"
    br=$(git rev-parse --abbrev-ref HEAD 2> /dev/null)
    commitid=$(git rev-parse --short HEAD)
    TAG="${TAG}_${br}_${commitid}"
fi

docker build -t ${IMAGE_NAME}:latest -f ${BASEDIR}/${COMPONENT}-dockerfile  .
docker tag  ${IMAGE_NAME}:latest registry.cn-hangzhou.aliyuncs.com/${NS}/${IMAGE_NAME}:${TAG}
echo "docker push registry.cn-hangzhou.aliyuncs.com/${NS}/${IMAGE_NAME}:${TAG}"
docker push registry.cn-hangzhou.aliyuncs.com/${NS}/${IMAGE_NAME}:${TAG}

