#!/bin/sh
# Author : Deniz Balcı
# Project Automation
mkdir logs
mkdir outputs
cd outputs
echo "proje ismi nedir ?"
read projectname
mkdir $projectname
cd  $projectname
echo -e "Bugün \c ";date
iwconfig
touch index.html
mkdir static
mkdir tests
mkdir logs
mkdir database
cd static 
mkdir css
cd css
touch style.css
cd ..
mkdir js
cd js
touch script.js
cd ..
mkdir img

