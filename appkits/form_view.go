package appkits

import (
	"github.com/hsiafan/cocoa/appkit"
)

// FormView ia an appkit.View that arrange form field name and controls.
type FormView interface {
	appkit.View
	// AddExpandRow add a row, expand to fill parent view height.
	AddExpandRow()
	// AddRow add a new form row
	AddRow(name string, control appkit.Control)
	// InsertRow insert a new form row at specific location
	InsertRow(index int, name string, control appkit.Control)
	// SetRowSpacing set spacing between rows.
	SetRowSpacing(spacing float64)
	// SetLabelWidth set width for labels
	SetLabelWidth(width float64)
	// SetLabelAlignment set label text alignment
	SetLabelAlignment(alignment LabelAlignment)
	// SetLabelFont set label font
	SetLabelFont(font appkit.Font)
	// SetLabelControlSpacing set spacing between label and control
	SetLabelControlSpacing(spacing float64)
}

type LabelAlignment int

const (
	LabelAlignmentLeading  = 0
	LabelAlignmentTrailing = 1
	LabelAlignmentCenter   = 2
)

// FormViewImpl implements FormView
type FormViewImpl struct {
	appkit.NSGridView

	labelFont appkit.Font
}

// NewFormView create new form view
func NewFormView() FormView {
	gv := appkit.GridViewWithNumberOfColumns_Rows(2, 0)
	gv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	gv.SetContentHuggingPriority_ForOrientation(appkit.LayoutPriorityDefaultHigh, appkit.LayoutConstraintOrientationHorizontal)
	gv.SetContentHuggingPriority_ForOrientation(appkit.LayoutPriorityDefaultHigh, appkit.LayoutConstraintOrientationVertical)

	return &FormViewImpl{
		NSGridView: gv,
	}
}

func (f *FormViewImpl) AddExpandRow() {
	v := appkit.AllocView().Init()
	v.SetTranslatesAutoresizingMaskIntoConstraints(false)
	f.AddRowWithViews([]appkit.View{NewView(), NewView()})
}

func (f *FormViewImpl) AddRow(name string, control appkit.Control) {
	label := f.newLabel(name)
	f.AddRowWithViews([]appkit.View{label, control})
}

func (f *FormViewImpl) InsertRow(index int, name string, control appkit.Control) {
	if index > f.NumberOfRows() {
		panic("index out of row range")
	}
	label := f.newLabel(name)
	f.InsertRowAtIndex_WithViews(index, []appkit.View{label, control})
}

func (f *FormViewImpl) SetRowSpacing(spacing float64) {
	f.SetRowSpacing(spacing)
}

func (f *FormViewImpl) SetLabelWidth(width float64) {
	f.ColumnAtIndex(0).SetWidth(width)
}

func (f *FormViewImpl) SetLabelAlignment(alignment LabelAlignment) {
	switch alignment {
	case LabelAlignmentLeading:
		f.ColumnAtIndex(0).SetXPlacement(appkit.GridCellPlacementLeading)
	case LabelAlignmentTrailing:
		f.ColumnAtIndex(0).SetXPlacement(appkit.GridCellPlacementTrailing)
	case LabelAlignmentCenter:
		f.ColumnAtIndex(0).SetXPlacement(appkit.GridCellPlacementCenter)
	}
}

func (f *FormViewImpl) SetLabelFont(font appkit.Font) {
	f.labelFont = font
	labelColumn := f.ColumnAtIndex(0)
	size := labelColumn.NumberOfCells()
	for i := 0; i < size; i++ {
		v := labelColumn.CellAtIndex(i).ContentView()
		label := appkit.MakeTextField(v.Ptr())
		label.SetFont(font)
	}
}

func (f *FormViewImpl) SetLabelControlSpacing(spacing float64) {
	f.SetColumnSpacing(spacing)
}

func (f *FormViewImpl) newLabel(name string) appkit.TextField {
	label := NewLabel(name)
	label.SetTranslatesAutoresizingMaskIntoConstraints(false)
	if f.labelFont != nil && f.labelFont.Ptr() != nil {
		label.SetFont(f.labelFont)
	}

	return label
}
