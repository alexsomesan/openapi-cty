package test

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/alexsomesan/openapi-cty/foundry"
	"github.com/davecgh/go-spew/spew"
	"github.com/hashicorp/go-cty/cty"
)

type testSample struct {
	id   string
	want cty.Type
}
type testSamples []testSample

var samples = testSamples{
	{
		id: "io.k8s.apimachinery.pkg.apis.meta.v1.ObjectMeta",
		want: cty.Object(map[string]cty.Type{
			"annotations":                cty.Map(cty.String),
			"clusterName":                cty.String,
			"creationTimestamp":          cty.String,
			"deletionGracePeriodSeconds": cty.Number,
			"deletionTimestamp":          cty.String,
			"finalizers":                 cty.List(cty.String),
			"generateName":               cty.String,
			"generation":                 cty.Number,
			"labels":                     cty.Map(cty.String),
			"managedFields": cty.List(
				cty.Object(map[string]cty.Type{
					"apiVersion": cty.String,
					"fieldsType": cty.String,
					"fieldsV1":   cty.DynamicPseudoType,
					"manager":    cty.String,
					"operation":  cty.String,
					"time":       cty.String,
				})),
			"name":      cty.String,
			"namespace": cty.String,
			"ownerReferences": cty.List(
				cty.Object(map[string]cty.Type{
					"apiVersion":         cty.String,
					"blockOwnerDeletion": cty.Bool,
					"controller":         cty.Bool,
					"kind":               cty.String,
					"name":               cty.String,
					"uid":                cty.String,
				})),
			"resourceVersion": cty.String,
			"selfLink":        cty.String,
			"uid":             cty.String,
		}),
	},
	{
		id: "io.k8s.api.core.v1.ServicePort",
		want: cty.Object(map[string]cty.Type{
			"appProtocol": cty.String,
			"name":        cty.String,
			"nodePort":    cty.Number,
			"port":        cty.Number,
			"protocol":    cty.String,
			"targetPort":  cty.DynamicPseudoType,
		}),
	},
	// {
	// 	id:   "io.k8s.apiextensions-apiserver.pkg.apis.apiextensions.v1.CustomResourceDefinition",
	// 	want: cty.Object(map[string]cty.Type{}),
	// },
}

func TestGetType(t *testing.T) {
	tf, err := buildFixtureFoundry()
	if err != nil {
		t.Skip()
	}
	for _, s := range samples {
		t.Run(s.id,
			func(t *testing.T) {
				rt, err := tf.GetTypeByID(s.id)
				if err != nil {
					t.Fatal(fmt.Errorf("GetTypeByID() failed: %s", err))
				}
				if !rt.Equals(s.want) {
					t.Fatalf("\nRETURNED %s\nEXPECTED: %s", spew.Sdump(rt), spew.Sdump(s.want))
				}
			})
	}
}

func TestFoundryOAPIv2(t *testing.T) {
	_, err := buildFixtureFoundry()
	if err != nil {
		t.Error(err)
	}
}

func buildFixtureFoundry() (foundry.Foundry, error) {
	sfile := filepath.Join("testdata", "k8s-swagger.json")

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
