package errs

import (
	"fmt"
	"github.com/n0rad/go-erlog/with"
	o "github.com/onsi/gomega"
	"testing"
)

func TestConstruction(t *testing.T) {
	o.RegisterTestingT(t)

	path := "/bin/toto42"
	err := fmt.Errorf("%s", "erf")

	err1 := With("salut").WithErr(err).WithField("path", path)
	err2 := WithSource(err).WithMessage("salut").WithField("path", path)
	err3 := With("salut").WithErr(err).WithField("path", path)
	err4 := fill(&EntryError{
		Message: "salut",
		Fields:  with.Field("path", path),
		Err:     err,
	})

	o.Expect(Is(err1, err2)).To(o.BeTrue())
	o.Expect(Is(err3, err4)).To(o.BeTrue())
	o.Expect(Is(err2, err3)).To(o.BeTrue())
}
