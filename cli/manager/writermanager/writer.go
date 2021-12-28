package writermanager

import (
	"bytes"
	"fmt"
	"os"

	"github.com/seeis/sensor-app/cli/client"
	"github.com/seeis/sensor-app/cli/constants"
	"github.com/seeis/sensor-app/cli/manager/flagsmanager"
	"github.com/seeis/sensor-app/cli/manager/httpmanager"
	"github.com/seeis/sensor-app/cli/manager/ioutilmanager"
	"github.com/seeis/sensor-app/cli/manager/logmanager"
)

// WriterManager struct: implements IWriterManager interface
type WriterManager struct{}

// WriteMsg function: writes the data to selected output
func (writerManager WriterManager) WriteMsg(data string) {
	var b bytes.Buffer

	fmt.Fprintln(&b, data)

	if flagsmanager.WebHookURLProvided() && data != constants.InfoMsgEnd && data != constants.ErrMsgPOST {
		fmt.Fprintln(&b, client.SendToServer(httpmanager.HTTPManager{}, ioutilmanager.IOUtilManager{}, logmanager.LogManager{}, &b))
	}

	b.WriteTo(os.Stdout)
}
