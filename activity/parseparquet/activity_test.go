package parseparquet

import (
	"testing"

	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"

	// old
	//"github.com/TIBCOSoftware/flogo-lib/core/support/test"
	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"

	"github.com/stretchr/testify/assert"
)

/*
func TestRegister(t *testing.T) {

	ref := activity.GetRef(&ParseParquetActivity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
} */

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {

	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("descriptor.json")
		if err != nil {
			panic("No Json Metadata found for descriptor.json path")
		}

		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	return activityMetadata
}

func TestCreate(t *testing.T) {

	act := NewActivity(getActivityMetadata())

	if act == nil {
		t.Error("Activity Not Created")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {

	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(getActivityMetadata())
	tc.SetInput("parquetFile", "/tmp/file.parquet")
	tc.SetInput("maxRows", 1000)
	tc.SetInput("initRow", 0)

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)

	//check result attr
	b, _ := json.Marshal(tc.GetOutput("output"))
	fmt.Println(string(b))
}
