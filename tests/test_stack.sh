#!/usr/bin/env bash

# ---------------------------------------------------------------------------
# PURPOSE

# This script will test the eris stack and the connection between eris cli
# and eris pm. **Generally, it should not be used in isolation.**

# ---------------------------------------------------------------------------
# REQUIREMENTS

# Docker installed locally
# Docker-Machine installed locally (if using remote boxes)
# eris' test_machines image (if testing against eris' test boxes)
# Eris installed locally

# ---------------------------------------------------------------------------
# USAGE

# test_stack.sh

# ----------------------------------------------------------------------------
# Set definitions and defaults

# Where are the Things
start=`pwd`
base=github.com/monax/cli
repo=$GOPATH/src/$base
fail="false"
if [ "$TRAVIS_BRANCH" ]
then
  ci=true
  osx=true
elif [ "$APPVEYOR_REPO_BRANCH" ]
then
  ci=true
  win=true
fi

export ERIS_PULL_APPROVE="true"
export ERIS_MIGRATE_APPROVE="true"
export SKIP_BUILD="true"

# ----------------------------------------------------------------------------
# Utility functions

check_and_exit() {
  if [ $fail -ne "false" ]
  then
    cd $start
    exit $test_exit
  fi
}

# ----------------------------------------------------------------------------
# Run [eris packages do] tests

time tests/test_jobs.sh
if [ $? -ne 0 ]; then fail="true"; fi
cd $start

# ----------------------------------------------------------------------------
# Run [eris chains make] tests

time tests/test_chains_make.sh
if [ $? -ne 0 ]; then fail="true"; fi
cd $start

# ----------------------------------------------------------------------------
# Cleanup
cd $start
check_and_exit
exit $test_exit