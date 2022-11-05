package stack_test

import (
	"github.com/edwintantawi/go-calculator/stack"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Stack", func() {
	Describe("String stack", func() {
		It("Should create a new stack correctly", func() {
			s := stack.NewStrings()
			Expect(len(s)).To(Equal(0))
		})

		It("Should add item to end of stack with push", func() {
			s := stack.NewStrings()
			s.Push("A")
			s.Push("B")
			s.Push("C")

			Expect(s[0]).To(Equal("A"))
			Expect(s[1]).To(Equal("B"))
			Expect(s[2]).To(Equal("C"))
		})

		It("Should remove item from end of stack and return it with pop", func() {
			s := stack.NewStrings()
			s.Push("X")
			s.Push("Gopher")

			item := s.Pop()
			Expect(item).To(Equal("Gopher"))
			Expect(len(s)).To(Equal(1))
			Expect(s[0]).To(Equal("X"))
		})
	})

	Describe("Float stack", func() {
		It("Should create a new stack correctly", func() {
			s := stack.NewFloats()
			Expect(len(s)).To(Equal(0))
		})

		It("Should add item to end of stack with push", func() {
			s := stack.NewFloats()
			s.Push(11.0)
			s.Push(12.0)
			s.Push(13.0)

			Expect(s[0]).To(Equal(11.0))
			Expect(s[1]).To(Equal(12.0))
			Expect(s[2]).To(Equal(13.0))
		})

		It("Should remove item from end of stack and return it with pop", func() {
			s := stack.NewFloats()
			s.Push(11.0)
			s.Push(12.0)

			item := s.Pop()
			Expect(item).To(Equal(12.0))
			Expect(len(s)).To(Equal(1))
			Expect(s[0]).To(Equal(11.0))
		})
	})
})
