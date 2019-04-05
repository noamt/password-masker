package masker_test

import (
	"github.com/noamt/password-masker"
	"testing"
)

func ExampleMask() {
	masker.Mask("password: abcdef123456789")
}

func ExampleMaskWith() {
	masker.MaskWith("password: abcdef123456789", "????")
}

func TestMaskColon(t *testing.T) {
	got := masker.Mask("password:abcdef123456789")
	if "password:****" != got {
		t.Errorf("Got '%s'; want 'password:****'", got)
	}
}

func TestMaskMultilineColon(t *testing.T) {
	got := masker.Mask(`
password:
	abcdef123456789
`)

	expected := `
password:
	****
`
	if expected != got {
		t.Errorf("Got '%s'; want '%s'", got, expected)
	}
}

func TestMaskMultilineWithFluff(t *testing.T) {
	got := masker.Mask(`
some: values

password:
	abcdef123456789

more:
    values
`)

	expected := `
some: values

password:
	****

more:
    values
`
	if expected != got {
		t.Errorf("Got '%s'; want '%s'", got, expected)
	}
}

func TestMaskMultilineWithMultipleOccurrences(t *testing.T) {
	got := masker.Mask(`
password:
	abcdef123456789

password:
	whynot
`)

	expected := `
password:
	****

password:
	****
`
	if expected != got {
		t.Errorf("Got '%s'; want '%s'", got, expected)
	}
}

func TestMaskColonSpace(t *testing.T) {
	got := masker.Mask("password: abcdef123456789")
	if "password: ****" != got {
		t.Errorf("Got '%s'; want 'password: ****'", got)
	}
}

func TestMaskSpaceColonSpace(t *testing.T) {
	got := masker.Mask("password : abcdef123456789")
	if "password : ****" != got {
		t.Errorf("Got '%s'; want 'password : ****'", got)
	}
}

func TestMaskEqual(t *testing.T) {
	got := masker.Mask("password=abcdef123456789")
	if "password=****" != got {
		t.Errorf("Got '%s'; want 'password=****'", got)
	}
}

func TestMaskEqualSpace(t *testing.T) {
	got := masker.Mask("password= abcdef123456789")
	if "password= ****" != got {
		t.Errorf("Got '%s'; want 'password= ****'", got)
	}
}

func TestMaskSpaceEqualSpace(t *testing.T) {
	got := masker.Mask("password = abcdef123456789")
	if "password = ****" != got {
		t.Errorf("Got '%s'; want 'password = ****'", got)
	}
}

func TestMaskWith(t *testing.T) {
	got := masker.MaskWith("password: abcdef123456789", "????")
	if "password: ????" != got {
		t.Errorf("Got '%s'; want 'password: ????'", got)
	}
}

func TestMaskPass(t *testing.T) {
	got := masker.Mask("pass: abcdef123456789")
	if "pass: ****" != got {
		t.Errorf("Got '%s'; want 'pass: ****'", got)
	}
}
