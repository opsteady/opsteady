#!/bin/bash

# Find the version of the component in the docker folder and update it.
# Uses SetDockerBuildInfo function in go to find the version

#git diff -s --exit-code main -- docker/base/Dockerfile
SOURCE_FILE="empty"

git diff -s --exit-code main -- docker/base/Dockerfile
if [ $? -eq 1 ]; then
  SOURCE_FILE="docker/base/cicd/component.go"
fi

git diff -s --exit-code main -- docker/cicd/Dockerfile
if [ $? -eq 1 ]; then
  SOURCE_FILE="docker/cicd/cicd/component.go"
fi

if [ "$SOURCE_FILE" == "empty" ]; then
  echo "No changes detected stopping"
  exit 0
fi

echo "Finding version in $SOURCE_FILE"
ORIGINAL_VERSION=$(cat $SOURCE_FILE | grep SetDockerBuildInfo | awk -F"," '{print substr($2, 2)}' | awk '{print substr($0, 2, length($0) - 2)}')

echo "Original version $ORIGINAL_VERSION"

# Split version into numbers
IFS='.' read -ra ver <<< "$ORIGINAL_VERSION"
#Take the minor and update
minor=${ver[1]}
minor=$((minor+1))

NEW_VERSION="${ver[0]}.$minor.${ver[2]}"

echo "New version $NEW_VERSION"
sed -i "s/$ORIGINAL_VERSION/$NEW_VERSION/" $SOURCE_FILE
