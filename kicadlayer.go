package gokicadlib

type KicadLayer int
type KicadLayerSlice []KicadLayer

const (
	Edge_Cuts KicadLayer = iota
	Margin
	Eco2_User
	Eco1_User
	Cmts_User
	Dwgs_User
	F_Fab
	F_CrtYd
	F_Adhes
	F_SilkS
	F_Paste
	F_Mask
	F_Cu
	In1_Cu
	In2_Cu
	In3_Cu
	In4_Cu
	In5_Cu
	In6_Cu
	In7_Cu
	In8_Cu
	In9_Cu
	In10_Cu
	In11_Cu
	In12_Cu
	In13_Cu
	In14_Cu
	In15_Cu
	In16_Cu
	In17_Cu
	In18_Cu
	In19_Cu
	In20_Cu
	In21_Cu
	In22_Cu
	In23_Cu
	In24_Cu
	In25_Cu
	In26_Cu
	In27_Cu
	In28_Cu
	In29_Cu
	In30_Cu
	B_Cu
	B_Mask
	B_Paste
	B_SilkS
	B_Adhes
	B_CrtYd
	B_Fab
)
