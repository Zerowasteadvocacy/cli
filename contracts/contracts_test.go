package contracts

import (
	"os"
	"testing"

	tests "github.com/eris-ltd/eris-cli/testutils"

	log "github.com/eris-ltd/eris-cli/Godeps/_workspace/src/github.com/Sirupsen/logrus"
	logger "github.com/eris-ltd/eris-cli/Godeps/_workspace/src/github.com/eris-ltd/common/go/log"
)

func TestMain(m *testing.M) {
	log.SetFormatter(logger.ErisFormatter{})

	log.SetLevel(log.ErrorLevel)
	// log.SetLevel(log.InfoLevel)
	// log.SetLevel(log.DebugLevel)

	tests.IfExit(testsInit())

	exitCode := m.Run()

	log.Info("Tearing tests down")
	if os.Getenv("TEST_IN_CIRCLE") != "true" {
		tests.IfExit(tests.TestsTearDown())
	}

	os.Exit(exitCode)
}

func TestContractsTest(t *testing.T) {
	// TODO: test more app types once we have
	// canonical apps + eth throwaway chains
}

func TestContractsDeploy(t *testing.T) {
	// TODO: finish. not worried about this too much now
	// since test will deploy.
}

func testsInit() error {
	if err := tests.TestsInit("contracts"); err != nil {
		return err
	}
	return nil
}
