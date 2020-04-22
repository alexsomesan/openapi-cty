package test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/alexsomesan/openapi-cty/foundry"
)

func buildFixtureFoundry() (foundry.Foundry, error) {
	sfile := filepath.Join("testdata","k8s-swagger.json")

	input, err := ioutil.ReadFile(sfile)
	if err != nil {
		return nil, fmt.Errorf("failed to load definition file: %s : %s", sfile, err)
	}
	
	tf, err := foundry.NewFoundryFromSpecV2(input)
	
	if err != nil {
		return nil, err
	}
	
	if tf == nil {
		return nil, fmt.Errorf("constructed foundry is nil")
	}
	
	return tf, nil
}

func TestFoundryOAPIv2(t *testing.T) {
	_, err := buildFixtureFoundry()
	if err != nil {
		t.Error(err)
	}
}

func TestGetPodSecurityPolicyType(t *testing.T) {
	tf, err := buildFixtureFoundry()
	if err != nil {
		t.Skip()
	}
	ty, err := tf.GetTypeById("io.k8s.api.policy.v1beta1.PodSecurityPolicy")
	if err != nil {
		t.Fatal(err)
	}
	if !ty.IsObjectType() {
		t.Fatalf("%s type returned - expected an Object", ty.FriendlyName())
	}
}