#!/bin/bash -eu

cat $1 | pup '.list-proposal > .type > .name text{}'
