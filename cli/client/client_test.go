package client_test

import (
	"bytes"
	"errors"
	"fmt"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/seeis/sensor-app/cli/client"
	"github.com/seeis/sensor-app/cli/constants"
	"github.com/seeis/sensor-app/cli/mocks"
)

var _ = Describe("Client", func() {
	Describe("Sends measurement data to server", func() {
		Context("When executing POST request with no error", func() {
			It("should return response body", func() {
				var b bytes.Buffer

				fmt.Fprintln(&b, "MockedString")

				result := client.SendToServer(mocks.MockHttp{
					Status: "200, OK",
					Error:  nil,
				}, mocks.MockIOUtil{
					Msg:   "MockedMsg",
					Error: nil,
				}, mocks.MockLogger{
					Writer: os.Stdout,
				}, &b)

				Expect(result).To(Equal("MockedMsg"))
			})
		})

		Context("Executing POST request results in error", func() {
			It("should return error message", func() {
				var b bytes.Buffer

				fmt.Fprintln(&b, "MockedString")

				result := client.SendToServer(mocks.MockHttp{
					Status: "404, NOT FOUND",
					Error:  errors.New("POST error"),
				}, mocks.MockIOUtil{
					Msg:   "MockedMsg",
					Error: nil,
				}, mocks.MockLogger{
					Writer: os.Stdout,
				}, &b)

				Expect(result).To(Equal(constants.ErrMsgPOST))
			})
		})

		Context("Reading response body results in error", func() {
			It("should return error message", func() {
				var b bytes.Buffer

				fmt.Fprintln(&b, "MockedString")

				result := client.SendToServer(mocks.MockHttp{
					Status: "200, OK",
					Error:  nil,
				}, mocks.MockIOUtil{
					Msg:   "MockedMsg",
					Error: errors.New("RespBody error"),
				}, mocks.MockLogger{
					Writer: os.Stdout,
				}, &b)

				Expect(result).To(Equal(constants.ErrMsgReadRespBody))
			})
		})
	})
})
