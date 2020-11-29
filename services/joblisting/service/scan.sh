#!/bin/sh

while getopts u:n: flag
do
    case $flag in
        u) url=$OPTARG;;
        n) name=$OPTARG;;
    esac
done

git clone $url $name;
cd $name && echo "sonar.joblistingKey=$name" > sonar-joblisting.properties && sonar-scanner && cd .. && rm -rf $name
