/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package config

import (
	"fmt"
	"os"
	"testing"

	api "github.com/tuxago/fabric-sdk-go/api"
	"github.com/spf13/viper"
)

var configImpl api.Config

func TestCAConfig(t *testing.T) {
	caname := configImpl.GetFabricCAName()
	if caname != "ca-org1" {
		t.Fatalf("caname doesn't match")
	}
}

func TestGetPeersConfig(t *testing.T) {
	pc, err := configImpl.GetPeersConfig()
	if err != nil {
		t.Fatalf(err.Error())
	}

	for _, value := range pc {
		if value.Host == "" {
			t.Fatalf("Host value is empty")
		}
		if value.Port == 0 {
			t.Fatalf("Port value is empty")
		}
		if value.EventHost == "" {
			t.Fatalf("EventHost value is empty")
		}
		if value.EventPort == 0 {
			t.Fatalf("EventPort value is empty")
		}

	}

}

// Test case to create a new viper instance to prevent conflict with existing
// viper instances in applications that use the SDK
func TestMultipleVipers(t *testing.T) {
	viper.SetConfigFile("./test.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err.Error())
	}
	testValue1 := viper.GetString("test.testKey")
	// Read initial value from test.yaml
	if testValue1 != "testvalue" {
		t.Fatalf("Expected testValue before config initialization got: %s", testValue1)
	}
	// initialize go sdk
	_, err = InitConfig("../../test/fixtures/config/config_test.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	// Make sure initial value is unaffected
	testValue2 := viper.GetString("test.testKey")
	if testValue2 != "testvalue" {
		t.Fatalf("Expected testvalue after config initialization")
	}
	// Make sure Go SDK config is unaffected
	testValue3 := myViper.GetBool("client.security.enabled")
	if testValue3 != true {
		t.Fatalf("Expected existing config value to remain unchanged")
	}
}

func TestEnvironmentVariablesDefaultCmdRoot(t *testing.T) {
	testValue := myViper.GetString("env.test")
	if testValue != "" {
		t.Fatalf("Expected environment variable value to be empty but got: %s", testValue)
	}

	err := os.Setenv("FABRIC_SDK_ENV_TEST", "123")
	defer os.Unsetenv("FABRIC_SDK_ENV_TEST")

	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = InitConfig("../../test/fixtures/config/config_test.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}

	testValue = myViper.GetString("env.test")
	if testValue != "123" {
		t.Fatalf("Expected environment variable value but got: %s", testValue)
	}
}

func TestEnvironmentVariablesSpecificCmdRoot(t *testing.T) {
	testValue := myViper.GetString("env.test")
	if testValue != "" {
		t.Fatalf("Expected environment variable value to be empty but got: %s", testValue)
	}

	err := os.Setenv("TEST_ROOT_ENV_TEST", "456")
	defer os.Unsetenv("TEST_ROOT_ENV_TEST")

	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = InitConfigWithCmdRoot("../../test/fixtures/config/config_test.yaml", "test_root")
	if err != nil {
		fmt.Println(err.Error())
	}

	testValue = myViper.GetString("env.test")
	if testValue != "456" {
		t.Fatalf("Expected environment variable value but got: %s", testValue)
	}
}

func TestMain(m *testing.M) {
	var err error
	configImpl, err = InitConfig("../../test/fixtures/config/config_test.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	os.Exit(m.Run())
}
