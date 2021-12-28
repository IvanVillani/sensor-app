package client

import (
	"bytes"

	"github.com/seeis/sensor-app/cli/constants"
	"github.com/seeis/sensor-app/cli/logger"
	"github.com/seeis/sensor-app/cli/manager/flagsmanager"
	"github.com/seeis/sensor-app/cli/manager/httpmanager"
	"github.com/seeis/sensor-app/cli/manager/ioutilmanager"
	"github.com/seeis/sensor-app/cli/manager/logmanager"
)

func SendToServer(httpManager httpmanager.IHTTPManager, ioUtilManager ioutilmanager.IIOUtilManager, logManager logmanager.ILogManager, buffer *bytes.Buffer) string {
	logManager.SetupLog()

	respBody := buffer

	url := flagsmanager.Opts.WebHookURL

	resp, err := httpManager.Post(url, "application/json", respBody)

	if err != nil {
		logger.Error.Printf("Cannot execute POST to server <%s> : %s\n", url, err)
		return constants.ErrMsgPOST
	}
	defer resp.Body.Close()

	body, err := ioUtilManager.ReadAll(resp.Body)
	if err != nil {
		logger.Error.Printf("Cannot read response body: %s\n", err)
		return constants.ErrMsgReadRespBody
	}
	return string(body)
}
