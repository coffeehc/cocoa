package layouts

import "github.com/hsiafan/cocoa/appkit"

// AddViewWithPadding add the only one sub view to parent view, with paddings.
func AddViewWithPadding(view appkit.View, subView appkit.View, left, top, right, bottom float64) {
	view.AddSubview(subView)
	view.LeadingAnchor().ConstraintEqualToAnchor_Constant(subView.LeadingAnchor(), -left).SetActive(true)
	view.TopAnchor().ConstraintEqualToAnchor_Constant(subView.TopAnchor(), -top).SetActive(true)
	view.TrailingAnchor().ConstraintEqualToAnchor_Constant(subView.TrailingAnchor(), right).SetActive(true)
	view.BottomAnchor().ConstraintEqualToAnchor_Constant(subView.BottomAnchor(), bottom).SetActive(true)
}

// SetPreferredWidth add and active width constraintEqual
func SetPreferredWidth(view appkit.View, width float64) {
	view.WidthAnchor().ConstraintEqualToConstant(width).SetActive(true)
}

// SetPreferredHeight add and active height constraintEqual
func SetPreferredHeight(view appkit.View, height float64) {
	view.HeightAnchor().ConstraintEqualToConstant(height).SetActive(true)
}

func SetPreferredMaxWidth(view appkit.View, width float64) {
	view.WidthAnchor().ConstraintLessThanOrEqualToConstant(width).SetActive(true)
}

func SetPreferredMaxHeight(view appkit.View, height float64) {
	view.HeightAnchor().ConstraintLessThanOrEqualToConstant(height).SetActive(true)
}

func SetPreferredMinWidth(view appkit.View, width float64) {
	view.WidthAnchor().ConstraintGreaterThanOrEqualToConstant(width).SetActive(true)
}

func SetPreferredMinHeight(view appkit.View, height float64) {
	view.HeightAnchor().ConstraintGreaterThanOrEqualToConstant(height).SetActive(true)
}
