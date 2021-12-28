package logger_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/seeis/sensor-app/cli/constants"
	"github.com/seeis/sensor-app/cli/logger"
	"github.com/seeis/sensor-app/cli/mocks"
)

var _ = Describe("Logger", func() {
	Describe("Configuring the logger and output meassages", func() {
		Context("When can't open/create log file", func() {
			It("should return error and log to os.Stderr", func() {
				logger.SetupLogger(mocks.MockOS{
					MockedWithErr: true,
				})

				infoWriter := logger.Info.Writer()
				warningWriter := logger.Warning.Writer()
				errorWriter := logger.Error.Writer()

				Expect(infoWriter).To(Equal(os.Stderr))
				Expect(warningWriter).To(Equal(os.Stderr))
				Expect(errorWriter).To(Equal(os.Stderr))
			})
		})

		Context("When opened/created log file", func() {
			It("should return file and log to it", func() {
				logger.SetupLogger(mocks.MockOS{
					MockedWithErr: false,
				})

				infoWriter := logger.Info.Writer()
				file, ok := infoWriter.(*os.File)
				if ok {
					Expect(file.Name()).To(Equal(constants.LogFileName))
				}

				warningWriter := logger.Warning.Writer()
				file, ok = warningWriter.(*os.File)
				if ok {
					Expect(file.Name()).To(Equal(constants.LogFileName))
				}

				errorWriter := logger.Error.Writer()
				file, ok = errorWriter.(*os.File)
				if ok {
					Expect(file.Name()).To(Equal(constants.LogFileName))
				}
			})
		})

		Context("When opened/created log file", func() {
			It("should return file and log to it", func() {
				logger.MockSetupLogger(os.Stdout)

				infoWriter := logger.Info.Writer()
				warningWriter := logger.Warning.Writer()
				errorWriter := logger.Error.Writer()

				Expect(infoWriter).To(Equal(os.Stdout))
				Expect(warningWriter).To(Equal(os.Stdout))
				Expect(errorWriter).To(Equal(os.Stdout))
			})
		})
	})
})
