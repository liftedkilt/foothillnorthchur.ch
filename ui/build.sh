#!/bin/bash

script_dir=$(dirname $0)
cd $script_dir

set -o verbose

mkdir -p dist/js dist/css

# Compiling CSS
sass --style compressed src/css/main.scss:dist/css/main.min.css

# Copying external CSS
sass --style compressed src/ext/css/normalize.css:dist/css/normalize.min.css

# Copying JS
cp src/js/main.js dist/js/main.js

# Copying external JS
cp src/ext/js/*.js dist/js/

# Copying HTML
cp src/index.html dist/index.html

# TEMP: link caddyfile
cd dist
ln -s ../src/caddyfile ./Caddyfile