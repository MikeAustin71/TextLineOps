package textlinebuilder



type BlankLinesSpec struct {
	NumBlankLines        int
}

func (blkLine *BlankLinesSpec) CopyOut() BlankLinesSpec {
	newBlankLineSpec := BlankLinesSpec{}

	newBlankLineSpec.NumBlankLines = blkLine.NumBlankLines

	return newBlankLineSpec
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextFieldSpec interface
func (blkLine BlankLinesSpec) TextTypeName() string {
	return "BlankLinesSpec"
}

func (blkLine BlankLinesSpec) NewPtr() *BlankLinesSpec {

	newBlkLine := BlankLinesSpec{}

	return &newBlkLine
}


type LineBreakSpec struct {
	CreateLineBreak      bool
	LeadingBlankLines    BlankLinesSpec
	LeftMargin           MarginSpec
	LeftSpacer           MarginSpec
	LineSpec             LineSpec
	RightSpacer          MarginSpec
	TerminateWithNewLine NewLineSpec
	TrailingBlankLines   BlankLinesSpec
}

func (lineBrk *LineBreakSpec) CopyOut() LineBreakSpec {

	newLineBreak := LineBreakSpec{}

	newLineBreak.CreateLineBreak = lineBrk.CreateLineBreak
	newLineBreak.LeadingBlankLines = lineBrk.LeadingBlankLines.CopyOut()
	newLineBreak.LeftMargin = lineBrk.LeftMargin.CopyOut()
	newLineBreak.LeftSpacer = lineBrk.LeftSpacer.CopyOut()
	newLineBreak.LineSpec = lineBrk.LineSpec.CopyOut()
	newLineBreak.RightSpacer = lineBrk.RightSpacer.CopyOut()
	newLineBreak.TerminateWithNewLine = lineBrk.TerminateWithNewLine.CopyOut()
	newLineBreak.TrailingBlankLines = lineBrk.TrailingBlankLines.CopyOut()

	return newLineBreak
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextFieldSpec interface
func (lineBrk LineBreakSpec) TextTypeName() string {
	return "LineBreakSpec"
}

// NewPtr - Returns a pointer to a new LineBreakSpec
// instance.
func (lineBrk LineBreakSpec) NewPtr() *LineBreakSpec {

	newLineBreak := LineBreakSpec{}

	return &newLineBreak
}


type TwoLabelStrLine struct {
	LeadingBlankLines  BlankLinesSpec
	TopLineBreak       LineBreakSpec
	Label1             StringFieldSpec
	Label2             StringFieldSpec
	BottomLineBreak    LineBreakSpec
	TrailingBlankLines BlankLinesSpec
}

// CopyOut - Returns a deep copy of the current TwoLabelStrLine
// instance.
func (twoLabelLine *TwoLabelStrLine) CopyOut() TwoLabelStrLine {

	newTwoLabelLine := TwoLabelStrLine{}

	newTwoLabelLine.LeadingBlankLines = twoLabelLine.LeadingBlankLines.CopyOut()
	newTwoLabelLine.TopLineBreak = twoLabelLine.TopLineBreak.CopyOut()
	newTwoLabelLine.Label1 = twoLabelLine.Label1.CopyOut()
	newTwoLabelLine.Label2 = twoLabelLine.Label2.CopyOut()
	newTwoLabelLine.BottomLineBreak = twoLabelLine.BottomLineBreak.CopyOut()
	newTwoLabelLine.TrailingBlankLines = twoLabelLine.TrailingBlankLines.CopyOut()

	return newTwoLabelLine
}

// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextFieldSpec interface
func (twoLabelLine TwoLabelStrLine) TextTypeName() string {
	return "TwoLabelStrLine"
}

// NewPtr - Returns a pointer to a new TwoLabelStrLine instance.
func (twoLabelLine TwoLabelStrLine) NewPtr() *TwoLabelStrLine {

	newTwoLabelLine := TwoLabelStrLine{}

	return &newTwoLabelLine
}


type OneLabelOneIntLine struct {
	LeadingBlankLines    BlankLinesSpec
	TopLineBreak         LineBreakSpec
	Label1               StringFieldSpec
	Number1              NumericIntFieldSpec
	BottomLineBreak      LineBreakSpec
	TrailingBlankLines   BlankLinesSpec
}

// CopyOut - Returns a deep copy of the current OneLabelOneIntLine
// instance.
func (labelIntLine *OneLabelOneIntLine) CopyOut() OneLabelOneIntLine {

	newLabelInt := OneLabelOneIntLine{}

	newLabelInt.LeadingBlankLines = labelIntLine.LeadingBlankLines
	newLabelInt.TopLineBreak = labelIntLine.TopLineBreak.CopyOut()
	newLabelInt.Label1 = labelIntLine.Label1.CopyOut()
	newLabelInt.Number1 = labelIntLine.Number1.CopyOut()
	newLabelInt.BottomLineBreak = labelIntLine.BottomLineBreak.CopyOut()
	newLabelInt.TrailingBlankLines = labelIntLine.TrailingBlankLines
	return newLabelInt
}


// TextTypeName - returns a string specifying the type
// of Text Field specification. This method fulfills
// requirements of TextFieldSpec interface
func (labelIntLine OneLabelOneIntLine) TextTypeName() string {
	return "OneLabelOneIntLine"
}

// NewPtr - Returns a pointer to a new OneLabelOneIntLine instance.
func (labelIntLine OneLabelOneIntLine) NewPtr() *OneLabelOneIntLine {

	newOneLabelOneIntLine := OneLabelOneIntLine{}

	return &newOneLabelOneIntLine
}
