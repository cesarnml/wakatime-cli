package heartbeat_test

import (
	"encoding/json"
	"fmt"
	"regexp"
	"testing"

	"github.com/wakatime/wakatime-cli/pkg/heartbeat"

	"github.com/alecthomas/chroma/lexers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func languageTests() map[string]heartbeat.Language {
	return map[string]heartbeat.Language{
		"ABAP":                        heartbeat.LanguageABAP,
		"ABNF":                        heartbeat.LanguageABNF,
		"ActionScript":                heartbeat.LanguageActionScript,
		"ActionScript 3":              heartbeat.LanguageActionScript3,
		"Ada":                         heartbeat.LanguageAda,
		"ADL":                         heartbeat.LanguageADL,
		"AdvPL":                       heartbeat.LanguageAdvPL,
		"Agda":                        heartbeat.LanguageAgda,
		"Aheui":                       heartbeat.LanguageAheui,
		"Alloy":                       heartbeat.LanguageAlloy,
		"AmbientTalk":                 heartbeat.LanguageAmbientTalk,
		"Ampl":                        heartbeat.LanguageAmpl,
		"Angular2":                    heartbeat.LanguageAngular2,
		"Ansible":                     heartbeat.LanguageAnsible,
		"ANTLR":                       heartbeat.LanguageANTLR,
		"APL":                         heartbeat.LanguageAPL,
		"AppleScript":                 heartbeat.LanguageAppleScript,
		"Apache Config":               heartbeat.LanguageApacheConfig,
		"Apex":                        heartbeat.LanguageApex,
		"Arc":                         heartbeat.LanguageArc,
		"Arduino":                     heartbeat.LanguageArduino,
		"Arrow":                       heartbeat.LanguageArrow,
		"ASP Classic":                 heartbeat.LanguageASPClassic,
		"ASP.NET":                     heartbeat.LanguageASPDotNet,
		"AspectJ":                     heartbeat.LanguageAspectJ,
		"aspx-cs":                     heartbeat.LanguageAspxCSharp,
		"aspx-vb":                     heartbeat.LanguageAspxVBNet,
		"Assembly":                    heartbeat.LanguageAssembly,
		"Asymptote":                   heartbeat.LanguageAsymptote,
		"Augeas":                      heartbeat.LanguageAugeas,
		"Autoconf":                    heartbeat.LanguageAutoconf,
		"AutoHotkey":                  heartbeat.LanguageAutoHotkey,
		"AutoIt":                      heartbeat.LanguageAutoIt,
		"AWK":                         heartbeat.LanguageAwk,
		"Ballerina":                   heartbeat.LanguageBallerina,
		"BARE":                        heartbeat.LanguageBARE,
		"Bash":                        heartbeat.LanguageBash,
		"Bash Session":                heartbeat.LanguageBashSession,
		"Batchfile":                   heartbeat.LanguageBatchfile,
		"Basic":                       heartbeat.LanguageBasic,
		"Batch Script":                heartbeat.LanguageBatchScript,
		"BBC Basic":                   heartbeat.LanguageBBCBasic,
		"BBCode":                      heartbeat.LanguageBBCode,
		"BC":                          heartbeat.LanguageBC,
		"Befunge":                     heartbeat.LanguageBefunge,
		"BibTeX":                      heartbeat.LanguageBibTeX,
		"Blade Template":              heartbeat.LanguageBladeTemplate,
		"Blazor":                      heartbeat.LanguageBlazor,
		"BlitzBasic":                  heartbeat.LanguageBlitzBasic,
		"BlitzMax":                    heartbeat.LanguageBlitzMax,
		"BNF":                         heartbeat.LanguageBNF,
		"Boa":                         heartbeat.LanguageBoa,
		"Boo":                         heartbeat.LanguageBoo,
		"Boogie":                      heartbeat.LanguageBoogie,
		"Brainfuck":                   heartbeat.LanguageBrainfuck,
		"BrightScript":                heartbeat.LanguageBrightScript,
		"Bro":                         heartbeat.LanguageBro,
		"BST":                         heartbeat.LanguageBST,
		"BUGS":                        heartbeat.LanguageBUGS,
		"C":                           heartbeat.LanguageC,
		"C++":                         heartbeat.LanguageCPP,
		"C#":                          heartbeat.LanguageCSharp,
		"ca65 assembler":              heartbeat.LanguageCa65Assembler,
		"Caddyfile":                   heartbeat.LanguageCaddyfile,
		"Caddyfile Directives":        heartbeat.LanguageCaddyfileDirectives,
		"cADL":                        heartbeat.LanguageCADL,
		"CAmkES":                      heartbeat.LanguageCAmkES,
		"CapDL":                       heartbeat.LanguageCapDL,
		"Cap'n Proto":                 heartbeat.LanguageCapNProto,
		"Cassandra CQL":               heartbeat.LanguageCassandraCQL,
		"CBM BASIC V2":                heartbeat.LanguageCBMBasicV2,
		"Ceylon":                      heartbeat.LanguageCeylon,
		"CFEngine3":                   heartbeat.LanguageCFEngine3,
		"cfstatement":                 heartbeat.LanguageCfstatement,
		"ChaiScript":                  heartbeat.LanguageChaiScript,
		"Chapel":                      heartbeat.LanguageChapel,
		"Charmci":                     heartbeat.LanguageCharmci,
		"Cheetah":                     heartbeat.LanguageCheetah,
		"Cirru":                       heartbeat.LanguageCirru,
		"Clay":                        heartbeat.LanguageClay,
		"Clean":                       heartbeat.LanguageClean,
		"Clojure":                     heartbeat.LanguageClojure,
		"ClojureScript":               heartbeat.LanguageClojureScript,
		"c-objdump":                   heartbeat.LanguageCObjdump,
		"CMake":                       heartbeat.LanguageCMake,
		"COBOL":                       heartbeat.LanguageCOBOL,
		"COBOLFree":                   heartbeat.LanguageCOBOLFree,
		"Cocoa":                       heartbeat.LanguageCocoa,
		"CoffeeScript":                heartbeat.LanguageCoffeeScript,
		"Coldfusion":                  heartbeat.LanguageColdfusionHTML,
		"Coldfusion CFC":              heartbeat.LanguageColdfusionCFC,
		"Component Pascal":            heartbeat.LanguageComponentPascal,
		"Common Lisp":                 heartbeat.LanguageCommonLisp,
		"Coq":                         heartbeat.LanguageCoq,
		"cperl":                       heartbeat.LanguageCPerl,
		"cpp-objdump":                 heartbeat.LanguageCppObjdump,
		"CPSA":                        heartbeat.LanguageCPSA,
		"Crmsh":                       heartbeat.LanguageCrmsh,
		"Croc":                        heartbeat.LanguageCroc,
		"Crontab":                     heartbeat.LanguageCrontab,
		"Cryptol":                     heartbeat.LanguageCryptol,
		"Crystal":                     heartbeat.LanguageCrystal,
		"CSON":                        heartbeat.LanguageCSON,
		"Csound Document":             heartbeat.LanguageCsoundDocument,
		"Csound Orchestra":            heartbeat.LanguageCsoundOrchestra,
		"Csound Score":                heartbeat.LanguageCsoundScore,
		"CSHTML":                      heartbeat.LanguageCSHTML,
		"CSS":                         heartbeat.LanguageCSS,
		"CSV":                         heartbeat.LanguageCSV,
		"CUDA":                        heartbeat.LanguageCUDA,
		"CVS":                         heartbeat.LanguageCVS,
		"Cypher":                      heartbeat.LanguageCypher,
		"Cython":                      heartbeat.LanguageCython,
		"D":                           heartbeat.LanguageD,
		"d-objdump":                   heartbeat.LanguageDObjdump,
		"Darcs Patch":                 heartbeat.LanguageDarcsPatch,
		"Dart":                        heartbeat.LanguageDart,
		"DASM16":                      heartbeat.LanguageDASM16,
		"DCL":                         heartbeat.LanguageDCL,
		"DCPU-16 ASM":                 heartbeat.LanguageDCPU16Asm,
		"Debian Control file":         heartbeat.LanguageDebianControlFile,
		"Delphi":                      heartbeat.LanguageDelphi,
		"Devicetree":                  heartbeat.LanguageDevicetree,
		"dg":                          heartbeat.LanguageDG,
		"Dhall":                       heartbeat.LanguageDhall,
		"Diff":                        heartbeat.LanguageDiff,
		"Django/Jinja":                heartbeat.LanguageDjangoJinja,
		"Docker":                      heartbeat.LanguageDocker,
		"DTD":                         heartbeat.LanguageDTD,
		"DocTeX":                      heartbeat.LanguageDocTeX,
		"Duel":                        heartbeat.LanguageDuel,
		"Dylan":                       heartbeat.LanguageDylan,
		"DylanLID":                    heartbeat.LanguageDylanLID,
		"Dylan session":               heartbeat.LanguageDylanSession,
		"DynASM":                      heartbeat.LanguageDynASM,
		"E-mail":                      heartbeat.LanguageEMail,
		"Earl Grey":                   heartbeat.LanguageEarlGrey,
		"Easytrieve":                  heartbeat.LanguageEasytrieve,
		"EBNF":                        heartbeat.LanguageEBNF,
		"eC":                          heartbeat.LanguageEC,
		"ECL":                         heartbeat.LanguageECL,
		"Eiffel":                      heartbeat.LanguageEiffel,
		"EJS":                         heartbeat.LanguageEJS,
		"Elixir":                      heartbeat.LanguageElixir,
		"Elixir iex session":          heartbeat.LanguageElixirIexSession,
		"Elm":                         heartbeat.LanguageElm,
		"Emacs Lisp":                  heartbeat.LanguageEmacsLisp,
		"ERB":                         heartbeat.LanguageERB,
		"Erlang":                      heartbeat.LanguageErlang,
		"Erlang erl session":          heartbeat.LanguageErlangErlSession,
		"Eshell":                      heartbeat.LanguageEshell,
		"Evoque":                      heartbeat.LanguageEvoque,
		"execline":                    heartbeat.LanguageExecline,
		"Ezhil":                       heartbeat.LanguageEzhil,
		"F#":                          heartbeat.LanguageFSharp,
		"Factor":                      heartbeat.LanguageFactor,
		"Fancy":                       heartbeat.LanguageFancy,
		"Fantom":                      heartbeat.LanguageFantom,
		"Felix":                       heartbeat.LanguageFelix,
		"Fennel":                      heartbeat.LanguageFennel,
		"Fish":                        heartbeat.LanguageFish,
		"Flatline":                    heartbeat.LanguageFlatline,
		"FloScript":                   heartbeat.LanguageFloScript,
		"Font":                        heartbeat.LanguageFont,
		"Forth":                       heartbeat.LanguageForth,
		"Fortran":                     heartbeat.LanguageFortran,
		"FortranFixed":                heartbeat.LanguageFortranFixed,
		"FoxPro":                      heartbeat.LanguageFoxPro,
		"Freefem":                     heartbeat.LanguageFreefem,
		"FStar":                       heartbeat.LanguageFStar,
		"GAP":                         heartbeat.LanguageGap,
		"GAS":                         heartbeat.LanguageGas,
		"GDNative":                    heartbeat.LanguageGDNative,
		"GDScript":                    heartbeat.LanguageGDScript,
		"Genshi":                      heartbeat.LanguageGenshi,
		"Genshi HTML":                 heartbeat.LanguageGenshiHTML,
		"Genshi Text":                 heartbeat.LanguageGenshiText,
		"Gettext Catalog":             heartbeat.LanguageGettextCatalog,
		"Gherkin":                     heartbeat.LanguageGherkin,
		"Git":                         heartbeat.LanguageGit,
		"Git Config":                  heartbeat.LanguageGitConfig,
		"GLSL":                        heartbeat.LanguageGLSL,
		"Gnuplot":                     heartbeat.LanguageGnuplot,
		"Go":                          heartbeat.LanguageGo,
		"Golo":                        heartbeat.LanguageGolo,
		"GoodData-CL":                 heartbeat.LanguageGoodDataCL,
		"Gosu":                        heartbeat.LanguageGosu,
		"Gosu Template":               heartbeat.LanguageGosuTemplate,
		"Gradle Config":               heartbeat.LanguageGradleConfig,
		"GraphQL":                     heartbeat.LanguageGraphQL,
		"Groff":                       heartbeat.LanguageGroff,
		"Groovy":                      heartbeat.LanguageGroovy,
		"Haml":                        heartbeat.LanguageHaml,
		"Handlebars":                  heartbeat.LanguageHandlebars,
		"Haskell":                     heartbeat.LanguageHaskell,
		"Haxe":                        heartbeat.LanguageHaxe,
		"HCL":                         heartbeat.LanguageHCL,
		"Hexdump":                     heartbeat.LanguageHexdump,
		"HLB":                         heartbeat.LanguageHLB,
		"HLSL":                        heartbeat.LanguageHLSL,
		"HSAIL":                       heartbeat.LanguageHSAIL,
		"Hspec":                       heartbeat.LanguageHspec,
		"HTML":                        heartbeat.LanguageHTML,
		"HTTP":                        heartbeat.LanguageHTTP,
		"Hxml":                        heartbeat.LanguageHxml,
		"Hy":                          heartbeat.LanguageHy,
		"Hybris":                      heartbeat.LanguageHybris,
		"Icon":                        heartbeat.LanguageIcon,
		"IDL":                         heartbeat.LanguageIDL,
		"Idris":                       heartbeat.LanguageIdris,
		"Igor":                        heartbeat.LanguageIgor,
		"Inform 6":                    heartbeat.LanguageInform6,
		"Inform 6 template":           heartbeat.LanguageInform6Template,
		"Inform 7":                    heartbeat.LanguageInform7,
		"Image (jpeg)":                heartbeat.LanguageImageJPEG,
		"Image (png)":                 heartbeat.LanguageImagePNG,
		"INI":                         heartbeat.LanguageINI,
		"Io":                          heartbeat.LanguageIo,
		"Ioke":                        heartbeat.LanguageIoke,
		"IRC Logs":                    heartbeat.LanguageIRCLogs,
		"Isabelle":                    heartbeat.LanguageIsabelle,
		"J":                           heartbeat.LanguageJ,
		"Jade":                        heartbeat.LanguageJade,
		"JAGS":                        heartbeat.LanguageJAGS,
		"Jasmin":                      heartbeat.LanguageJasmin,
		"Java":                        heartbeat.LanguageJava,
		"JavaScript":                  heartbeat.LanguageJavaScript,
		"JCL":                         heartbeat.LanguageJCL,
		"JSGF":                        heartbeat.LanguageJSGF,
		"JSON":                        heartbeat.LanguageJSON,
		"JSON-LD":                     heartbeat.LanguageJSONLD,
		"Java Server Page":            heartbeat.LanguageJSP,
		"JSX":                         heartbeat.LanguageJSX,
		"Julia":                       heartbeat.LanguageJulia,
		"Julia console":               heartbeat.LanguageJuliaConsole,
		"Jungle":                      heartbeat.LanguageJungle,
		"Juttle":                      heartbeat.LanguageJuttle,
		"Kal":                         heartbeat.LanguageKal,
		"Kconfig":                     heartbeat.LanguageKconfig,
		"Kernel log":                  heartbeat.LanguageKernelLog,
		"Koka":                        heartbeat.LanguageKoka,
		"Kotlin":                      heartbeat.LanguageKotlin,
		"Laravel Template":            heartbeat.LanguageLaravelTemplate,
		"Lasso":                       heartbeat.LanguageLasso,
		"LaTeX":                       heartbeat.LanguageLaTeX,
		"Latte":                       heartbeat.LanguageLatte,
		"Lean":                        heartbeat.LanguageLean,
		"LESS":                        heartbeat.LanguageLess,
		"Lighttpd configuration file": heartbeat.LanguageLighttpd,
		"Limbo":                       heartbeat.LanguageLimbo,
		"Linker Script":               heartbeat.LanguageLinkerScript,
		"Liquid":                      heartbeat.LanguageLiquid,
		"Literate Agda":               heartbeat.LanguageLiterateAgda,
		"Literate Cryptol":            heartbeat.LanguageLiterateCryptol,
		"Literate Haskell":            heartbeat.LanguageLiterateHaskell,
		"Literate Idris":              heartbeat.LanguageLiterateIdris,
		"LiveScript":                  heartbeat.LanguageLiveScript,
		"LLVM":                        heartbeat.LanguageLLVM,
		"LLVM-MIR":                    heartbeat.LanguageLLVMMIR,
		"LLVM-MIR Body":               heartbeat.LanguageLLVMMIRBody,
		"Log File":                    heartbeat.LanguageLogFile,
		"Logos":                       heartbeat.LanguageLogos,
		"Logtalk":                     heartbeat.LanguageLogtalk,
		"LSL":                         heartbeat.LanguageLSL,
		"Lua":                         heartbeat.LanguageLua,
		"Makefile":                    heartbeat.LanguageMakefile,
		"Mako":                        heartbeat.LanguageMako,
		"Man":                         heartbeat.LanguageMan,
		"MAQL":                        heartbeat.LanguageMAQL,
		"Markdown":                    heartbeat.LanguageMarkdown,
		"Marko":                       heartbeat.LanguageMarko,
		"Mask":                        heartbeat.LanguageMask,
		"Mason":                       heartbeat.LanguageMason,
		"Mathematica":                 heartbeat.LanguageMathematica,
		"Matlab":                      heartbeat.LanguageMatlab,
		"Matlab session":              heartbeat.LanguageMatlabSession,
		"Max":                         heartbeat.LanguageMax,
		"Max/MSP":                     heartbeat.LanguageMaxMSP,
		"Meson":                       heartbeat.LanguageMeson,
		"Metafont":                    heartbeat.LanguageMetafont,
		"Metapost":                    heartbeat.LanguageMetapost,
		"MIME":                        heartbeat.LanguageMIME,
		"MiniD":                       heartbeat.LanguageMiniD,
		"MiniScript":                  heartbeat.LanguageMiniScript,
		"MiniZinc":                    heartbeat.LanguageMiniZinc,
		"Mirah":                       heartbeat.LanguageMirah,
		"MLIR":                        heartbeat.LanguageMLIR,
		"Modelica":                    heartbeat.LanguageModelica,
		"Modula-2":                    heartbeat.LanguageModula2,
		"MoinMoin/Trac Wiki markup":   heartbeat.LanguageMoinWiki,
		"Monkey":                      heartbeat.LanguageMonkey,
		"MonkeyC":                     heartbeat.LanguageMonkeyC,
		"Monte":                       heartbeat.LanguageMonte,
		"MOOCode":                     heartbeat.LanguageMOOCode,
		"MoonScript":                  heartbeat.LanguageMoonScript,
		"MorrowindScript":             heartbeat.LanguageMorrowindScript,
		"Mosel":                       heartbeat.LanguageMosel,
		"mozhashpreproc":              heartbeat.LanguageMozPreprocHash,
		"mozpercentpreproc":           heartbeat.LanguageMozPreprocPercent,
		"MQL":                         heartbeat.LanguageMQL,
		"Mscgen":                      heartbeat.LanguageMscgen,
		"MSDOS Session":               heartbeat.LanguageMSDOSSession,
		"MuPAD":                       heartbeat.LanguageMuPAD,
		"Mustache":                    heartbeat.LanguageMustache,
		"MXML":                        heartbeat.LanguageMXML,
		"Myghty":                      heartbeat.LanguageMyghty,
		"MySQL":                       heartbeat.LanguageMySQL,
		"NASM":                        heartbeat.LanguageNASM,
		"NCL":                         heartbeat.LanguageNCL,
		"Nemerle":                     heartbeat.LanguageNemerle,
		"Neon":                        heartbeat.LanguageNeon,
		"nesC":                        heartbeat.LanguageNesC,
		"newLisp":                     heartbeat.LanguageNewLisp,
		"Newspeak":                    heartbeat.LanguageNewspeak,
		"Nginx":                       heartbeat.LanguageNginx,
		"Nginx configuration file":    heartbeat.LanguageNginxConfig,
		"Nimrod":                      heartbeat.LanguageNimrod,
		"Nit":                         heartbeat.LanguageNit,
		"Nix":                         heartbeat.LanguageNix,
		"Notmuch":                     heartbeat.LanguageNotmuch,
		"Nu":                          heartbeat.LanguageNu,
		"NSIS":                        heartbeat.LanguageNSIS,
		"NumPy":                       heartbeat.LanguageNumPy,
		"NuSMV":                       heartbeat.LanguageNuSMV,
		"objdump":                     heartbeat.LanguageObjdump,
		"objdump-nasm":                heartbeat.LanguageNASMObjdump,
		"Objective-C":                 heartbeat.LanguageObjectiveC,
		"Objective-C++":               heartbeat.LanguageObjectiveCPP,
		"Objective-J":                 heartbeat.LanguageObjectiveJ,
		"OCaml":                       heartbeat.LanguageOCaml,
		"Octave":                      heartbeat.LanguageOctave,
		"ODIN":                        heartbeat.LanguageODIN,
		"ooc":                         heartbeat.LanguageOoc,
		"Opa":                         heartbeat.LanguageOpa,
		"OpenEdge ABL":                heartbeat.LanguageOpenEdgeABL,
		"OpenSCAD":                    heartbeat.LanguageOpenSCAD,
		"Org":                         heartbeat.LanguageOrg,
		"Pascal":                      heartbeat.LanguagePascal,
		"Pawn":                        heartbeat.LanguagePawn,
		"Perl":                        heartbeat.LanguagePerl,
		"PHP":                         heartbeat.LanguagePHP,
		"PostScript":                  heartbeat.LanguagePostScript,
		"POVRay":                      heartbeat.LanguagePOVRay,
		"PowerShell":                  heartbeat.LanguagePowerShell,
		"Prolog":                      heartbeat.LanguageProlog,
		"Protocol Buffer":             heartbeat.LanguageProtocolBuffer,
		"Pug":                         heartbeat.LanguagePug,
		"Puppet":                      heartbeat.LanguagePuppet,
		"PureScript":                  heartbeat.LanguagePureScript,
		"Python":                      heartbeat.LanguagePython,
		"QML":                         heartbeat.LanguageQML,
		"R":                           heartbeat.LanguageR,
		"ReasonML":                    heartbeat.LanguageReasonML,
		"reStructuredText":            heartbeat.LanguageReStructuredText,
		"RPMSpec":                     heartbeat.LanguageRPMSpec,
		"Ruby":                        heartbeat.LanguageRuby,
		"Rust":                        heartbeat.LanguageRust,
		"Salt":                        heartbeat.LanguageSalt,
		"Sass":                        heartbeat.LanguageSass,
		"Scala":                       heartbeat.LanguageScala,
		"Scheme":                      heartbeat.LanguageScheme,
		"Scribe":                      heartbeat.LanguageScribe,
		"SCSS":                        heartbeat.LanguageSCSS,
		"SGML":                        heartbeat.LanguageSGML,
		"Shell":                       heartbeat.LanguageShell,
		"Simula":                      heartbeat.LanguageSimula,
		"Singularity":                 heartbeat.LanguageSingularity,
		"Sketch Drawing":              heartbeat.LanguageSketchDrawing,
		"SKILL":                       heartbeat.LanguageSKILL,
		"Slim":                        heartbeat.LanguageSlim,
		"Smali":                       heartbeat.LanguageSmali,
		"Smalltalk":                   heartbeat.LanguageSmalltalk,
		"S/MIME":                      heartbeat.LanguageSMIME,
		"SourcePawn":                  heartbeat.LanguageSourcePawn,
		"SQL":                         heartbeat.LanguageSQL,
		"Sublime Text Config":         heartbeat.LanguageSublimeTextConfig,
		"Svelte":                      heartbeat.LanguageSvelte,
		"Swift":                       heartbeat.LanguageSwift,
		"SWIG":                        heartbeat.LanguageSWIG,
		"systemverilog":               heartbeat.LanguageSystemVerilog,
		"TeX":                         heartbeat.LanguageTeX,
		"Text":                        heartbeat.LanguageText,
		"Thrift":                      heartbeat.LanguageThrift,
		"TOML":                        heartbeat.LanguageTOML,
		"Turing":                      heartbeat.LanguageTuring,
		"Twig":                        heartbeat.LanguageTwig,
		"TypeScript":                  heartbeat.LanguageTypeScript,
		"TypoScript":                  heartbeat.LanguageTypoScript,
		"VB":                          heartbeat.LanguageVB,
		"VB.net":                      heartbeat.LanguageVBNet,
		"VCL":                         heartbeat.LanguageVCL,
		"Velocity":                    heartbeat.LanguageVelocity,
		"Verilog":                     heartbeat.LanguageVerilog,
		"vhdl":                        heartbeat.LanguageVHDL,
		"VimL":                        heartbeat.LanguageVimL,
		"Vue.js":                      heartbeat.LanguageVueJS,
		"XAML":                        heartbeat.LanguageXAML,
		"XML":                         heartbeat.LanguageXML,
		"XSLT":                        heartbeat.LanguageXSLT,
		"YAML":                        heartbeat.LanguageYAML,
		"Zig":                         heartbeat.LanguageZig,
	}
}

func TestParseLanguage(t *testing.T) {
	for value, language := range languageTests() {
		t.Run(value, func(t *testing.T) {
			parsed, ok := heartbeat.ParseLanguage(value)
			assert.True(t, ok)

			assert.Equal(t, language, parsed, fmt.Sprintf("Got: %q, want: %q", parsed, language))
		})
	}

	t.Run("lower case", func(t *testing.T) {
		parsed, ok := heartbeat.ParseLanguage("go")
		assert.True(t, ok)

		assert.Equal(t, heartbeat.LanguageGo, parsed)
	})

	t.Run("hash", func(t *testing.T) {
		parsed, ok := heartbeat.ParseLanguage("CSharp")
		assert.True(t, ok)

		assert.Equal(t, heartbeat.LanguageCSharp, parsed)
	})

	t.Run("plus sign", func(t *testing.T) {
		parsed, ok := heartbeat.ParseLanguage("CPP")
		assert.True(t, ok)

		assert.Equal(t, heartbeat.LanguageCPP, parsed)
	})

	t.Run("leading space", func(t *testing.T) {
		parsed, ok := heartbeat.ParseLanguage(" Go")
		assert.True(t, ok)

		assert.Equal(t, heartbeat.LanguageGo, parsed)
	})

	t.Run("trailing space", func(t *testing.T) {
		parsed, ok := heartbeat.ParseLanguage("Go ")
		assert.True(t, ok)

		assert.Equal(t, heartbeat.LanguageGo, parsed)
	})

	t.Run("missing hyphen", func(t *testing.T) {
		parsed, ok := heartbeat.ParseLanguage("ObjectiveC")
		assert.True(t, ok)

		assert.Equal(t, heartbeat.LanguageObjectiveC, parsed)
	})

	t.Run("missing space", func(t *testing.T) {
		parsed, ok := heartbeat.ParseLanguage("Sublime Text Config")
		assert.True(t, ok)

		assert.Equal(t, heartbeat.LanguageSublimeTextConfig, parsed)
	})
}

func TestParseLanguage_Unknown(t *testing.T) {
	parsed, ok := heartbeat.ParseLanguage("invalid")

	assert.False(t, ok)
	assert.Equal(t, heartbeat.LanguageUnknown, parsed)
}

func TestParseLanguageFromChroma(t *testing.T) {
	tests := map[string]heartbeat.Language{
		"Base Makefile":    heartbeat.LanguageMakefile,
		"Coldfusion HTML":  heartbeat.LanguageColdfusionHTML,
		"EmacsLisp":        heartbeat.LanguageEmacsLisp,
		"Go HTML Template": heartbeat.LanguageGo,
		"Go Text Template": heartbeat.LanguageGo,
		"FSharp":           heartbeat.LanguageFSharp,
		"GAS":              heartbeat.LanguageAssembly,
		"Gosu Template":    heartbeat.LanguageGosu,
		"LessCss":          heartbeat.LanguageLess,
		"liquid":           heartbeat.LanguageLiquid,
		"markdown":         heartbeat.LanguageMarkdown,
		"NewLisp":          heartbeat.LanguageNewLisp,
		"Nim":              heartbeat.LanguageNimrod,
		"Org Mode":         heartbeat.LanguageOrg,
		"plaintext":        heartbeat.LanguageText,
		"react":            heartbeat.LanguageJSX,
		"VHDL":             heartbeat.LanguageVHDL,
		"vue":              heartbeat.LanguageVueJS,
		// lowercase
		"zig": heartbeat.LanguageZig,
		// missing blank space
		"ProtocolBuffer": heartbeat.LanguageProtocolBuffer,
		// missing hyphen
		"ObjectiveC": heartbeat.LanguageObjectiveC,
		// plus sign
		"CPP": heartbeat.LanguageCPP,
		// hash
		"CSharp": heartbeat.LanguageCSharp,
	}

	for lexerName, language := range tests {
		t.Run(lexerName, func(t *testing.T) {
			parsed, ok := heartbeat.ParseLanguageFromChroma(lexerName)

			assert.True(t, ok)
			assert.Equal(t, language, parsed, fmt.Sprintf("Got: %q, want: %q", parsed, language))
		})
	}
}

func TestParseLanguageFromChroma_Unknown(t *testing.T) {
	parsed, ok := heartbeat.ParseLanguageFromChroma("invalid")

	assert.False(t, ok)
	assert.Equal(t, heartbeat.LanguageUnknown, parsed)
}

func TestParseLanguageFromChroma_AllLexersSupported(t *testing.T) {
	for _, lexer := range lexers.Registry.Lexers {
		config := lexer.Config()

		// TODO: This condition restricts testing to lexers starting with particular
		// letters. Currently only lexers are testsed, which start with letters, where
		// language support was already ensured. Has to be adjust to cover more letters,
		// once another issue is resolved. Has to be removed finally, once all issues
		// are done. Issues:
		// - https://github.com/wakatime/wakatime-cli/issues/232
		// - https://github.com/wakatime/wakatime-cli/issues/233
		// - https://github.com/wakatime/wakatime-cli/issues/234
		// - https://github.com/wakatime/wakatime-cli/issues/238
		// - https://github.com/wakatime/wakatime-cli/issues/239
		rgx := regexp.MustCompile(`^[a-oA-O]`)
		if !rgx.MatchString(config.Name) {
			continue
		}

		parsed, ok := heartbeat.ParseLanguageFromChroma(config.Name)

		assert.True(t, ok, fmt.Sprintf("Failed parsing language from lexer %q", config.Name))
		assert.NotEqual(t, heartbeat.LanguageUnknown, parsed, fmt.Sprintf(
			"Parsed language.Unknown. Failed parsing language from lexer %q",
			config.Name,
		))
	}
}

func TestLanguage_MarshalJSON(t *testing.T) {
	for value, language := range languageTests() {
		t.Run(value, func(t *testing.T) {
			data, err := json.Marshal(language)
			require.NoError(t, err)

			assert.JSONEq(t, `"`+value+`"`, string(data))
		})
	}
}

func TestLanguage_MarshalJSON_UnknownLanguage(t *testing.T) {
	data, err := json.Marshal(heartbeat.LanguageUnknown)
	require.NoError(t, err)

	assert.JSONEq(t, `null`, string(data))
}

func TestLanguage_UnmarshalJSON(t *testing.T) {
	for value, language := range languageTests() {
		t.Run(value, func(t *testing.T) {
			var l heartbeat.Language
			require.NoError(t, json.Unmarshal([]byte(`"`+value+`"`), &l))

			assert.Equal(t, language, l)
		})
	}
}

func TestLanguage_String(t *testing.T) {
	for value, language := range languageTests() {
		t.Run(value, func(t *testing.T) {
			assert.Equal(t, value, language.String())
		})
	}
}

func TestLanguage_String_UnknownLanguage(t *testing.T) {
	assert.Equal(t, "Unknown", heartbeat.LanguageUnknown.String())
}
