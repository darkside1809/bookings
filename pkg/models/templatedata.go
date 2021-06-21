package models

// TemplateData renders the different data types
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap	 map[string]float32
	Data		 map[string]interface{}
	Token		 string
	Flash		 string
	Warning	 string
	Error		 string
}