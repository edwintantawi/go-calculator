package calculator_test

import (
	"github.com/edwintantawi/go-calculator/calculator"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator", func() {
	It("Should return error when statement is empty (\"\")", func() {
		c, err := calculator.Calculate("")
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal("statement cannot be empty"))
		Expect(c).To(Equal(0.0))
	})

	It("Should return error when statement is only whitespace", func() {
		c, err := calculator.Calculate("  ")
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal("statement cannot be empty"))
		Expect(c).To(Equal(0.0))
	})

	It("Should return error when operand is invalid (a + b)", func() {
		c, err := calculator.Calculate("a + b")
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal("invalid operand"))
		Expect(c).To(Equal(0.0))
	})

	It("Should return error + 2 calculation", func() {
		c, err := calculator.Calculate("+ 2")
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal("missing required operand"))
		Expect(c).To(Equal(0.0))
	})

	It("Should return error 2 + calculation", func() {
		c, err := calculator.Calculate("2 +")
		Expect(err).ToNot(BeNil())
		Expect(err.Error()).To(Equal("missing required operand"))
		Expect(c).To(Equal(0.0))
	})

	It("Should return 1 + 2 calculation result correctly", func() {
		c, err := calculator.Calculate("1 + 2")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(3.0))
	})

	It("Should return \"  1 + 2  \" calculation result correctly with whitespace", func() {
		c, err := calculator.Calculate("  1 + 2  ")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(3.0))
	})

	It("Should return 1 + 2 * 3 calculation result correctly", func() {
		c, err := calculator.Calculate("1 + 2 * 3")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(7.0))
	})

	It("Should return 1 * 2 + 3 calculation result correctly", func() {
		c, err := calculator.Calculate("1 * 2 + 3")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(5.0))
	})

	It("Should return 1 + 2 / 3 * 4 calculation result correctly", func() {
		c, err := calculator.Calculate("1 + 2 / 3 * 4")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(3.6666666666666665))
	})

	It("Should return 1 * 2 - 4 / 2 calculation result correctly", func() {
		c, err := calculator.Calculate("1 * 2 - 4 / 2")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(0.0))
	})

	It("Should return 2 ^ 2 calculation result correctly", func() {
		c, err := calculator.Calculate("2 ^ 2")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(4.0))
	})

	It("Should return 2 ^ 3 calculation result correctly", func() {
		c, err := calculator.Calculate("2 ^ 3")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(8.0))
	})

	It("Should return 2 ^ 2 ^ 3 calculation result correctly", func() {
		c, err := calculator.Calculate("2 ^ 2 ^ 3")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(256.0))
	})

	It("Should return 2 ^ 2 ^ 3 calculation result correctly", func() {
		c, err := calculator.Calculate("2 ^ 2 ^ 3")
		Expect(err).To(BeNil())
		Expect(c).To(Equal(256.0))
	})
})
