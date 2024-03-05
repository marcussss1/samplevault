#!/bin/bash

# Заменяем значение image в deployment.yml на сгенерированную версию образа
sed -i "s|image: marcussss/simplevault:.*|image: marcussss/simplevault:$(echo $GITHUB_SHA | cut -c -7)|" deployment.yml
