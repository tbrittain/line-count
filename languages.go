package main

type language struct {
	name       string
	extensions []string
	color      string
}

var SupportedLanguages = []language{
	{
		"TypeScript",
		[]string{"ts", "tsx"},
		"#2b7489",
	},
	{
		"Objective-C",
		[]string{"m, mm"},
		"#6866fb",
	},
	{
		"Dart",
		[]string{"dart"},
		"#00B4AB",
	},
	{
		"Shell",
		[]string{"sh", "bash", "zsh", "ksh", "fish"},
		"#89e051",
	},
	{
		"Elixir",
		[]string{"ex", "exs"},
		"#6e4a7e",
	},
	{
		"Kotlin",
		[]string{"kt", "kts", "ktm"},
		"#F18E33",
	},
	{
		"Ruby",
		[]string{"rb"},
		"#701516",
	},
	{
		"Go",
		[]string{"go"},
		"#375eab",
	},
	{
		"Visual Basic",
		[]string{"vb"},
		"#945db7",
	},
	{
		"PHP",
		[]string{"php"},
		"#4F5D95",
	},
	{
		"Java",
		[]string{"java"},
		"#b07219",
	},
	{
		"Scala",
		[]string{"scala", "sc"},
		"#c22d40",
	},
	{
		"Makefile",
		[]string{"makefile", "mk", "mak"},
		"#427819",
	},
	{
		"Perl",
		[]string{"pl", "plx", "pm", "t", "xs"},
		"#0298c3",
	},
	{
		"Lua",
		[]string{"lua"},
		"#000080",
	},
	{
		"CoffeeScript",
		[]string{"coffee"},
		"#244776",
	},
	{
		"HTML",
		[]string{"html", "htm"},
		"#e34c26",
	},
	{
		"Swift",
		[]string{"swift"},
		"#ffac45",
	},
	{
		"C",
		[]string{"c", "h"},
		"#555",
	},
	{
		"Clojure",
		[]string{"clj", "cljs", "cljc"},
		"#db5855",
	},
	{
		"Rust",
		[]string{"rs"},
		"#dea584",
	},
	{
		"C#",
		[]string{"cs"},
		"#178600",
	},
	{
		"CSS",
		[]string{"css"},
		"#563d7c",
	},
	{
		"F#",
		[]string{"fs", "fsi", "fsx", "fsscript"},
		"#b845fc",
	},
	{
		"Smalltalk",
		[]string{"st"},
		"#596706",
	},
	{
		"JavaScript",
		[]string{"js", "jsx"},
		"#f1e05a",
	},
	{
		"R",
		[]string{"r"},
		"#198ce7",
	},
	{
		"Erlang",
		[]string{"erl", "hrl"},
		"#949e0e",
	},
	{
		"Python",
		[]string{"py", "pyc", "pyd", "pyi", "pyo", "pyx"},
		"#3572A5",
	},
	{
		"Haskell",
		[]string{"hs", "lhs"},
		"#29b544",
	},
	{
		"C++",
		[]string{"cpp", "cc", "cxx", "hh", "hpp", "hxx"},
		"#f34b7d",
	},
	{
		"SCSS",
		[]string{"scss", "sass"},
		"#c6538c",
	},
}
