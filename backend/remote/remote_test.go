package remote

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/logging"
)

func TestMain(m *testing.M) {
	flag.Parse()

	if testing.Verbose() {
		// if we're verbose, use the logging requested by TF_LOG
		logging.SetOutput()
	} else {
		// otherwise silence all logs
		log.SetOutput(ioutil.Discard)
	}

	// Make sure TF_FORCE_LOCAL_BACKEND is unset
	os.Unsetenv("TF_FORCE_LOCAL_BACKEND")

	os.Exit(m.Run())
}
