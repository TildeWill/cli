package commands_test

import (
	. "cf/commands"
	"github.com/stretchr/testify/assert"
	"testhelpers"
	"testing"
)

func TestCreateSpaceFailsWithUsage(t *testing.T) {
	reqFactory := &testhelpers.FakeReqFactory{}
	spaceRepo := &testhelpers.FakeSpaceRepository{}

	fakeUI := callCreateSpace([]string{}, reqFactory, spaceRepo)
	assert.True(t, fakeUI.FailedWithUsage)

	fakeUI = callCreateSpace([]string{"my-space"}, reqFactory, spaceRepo)
	assert.False(t, fakeUI.FailedWithUsage)
}

func TestCreateSpaceRequirements(t *testing.T) {
	spaceRepo := &testhelpers.FakeSpaceRepository{}

	reqFactory := &testhelpers.FakeReqFactory{LoginSuccess: true, OrgSuccess: true}
	callCreateSpace([]string{"my-space"}, reqFactory, spaceRepo)
	assert.True(t, testhelpers.CommandDidPassRequirements)

	reqFactory = &testhelpers.FakeReqFactory{LoginSuccess: true, OrgSuccess: false}
	callCreateSpace([]string{"my-space"}, reqFactory, spaceRepo)
	assert.False(t, testhelpers.CommandDidPassRequirements)

	reqFactory = &testhelpers.FakeReqFactory{LoginSuccess: false, OrgSuccess: true}
	callCreateSpace([]string{"my-space"}, reqFactory, spaceRepo)
	assert.False(t, testhelpers.CommandDidPassRequirements)

}

func TestCreateSpace(t *testing.T) {
	spaceRepo := &testhelpers.FakeSpaceRepository{}

	reqFactory := &testhelpers.FakeReqFactory{LoginSuccess: true, OrgSuccess: true}
	fakeUI := callCreateSpace([]string{"my-space"}, reqFactory, spaceRepo)

	assert.Contains(t, fakeUI.Outputs[0], "Creating space")
	assert.Contains(t, fakeUI.Outputs[0], "my-space")
	assert.Equal(t, spaceRepo.CreateSpaceName, "my-space")
	assert.Contains(t, fakeUI.Outputs[1], "OK")
	assert.Contains(t, fakeUI.Outputs[2], "TIP")
}

func callCreateSpace(args []string, reqFactory *testhelpers.FakeReqFactory, spaceRepo *testhelpers.FakeSpaceRepository) (ui *testhelpers.FakeUI) {
	ui = new(testhelpers.FakeUI)
	ctxt := testhelpers.NewContext("create-space", args)

	cmd := NewCreateSpace(ui, spaceRepo)
	testhelpers.RunCommand(cmd, ctxt, reqFactory)
	return
}