package expconv_test

import (
	"github.com/edwintantawi/go-calculator/expconv"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Itop", func() {
	It("Should convert infix \"\" to postfix correctly", func() {
		r := expconv.ItoP("")
		Expect(r).To(Equal(""))
	})

	It("Should convert infix 2 to postfix correctly", func() {
		r := expconv.ItoP("2")
		Expect(r).To(Equal("2"))
	})

	It("Should convert infix + 2 to postfix correctly", func() {
		r := expconv.ItoP("+ 2")
		Expect(r).To(Equal("2 +"))
	})

	It("Should convert infix 2 + to postfix correctly", func() {
		r := expconv.ItoP("2 +")
		Expect(r).To(Equal("2 +"))
	})

	It("Should convert infix 1 + 2 to postfix correctly", func() {
		r := expconv.ItoP("1 + 2")
		Expect(r).To(Equal("1 2 +"))
	})

	It("Should convert infix 1 + 2 - 3 to postfix correctly", func() {
		r := expconv.ItoP("1 + 2 - 3")
		Expect(r).To(Equal("1 2 + 3 -"))
	})

	It("Should convert infix 2 * 2 + 3 to postfix correctly", func() {
		r := expconv.ItoP("2 * 2 + 3")
		Expect(r).To(Equal("2 2 * 3 +"))
	})

	It("Should convert infix 2 + 2 * 3 to postfix correctly", func() {
		r := expconv.ItoP("2 + 2 * 3")
		Expect(r).To(Equal("2 2 3 * +"))
	})

	It("Should convert infix 2 / 2 * 3 to postfix correctly", func() {
		r := expconv.ItoP("2 / 2 * 3")
		Expect(r).To(Equal("2 2 / 3 *"))
	})

	It("Should convert infix 2 + 2 * 3 / 2 to postfix correctly", func() {
		r := expconv.ItoP("2 + 2 * 3 / 2")
		Expect(r).To(Equal("2 2 3 * 2 / +"))
	})

	It("Should convert infix 2 ^ 2 ^ 3 to postfix correctly", func() {
		r := expconv.ItoP("2 ^ 2 ^ 3")
		Expect(r).To(Equal("2 2 3 ^ ^"))
	})

	It("Should convert infix 2 + 2 ^ 3 to postfix correctly", func() {
		r := expconv.ItoP("2 + 2 ^ 3")
		Expect(r).To(Equal("2 2 3 ^ +"))
	})

	It("Should convert infix 2 * 2 ^ 3 to postfix correctly", func() {
		r := expconv.ItoP("2 * 2 ^ 3")
		Expect(r).To(Equal("2 2 3 ^ *"))
	})

	It("Should convert infix 2 ^ 2 * 3 to postfix correctly", func() {
		r := expconv.ItoP("2 ^ 2 * 3")
		Expect(r).To(Equal("2 2 ^ 3 *"))
	})
})
