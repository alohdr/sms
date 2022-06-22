#!/bin/bash

# Put your modules in here
modules=(
  './internal'
)


for idx in "${!modules[@]}"
do
  {
    module=${modules[$idx]}
    echo 'Sqlc generate for '.$module
    (cd ${module}/repository && sqlc generate)

  } || true
done