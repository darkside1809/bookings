package models

import (
	"github.com/darkside1809/bookings/pkg/forms"
)

// TemplateData renders the different data types to client
type TemplateData struct {
	StringMap 			map[string]string
	IntMap    			map[string]int
	FloatMap  			map[string]float32
	Data      			map[string]interface{}
	CSRFToken 			string
	Flash     			string
	Warning   			string
	Error     			string
	Form      			*forms.Form
	IsAuthenticated	int
}
