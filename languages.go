package main

type language struct {
	name       string
	extensions []string
	color      string
}

// https://github.com/doda-zz/github-language-colors/blob/master/colors.json

var SupportedLanguages = []language{
	{
		"Mercury",
		[]string{"m"},
		"#ff2b2b",
	},
	{
		"TypeScript",
		[]string{"ts", "tsx"},
		"#2b7489",
	},
	{
		"PureBasic",
		[]string{"pb", "pbi", "pbf", "pbp", "pbv"},
		"#5a6986",
	},
	{
		"Objective-C++",
		[]string{"mm"},
		"#6866fb",
	},
}
