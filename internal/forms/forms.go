package forms

import (
	// built in Golang packages
	"fmt"
	"net/url"
	"strings"
	"regexp"
	"strconv"
	// External packages/dependencies
	"github.com/asaskevich/govalidator"
)

// Number validation variables
var (
	validAlpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	validPunctuation = "-x\u2010-\u2015\u2212\u30FC\uFF0D-\uFF0F " + "\u00A0\u00AD\u200B\u2060\u3000()\uFF08\uFF09\uFF3B\uFF3D." + "\\[\\]/~\u2053\u223C\uFF5E"
	regexpValidNumber = regexp.MustCompile("^(" + validPhoneNumber + "(?:" + extnPatternForParsing + ")?)$")
	validPhoneNumber = "\\p{Nd}" + "{" + strconv.Itoa(2) + "}" + "|" + "[" + "+\uFF0B" + "]*(?:[" + validPunctuation + string('*') + "]*" + "\\p{Nd}" + "){3,}[" +
	validPunctuation + string('*') + validAlpha + "\\p{Nd}" + "]*"
	extnPatternForParsing = ";ext=" + "(" + "\\p{Nd}" + "{1,7})" + "|" + "[ \u00A0\\t,]*" + "(?:e?xt(?:ensi(?:o\u0301?|\u00F3))?n?|\uFF45?\uFF58\uFF54\uFF4E?|" + "[;,x\uFF58#\uFF03~\uFF5E]|int|anexo|\uFF49\uFF4E\uFF54)" + "[:\\.\uFF0E]?[ \u00A0\\t,-]*" + "(" + "\\p{Nd}" + "{1,7})" + "#?|" + "[- ]+(" + "\\p{Nd}" + "{1,5})#"
)

// Form creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors
}

// Valid returns true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New is a constructor for Form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Required iterates through fields and checks is a definite filed empty or not
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		val := f.Get(field)
		if strings.TrimSpace(val) == "" {
			f.Errors.Add(field, "This filed cannot be blank")
		}
	}
}

// Has checks is form field in post and not empty
func (f *Form) Has(field string) bool {
	x := f.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}
	return true
}

// MinLength checks for sttring minimum length of user's credentials
func (f *Form) MinLength(field string, length int) bool {
	x := f.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}
	return true
}

//  CheckEmail checks for valid email address entered by user
func (f *Form) CheckEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Ivalid email address")
	}
}

// ValidatePhoneNumber checks for valid phone number, if it contains letter, return false
func (f *Form) ValidatePhoneNumber(number string) bool {
	if len(number) < 2 {
		f.Errors.Add(number, "Ivalid phone number")
	}
	if !regexpValidNumber.MatchString(number) {
		f.Errors.Add(number, "Ivalid phone number")
	}
	return true
}

