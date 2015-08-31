#!/bin/bash

if [ -e "./vendor/github.com/interlock/some_lib/vendor" ]; then
  rm -rf "./vendor/github.com/interlock/some_lib/vendor"
fi