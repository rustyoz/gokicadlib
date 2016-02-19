package gokicadlib

type Layer string
type LayerSlice []Layer

const (
	Edge_Cuts Layer = "Edge.Cuts"
	Margin          = "Margin"
	Eco2_User       = "Eco2.User"
	Eco1_User       = "Eco1.User"
	Cmts_User       = "Cmts.User"
	Dwgs_User       = "Dwgs.User"
	F_Fab           = "F.Fab"
	F_CrtYd         = "F.CrtYd"
	F_Adhes         = "F.Adhes"
	F_SilkS         = "F.SilkS"
	F_Paste         = "F.Paste"
	F_Mask          = "F.Mask"
	F_Cu            = "F.Cu"
	A_Cu            = "*.Cu"
	A_Mask          = "*.Mask"
	In1_Cu          = "In1.Cu"
	In2_Cu          = "In2.Cu"
	In3_Cu          = "In3.Cu"
	In4_Cu          = "In4.Cu"
	In5_Cu          = "In5.Cu"
	In6_Cu          = "In6.Cu"
	In7_Cu          = "In7.Cu"
	In8_Cu          = "In8.Cu"
	In9_Cu          = "In9.Cu"
	In10_Cu         = "In10.Cu"
	In11_Cu         = "In11.Cu"
	In12_Cu         = "In12.Cu"
	In13_Cu         = "In13.Cu"
	In14_Cu         = "In14.Cu"
	In15_Cu         = "In15.Cu"
	In16_Cu         = "In16.Cu"
	In17_Cu         = "In17.Cu"
	In18_Cu         = "In18.Cu"
	In19_Cu         = "In19.Cu"
	In20_Cu         = "In20.Cu"
	In21_Cu         = "In21.Cu"
	In22_Cu         = "In22.Cu"
	In23_Cu         = "In23.Cu"
	In24_Cu         = "In24.Cu"
	In25_Cu         = "In25.Cu"
	In26_Cu         = "In26.Cu"
	In27_Cu         = "In27.Cu"
	In28_Cu         = "In28.Cu"
	In29_Cu         = "In29.Cu"
	In30_Cu         = "In30.Cu"
	B_Cu            = "B.Cu"
	B_Mask          = "B.Mask"
	B_Paste         = "B.Paste"
	B_SilkS         = "B.SilkS"
	B_Adhes         = "B.Adhes"
	B_CrtYd         = "B.CrtYd"
	B_Fab           = "B.Fab"
)

func (ls LayerSlice) String() []string {
	var ss []string
	for _, l := range ls {
		ss = append(ss, string(l))
	}
	return ss
}
