#
# Copyright SecureKey Technologies Inc. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0
#

# Supported Targets:
# all : runs unit and integration tests
# depend: installs test dependencies
# unit-test: runs all the unit tests
# integration-test: runs all the integration tests
# checks: runs all check conditions (license, spelling, linting)
# clean: stops docker conatainers used for integration testing
# mock-gen: generate mocks needed for testing (using mockgen)
#

export ARCH=$(shell uname -m)
export LDFLAGS=-ldflags=-s

all: checks unit-test integration-test

depend:
	@test/scripts/dependencies.sh

checks: depend license lint spelling

.PHONY: license
license:
	@test/scripts/check_license.sh

lint:
	@test/scripts/check_lint.sh

spelling:
	@test/scripts/check_spelling.sh

unit-test: checks
	@test/scripts/unit.sh

unit-tests: unit-test

integration-test: clean depend
	@test/scripts/integration.sh

integration-tests: integration-test

mock-gen:
	go get -u github.com/golang/mock/gomock
	go get -u github.com/golang/mock/mockgen
	mockgen -source=pkg/fabric-client/peer/peer.go -destination=pkg/fabric-client/peer/mocks/mockpeer.gen.go
	mockgen -build_flags '$(LDFLAGS)' github.com/tuxago/fabric-sdk-go/api Config | sed "s/github.com\/hyperledger\/fabric-sdk-go\/vendor\///g"  > api/mocks/mockconfig.gen.go

clean:
	rm -Rf /tmp/enroll_user /tmp/msp /tmp/keyvaluestore
	rm -f integration-report.xml report.xml
	cd test/fixtures && docker-compose down
