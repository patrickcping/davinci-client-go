package davinci

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

type envs struct {
	DAVINCI_USERNAME       string `json:"DAVINCI_USERNAME"`
	DAVINCI_PASSWORD       string `json:"DAVINCI_PASSWORD"`
	DAVINCI_COMPANY_ID     string `json:"DAVINCI_COMPANY_ID"`
	DAVINCI_BASE_URL       string `json:"DAVINCI_BASE_URL"`
	PING_ONE_ADMIN_ENV_ID  string `json:"PING_ONE_ADMIN_ENV_ID"`
	PING_ONE_TARGET_ENV_ID string `json:"PING_ONE_TARGET_ENV_ID"`
}

func TestNewClient_GA(t *testing.T) {
	var host, username, password string
	jsonFile, err := os.Open("../local/env-ga.json")
	// jsonFile, err := os.Open("../local/env-ga.json")
	// if we os.Open returns an error then handle it
	var envs envs
	if err == nil {
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &envs)
		username = envs.DAVINCI_USERNAME
		password = envs.DAVINCI_PASSWORD
	} else {
		fmt.Println("File: ./local/env.json not found, \n trying env vars for DAVINCI_USERNAME/DAVINCI_PASSWORD")
		username = os.Getenv("DAVINCI_USERNAME")
		password = os.Getenv("DAVINCI_PASSWORD")
	}
	// defer the closing of our jsonFile so that we can parse it later on
	var tests = map[string]struct {
		host string
	}{
		"default":     {"https://api.singularkey.com/v1"},
		"nil":         {},
		"emptystring": {""},
		"testNeg":     {"https://badhost.io/v1"},
	}
	for name, hostStruct := range tests {
		testName := name
		t.Run(testName, func(t *testing.T) {
			cInput := ClientInput{
				HostURL:  hostStruct.host,
				Username: username,
				Password: password,
			}
			_, err := NewClient(&cInput)
			msg := fmt.Sprintf("\nGot client successfully, for test: %v\n", testName)
			if err != nil {
				fmt.Println(err.Error())
				msg = fmt.Sprint("Failed Successfully\n")
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(testName, "neg")) && !(strings.Contains(testName, "Neg")) {
					msg = fmt.Sprintf("failed to make client with host: %v \n Error is: %v", host, err)
					t.Fail()
				}
			}
			fmt.Printf(msg)
		})
	}
}
func TestNewClient_V2(t *testing.T) {
	var host, username, password string
	jsonFile, err := os.Open("../local/env-v2.json")
	// jsonFile, err := os.Open("../local/env-ga.json")
	// if we os.Open returns an error then handle it
	var envs envs
	if err == nil {
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &envs)
		username = envs.DAVINCI_USERNAME
		password = envs.DAVINCI_PASSWORD
		host = envs.DAVINCI_BASE_URL
	} else {
		fmt.Println("File: ./local/env.json not found, \n trying env vars for DAVINCI_USERNAME/DAVINCI_PASSWORD")
		username = os.Getenv("DAVINCI_USERNAME")
		password = os.Getenv("DAVINCI_PASSWORD")
		host = os.Getenv("DAVINCI_BASE_URL")
	}
	// defer the closing of our jsonFile so that we can parse it later on
	var tests = map[string]struct {
		host string
	}{
		"default":        {"https://orchestrate-api.pingone.com/v1"},
		"fromEnv":        {host},
		"emptystringNeg": {""},
		"testNeg":        {"https://badhost.io/v1"},
	}
	for name, hostStruct := range tests {
		testName := name
		t.Run(testName, func(t *testing.T) {
			cInput := ClientInput{
				HostURL:  hostStruct.host,
				Username: username,
				Password: password,
			}
			_, err := NewClient(&cInput)
			msg := fmt.Sprintf("\nGot client successfully, for test: %v\n", testName)
			if err != nil {
				fmt.Println(err.Error())
				msg = fmt.Sprint("Failed Successfully\n")
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(testName, "neg")) && !(strings.Contains(testName, "Neg")) {
					msg = fmt.Sprintf("failed to make client with host: %v \n Error is: %v", host, err)
					t.Fail()
				}
			}
			fmt.Printf(msg)
		})
	}
}
func TestNewClient_V2_SSO(t *testing.T) {
	var host, username, password, p1AdminEnv, p1TargetEnv, companyId string
	jsonFile, err := os.Open("../local/env-v2-sso.json")
	// jsonFile, err := os.Open("../local/env-v2-sso.json")
	// if we os.Open returns an error then handle it
	var envs envs
	if err == nil {
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &envs)
		username = envs.DAVINCI_USERNAME
		password = envs.DAVINCI_PASSWORD
		host = envs.DAVINCI_BASE_URL
		companyId = envs.DAVINCI_COMPANY_ID
		p1AdminEnv = envs.PING_ONE_ADMIN_ENV_ID
		p1TargetEnv = envs.PING_ONE_TARGET_ENV_ID
	} else {
		fmt.Println("File: ./local/env-v2-sso.json not found, \n trying env vars for DAVINCI_USERNAME/DAVINCI_PASSWORD")
		username = os.Getenv("DAVINCI_USERNAME")
		password = os.Getenv("DAVINCI_PASSWORD")
		host = os.Getenv("DAVINCI_BASE_URL")
		companyId = os.Getenv("DAVINCI_COMPANY_ID")
		p1AdminEnv = os.Getenv("PING_ONE_ADMIN_ENV_ID")
		p1TargetEnv = os.Getenv("PING_ONE_TARGET_ENV_ID")
	}
	// defer the closing of our jsonFile so that we can parse it later on
	var tests = map[string]ClientInput{
		"correct": {
			HostURL:  "https://orchestrate-api.pingone.com/v1",
			Username: username,
			Password: password,
			AuthP1SSO: AuthP1SSO{
				PingOneAdminEnvId:  p1AdminEnv,
				PingOneTargetEnvId: p1TargetEnv,
			},
		},
		"fromEnv": {
			HostURL:  host,
			Username: username,
			Password: password,
			AuthP1SSO: AuthP1SSO{
				PingOneAdminEnvId:  p1AdminEnv,
				PingOneTargetEnvId: p1TargetEnv,
			},
		},
		"emptyStringNeg": {
			HostURL:  "host",
			Username: username,
			Password: password,
			AuthP1SSO: AuthP1SSO{
				PingOneAdminEnvId:  p1AdminEnv,
				PingOneTargetEnvId: p1TargetEnv,
			},
		},
		"badhostNeg": {
			HostURL:  "https://badhost.io/v1",
			Username: username,
			Password: password,
			AuthP1SSO: AuthP1SSO{
				PingOneAdminEnvId:  p1AdminEnv,
				PingOneTargetEnvId: p1TargetEnv,
			},
		},
	}
	for name, inputStruct := range tests {
		testName := name
		t.Run(testName, func(t *testing.T) {
			client, err := NewClient(&inputStruct)
			if companyId != "" {
				client.CompanyID = companyId
			}
			msg := fmt.Sprintf("\nGot client successfully, for test: %v\n", testName)
			if err != nil {
				fmt.Println(err.Error())
				msg = fmt.Sprint("Failed Successfully\n")
				// if it's not a negative test, consider it an actual failure.
				if !(strings.Contains(testName, "neg")) && !(strings.Contains(testName, "Neg")) {
					msg = fmt.Sprintf("failed to make client with host: %v \n Error is: %v", host, err)
					t.Fail()
				}
			}
			fmt.Println(msg)
		})
	}
}

func newTestClient() (*Client, error) {
	var host, username, password, p1AdminEnv, p1TargetEnv, companyId string
	jsonFile, err := os.Open("../local/env-v2-sso.json")
	// jsonFile, err := os.Open("../local/env-v2-sso.json")
	// if we os.Open returns an error then handle it
	var envs envs
	if err == nil {
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		json.Unmarshal(byteValue, &envs)
		username = envs.DAVINCI_USERNAME
		password = envs.DAVINCI_PASSWORD
		host = envs.DAVINCI_BASE_URL
		companyId = envs.DAVINCI_COMPANY_ID
		p1AdminEnv = envs.PING_ONE_ADMIN_ENV_ID
		p1TargetEnv = envs.PING_ONE_TARGET_ENV_ID
	} else {
		fmt.Println("File: ./local/env-v2-sso.json not found, \n trying env vars for DAVINCI_USERNAME/DAVINCI_PASSWORD")
		username = os.Getenv("DAVINCI_USERNAME")
		password = os.Getenv("DAVINCI_PASSWORD")
		host = os.Getenv("DAVINCI_BASE_URL")
		companyId = os.Getenv("DAVINCI_COMPANY_ID")
		p1AdminEnv = os.Getenv("PING_ONE_ADMIN_ENV_ID")
		p1TargetEnv = os.Getenv("PING_ONE_TARGET_ENV_ID")
	}
	cInput := ClientInput{
		HostURL:  host,
		Username: username,
		Password: password,
		AuthP1SSO: AuthP1SSO{
			PingOneAdminEnvId:  p1AdminEnv,
			PingOneTargetEnvId: p1TargetEnv,
		},
	}
	client, err := NewClient(&cInput)
	if companyId != "" {
		client.CompanyID = companyId
	}
	if err != nil {
		log.Fatalf("failed to make client %v: ", err)
	}
	return client, nil
}
