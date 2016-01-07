package with
import (
	"testing"
	o "github.com/onsi/gomega"
)


type Options struct {
	Paginate bool
	Count    uint
	MaxPages uint
}

func (o Options) ToFields() Fields {
	fields := make(Fields, 3)
	fields["paginate"] = o.Paginate
	fields["count"] = o.Count
	fields["maxpages"] = o.MaxPages
	return fields
}

func TestWithAllInterface(t *testing.T) {
	o.RegisterTestingT(t)
	ops := &Options{Paginate: true, Count: 42, MaxPages:3}
	with := Fields(ops)

	o.Expect(with).To(o.HaveLen(3))
	o.Expect(with).To(o.HaveKeyWithValue("paginate", true))
	o.Expect(with).To(o.HaveKeyWithValue("count", 42))
	o.Expect(with).To(o.HaveKeyWithValue("maxpages", 3))
}

func TestWith(t *testing.T) {
	o.RegisterTestingT(t)
	dir := Field("dir", "there")

	o.Expect(dir["dir"].(string)).To(o.Equal("there"))
}

func TestImmutable(t *testing.T) {
	o.RegisterTestingT(t)
	dir := Field("dir", "there")
	other := dir.With("other", "here")

	o.Expect(dir).To(o.HaveLen(1))
	o.Expect(other).To(o.HaveLen(2))
	o.Expect(other["dir"]).To(o.Equal("there"))
	o.Expect(other["other"]).To(o.Equal("here"))

}