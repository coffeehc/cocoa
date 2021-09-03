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

// UseSameWidth set width constant for multi controls, using the max width for width
func UseSameWidth(controls ...appkit.Control) {
	var maxWidth float64
	for _, c := range controls {
		c.SizeToFit()
		width := c.Bounds().Size.Width
		if maxWidth < width {
			maxWidth = width
		}
	}
	for _, c := range controls {
		c.WidthAnchor().ConstraintEqualToConstant(maxWidth).SetActive(true)
	}
}

// UseSameHeight set height constant for multi controls, using the max height for height
func UseSameHeight(controls ...appkit.Control) {
	var maxHeight float64
	for _, c := range controls {
		c.SizeToFit()
		height := c.Bounds().Size.Height
		if maxHeight < height {
			maxHeight = height
		}
	}
	for _, c := range controls {
		c.HeightAnchor().ConstraintEqualToConstant(maxHeight).SetActive(true)
	}
}
