package engine_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gbytes"
	"github.com/seeis/sensor-app/cli/constants"
	"github.com/seeis/sensor-app/cli/engine"
	"github.com/seeis/sensor-app/cli/mocks"
)

var _ = Describe("Engine", func() {
	Describe("Starts the engine for measurements", func() {
		Context("When uninterrupted - one measurement", func() {
			It("should return measurement meassage", func() {
				buffer := gbytes.NewBuffer()

				engine.Start(mocks.MockWriter{
					Output: buffer,
				}, mocks.MockOS{
					MockedWithErr: false,
				}, mocks.MockFlags{
					MockedWithErr:    false,
					ErrorIsFromFlags: false,
					ErrorIsHelpError: false,
				}, mocks.MockMeasure{
					Uninterrupted: true,
				})

				Expect(buffer).To(gbytes.Say("Mocked measurement"))
			})
		})

		Context("When interrupted during sensor reading", func() {
			It("should return terminal error meassage", func() {
				buffer := gbytes.NewBuffer()

				engine.Start(mocks.MockWriter{
					Output: buffer,
				}, mocks.MockOS{
					MockedWithErr: false,
				}, mocks.MockFlags{
					MockedWithErr:    false,
					ErrorIsFromFlags: false,
					ErrorIsHelpError: false,
				}, mocks.MockMeasure{
					Uninterrupted: false,
				})

				Expect(buffer).To(gbytes.Say(constants.ErrMsgRead))
			})
		})

		Context("When interrupted during flags parsing", func() {
			It("should return terminal error meassage", func() {
				buffer := gbytes.NewBuffer()

				engine.Start(mocks.MockWriter{
					Output: buffer,
				}, mocks.MockOS{
					MockedWithErr: false,
				}, mocks.MockFlags{
					MockedWithErr:    true,
					ErrorIsFromFlags: false,
					ErrorIsHelpError: false,
				}, mocks.MockMeasure{
					Uninterrupted: false,
				})

				Expect(buffer).To(gbytes.Say(constants.ErrMsgParse))
			})
		})
	})
})
