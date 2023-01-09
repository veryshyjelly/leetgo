package lang

var (
	golangGen = golang{
		baseLang{
			name:              "Go",
			slug:              "golang",
			shortName:         "go",
			extension:         ".go",
			lineComment:       "//",
			blockCommentStart: "/*",
			blockCommentEnd:   "*/",
		},
	}
	python3Gen = python{
		baseLang{
			name:              "Python",
			slug:              "python3",
			shortName:         "py",
			extension:         ".py",
			lineComment:       "#",
			blockCommentStart: `"""`,
			blockCommentEnd:   `"""`,
		},
	}
	cppGen = baseLang{
		name:              "C++",
		slug:              "cpp",
		shortName:         "cpp",
		extension:         ".cpp",
		lineComment:       "//",
		blockCommentStart: "/*",
		blockCommentEnd:   "*/",
	}
	rustGen = baseLang{
		name:              "Rust",
		slug:              "rust",
		shortName:         "rs",
		extension:         ".rs",
		lineComment:       "//",
		blockCommentStart: "/*",
		blockCommentEnd:   "*/",
	}
	javaGen = baseLang{
		name:              "Java",
		slug:              "java",
		shortName:         "java",
		extension:         ".java",
		lineComment:       "//",
		blockCommentStart: "/*",
		blockCommentEnd:   "*/",
	}
	cGen = baseLang{
		name:              "C",
		slug:              "c",
		shortName:         "c",
		extension:         ".c",
		lineComment:       "//",
		blockCommentStart: "/*",
		blockCommentEnd:   "*/",
	}
	csharpGen = baseLang{
		name:              "C#",
		slug:              "csharp",
		shortName:         "cs",
		extension:         ".cs",
		lineComment:       "//",
		blockCommentStart: "/*",
		blockCommentEnd:   "*/",
	}
	jsGen = baseLang{
		name:              "JavaScript",
		slug:              "javascript",
		shortName:         "js",
		extension:         ".js",
		lineComment:       "//",
		blockCommentStart: "/*",
		blockCommentEnd:   "*/",
	}
	phpGen = baseLang{
		name:              "PHP",
		slug:              "php",
		shortName:         "php",
		extension:         ".php",
		lineComment:       "//",
		blockCommentStart: "/*",
		blockCommentEnd:   "*/",
	}
	rubyGen = baseLang{
		name:              "Ruby",
		slug:              "ruby",
		shortName:         "rb",
		extension:         ".rb",
		lineComment:       "#",
		blockCommentStart: "=begin",
		blockCommentEnd:   "=end",
	}
	swiftGen = baseLang{
		name:              "Swift",
		slug:              "swift",
		shortName:         "swift",
		extension:         ".swift",
		lineComment:       "//",
		blockCommentStart: "/*",
		blockCommentEnd:   "*/",
	}
	kotlinGen = baseLang{
		name:              "Kotlin",
		slug:              "kotlin",
		shortName:         "kt",
		extension:         ".kt",
		lineComment:       "//",
		blockCommentStart: "/*",
		blockCommentEnd:   "*/",
	}
	// TODO scala, typescript, erlang, dart, racket
	SupportedLangs = []Lang{
		golangGen,
		python3Gen,
		cppGen,
		rustGen,
		javaGen,
		jsGen,
		phpGen,
		cGen,
		csharpGen,
		rubyGen,
		swiftGen,
		kotlinGen,
	}
)
