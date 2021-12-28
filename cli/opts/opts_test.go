package opts_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/seeis/sensor-app/cli/constants"
	"github.com/seeis/sensor-app/cli/logger"
	"github.com/seeis/sensor-app/cli/mocks"
	"github.com/seeis/sensor-app/cli/opts"
)

var _ = Describe("Opts", func() {
	Describe("Parsing the input options into commands", func() {
		Context("When no errors during parsing", func() {
			It("should return nil error and empty string", func() {
				res, err := opts.ParseOpts(mocks.MockFlags{
					MockedWithErr:    false,
					ErrorIsFromFlags: false,
					ErrorIsHelpError: false,
				})

				Expect(res).To(Equal(""))
				Expect(err).To(BeNil())
			})
		})

		Context("When error is unknown", func() {
			It("should return unknown error and terminal message", func() {
				buffer := gbytes.NewBuffer()

				logger.MockSetupLogger(buffer)

				res, err := opts.ParseOpts(mocks.MockFlags{
					MockedWithErr:    true,
					ErrorIsFromFlags: false,
					ErrorIsHelpError: false,
				})

				Expect(res).To(Equal(constants.ErrMsgParse))
				Expect(err.Error()).To(Equal("unknown error"))
				Expect(buffer).To(gbytes.Say("ERROR"))
			})
		})

		Context("When error is from flags but is not of type ErrHelp", func() {
			It("should return flags error and empty string", func() {
				buffer := gbytes.NewBuffer()

				logger.MockSetupLogger(buffer)

				res, err := opts.ParseOpts(mocks.MockFlags{
					MockedWithErr:    true,
					ErrorIsFromFlags: true,
					ErrorIsHelpError: false,
				})

				Expect(res).To(Equal(""))
				Expect(err.Error()).To(Equal("Error flag message from MockFlags"))
				Expect(buffer).To(gbytes.Say("WARNING"))
			})
		})

		Context("When error is from flags and is of type ErrHelp", func() {
			It("should return help error and terminal message", func() {
				buffer := gbytes.NewBuffer()

				logger.MockSetupLogger(buffer)

				res, err := opts.ParseOpts(mocks.MockFlags{
					MockedWithErr:    true,
					ErrorIsFromFlags: true,
					ErrorIsHelpError: true,
				})

				Expect(res).To(Equal(constants.ErrMsgFlags))
				Expect(err.Error()).To(Equal("Error help message from MockFlags"))
				Expect(buffer).To(gbytes.Say("INFO"))
			})
		})
	})
})
