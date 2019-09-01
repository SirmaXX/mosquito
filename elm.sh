#!/bin/sh
# Author : Deniz Balcı
# Project: Elm Project Creator
# Date : 02/09/2019
echo "proje ismi nedir ?"
read projectname
mkdir $projectname
cd  $projectname
elm init
elm install elm/http
elm install elm/json
elm install mdgriffith/elm-ui
cd src 
touch Main.elm
touch View.elm
touch State.elm
touch Rest.elm
touch Types.elm
touch Colors.elm
mkdir Page
