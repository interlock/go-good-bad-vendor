#!/bin/bash

install -d ./vendor/github.com/interlock/some_lib/vendor/github.com/interlock/bad_lib
cp -R ./vendor/github.com/interlock/bad_lib ./vendor/github.com/interlock/some_lib/vendor/github.com/interlock/

install -d ./vendor/github.com/interlock/some_lib/vendor/github.com/interlock/good_lib
cp -R ./vendor/github.com/interlock/good_lib ./vendor/github.com/interlock/some_lib/vendor/github.com/interlock/
