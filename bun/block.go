package bun

import (
	assets "thdwb/assets"
	gg "thdwb/gg"
	structs "thdwb/structs"
)

func paintBlockElement(ctx *gg.Context, node *structs.NodeDOM) {
	ctx.DrawRectangle(node.RenderBox.Left, node.RenderBox.Top, node.RenderBox.Width, node.RenderBox.Height)
	ctx.SetRGBA(node.Style.BackgroundColor.R, node.Style.BackgroundColor.G, node.Style.BackgroundColor.B, node.Style.BackgroundColor.A)
	ctx.Fill()

	ctx.SetRGBA(node.Style.Color.R, node.Style.Color.G, node.Style.Color.B, node.Style.Color.A)
	ctx.LoadAssetFont(assets.SansSerif(), node.Style.FontSize)
	ctx.DrawStringWrapped(node.Content, node.RenderBox.Left, node.RenderBox.Top+1, 0, 0, node.RenderBox.Width, 1.5, gg.AlignLeft)
	ctx.Fill()
}

func calculateBlockLayout(ctx *gg.Context, node *structs.NodeDOM, childIdx int) {
	if node.Style.Width == 0 {
		node.RenderBox.Width = node.Parent.RenderBox.Width
	}

	if node.Style.Height == 0 {
		ctx.LoadAssetFont(assets.SansSerif(), node.Style.FontSize)
		node.RenderBox.Height = ctx.MeasureStringWrapped(node.Content, node.RenderBox.Width, 1.5) + 2 + ctx.FontHeight()*.5
	}

	if childIdx > 0 {
		prev := node.Parent.Children[childIdx-1]

		if prev.Style.Display != "inline" {
			node.RenderBox.Top = prev.RenderBox.Top + prev.RenderBox.Height
		} else {
			node.RenderBox.Top = prev.RenderBox.Top
		}
	} else {
		node.RenderBox.Top = node.Parent.RenderBox.Top
	}
}
