package orchestration

import (
	"os"
	"testing"

	"github.com/eris-ltd/eris/log"
	"github.com/eris-ltd/eris/testutil"
	//docker "github.com/fsouza/go-dockerclient"
)

func TestMain(m *testing.M) {
	log.SetLevel(log.ErrorLevel)
	// log.SetLevel(log.InfoLevel)
	// log.SetLevel(log.DebugLevel)

	testutil.IfExit(testutil.Init(testutil.Pull{
		Images:   []string{"data", "keys", "ipfs"},
		Services: []string{"keys", "ipfs"},
	}))

	testutil.RemoveAllContainers()

	exitCode := m.Run()
	testutil.IfExit(testutil.TearDown())
	os.Exit(exitCode)
}

func TestCreateBase(t *testing.T) {
	defer testutil.RemoveAllContainers()
	base, err := CreateBase()
	if err != nil {
		t.Fatal(err)
	}
}

func TestPullImage(t *testing.T) {
	defer testutil.RemoveAllContainers()
	base, err := CreateBase()
	if err != nil {
		t.Fatal(err)
	}
	base.PullImageOptions.Repository = "ethereum/solc"
	base.PullImageOptions.Tag = "stable"
	if err := base.Pull(); err != nil {
		t.Fatal(err)
	}
}
