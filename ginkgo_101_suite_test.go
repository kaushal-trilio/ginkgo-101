package ginkgo_101_test

import (
	ginkgo_101 "ginkgo-101"
	"log"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGinkgo101(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Ginkgo101 Suite")
}

var _ = Describe("Person.IsChild()", func() {
	Context("when the person is a child", func() {
		It("returns true", func() {
			person := ginkgo_101.Person{Age: 17}

			response := person.IsChild()

			Expect(response).To(BeTrue())
		})
	})

	Context("when the person is an adult", func() {
		BeforeEach(func() {
			log.Print("not a child")
		})
		It("returns true", func() {
			person := ginkgo_101.Person{Age: 100}

			response := person.IsChild()

			Expect(response).To(BeFalse())
		})
	})

	DescribeTable("isChild table test", func(age int, expectedResponse bool) {
		p := ginkgo_101.Person{Age: age}

		Expect(p.IsChild()).To(Equal(expectedResponse))
	},
		Entry("when is a child", 10, true),
		Entry("when is a child", 19, false),
	)
})
