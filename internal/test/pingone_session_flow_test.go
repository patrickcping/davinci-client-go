package test

import (
	"fmt"
	// dv "github.com/samir-gandhi/davinci-client-go/davinci"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPingOneSessionFlowApp(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	apps := makeTestApps(RandString(10))
	app := apps["plain"]
	policy := testDataAppPolicy["basePolicy"]
	app.Policies = append(app.Policies, policy)
	// using flow that is manually created in feature flag env.
	app.Policies[0].PolicyFlows[0].FlowID = "5c32a89d4093b0eba7292ecafdd6b0e9"
	resp, err := c.CreateInitializedApplication(&c.CompanyID, &app)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, "AUTHENTICATION", resp.Policies[0].Trigger.TriggerType)
}

func TestNoPolicyApp(t *testing.T) {
	c, err := newTestClient()
	if err != nil {
		panic(err)
	}
	apps := makeTestApps(RandString(10))
	app := apps["noPolicy"]
	_, err = c.CreateInitializedApplication(&c.CompanyID, &app)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

}
