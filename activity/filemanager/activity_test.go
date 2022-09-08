package filemanager

import (
	"fmt"
	"testing"

	"github.com/project-flogo/core/activity"
	//"github.com/project-flogo/core/data/mapper"
	//"github.com/project-flogo/core/data/resolve"

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

	//settings := &Settings{Action: "LIST"}
	//mf := mapper.NewFactory(resolve.GetBasicResolver())

	//iCtx := test.NewActivityInitContext(settings, mf)
	//act, err := New(iCtx)
	//assert.Nil(t, err)

	tc := test.NewActivityContext(act.Metadata())
	tc.SetInput("file", "/")

	done, err := act.Eval(tc)
	assert.True(t, done)
	assert.Nil(t, err)

	//check result attr
	aOutput := &Output{}
	err = tc.GetOutputObject(aOutput)
	assert.Nil(t, err)

	fmt.Println(aOutput.Result)
	fmt.Println(aOutput.Files)
}
