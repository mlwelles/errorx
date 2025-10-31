package errorx_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/noho-digital/insurews/pkg/suite"
	"github.com/noho-digital/insurews/pkg/test/reporter"

	"github.com/franela/goblin"
	. "github.com/noho-digital/insurews/pkg/errors"
	"github.com/onsi/gomega"
)

type ErrorsSuite struct {
	suite.StandaloneSuite
	g *goblin.G
}

func TestErrors(t *testing.T) {
	suite.Run(t, new(ErrorsSuite))
}

func (s *ErrorsSuite) SetupTest() {
	s.g = goblin.Goblin(s.T())
	s.g.SetReporter(reporter.NewReporter())
	gomega.RegisterFailHandler(func(m string, _ ...int) { s.g.Fail(m) })
}

// tests follow
func (s *ErrorsSuite) SkipTestWrapVsMatchError() {
	err := ErrMinAgeAnswer18
	s.g.Describe(fmt.Sprintf("Wrapped error '%v'", err), func() {
		s.g.It("should match themselves", func() {
			gomega.Expect(err).To(gomega.MatchError(err))
		})
		s.g.It("should MatchError with ErrAgeAnswer", func() {
			gomega.Expect(err).To(gomega.MatchError(ErrInvalidAgeAnswer))
		})
		s.g.It("should MatchError with ErrInvalidAnswer", func() {
			gomega.Expect(err).To(gomega.MatchError(ErrInvalidAnswer))
		})
	})
}

func (s *ErrorsSuite) TestErrorMatch() {
	err := ErrMinAgeAnswer18
	s.g.Describe(fmt.Sprintf("Wrapped error '%v'", err), func() {
		s.g.It("should Is themselves", func() {
			s.g.Assert(Match(err, err)).IsTrue()
		})
		s.g.It("should also be an ErrAgeAnswer", func() {
			s.g.Assert(Match(err, ErrInvalidAgeAnswer)).IsTrue()
		})
		s.g.It("should also be an ErrInvalidAnswer", func() {
			s.g.Assert(Match(err, ErrInvalidAnswer)).IsTrue()
		})
	})
}

func (s *ErrorsSuite) TestStdErrorIs() {
	err := ErrMinAgeAnswer18
	s.g.Describe(fmt.Sprintf("Wrapped error '%v'", err), func() {
		s.g.It("should errors.Is themselves", func() {
			s.g.Assert(errors.Is(err, err)).IsTrue()
		})
		s.g.It("should also be an ErrAgeAnswer", func() {
			s.g.Assert(errors.Is(err, ErrInvalidAgeAnswer)).IsTrue()
		})
		s.g.It("should also be an ErrInvalidAnswer", func() {
			s.g.Assert(errors.Is(err, ErrInvalidAnswer)).IsTrue()
		})
	})
}

func (s *ErrorsSuite) TestFormatting() {
	e1 := ErrInvalidAgeAnswer
	e2 := Wrap(e1, "just testing")
	expected := "just testing: Please enter your correct birthday.: invalid answer provided"
	s.g.Describe(fmt.Sprintf("Wrapped error '%v'", e2), func() {
		s.g.It("should format %s as expected", func() {
			s.g.Assert(fmt.Sprintf("%s", e2)).Equal(expected)
		})
		s.g.It("should format %v as expected", func() {
			s.g.Assert(fmt.Sprintf("%v", e2)).Equal(expected)
		})

	})
}

func (s *ErrorsSuite) TestMatch() {
	matches := []error{NoToken(), InvalidToken(), NotFound(), ErrAccountNotFound}
	err := WrapError(NoToken(), ErrAccountNotFound)
	for _, expected := range matches {
		s.Assert().True(Match(err, expected), "%q should match %q", err.Error(), expected.Error())
	}
}
