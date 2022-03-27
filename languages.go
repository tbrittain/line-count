package main

type Language struct {
	Name  string
	Color string
}

var typeScript = Language{
	"TypeScript",
	"#2b7489",
}

var objectiveC = Language{
	"Objective-C",
	"#6866fb",
}

var dart = Language{
	"Dart",
	"#00B4AB",
}

var shell = Language{
	"Shell",
	"#89e051",
}

var elixir = Language{
	"Elixir",
	"#6e4a7e",
}

var kotlin = Language{
	"Kotlin",
	"#F18E33",
}

var ruby = Language{
	"Ruby",
	"#701516",
}

var golang = Language{
	"Go",
	"#375eab",
}

var visualBasic = Language{
	"Visual Basic",
	"#945db7",
}

var php = Language{
	"PHP",
	"#4F5D95",
}

var java = Language{
	"Java",
	"#b07219",
}

var scala = Language{
	"Scala",
	"#c22d40",
}

var makefile = Language{
	"Makefile",
	"#427819",
}

var batchfile = Language{
	"Batchfile",
	"#c1f12e",
}

var perl = Language{
	"Perl",
	"#0298c3",
}

var lua = Language{
	"Lua",
	"#000080",
}

var coffeeScript = Language{
	"CoffeeScript",
	"244776",
}

var html = Language{
	"HTML",
	"#e34c26",
}

var swift = Language{
	"Swift",
	"#ffac45",
}

var c = Language{
	"C",
	"#555",
}

var clojure = Language{
	"Clojure",
	"#db5855",
}

var rust = Language{
	"Rust",
	"#dea584",
}

var cSharp = Language{
	"C#",
	"#178600",
}

var css = Language{
	"CSS",
	"#563d7c",
}

var fSharp = Language{
	"F#",
	"#b845fc",
}

var smalltalk = Language{
	"Smalltalk",
	"#596706",
}

var javaScript = Language{
	"JavaScript",
	"#f1e05a",
}

var r = Language{
	"R",
	"#198ce7",
}

var erlang = Language{
	"Erlang",
	"#949e0e",
}

var python = Language{
	"Python",
	"#3572A5",
}

var haskell = Language{
	"Haskell",
	"#29b544",
}

var cpp = Language{
	"C++",
	"#f34b7d",
}

var scss = Language{
	"SCSS",
	"#c6538c",
}

var LanguageMap = map[string]Language{
	".ts":       typeScript,
	".tsx":      typeScript,
	".m":        objectiveC,
	".mm":       objectiveC,
	".dart":     dart,
	".sh":       shell,
	".bash":     shell,
	".zsh":      shell,
	".ksh":      shell,
	".ex":       elixir,
	".exs":      elixir,
	".kt":       kotlin,
	".kts":      kotlin,
	".ktm":      kotlin,
	".rb":       ruby,
	".go":       golang,
	".vb":       visualBasic,
	".php":      php,
	".java":     java,
	".scala":    scala,
	".sc":       scala,
	".makefile": makefile,
	".mk":       makefile,
	".mak":      makefile,
	".bat":      batchfile,
	".pl":       perl,
	".plx":      perl,
	".pm":       perl,
	".lua":      lua,
	".coffee":   coffeeScript,
	".html":     html,
	".htm":      html,
	".swift":    swift,
	".c":        c,
	".h":        c,
	".clj":      clojure,
	".cljs":     clojure,
	".cljc":     clojure,
	".rs":       rust,
	".cs":       cSharp,
	".css":      css,
	".fs":       fSharp,
	".fsi":      fSharp,
	".fsx":      fSharp,
	".fsscript": fSharp,
	".st":       smalltalk,
	".js":       javaScript,
	".jsx":      javaScript,
	".r":        r,
	".erl":      erlang,
	".hrl":      erlang,
	".py":       python,
	".pyd":      python,
	".pyi":      python,
	".hs":       haskell,
	".lhs":      haskell,
	".cpp":      cpp,
	".cc":       cpp,
	".cxx":      cpp,
	".hh":       cpp,
	".hpp":      cpp,
	".hxx":      cpp,
	".scss":     scss,
	".sass":     scss,
}
