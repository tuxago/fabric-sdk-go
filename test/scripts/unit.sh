#!/bin/bash
#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

set -e

# Packages to exclude
PKGS=`go list github.com/tuxago/fabric-sdk-go/... 2> /dev/null | \
                                                  grep -v /vendor/ | \
                                                  grep -v /test/`
echo "Running tests..."
gocov test $LDFLAGS $PKGS -p 1 -timeout=5m | gocov-xml > report.xml
