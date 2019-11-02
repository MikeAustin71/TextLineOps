package textlinebuilder

type TextFieldSpec interface {
	TextTypeName() string
}

type IntegerSpec struct {
	NumericValue          int
	NumericFieldSpec      string
	NumericFieldLength    int
	NumericPadChar        rune
	NumericPosition       FieldPositionSpec
}

func (intSpec *IntegerSpec) CopyOut() IntegerSpec {

	newInt := IntegerSpec{}
	newInt.NumericValue = intSpec.NumericValue
	newInt.NumericFieldSpec = intSpec.NumericFieldSpec
	newInt.NumericFieldLength = intSpec.NumericFieldLength
	newInt.NumericPadChar = intSpec.NumericPadChar

	return newInt
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextFieldSpec interface
func (intSpec IntegerSpec) TextTypeName() string {
	return "IntegerSpec"
}

// NewPtr - Returns a pointer to a new LineSpec
// instance.
func (intSpec IntegerSpec) NewPtr() *IntegerSpec {

	anotherIntTxtSpec := IntegerSpec{}

	return &anotherIntTxtSpec
}


type LineSpec struct {
	LineChar         rune
	LineLength       int
}

// CopyOut - Returns a deep copy of LineSpec
// data fields.
func (lineSpec *LineSpec) CopyOut() LineSpec {

	newLineSpec := LineSpec{}
	newLineSpec.LineChar = lineSpec.LineChar
	newLineSpec.LineLength = lineSpec.LineLength

	return newLineSpec
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextFieldSpec interface
func (lineSpec LineSpec) TextTypeName() string {
	return "LineSpec"
}

// NewPtr - Returns a pointer to a new LineSpec
// instance.
func (lineSpec LineSpec) NewPtr() *LineSpec {

	anotherLineSpec := LineSpec{}

	return &anotherLineSpec
}


type MarginSpec struct {
	MarginStr           string
	MarginLength        int
	MarginChar          rune
}

func (margin *MarginSpec) CopyOut() MarginSpec {

	newMargin := MarginSpec{}

	newMargin.MarginStr = margin.MarginStr
	newMargin.MarginChar = margin.MarginChar
	newMargin.MarginLength = margin.MarginLength

	return newMargin
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextFieldSpec interface
func (margin MarginSpec) TextTypeName() string {
	return "MarginSpec"
}

// NewPtr - Returns a pointer to a new MarginSpec
// instance.
func (margin MarginSpec) NewPtr() *MarginSpec {

	newMarginField := MarginSpec{}

	return &newMarginField
}


type NewLineSpec struct {
	AddNewLine        bool
}

func (newLine *NewLineSpec) CopyOut() NewLineSpec {
	newNewLine := NewLineSpec{}

	newNewLine.AddNewLine = newLine.AddNewLine

	return newNewLine
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextFieldSpec interface
func (newLine NewLineSpec) TextTypeName() string {
	return "NewLineSpec"
}


// NewPtr - Returns a pointer to a new NewLineSpec
// instance.
func (newLine NewLineSpec) NewPtr() *NewLineSpec {

	newLineSpec := NewLineSpec{}

	return &newLineSpec
}

type NumericIntFieldSpec struct {
	LeftMargin           MarginSpec
	LeftSpacer           MarginSpec
	NumberSpec           IntegerSpec
	RightSpacer          MarginSpec
	TerminateWithNewLine NewLineSpec
}

// CopyOut - Returns a deep copy of the current NumericIntFieldSpec
// instance.
func (numIntField *NumericIntFieldSpec) CopyOut() NumericIntFieldSpec {

	newNumField := NumericIntFieldSpec{}

	newNumField.LeftMargin = numIntField.LeftMargin.CopyOut()
	newNumField.LeftSpacer = numIntField.LeftSpacer.CopyOut()
	newNumField.NumberSpec = numIntField.NumberSpec.CopyOut()
	newNumField.RightSpacer = numIntField.RightSpacer.CopyOut()
	newNumField.TerminateWithNewLine = numIntField.TerminateWithNewLine.CopyOut()

	return newNumField
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextFieldSpec interface
func (numIntField NumericIntFieldSpec) TextTypeName() string {
	return "NumericIntFieldSpec"
}

// NewPtr - Returns a pointer to a new NumericIntFieldSpec instance.
func (numIntField NumericIntFieldSpec) NewPtr() *NumericIntFieldSpec {

	newNumIntSpec := NumericIntFieldSpec{}

	return &newNumIntSpec
}


type StringFieldSpec struct {
	LeftMargin           MarginSpec
	LeftSpacer           MarginSpec
	StrTxtSpec           StringSpec
	RightSpacer          MarginSpec
	TerminateWithNewLine NewLineSpec
}

func (strField *StringFieldSpec) CopyOut() StringFieldSpec {

	newStrField := StringFieldSpec{}

	newStrField.LeftMargin = strField.LeftMargin.CopyOut()
	newStrField.LeftSpacer = strField.LeftSpacer.CopyOut()
	newStrField.StrTxtSpec = strField.StrTxtSpec.CopyOut()
	newStrField.RightSpacer = strField.RightSpacer.CopyOut()
	newStrField.TerminateWithNewLine = strField.TerminateWithNewLine

	return newStrField
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextFieldSpec interface
func (strField StringFieldSpec) TextTypeName() string {
	return "StringFieldSpec"
}


// NewPtr - Returns a pointer to a new StringFieldSpec instance.
func (strField StringFieldSpec) NewPtr() *StringFieldSpec {

	newStrFieldSpec := StringFieldSpec{}

	return &newStrFieldSpec
}


type StringSpec struct {
	StrValue              string
	StrFieldLength        int
	StrPadChar            rune
	StrPosition           FieldPositionSpec
}

// CopyOut - Returns a deep copy of the current StringSpec
// instance.
func (strTxtSpec *StringSpec) CopyOut() StringSpec {

	newStrSpec := StringSpec{}
	newStrSpec.StrValue = strTxtSpec.StrValue
	newStrSpec.StrFieldLength = strTxtSpec.StrFieldLength
	newStrSpec.StrPadChar = strTxtSpec.StrPadChar

	return newStrSpec
}


// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextFieldSpec interface
func (strTxtSpec StringSpec) TextTypeName() string {
	return "StringSpec"
}

// NewPtr - Returns a pointer to a new StringSpec instance.
func (strTxtSpec StringSpec) NewPtr() *StringSpec {

	newStrTxtSpec := StringSpec{}

	return &newStrTxtSpec
}
