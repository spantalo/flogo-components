package parquetfileparser

import (
	"testing"

	"encoding/json"
	"fmt"

	"github.com/project-flogo/core/activity"
	"github.com/project-flogo/core/support/test"
	"github.com/stretchr/testify/assert"
)

func TestRegister(t *testing.T) {

	ref := activity.GetRef(&Activity{})
	act := activity.Get(ref)

	assert.NotNil(t, act)
}

func TestEval(t *testing.T) {

	defer func() {
		if r := recover(); r != nil {
			t.Failed()
			t.Errorf("Error during execution: %v", r)
		}
	}()

	act := &Activity{}
	tc := test.NewActivityContext(act.Metadata())

	tc.SetInput("filename", "/tmp/file.parquet")
	tc.SetInput("maxrows", 1000)
	tc.SetInput("initrow", 0)

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)

	//check result attr
	aOutput := &Output{}
	err = tc.GetOutputObject(aOutput)
	assert.Nil(t, err)

	b, _ := json.Marshal(aOutput.Result)
	fmt.Println(string(b))
}
