package layouts

import (
	"github.com/hsiafan/cocoa/appkit"
)

// FormView ia an appkit.View that arrange form field name and controls.
type FormView struct {
	appkit.View
	stackView appkit.StackView
	rows      []*formRowView

	labelWidth          float64
	labelAlignment      appkit.TextAlignment
	labelControlSpacing float64
	labelFont           appkit.Font
}

// NewFormView create new form view
func NewFormView() *FormView {
	sv := appkit.NewVerticalStackView()
	sv.SetDistribution(appkit.StackViewDistributionFillEqually)
	sv.SetAlignment(appkit.LayoutAttributeBottom)
	sv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	sv.SetHuggingPriority_ForOrientation(appkit.LayoutPriorityRequired, appkit.LayoutConstraintOrientationVertical)
	sv.SetAlignment(appkit.LayoutAttributeLeft)

	return &FormView{
		View:                sv,
		stackView:           sv,
		labelWidth:          0,
		labelAlignment:      appkit.TextAlignmentRight,
		labelControlSpacing: 10,
	}
}

// AddExpandRow add a row, expand to fill parent view height.
func (f *FormView) AddExpandRow() {
	v := appkit.AllocView().Init()
	v.SetTranslatesAutoresizingMaskIntoConstraints(false)
	f.stackView.AddView_InGravity(v, appkit.StackViewGravityTop)
}

// AddRow add a new form row
func (f *FormView) AddRow(name string, control appkit.Control) {
	r := f.newFormRowView(name, control)
	r.SetSpacing(f.labelControlSpacing)
	f.rows = append(f.rows, r)

	f.stackView.AddView_InGravity(r, appkit.StackViewGravityTop)

}

// InsertRow insert a new form row at specific location
func (f *FormView) InsertRow(index uint, name string, control appkit.Control) {
	if index > uint(len(f.rows)) {
		panic("index out of range")
	}

	r := f.newFormRowView(name, control)
	r.SetSpacing(f.labelControlSpacing)
	f.rows = append(append(f.rows[:index], r), f.rows[index:]...)
	f.stackView.InsertView_AtIndex_InGravity(r, index, appkit.StackViewGravityTop)
}

// SetRowSpacing set spacing between rows.
func (f *FormView) SetRowSpacing(spacing float64) {
	f.stackView.SetSpacing(spacing)
}

// SetLabelWidth set width for labels
func (f *FormView) SetLabelWidth(width float64) {
	f.labelWidth = width
	for _, r := range f.rows {
		r.labelWidthConstant.SetConstant(f.labelWidth)
	}
}

// SetLabelAlignment set label text alignment
func (f *FormView) SetLabelAlignment(alignment appkit.TextAlignment) {
	f.labelAlignment = alignment
	for _, r := range f.rows {
		r.label.SetAlignment(alignment)
	}
}

// SetLabelFont set label font
func (f *FormView) SetLabelFont(font appkit.Font) {
	f.labelFont = font
	for _, r := range f.rows {
		r.label.SetFont(font)
	}
}

// SetLabelControlSpacing set spacing between label and control
func (f *FormView) SetLabelControlSpacing(spacing float64) {
	f.labelControlSpacing = spacing
	for _, r := range f.rows {
		r.SetSpacing(spacing)
	}
}

// formRowView is a row of FormView
type formRowView struct {
	appkit.StackView
	name               string
	label              appkit.TextField
	control            appkit.Control
	labelWidthConstant appkit.LayoutConstraint
}

func (f *FormView) newFormRowView(name string, control appkit.Control) *formRowView {
	row := appkit.AllocStackView().Init()
	row.SetTranslatesAutoresizingMaskIntoConstraints(false)
	row.SetOrientation(appkit.UserInterfaceLayoutOrientationHorizontal)
	row.SetDistribution(appkit.StackViewDistributionFillProportionally)
	row.SetAlignment(appkit.LayoutAttributeTop)

	label := appkit.NewLabel(name)
	label.SetTranslatesAutoresizingMaskIntoConstraints(false)
	label.SizeToFit()
	labelFitWidth := label.Bounds().Size.Width
	if f.labelWidth < labelFitWidth {
		f.SetLabelWidth(labelFitWidth)
	}
	labelWidthConstant := label.WidthAnchor().ConstraintEqualToConstant(f.labelWidth)
	labelWidthConstant.SetActive(true)
	label.SetAlignment(f.labelAlignment)
	if f.labelFont != nil && f.labelFont.Ptr() != nil {
		label.SetFont(label.Font())
	}
	row.AddView_InGravity(label, appkit.StackViewGravityLeading)

	control.SetTranslatesAutoresizingMaskIntoConstraints(false)
	row.AddView_InGravity(control, appkit.StackViewGravityLeading)

	row.SetHuggingPriority_ForOrientation(1000, appkit.LayoutConstraintOrientationVertical)

	return &formRowView{
		StackView:          row,
		name:               name,
		label:              label,
		control:            control,
		labelWidthConstant: labelWidthConstant,
	}
}
