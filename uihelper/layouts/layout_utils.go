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

// SetWidthConstraint add and active width constraintEqual
func SetWidthConstraint(view appkit.View, width float64) {
	view.WidthAnchor().ConstraintEqualToConstant(width).SetActive(true)
}

// SetHeightConstraint add and active height constraintEqual
func SetHeightConstraint(view appkit.View, height float64) {
	view.HeightAnchor().ConstraintEqualToConstant(height).SetActive(true)
}
