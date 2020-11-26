#!/bin/sh

while getopts u:n: flag
do
    case $flag in
        u) url=$OPTARG;;
        n) name=$OPTARG;;
    esac
done

git clone $url $name;
cd $name && echo "sonar.projectKey=$name" > sonar-project.properties && sonar-scanner && rm -rf $name
