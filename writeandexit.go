//Package writeerror is for writing to STDERR and logging "github.com/juju/errors" errors,
//and finally os.Exit with an exit code
package writeerror

import (
	"fmt"
	"log"
	"os"

	"github.com/juju/errors"
)

//AndExit adds a "github.com/juju/errors".Trace to the error, writethe error to STDERR, log the error, and finally
//os.Exit with the specified exitCode
func AndExit(err error, exitCode int) {
	annotatedErr := errors.Trace(err)
	s := errors.ErrorStack(annotatedErr)
	fmt.Fprintf(os.Stderr, s)
	log.Println(s)
	os.Exit(exitCode)
}
