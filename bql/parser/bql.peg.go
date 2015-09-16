package parser

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

const end_symbol rune = 1114112

/* The rule types inferred from the grammar are below. */
type pegRule uint16

const (
	ruleUnknown pegRule = iota
	ruleSingleStatement
	ruleStatementWithRest
	ruleStatementWithoutRest
	ruleStatement
	ruleSourceStmt
	ruleSinkStmt
	ruleStateStmt
	ruleStreamStmt
	ruleSelectStmt
	ruleSelectUnionStmt
	ruleCreateStreamAsSelectStmt
	ruleCreateStreamAsSelectUnionStmt
	ruleCreateSourceStmt
	ruleCreateSinkStmt
	ruleCreateStateStmt
	ruleUpdateStateStmt
	ruleUpdateSourceStmt
	ruleUpdateSinkStmt
	ruleInsertIntoSelectStmt
	ruleInsertIntoFromStmt
	rulePauseSourceStmt
	ruleResumeSourceStmt
	ruleRewindSourceStmt
	ruleDropSourceStmt
	ruleDropStreamStmt
	ruleDropSinkStmt
	ruleDropStateStmt
	ruleLoadStateStmt
	ruleLoadStateOrCreateStmt
	ruleSaveStateStmt
	ruleEvalStmt
	ruleEmitter
	ruleEmitterOptions
	ruleEmitterOptionCombinations
	ruleEmitterLimit
	ruleEmitterSample
	ruleCountBasedSampling
	ruleRandomizedSampling
	ruleTimeBasedSampling
	ruleTimeBasedSamplingSeconds
	ruleTimeBasedSamplingMilliseconds
	ruleProjections
	ruleProjection
	ruleAliasExpression
	ruleWindowedFrom
	ruleInterval
	ruleTimeInterval
	ruleTuplesInterval
	ruleRelations
	ruleFilter
	ruleGrouping
	ruleGroupList
	ruleHaving
	ruleRelationLike
	ruleAliasedStreamWindow
	ruleStreamWindow
	ruleStreamLike
	ruleUDSFFuncApp
	ruleSourceSinkSpecs
	ruleUpdateSourceSinkSpecs
	ruleSetOptSpecs
	ruleSourceSinkParam
	ruleSourceSinkParamVal
	ruleParamLiteral
	ruleParamArrayExpr
	rulePausedOpt
	ruleExpressionOrWildcard
	ruleExpression
	ruleorExpr
	ruleandExpr
	rulenotExpr
	rulecomparisonExpr
	ruleotherOpExpr
	ruleisExpr
	ruletermExpr
	ruleproductExpr
	ruleminusExpr
	rulecastExpr
	rulebaseExpr
	ruleFuncTypeCast
	ruleFuncApp
	ruleFuncAppWithOrderBy
	ruleFuncAppWithoutOrderBy
	ruleFuncParams
	ruleParamsOrder
	ruleSortedExpression
	ruleOrderDirectionOpt
	ruleArrayExpr
	ruleMapExpr
	ruleKeyValuePair
	ruleLiteral
	ruleComparisonOp
	ruleOtherOp
	ruleIsOp
	rulePlusMinusOp
	ruleMultDivOp
	ruleStream
	ruleRowMeta
	ruleRowTimestamp
	ruleRowValue
	ruleNumericLiteral
	ruleFloatLiteral
	ruleFunction
	ruleNullLiteral
	ruleBooleanLiteral
	ruleTRUE
	ruleFALSE
	ruleWildcard
	ruleStringLiteral
	ruleISTREAM
	ruleDSTREAM
	ruleRSTREAM
	ruleTUPLES
	ruleSECONDS
	ruleMILLISECONDS
	ruleStreamIdentifier
	ruleSourceSinkType
	ruleSourceSinkParamKey
	rulePaused
	ruleUnpaused
	ruleAscending
	ruleDescending
	ruleType
	ruleBool
	ruleInt
	ruleFloat
	ruleString
	ruleBlob
	ruleTimestamp
	ruleArray
	ruleMap
	ruleOr
	ruleAnd
	ruleNot
	ruleEqual
	ruleLess
	ruleLessOrEqual
	ruleGreater
	ruleGreaterOrEqual
	ruleNotEqual
	ruleConcat
	ruleIs
	ruleIsNot
	rulePlus
	ruleMinus
	ruleMultiply
	ruleDivide
	ruleModulo
	ruleUnaryMinus
	ruleIdentifier
	ruleTargetIdentifier
	ruleident
	rulejsonGetPath
	rulejsonSetPath
	rulejsonPathHead
	rulejsonGetPathNonHead
	rulejsonSetPathNonHead
	rulejsonMapSingleLevel
	rulejsonMapMultipleLevel
	rulejsonMapAccessString
	rulejsonMapAccessBracket
	rulesingleQuotedString
	rulejsonArrayAccess
	rulejsonNonNegativeArrayAccess
	rulejsonArraySlice
	rulejsonArrayPartialSlice
	rulejsonArrayFullSlice
	rulespElem
	rulesp
	rulespOpt
	rulecomment
	rulefinalComment
	rulePegText
	ruleAction0
	ruleAction1
	ruleAction2
	ruleAction3
	ruleAction4
	ruleAction5
	ruleAction6
	ruleAction7
	ruleAction8
	ruleAction9
	ruleAction10
	ruleAction11
	ruleAction12
	ruleAction13
	ruleAction14
	ruleAction15
	ruleAction16
	ruleAction17
	ruleAction18
	ruleAction19
	ruleAction20
	ruleAction21
	ruleAction22
	ruleAction23
	ruleAction24
	ruleAction25
	ruleAction26
	ruleAction27
	ruleAction28
	ruleAction29
	ruleAction30
	ruleAction31
	ruleAction32
	ruleAction33
	ruleAction34
	ruleAction35
	ruleAction36
	ruleAction37
	ruleAction38
	ruleAction39
	ruleAction40
	ruleAction41
	ruleAction42
	ruleAction43
	ruleAction44
	ruleAction45
	ruleAction46
	ruleAction47
	ruleAction48
	ruleAction49
	ruleAction50
	ruleAction51
	ruleAction52
	ruleAction53
	ruleAction54
	ruleAction55
	ruleAction56
	ruleAction57
	ruleAction58
	ruleAction59
	ruleAction60
	ruleAction61
	ruleAction62
	ruleAction63
	ruleAction64
	ruleAction65
	ruleAction66
	ruleAction67
	ruleAction68
	ruleAction69
	ruleAction70
	ruleAction71
	ruleAction72
	ruleAction73
	ruleAction74
	ruleAction75
	ruleAction76
	ruleAction77
	ruleAction78
	ruleAction79
	ruleAction80
	ruleAction81
	ruleAction82
	ruleAction83
	ruleAction84
	ruleAction85
	ruleAction86
	ruleAction87
	ruleAction88
	ruleAction89
	ruleAction90
	ruleAction91
	ruleAction92
	ruleAction93
	ruleAction94
	ruleAction95
	ruleAction96
	ruleAction97
	ruleAction98
	ruleAction99
	ruleAction100
	ruleAction101
	ruleAction102
	ruleAction103
	ruleAction104
	ruleAction105
	ruleAction106
	ruleAction107
	ruleAction108
	ruleAction109
	ruleAction110
	ruleAction111
	ruleAction112
	ruleAction113
	ruleAction114
	ruleAction115
	ruleAction116
	ruleAction117
	ruleAction118
	ruleAction119
	ruleAction120
	ruleAction121

	rulePre_
	rule_In_
	rule_Suf
)

var rul3s = [...]string{
	"Unknown",
	"SingleStatement",
	"StatementWithRest",
	"StatementWithoutRest",
	"Statement",
	"SourceStmt",
	"SinkStmt",
	"StateStmt",
	"StreamStmt",
	"SelectStmt",
	"SelectUnionStmt",
	"CreateStreamAsSelectStmt",
	"CreateStreamAsSelectUnionStmt",
	"CreateSourceStmt",
	"CreateSinkStmt",
	"CreateStateStmt",
	"UpdateStateStmt",
	"UpdateSourceStmt",
	"UpdateSinkStmt",
	"InsertIntoSelectStmt",
	"InsertIntoFromStmt",
	"PauseSourceStmt",
	"ResumeSourceStmt",
	"RewindSourceStmt",
	"DropSourceStmt",
	"DropStreamStmt",
	"DropSinkStmt",
	"DropStateStmt",
	"LoadStateStmt",
	"LoadStateOrCreateStmt",
	"SaveStateStmt",
	"EvalStmt",
	"Emitter",
	"EmitterOptions",
	"EmitterOptionCombinations",
	"EmitterLimit",
	"EmitterSample",
	"CountBasedSampling",
	"RandomizedSampling",
	"TimeBasedSampling",
	"TimeBasedSamplingSeconds",
	"TimeBasedSamplingMilliseconds",
	"Projections",
	"Projection",
	"AliasExpression",
	"WindowedFrom",
	"Interval",
	"TimeInterval",
	"TuplesInterval",
	"Relations",
	"Filter",
	"Grouping",
	"GroupList",
	"Having",
	"RelationLike",
	"AliasedStreamWindow",
	"StreamWindow",
	"StreamLike",
	"UDSFFuncApp",
	"SourceSinkSpecs",
	"UpdateSourceSinkSpecs",
	"SetOptSpecs",
	"SourceSinkParam",
	"SourceSinkParamVal",
	"ParamLiteral",
	"ParamArrayExpr",
	"PausedOpt",
	"ExpressionOrWildcard",
	"Expression",
	"orExpr",
	"andExpr",
	"notExpr",
	"comparisonExpr",
	"otherOpExpr",
	"isExpr",
	"termExpr",
	"productExpr",
	"minusExpr",
	"castExpr",
	"baseExpr",
	"FuncTypeCast",
	"FuncApp",
	"FuncAppWithOrderBy",
	"FuncAppWithoutOrderBy",
	"FuncParams",
	"ParamsOrder",
	"SortedExpression",
	"OrderDirectionOpt",
	"ArrayExpr",
	"MapExpr",
	"KeyValuePair",
	"Literal",
	"ComparisonOp",
	"OtherOp",
	"IsOp",
	"PlusMinusOp",
	"MultDivOp",
	"Stream",
	"RowMeta",
	"RowTimestamp",
	"RowValue",
	"NumericLiteral",
	"FloatLiteral",
	"Function",
	"NullLiteral",
	"BooleanLiteral",
	"TRUE",
	"FALSE",
	"Wildcard",
	"StringLiteral",
	"ISTREAM",
	"DSTREAM",
	"RSTREAM",
	"TUPLES",
	"SECONDS",
	"MILLISECONDS",
	"StreamIdentifier",
	"SourceSinkType",
	"SourceSinkParamKey",
	"Paused",
	"Unpaused",
	"Ascending",
	"Descending",
	"Type",
	"Bool",
	"Int",
	"Float",
	"String",
	"Blob",
	"Timestamp",
	"Array",
	"Map",
	"Or",
	"And",
	"Not",
	"Equal",
	"Less",
	"LessOrEqual",
	"Greater",
	"GreaterOrEqual",
	"NotEqual",
	"Concat",
	"Is",
	"IsNot",
	"Plus",
	"Minus",
	"Multiply",
	"Divide",
	"Modulo",
	"UnaryMinus",
	"Identifier",
	"TargetIdentifier",
	"ident",
	"jsonGetPath",
	"jsonSetPath",
	"jsonPathHead",
	"jsonGetPathNonHead",
	"jsonSetPathNonHead",
	"jsonMapSingleLevel",
	"jsonMapMultipleLevel",
	"jsonMapAccessString",
	"jsonMapAccessBracket",
	"singleQuotedString",
	"jsonArrayAccess",
	"jsonNonNegativeArrayAccess",
	"jsonArraySlice",
	"jsonArrayPartialSlice",
	"jsonArrayFullSlice",
	"spElem",
	"sp",
	"spOpt",
	"comment",
	"finalComment",
	"PegText",
	"Action0",
	"Action1",
	"Action2",
	"Action3",
	"Action4",
	"Action5",
	"Action6",
	"Action7",
	"Action8",
	"Action9",
	"Action10",
	"Action11",
	"Action12",
	"Action13",
	"Action14",
	"Action15",
	"Action16",
	"Action17",
	"Action18",
	"Action19",
	"Action20",
	"Action21",
	"Action22",
	"Action23",
	"Action24",
	"Action25",
	"Action26",
	"Action27",
	"Action28",
	"Action29",
	"Action30",
	"Action31",
	"Action32",
	"Action33",
	"Action34",
	"Action35",
	"Action36",
	"Action37",
	"Action38",
	"Action39",
	"Action40",
	"Action41",
	"Action42",
	"Action43",
	"Action44",
	"Action45",
	"Action46",
	"Action47",
	"Action48",
	"Action49",
	"Action50",
	"Action51",
	"Action52",
	"Action53",
	"Action54",
	"Action55",
	"Action56",
	"Action57",
	"Action58",
	"Action59",
	"Action60",
	"Action61",
	"Action62",
	"Action63",
	"Action64",
	"Action65",
	"Action66",
	"Action67",
	"Action68",
	"Action69",
	"Action70",
	"Action71",
	"Action72",
	"Action73",
	"Action74",
	"Action75",
	"Action76",
	"Action77",
	"Action78",
	"Action79",
	"Action80",
	"Action81",
	"Action82",
	"Action83",
	"Action84",
	"Action85",
	"Action86",
	"Action87",
	"Action88",
	"Action89",
	"Action90",
	"Action91",
	"Action92",
	"Action93",
	"Action94",
	"Action95",
	"Action96",
	"Action97",
	"Action98",
	"Action99",
	"Action100",
	"Action101",
	"Action102",
	"Action103",
	"Action104",
	"Action105",
	"Action106",
	"Action107",
	"Action108",
	"Action109",
	"Action110",
	"Action111",
	"Action112",
	"Action113",
	"Action114",
	"Action115",
	"Action116",
	"Action117",
	"Action118",
	"Action119",
	"Action120",
	"Action121",

	"Pre_",
	"_In_",
	"_Suf",
}

type tokenTree interface {
	Print()
	PrintSyntax()
	PrintSyntaxTree(buffer string)
	Add(rule pegRule, begin, end, next uint32, depth int)
	Expand(index int) tokenTree
	Tokens() <-chan token32
	AST() *node32
	Error() []token32
	trim(length int)
}

type node32 struct {
	token32
	up, next *node32
}

func (node *node32) print(depth int, buffer string) {
	for node != nil {
		for c := 0; c < depth; c++ {
			fmt.Printf(" ")
		}
		fmt.Printf("\x1B[34m%v\x1B[m %v\n", rul3s[node.pegRule], strconv.Quote(string(([]rune(buffer)[node.begin:node.end]))))
		if node.up != nil {
			node.up.print(depth+1, buffer)
		}
		node = node.next
	}
}

func (ast *node32) Print(buffer string) {
	ast.print(0, buffer)
}

type element struct {
	node *node32
	down *element
}

/* ${@} bit structure for abstract syntax tree */
type token32 struct {
	pegRule
	begin, end, next uint32
}

func (t *token32) isZero() bool {
	return t.pegRule == ruleUnknown && t.begin == 0 && t.end == 0 && t.next == 0
}

func (t *token32) isParentOf(u token32) bool {
	return t.begin <= u.begin && t.end >= u.end && t.next > u.next
}

func (t *token32) getToken32() token32 {
	return token32{pegRule: t.pegRule, begin: uint32(t.begin), end: uint32(t.end), next: uint32(t.next)}
}

func (t *token32) String() string {
	return fmt.Sprintf("\x1B[34m%v\x1B[m %v %v %v", rul3s[t.pegRule], t.begin, t.end, t.next)
}

type tokens32 struct {
	tree    []token32
	ordered [][]token32
}

func (t *tokens32) trim(length int) {
	t.tree = t.tree[0:length]
}

func (t *tokens32) Print() {
	for _, token := range t.tree {
		fmt.Println(token.String())
	}
}

func (t *tokens32) Order() [][]token32 {
	if t.ordered != nil {
		return t.ordered
	}

	depths := make([]int32, 1, math.MaxInt16)
	for i, token := range t.tree {
		if token.pegRule == ruleUnknown {
			t.tree = t.tree[:i]
			break
		}
		depth := int(token.next)
		if length := len(depths); depth >= length {
			depths = depths[:depth+1]
		}
		depths[depth]++
	}
	depths = append(depths, 0)

	ordered, pool := make([][]token32, len(depths)), make([]token32, len(t.tree)+len(depths))
	for i, depth := range depths {
		depth++
		ordered[i], pool, depths[i] = pool[:depth], pool[depth:], 0
	}

	for i, token := range t.tree {
		depth := token.next
		token.next = uint32(i)
		ordered[depth][depths[depth]] = token
		depths[depth]++
	}
	t.ordered = ordered
	return ordered
}

type state32 struct {
	token32
	depths []int32
	leaf   bool
}

func (t *tokens32) AST() *node32 {
	tokens := t.Tokens()
	stack := &element{node: &node32{token32: <-tokens}}
	for token := range tokens {
		if token.begin == token.end {
			continue
		}
		node := &node32{token32: token}
		for stack != nil && stack.node.begin >= token.begin && stack.node.end <= token.end {
			stack.node.next = node.up
			node.up = stack.node
			stack = stack.down
		}
		stack = &element{node: node, down: stack}
	}
	return stack.node
}

func (t *tokens32) PreOrder() (<-chan state32, [][]token32) {
	s, ordered := make(chan state32, 6), t.Order()
	go func() {
		var states [8]state32
		for i, _ := range states {
			states[i].depths = make([]int32, len(ordered))
		}
		depths, state, depth := make([]int32, len(ordered)), 0, 1
		write := func(t token32, leaf bool) {
			S := states[state]
			state, S.pegRule, S.begin, S.end, S.next, S.leaf = (state+1)%8, t.pegRule, t.begin, t.end, uint32(depth), leaf
			copy(S.depths, depths)
			s <- S
		}

		states[state].token32 = ordered[0][0]
		depths[0]++
		state++
		a, b := ordered[depth-1][depths[depth-1]-1], ordered[depth][depths[depth]]
	depthFirstSearch:
		for {
			for {
				if i := depths[depth]; i > 0 {
					if c, j := ordered[depth][i-1], depths[depth-1]; a.isParentOf(c) &&
						(j < 2 || !ordered[depth-1][j-2].isParentOf(c)) {
						if c.end != b.begin {
							write(token32{pegRule: rule_In_, begin: c.end, end: b.begin}, true)
						}
						break
					}
				}

				if a.begin < b.begin {
					write(token32{pegRule: rulePre_, begin: a.begin, end: b.begin}, true)
				}
				break
			}

			next := depth + 1
			if c := ordered[next][depths[next]]; c.pegRule != ruleUnknown && b.isParentOf(c) {
				write(b, false)
				depths[depth]++
				depth, a, b = next, b, c
				continue
			}

			write(b, true)
			depths[depth]++
			c, parent := ordered[depth][depths[depth]], true
			for {
				if c.pegRule != ruleUnknown && a.isParentOf(c) {
					b = c
					continue depthFirstSearch
				} else if parent && b.end != a.end {
					write(token32{pegRule: rule_Suf, begin: b.end, end: a.end}, true)
				}

				depth--
				if depth > 0 {
					a, b, c = ordered[depth-1][depths[depth-1]-1], a, ordered[depth][depths[depth]]
					parent = a.isParentOf(b)
					continue
				}

				break depthFirstSearch
			}
		}

		close(s)
	}()
	return s, ordered
}

func (t *tokens32) PrintSyntax() {
	tokens, ordered := t.PreOrder()
	max := -1
	for token := range tokens {
		if !token.leaf {
			fmt.Printf("%v", token.begin)
			for i, leaf, depths := 0, int(token.next), token.depths; i < leaf; i++ {
				fmt.Printf(" \x1B[36m%v\x1B[m", rul3s[ordered[i][depths[i]-1].pegRule])
			}
			fmt.Printf(" \x1B[36m%v\x1B[m\n", rul3s[token.pegRule])
		} else if token.begin == token.end {
			fmt.Printf("%v", token.begin)
			for i, leaf, depths := 0, int(token.next), token.depths; i < leaf; i++ {
				fmt.Printf(" \x1B[31m%v\x1B[m", rul3s[ordered[i][depths[i]-1].pegRule])
			}
			fmt.Printf(" \x1B[31m%v\x1B[m\n", rul3s[token.pegRule])
		} else {
			for c, end := token.begin, token.end; c < end; c++ {
				if i := int(c); max+1 < i {
					for j := max; j < i; j++ {
						fmt.Printf("skip %v %v\n", j, token.String())
					}
					max = i
				} else if i := int(c); i <= max {
					for j := i; j <= max; j++ {
						fmt.Printf("dupe %v %v\n", j, token.String())
					}
				} else {
					max = int(c)
				}
				fmt.Printf("%v", c)
				for i, leaf, depths := 0, int(token.next), token.depths; i < leaf; i++ {
					fmt.Printf(" \x1B[34m%v\x1B[m", rul3s[ordered[i][depths[i]-1].pegRule])
				}
				fmt.Printf(" \x1B[34m%v\x1B[m\n", rul3s[token.pegRule])
			}
			fmt.Printf("\n")
		}
	}
}

func (t *tokens32) PrintSyntaxTree(buffer string) {
	tokens, _ := t.PreOrder()
	for token := range tokens {
		for c := 0; c < int(token.next); c++ {
			fmt.Printf(" ")
		}
		fmt.Printf("\x1B[34m%v\x1B[m %v\n", rul3s[token.pegRule], strconv.Quote(string(([]rune(buffer)[token.begin:token.end]))))
	}
}

func (t *tokens32) Add(rule pegRule, begin, end, depth uint32, index int) {
	t.tree[index] = token32{pegRule: rule, begin: uint32(begin), end: uint32(end), next: uint32(depth)}
}

func (t *tokens32) Tokens() <-chan token32 {
	s := make(chan token32, 16)
	go func() {
		for _, v := range t.tree {
			s <- v.getToken32()
		}
		close(s)
	}()
	return s
}

func (t *tokens32) Error() []token32 {
	ordered := t.Order()
	length := len(ordered)
	tokens, length := make([]token32, length), length-1
	for i, _ := range tokens {
		o := ordered[length-i]
		if len(o) > 1 {
			tokens[i] = o[len(o)-2].getToken32()
		}
	}
	return tokens
}

/*func (t *tokens16) Expand(index int) tokenTree {
	tree := t.tree
	if index >= len(tree) {
		expanded := make([]token32, 2 * len(tree))
		for i, v := range tree {
			expanded[i] = v.getToken32()
		}
		return &tokens32{tree: expanded}
	}
	return nil
}*/

func (t *tokens32) Expand(index int) tokenTree {
	tree := t.tree
	if index >= len(tree) {
		expanded := make([]token32, 2*len(tree))
		copy(expanded, tree)
		t.tree = expanded
	}
	return nil
}

type bqlPegBackend struct {
	parseStack

	Buffer string
	buffer []rune
	rules  [296]func() bool
	Parse  func(rule ...int) error
	Reset  func()
	tokenTree
}

type textPosition struct {
	line, symbol int
}

type textPositionMap map[int]textPosition

func translatePositions(buffer string, positions []int) textPositionMap {
	length, translations, j, line, symbol := len(positions), make(textPositionMap, len(positions)), 0, 1, 0
	sort.Ints(positions)

search:
	for i, c := range []rune(buffer) {
		if c == '\n' {
			line, symbol = line+1, 0
		} else {
			symbol++
		}
		if i == positions[j] {
			translations[positions[j]] = textPosition{line, symbol}
			for j++; j < length; j++ {
				if i != positions[j] {
					continue search
				}
			}
			break search
		}
	}

	return translations
}

type parseError struct {
	p *bqlPegBackend
}

func (e *parseError) Error() string {
	tokens, error := e.p.tokenTree.Error(), "\n"
	positions, p := make([]int, 2*len(tokens)), 0
	for _, token := range tokens {
		positions[p], p = int(token.begin), p+1
		positions[p], p = int(token.end), p+1
	}
	translations := translatePositions(e.p.Buffer, positions)
	for _, token := range tokens {
		begin, end := int(token.begin), int(token.end)
		error += fmt.Sprintf("parse error near \x1B[34m%v\x1B[m (line %v symbol %v - line %v symbol %v):\n%v\n",
			rul3s[token.pegRule],
			translations[begin].line, translations[begin].symbol,
			translations[end].line, translations[end].symbol,
			/*strconv.Quote(*/ e.p.Buffer[begin:end] /*)*/)
	}

	return error
}

func (p *bqlPegBackend) PrintSyntaxTree() {
	p.tokenTree.PrintSyntaxTree(p.Buffer)
}

func (p *bqlPegBackend) Highlighter() {
	p.tokenTree.PrintSyntax()
}

func (p *bqlPegBackend) Execute() {
	buffer, _buffer, text, begin, end := p.Buffer, p.buffer, "", 0, 0
	for token := range p.tokenTree.Tokens() {
		switch token.pegRule {

		case rulePegText:
			begin, end = int(token.begin), int(token.end)
			text = string(_buffer[begin:end])

		case ruleAction0:

			p.IncludeTrailingWhitespace(begin, end)

		case ruleAction1:

			p.IncludeTrailingWhitespace(begin, end)

		case ruleAction2:

			p.AssembleSelect()

		case ruleAction3:

			p.AssembleSelectUnion(begin, end)

		case ruleAction4:

			p.AssembleCreateStreamAsSelect()

		case ruleAction5:

			p.AssembleCreateStreamAsSelectUnion()

		case ruleAction6:

			p.AssembleCreateSource()

		case ruleAction7:

			p.AssembleCreateSink()

		case ruleAction8:

			p.AssembleCreateState()

		case ruleAction9:

			p.AssembleUpdateState()

		case ruleAction10:

			p.AssembleUpdateSource()

		case ruleAction11:

			p.AssembleUpdateSink()

		case ruleAction12:

			p.AssembleInsertIntoSelect()

		case ruleAction13:

			p.AssembleInsertIntoFrom()

		case ruleAction14:

			p.AssemblePauseSource()

		case ruleAction15:

			p.AssembleResumeSource()

		case ruleAction16:

			p.AssembleRewindSource()

		case ruleAction17:

			p.AssembleDropSource()

		case ruleAction18:

			p.AssembleDropStream()

		case ruleAction19:

			p.AssembleDropSink()

		case ruleAction20:

			p.AssembleDropState()

		case ruleAction21:

			p.AssembleLoadState()

		case ruleAction22:

			p.AssembleLoadStateOrCreate()

		case ruleAction23:

			p.AssembleSaveState()

		case ruleAction24:

			p.AssembleEval(begin, end)

		case ruleAction25:

			p.AssembleEmitter()

		case ruleAction26:

			p.AssembleEmitterOptions(begin, end)

		case ruleAction27:

			p.AssembleEmitterLimit()

		case ruleAction28:

			p.AssembleEmitterSampling(CountBasedSampling, 1)

		case ruleAction29:

			p.AssembleEmitterSampling(RandomizedSampling, 1)

		case ruleAction30:

			p.AssembleEmitterSampling(TimeBasedSampling, 1)

		case ruleAction31:

			p.AssembleEmitterSampling(TimeBasedSampling, 0.001)

		case ruleAction32:

			p.AssembleProjections(begin, end)

		case ruleAction33:

			p.AssembleAlias()

		case ruleAction34:

			// This is *always* executed, even if there is no
			// FROM clause present in the statement.
			p.AssembleWindowedFrom(begin, end)

		case ruleAction35:

			p.AssembleInterval()

		case ruleAction36:

			p.AssembleInterval()

		case ruleAction37:

			// This is *always* executed, even if there is no
			// WHERE clause present in the statement.
			p.AssembleFilter(begin, end)

		case ruleAction38:

			// This is *always* executed, even if there is no
			// GROUP BY clause present in the statement.
			p.AssembleGrouping(begin, end)

		case ruleAction39:

			// This is *always* executed, even if there is no
			// HAVING clause present in the statement.
			p.AssembleHaving(begin, end)

		case ruleAction40:

			p.EnsureAliasedStreamWindow()

		case ruleAction41:

			p.AssembleAliasedStreamWindow()

		case ruleAction42:

			p.AssembleStreamWindow()

		case ruleAction43:

			p.AssembleUDSFFuncApp()

		case ruleAction44:

			p.AssembleSourceSinkSpecs(begin, end)

		case ruleAction45:

			p.AssembleSourceSinkSpecs(begin, end)

		case ruleAction46:

			p.AssembleSourceSinkSpecs(begin, end)

		case ruleAction47:

			p.AssembleSourceSinkParam()

		case ruleAction48:

			p.AssembleExpressions(begin, end)
			p.AssembleArray()

		case ruleAction49:

			p.EnsureKeywordPresent(begin, end)

		case ruleAction50:

			p.AssembleBinaryOperation(begin, end)

		case ruleAction51:

			p.AssembleBinaryOperation(begin, end)

		case ruleAction52:

			p.AssembleUnaryPrefixOperation(begin, end)

		case ruleAction53:

			p.AssembleBinaryOperation(begin, end)

		case ruleAction54:

			p.AssembleBinaryOperation(begin, end)

		case ruleAction55:

			p.AssembleBinaryOperation(begin, end)

		case ruleAction56:

			p.AssembleBinaryOperation(begin, end)

		case ruleAction57:

			p.AssembleBinaryOperation(begin, end)

		case ruleAction58:

			p.AssembleUnaryPrefixOperation(begin, end)

		case ruleAction59:

			p.AssembleTypeCast(begin, end)

		case ruleAction60:

			p.AssembleTypeCast(begin, end)

		case ruleAction61:

			p.AssembleFuncApp()

		case ruleAction62:

			p.AssembleExpressions(begin, end)
			p.AssembleFuncApp()

		case ruleAction63:

			p.AssembleExpressions(begin, end)

		case ruleAction64:

			p.AssembleExpressions(begin, end)

		case ruleAction65:

			p.AssembleSortedExpression()

		case ruleAction66:

			p.EnsureKeywordPresent(begin, end)

		case ruleAction67:

			p.AssembleExpressions(begin, end)
			p.AssembleArray()

		case ruleAction68:

			p.AssembleMap(begin, end)

		case ruleAction69:

			p.AssembleKeyValuePair()

		case ruleAction70:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, NewStream(substr))

		case ruleAction71:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, NewRowMeta(substr, TimestampMeta))

		case ruleAction72:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, NewRowValue(substr))

		case ruleAction73:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, NewNumericLiteral(substr))

		case ruleAction74:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, NewFloatLiteral(substr))

		case ruleAction75:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, FuncName(substr))

		case ruleAction76:

			p.PushComponent(begin, end, NewNullLiteral())

		case ruleAction77:

			p.PushComponent(begin, end, NewBoolLiteral(true))

		case ruleAction78:

			p.PushComponent(begin, end, NewBoolLiteral(false))

		case ruleAction79:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, NewWildcard(substr))

		case ruleAction80:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, NewStringLiteral(substr))

		case ruleAction81:

			p.PushComponent(begin, end, Istream)

		case ruleAction82:

			p.PushComponent(begin, end, Dstream)

		case ruleAction83:

			p.PushComponent(begin, end, Rstream)

		case ruleAction84:

			p.PushComponent(begin, end, Tuples)

		case ruleAction85:

			p.PushComponent(begin, end, Seconds)

		case ruleAction86:

			p.PushComponent(begin, end, Milliseconds)

		case ruleAction87:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, StreamIdentifier(substr))

		case ruleAction88:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, SourceSinkType(substr))

		case ruleAction89:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, SourceSinkParamKey(substr))

		case ruleAction90:

			p.PushComponent(begin, end, Yes)

		case ruleAction91:

			p.PushComponent(begin, end, No)

		case ruleAction92:

			p.PushComponent(begin, end, Yes)

		case ruleAction93:

			p.PushComponent(begin, end, No)

		case ruleAction94:

			p.PushComponent(begin, end, Bool)

		case ruleAction95:

			p.PushComponent(begin, end, Int)

		case ruleAction96:

			p.PushComponent(begin, end, Float)

		case ruleAction97:

			p.PushComponent(begin, end, String)

		case ruleAction98:

			p.PushComponent(begin, end, Blob)

		case ruleAction99:

			p.PushComponent(begin, end, Timestamp)

		case ruleAction100:

			p.PushComponent(begin, end, Array)

		case ruleAction101:

			p.PushComponent(begin, end, Map)

		case ruleAction102:

			p.PushComponent(begin, end, Or)

		case ruleAction103:

			p.PushComponent(begin, end, And)

		case ruleAction104:

			p.PushComponent(begin, end, Not)

		case ruleAction105:

			p.PushComponent(begin, end, Equal)

		case ruleAction106:

			p.PushComponent(begin, end, Less)

		case ruleAction107:

			p.PushComponent(begin, end, LessOrEqual)

		case ruleAction108:

			p.PushComponent(begin, end, Greater)

		case ruleAction109:

			p.PushComponent(begin, end, GreaterOrEqual)

		case ruleAction110:

			p.PushComponent(begin, end, NotEqual)

		case ruleAction111:

			p.PushComponent(begin, end, Concat)

		case ruleAction112:

			p.PushComponent(begin, end, Is)

		case ruleAction113:

			p.PushComponent(begin, end, IsNot)

		case ruleAction114:

			p.PushComponent(begin, end, Plus)

		case ruleAction115:

			p.PushComponent(begin, end, Minus)

		case ruleAction116:

			p.PushComponent(begin, end, Multiply)

		case ruleAction117:

			p.PushComponent(begin, end, Divide)

		case ruleAction118:

			p.PushComponent(begin, end, Modulo)

		case ruleAction119:

			p.PushComponent(begin, end, UnaryMinus)

		case ruleAction120:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, Identifier(substr))

		case ruleAction121:

			substr := string([]rune(buffer)[begin:end])
			p.PushComponent(begin, end, Identifier(substr))

		}
	}
	_, _, _, _ = buffer, text, begin, end
}

func (p *bqlPegBackend) Init() {
	p.buffer = []rune(p.Buffer)
	if len(p.buffer) == 0 || p.buffer[len(p.buffer)-1] != end_symbol {
		p.buffer = append(p.buffer, end_symbol)
	}

	var tree tokenTree = &tokens32{tree: make([]token32, math.MaxInt16)}
	position, depth, tokenIndex, buffer, _rules := uint32(0), uint32(0), 0, p.buffer, p.rules

	p.Parse = func(rule ...int) error {
		r := 1
		if len(rule) > 0 {
			r = rule[0]
		}
		matches := p.rules[r]()
		p.tokenTree = tree
		if matches {
			p.tokenTree.trim(tokenIndex)
			return nil
		}
		return &parseError{p}
	}

	p.Reset = func() {
		position, tokenIndex, depth = 0, 0, 0
	}

	add := func(rule pegRule, begin uint32) {
		if t := tree.Expand(tokenIndex); t != nil {
			tree = t
		}
		tree.Add(rule, begin, position, depth, tokenIndex)
		tokenIndex++
	}

	matchDot := func() bool {
		if buffer[position] != end_symbol {
			position++
			return true
		}
		return false
	}

	/*matchChar := func(c byte) bool {
		if buffer[position] == c {
			position++
			return true
		}
		return false
	}*/

	/*matchRange := func(lower byte, upper byte) bool {
		if c := buffer[position]; c >= lower && c <= upper {
			position++
			return true
		}
		return false
	}*/

	_rules = [...]func() bool{
		nil,
		/* 0 SingleStatement <- <(spOpt (StatementWithRest / StatementWithoutRest) !.)> */
		func() bool {
			position0, tokenIndex0, depth0 := position, tokenIndex, depth
			{
				position1 := position
				depth++
				if !_rules[rulespOpt]() {
					goto l0
				}
				{
					position2, tokenIndex2, depth2 := position, tokenIndex, depth
					if !_rules[ruleStatementWithRest]() {
						goto l3
					}
					goto l2
				l3:
					position, tokenIndex, depth = position2, tokenIndex2, depth2
					if !_rules[ruleStatementWithoutRest]() {
						goto l0
					}
				}
			l2:
				{
					position4, tokenIndex4, depth4 := position, tokenIndex, depth
					if !matchDot() {
						goto l4
					}
					goto l0
				l4:
					position, tokenIndex, depth = position4, tokenIndex4, depth4
				}
				depth--
				add(ruleSingleStatement, position1)
			}
			return true
		l0:
			position, tokenIndex, depth = position0, tokenIndex0, depth0
			return false
		},
		/* 1 StatementWithRest <- <(<(Statement spOpt ';' spOpt)> .* Action0)> */
		func() bool {
			position5, tokenIndex5, depth5 := position, tokenIndex, depth
			{
				position6 := position
				depth++
				{
					position7 := position
					depth++
					if !_rules[ruleStatement]() {
						goto l5
					}
					if !_rules[rulespOpt]() {
						goto l5
					}
					if buffer[position] != rune(';') {
						goto l5
					}
					position++
					if !_rules[rulespOpt]() {
						goto l5
					}
					depth--
					add(rulePegText, position7)
				}
			l8:
				{
					position9, tokenIndex9, depth9 := position, tokenIndex, depth
					if !matchDot() {
						goto l9
					}
					goto l8
				l9:
					position, tokenIndex, depth = position9, tokenIndex9, depth9
				}
				if !_rules[ruleAction0]() {
					goto l5
				}
				depth--
				add(ruleStatementWithRest, position6)
			}
			return true
		l5:
			position, tokenIndex, depth = position5, tokenIndex5, depth5
			return false
		},
		/* 2 StatementWithoutRest <- <(<(Statement spOpt)> Action1)> */
		func() bool {
			position10, tokenIndex10, depth10 := position, tokenIndex, depth
			{
				position11 := position
				depth++
				{
					position12 := position
					depth++
					if !_rules[ruleStatement]() {
						goto l10
					}
					if !_rules[rulespOpt]() {
						goto l10
					}
					depth--
					add(rulePegText, position12)
				}
				if !_rules[ruleAction1]() {
					goto l10
				}
				depth--
				add(ruleStatementWithoutRest, position11)
			}
			return true
		l10:
			position, tokenIndex, depth = position10, tokenIndex10, depth10
			return false
		},
		/* 3 Statement <- <(SelectUnionStmt / SelectStmt / SourceStmt / SinkStmt / StateStmt / StreamStmt / EvalStmt)> */
		func() bool {
			position13, tokenIndex13, depth13 := position, tokenIndex, depth
			{
				position14 := position
				depth++
				{
					position15, tokenIndex15, depth15 := position, tokenIndex, depth
					if !_rules[ruleSelectUnionStmt]() {
						goto l16
					}
					goto l15
				l16:
					position, tokenIndex, depth = position15, tokenIndex15, depth15
					if !_rules[ruleSelectStmt]() {
						goto l17
					}
					goto l15
				l17:
					position, tokenIndex, depth = position15, tokenIndex15, depth15
					if !_rules[ruleSourceStmt]() {
						goto l18
					}
					goto l15
				l18:
					position, tokenIndex, depth = position15, tokenIndex15, depth15
					if !_rules[ruleSinkStmt]() {
						goto l19
					}
					goto l15
				l19:
					position, tokenIndex, depth = position15, tokenIndex15, depth15
					if !_rules[ruleStateStmt]() {
						goto l20
					}
					goto l15
				l20:
					position, tokenIndex, depth = position15, tokenIndex15, depth15
					if !_rules[ruleStreamStmt]() {
						goto l21
					}
					goto l15
				l21:
					position, tokenIndex, depth = position15, tokenIndex15, depth15
					if !_rules[ruleEvalStmt]() {
						goto l13
					}
				}
			l15:
				depth--
				add(ruleStatement, position14)
			}
			return true
		l13:
			position, tokenIndex, depth = position13, tokenIndex13, depth13
			return false
		},
		/* 4 SourceStmt <- <(CreateSourceStmt / UpdateSourceStmt / DropSourceStmt / PauseSourceStmt / ResumeSourceStmt / RewindSourceStmt)> */
		func() bool {
			position22, tokenIndex22, depth22 := position, tokenIndex, depth
			{
				position23 := position
				depth++
				{
					position24, tokenIndex24, depth24 := position, tokenIndex, depth
					if !_rules[ruleCreateSourceStmt]() {
						goto l25
					}
					goto l24
				l25:
					position, tokenIndex, depth = position24, tokenIndex24, depth24
					if !_rules[ruleUpdateSourceStmt]() {
						goto l26
					}
					goto l24
				l26:
					position, tokenIndex, depth = position24, tokenIndex24, depth24
					if !_rules[ruleDropSourceStmt]() {
						goto l27
					}
					goto l24
				l27:
					position, tokenIndex, depth = position24, tokenIndex24, depth24
					if !_rules[rulePauseSourceStmt]() {
						goto l28
					}
					goto l24
				l28:
					position, tokenIndex, depth = position24, tokenIndex24, depth24
					if !_rules[ruleResumeSourceStmt]() {
						goto l29
					}
					goto l24
				l29:
					position, tokenIndex, depth = position24, tokenIndex24, depth24
					if !_rules[ruleRewindSourceStmt]() {
						goto l22
					}
				}
			l24:
				depth--
				add(ruleSourceStmt, position23)
			}
			return true
		l22:
			position, tokenIndex, depth = position22, tokenIndex22, depth22
			return false
		},
		/* 5 SinkStmt <- <(CreateSinkStmt / UpdateSinkStmt / DropSinkStmt)> */
		func() bool {
			position30, tokenIndex30, depth30 := position, tokenIndex, depth
			{
				position31 := position
				depth++
				{
					position32, tokenIndex32, depth32 := position, tokenIndex, depth
					if !_rules[ruleCreateSinkStmt]() {
						goto l33
					}
					goto l32
				l33:
					position, tokenIndex, depth = position32, tokenIndex32, depth32
					if !_rules[ruleUpdateSinkStmt]() {
						goto l34
					}
					goto l32
				l34:
					position, tokenIndex, depth = position32, tokenIndex32, depth32
					if !_rules[ruleDropSinkStmt]() {
						goto l30
					}
				}
			l32:
				depth--
				add(ruleSinkStmt, position31)
			}
			return true
		l30:
			position, tokenIndex, depth = position30, tokenIndex30, depth30
			return false
		},
		/* 6 StateStmt <- <(CreateStateStmt / UpdateStateStmt / DropStateStmt / LoadStateOrCreateStmt / LoadStateStmt / SaveStateStmt)> */
		func() bool {
			position35, tokenIndex35, depth35 := position, tokenIndex, depth
			{
				position36 := position
				depth++
				{
					position37, tokenIndex37, depth37 := position, tokenIndex, depth
					if !_rules[ruleCreateStateStmt]() {
						goto l38
					}
					goto l37
				l38:
					position, tokenIndex, depth = position37, tokenIndex37, depth37
					if !_rules[ruleUpdateStateStmt]() {
						goto l39
					}
					goto l37
				l39:
					position, tokenIndex, depth = position37, tokenIndex37, depth37
					if !_rules[ruleDropStateStmt]() {
						goto l40
					}
					goto l37
				l40:
					position, tokenIndex, depth = position37, tokenIndex37, depth37
					if !_rules[ruleLoadStateOrCreateStmt]() {
						goto l41
					}
					goto l37
				l41:
					position, tokenIndex, depth = position37, tokenIndex37, depth37
					if !_rules[ruleLoadStateStmt]() {
						goto l42
					}
					goto l37
				l42:
					position, tokenIndex, depth = position37, tokenIndex37, depth37
					if !_rules[ruleSaveStateStmt]() {
						goto l35
					}
				}
			l37:
				depth--
				add(ruleStateStmt, position36)
			}
			return true
		l35:
			position, tokenIndex, depth = position35, tokenIndex35, depth35
			return false
		},
		/* 7 StreamStmt <- <(CreateStreamAsSelectUnionStmt / CreateStreamAsSelectStmt / DropStreamStmt / InsertIntoSelectStmt / InsertIntoFromStmt)> */
		func() bool {
			position43, tokenIndex43, depth43 := position, tokenIndex, depth
			{
				position44 := position
				depth++
				{
					position45, tokenIndex45, depth45 := position, tokenIndex, depth
					if !_rules[ruleCreateStreamAsSelectUnionStmt]() {
						goto l46
					}
					goto l45
				l46:
					position, tokenIndex, depth = position45, tokenIndex45, depth45
					if !_rules[ruleCreateStreamAsSelectStmt]() {
						goto l47
					}
					goto l45
				l47:
					position, tokenIndex, depth = position45, tokenIndex45, depth45
					if !_rules[ruleDropStreamStmt]() {
						goto l48
					}
					goto l45
				l48:
					position, tokenIndex, depth = position45, tokenIndex45, depth45
					if !_rules[ruleInsertIntoSelectStmt]() {
						goto l49
					}
					goto l45
				l49:
					position, tokenIndex, depth = position45, tokenIndex45, depth45
					if !_rules[ruleInsertIntoFromStmt]() {
						goto l43
					}
				}
			l45:
				depth--
				add(ruleStreamStmt, position44)
			}
			return true
		l43:
			position, tokenIndex, depth = position43, tokenIndex43, depth43
			return false
		},
		/* 8 SelectStmt <- <(('s' / 'S') ('e' / 'E') ('l' / 'L') ('e' / 'E') ('c' / 'C') ('t' / 'T') Emitter Projections WindowedFrom Filter Grouping Having Action2)> */
		func() bool {
			position50, tokenIndex50, depth50 := position, tokenIndex, depth
			{
				position51 := position
				depth++
				{
					position52, tokenIndex52, depth52 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l53
					}
					position++
					goto l52
				l53:
					position, tokenIndex, depth = position52, tokenIndex52, depth52
					if buffer[position] != rune('S') {
						goto l50
					}
					position++
				}
			l52:
				{
					position54, tokenIndex54, depth54 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l55
					}
					position++
					goto l54
				l55:
					position, tokenIndex, depth = position54, tokenIndex54, depth54
					if buffer[position] != rune('E') {
						goto l50
					}
					position++
				}
			l54:
				{
					position56, tokenIndex56, depth56 := position, tokenIndex, depth
					if buffer[position] != rune('l') {
						goto l57
					}
					position++
					goto l56
				l57:
					position, tokenIndex, depth = position56, tokenIndex56, depth56
					if buffer[position] != rune('L') {
						goto l50
					}
					position++
				}
			l56:
				{
					position58, tokenIndex58, depth58 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l59
					}
					position++
					goto l58
				l59:
					position, tokenIndex, depth = position58, tokenIndex58, depth58
					if buffer[position] != rune('E') {
						goto l50
					}
					position++
				}
			l58:
				{
					position60, tokenIndex60, depth60 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l61
					}
					position++
					goto l60
				l61:
					position, tokenIndex, depth = position60, tokenIndex60, depth60
					if buffer[position] != rune('C') {
						goto l50
					}
					position++
				}
			l60:
				{
					position62, tokenIndex62, depth62 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l63
					}
					position++
					goto l62
				l63:
					position, tokenIndex, depth = position62, tokenIndex62, depth62
					if buffer[position] != rune('T') {
						goto l50
					}
					position++
				}
			l62:
				if !_rules[ruleEmitter]() {
					goto l50
				}
				if !_rules[ruleProjections]() {
					goto l50
				}
				if !_rules[ruleWindowedFrom]() {
					goto l50
				}
				if !_rules[ruleFilter]() {
					goto l50
				}
				if !_rules[ruleGrouping]() {
					goto l50
				}
				if !_rules[ruleHaving]() {
					goto l50
				}
				if !_rules[ruleAction2]() {
					goto l50
				}
				depth--
				add(ruleSelectStmt, position51)
			}
			return true
		l50:
			position, tokenIndex, depth = position50, tokenIndex50, depth50
			return false
		},
		/* 9 SelectUnionStmt <- <(<(SelectStmt (sp (('u' / 'U') ('n' / 'N') ('i' / 'I') ('o' / 'O') ('n' / 'N')) sp (('a' / 'A') ('l' / 'L') ('l' / 'L')) sp SelectStmt)+)> Action3)> */
		func() bool {
			position64, tokenIndex64, depth64 := position, tokenIndex, depth
			{
				position65 := position
				depth++
				{
					position66 := position
					depth++
					if !_rules[ruleSelectStmt]() {
						goto l64
					}
					if !_rules[rulesp]() {
						goto l64
					}
					{
						position69, tokenIndex69, depth69 := position, tokenIndex, depth
						if buffer[position] != rune('u') {
							goto l70
						}
						position++
						goto l69
					l70:
						position, tokenIndex, depth = position69, tokenIndex69, depth69
						if buffer[position] != rune('U') {
							goto l64
						}
						position++
					}
				l69:
					{
						position71, tokenIndex71, depth71 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l72
						}
						position++
						goto l71
					l72:
						position, tokenIndex, depth = position71, tokenIndex71, depth71
						if buffer[position] != rune('N') {
							goto l64
						}
						position++
					}
				l71:
					{
						position73, tokenIndex73, depth73 := position, tokenIndex, depth
						if buffer[position] != rune('i') {
							goto l74
						}
						position++
						goto l73
					l74:
						position, tokenIndex, depth = position73, tokenIndex73, depth73
						if buffer[position] != rune('I') {
							goto l64
						}
						position++
					}
				l73:
					{
						position75, tokenIndex75, depth75 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l76
						}
						position++
						goto l75
					l76:
						position, tokenIndex, depth = position75, tokenIndex75, depth75
						if buffer[position] != rune('O') {
							goto l64
						}
						position++
					}
				l75:
					{
						position77, tokenIndex77, depth77 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l78
						}
						position++
						goto l77
					l78:
						position, tokenIndex, depth = position77, tokenIndex77, depth77
						if buffer[position] != rune('N') {
							goto l64
						}
						position++
					}
				l77:
					if !_rules[rulesp]() {
						goto l64
					}
					{
						position79, tokenIndex79, depth79 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l80
						}
						position++
						goto l79
					l80:
						position, tokenIndex, depth = position79, tokenIndex79, depth79
						if buffer[position] != rune('A') {
							goto l64
						}
						position++
					}
				l79:
					{
						position81, tokenIndex81, depth81 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l82
						}
						position++
						goto l81
					l82:
						position, tokenIndex, depth = position81, tokenIndex81, depth81
						if buffer[position] != rune('L') {
							goto l64
						}
						position++
					}
				l81:
					{
						position83, tokenIndex83, depth83 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l84
						}
						position++
						goto l83
					l84:
						position, tokenIndex, depth = position83, tokenIndex83, depth83
						if buffer[position] != rune('L') {
							goto l64
						}
						position++
					}
				l83:
					if !_rules[rulesp]() {
						goto l64
					}
					if !_rules[ruleSelectStmt]() {
						goto l64
					}
				l67:
					{
						position68, tokenIndex68, depth68 := position, tokenIndex, depth
						if !_rules[rulesp]() {
							goto l68
						}
						{
							position85, tokenIndex85, depth85 := position, tokenIndex, depth
							if buffer[position] != rune('u') {
								goto l86
							}
							position++
							goto l85
						l86:
							position, tokenIndex, depth = position85, tokenIndex85, depth85
							if buffer[position] != rune('U') {
								goto l68
							}
							position++
						}
					l85:
						{
							position87, tokenIndex87, depth87 := position, tokenIndex, depth
							if buffer[position] != rune('n') {
								goto l88
							}
							position++
							goto l87
						l88:
							position, tokenIndex, depth = position87, tokenIndex87, depth87
							if buffer[position] != rune('N') {
								goto l68
							}
							position++
						}
					l87:
						{
							position89, tokenIndex89, depth89 := position, tokenIndex, depth
							if buffer[position] != rune('i') {
								goto l90
							}
							position++
							goto l89
						l90:
							position, tokenIndex, depth = position89, tokenIndex89, depth89
							if buffer[position] != rune('I') {
								goto l68
							}
							position++
						}
					l89:
						{
							position91, tokenIndex91, depth91 := position, tokenIndex, depth
							if buffer[position] != rune('o') {
								goto l92
							}
							position++
							goto l91
						l92:
							position, tokenIndex, depth = position91, tokenIndex91, depth91
							if buffer[position] != rune('O') {
								goto l68
							}
							position++
						}
					l91:
						{
							position93, tokenIndex93, depth93 := position, tokenIndex, depth
							if buffer[position] != rune('n') {
								goto l94
							}
							position++
							goto l93
						l94:
							position, tokenIndex, depth = position93, tokenIndex93, depth93
							if buffer[position] != rune('N') {
								goto l68
							}
							position++
						}
					l93:
						if !_rules[rulesp]() {
							goto l68
						}
						{
							position95, tokenIndex95, depth95 := position, tokenIndex, depth
							if buffer[position] != rune('a') {
								goto l96
							}
							position++
							goto l95
						l96:
							position, tokenIndex, depth = position95, tokenIndex95, depth95
							if buffer[position] != rune('A') {
								goto l68
							}
							position++
						}
					l95:
						{
							position97, tokenIndex97, depth97 := position, tokenIndex, depth
							if buffer[position] != rune('l') {
								goto l98
							}
							position++
							goto l97
						l98:
							position, tokenIndex, depth = position97, tokenIndex97, depth97
							if buffer[position] != rune('L') {
								goto l68
							}
							position++
						}
					l97:
						{
							position99, tokenIndex99, depth99 := position, tokenIndex, depth
							if buffer[position] != rune('l') {
								goto l100
							}
							position++
							goto l99
						l100:
							position, tokenIndex, depth = position99, tokenIndex99, depth99
							if buffer[position] != rune('L') {
								goto l68
							}
							position++
						}
					l99:
						if !_rules[rulesp]() {
							goto l68
						}
						if !_rules[ruleSelectStmt]() {
							goto l68
						}
						goto l67
					l68:
						position, tokenIndex, depth = position68, tokenIndex68, depth68
					}
					depth--
					add(rulePegText, position66)
				}
				if !_rules[ruleAction3]() {
					goto l64
				}
				depth--
				add(ruleSelectUnionStmt, position65)
			}
			return true
		l64:
			position, tokenIndex, depth = position64, tokenIndex64, depth64
			return false
		},
		/* 10 CreateStreamAsSelectStmt <- <(('c' / 'C') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('t' / 'T') ('e' / 'E') sp (('s' / 'S') ('t' / 'T') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('m' / 'M')) sp StreamIdentifier sp (('a' / 'A') ('s' / 'S')) sp SelectStmt Action4)> */
		func() bool {
			position101, tokenIndex101, depth101 := position, tokenIndex, depth
			{
				position102 := position
				depth++
				{
					position103, tokenIndex103, depth103 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l104
					}
					position++
					goto l103
				l104:
					position, tokenIndex, depth = position103, tokenIndex103, depth103
					if buffer[position] != rune('C') {
						goto l101
					}
					position++
				}
			l103:
				{
					position105, tokenIndex105, depth105 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l106
					}
					position++
					goto l105
				l106:
					position, tokenIndex, depth = position105, tokenIndex105, depth105
					if buffer[position] != rune('R') {
						goto l101
					}
					position++
				}
			l105:
				{
					position107, tokenIndex107, depth107 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l108
					}
					position++
					goto l107
				l108:
					position, tokenIndex, depth = position107, tokenIndex107, depth107
					if buffer[position] != rune('E') {
						goto l101
					}
					position++
				}
			l107:
				{
					position109, tokenIndex109, depth109 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l110
					}
					position++
					goto l109
				l110:
					position, tokenIndex, depth = position109, tokenIndex109, depth109
					if buffer[position] != rune('A') {
						goto l101
					}
					position++
				}
			l109:
				{
					position111, tokenIndex111, depth111 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l112
					}
					position++
					goto l111
				l112:
					position, tokenIndex, depth = position111, tokenIndex111, depth111
					if buffer[position] != rune('T') {
						goto l101
					}
					position++
				}
			l111:
				{
					position113, tokenIndex113, depth113 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l114
					}
					position++
					goto l113
				l114:
					position, tokenIndex, depth = position113, tokenIndex113, depth113
					if buffer[position] != rune('E') {
						goto l101
					}
					position++
				}
			l113:
				if !_rules[rulesp]() {
					goto l101
				}
				{
					position115, tokenIndex115, depth115 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l116
					}
					position++
					goto l115
				l116:
					position, tokenIndex, depth = position115, tokenIndex115, depth115
					if buffer[position] != rune('S') {
						goto l101
					}
					position++
				}
			l115:
				{
					position117, tokenIndex117, depth117 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l118
					}
					position++
					goto l117
				l118:
					position, tokenIndex, depth = position117, tokenIndex117, depth117
					if buffer[position] != rune('T') {
						goto l101
					}
					position++
				}
			l117:
				{
					position119, tokenIndex119, depth119 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l120
					}
					position++
					goto l119
				l120:
					position, tokenIndex, depth = position119, tokenIndex119, depth119
					if buffer[position] != rune('R') {
						goto l101
					}
					position++
				}
			l119:
				{
					position121, tokenIndex121, depth121 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l122
					}
					position++
					goto l121
				l122:
					position, tokenIndex, depth = position121, tokenIndex121, depth121
					if buffer[position] != rune('E') {
						goto l101
					}
					position++
				}
			l121:
				{
					position123, tokenIndex123, depth123 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l124
					}
					position++
					goto l123
				l124:
					position, tokenIndex, depth = position123, tokenIndex123, depth123
					if buffer[position] != rune('A') {
						goto l101
					}
					position++
				}
			l123:
				{
					position125, tokenIndex125, depth125 := position, tokenIndex, depth
					if buffer[position] != rune('m') {
						goto l126
					}
					position++
					goto l125
				l126:
					position, tokenIndex, depth = position125, tokenIndex125, depth125
					if buffer[position] != rune('M') {
						goto l101
					}
					position++
				}
			l125:
				if !_rules[rulesp]() {
					goto l101
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l101
				}
				if !_rules[rulesp]() {
					goto l101
				}
				{
					position127, tokenIndex127, depth127 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l128
					}
					position++
					goto l127
				l128:
					position, tokenIndex, depth = position127, tokenIndex127, depth127
					if buffer[position] != rune('A') {
						goto l101
					}
					position++
				}
			l127:
				{
					position129, tokenIndex129, depth129 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l130
					}
					position++
					goto l129
				l130:
					position, tokenIndex, depth = position129, tokenIndex129, depth129
					if buffer[position] != rune('S') {
						goto l101
					}
					position++
				}
			l129:
				if !_rules[rulesp]() {
					goto l101
				}
				if !_rules[ruleSelectStmt]() {
					goto l101
				}
				if !_rules[ruleAction4]() {
					goto l101
				}
				depth--
				add(ruleCreateStreamAsSelectStmt, position102)
			}
			return true
		l101:
			position, tokenIndex, depth = position101, tokenIndex101, depth101
			return false
		},
		/* 11 CreateStreamAsSelectUnionStmt <- <(('c' / 'C') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('t' / 'T') ('e' / 'E') sp (('s' / 'S') ('t' / 'T') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('m' / 'M')) sp StreamIdentifier sp (('a' / 'A') ('s' / 'S')) sp SelectUnionStmt Action5)> */
		func() bool {
			position131, tokenIndex131, depth131 := position, tokenIndex, depth
			{
				position132 := position
				depth++
				{
					position133, tokenIndex133, depth133 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l134
					}
					position++
					goto l133
				l134:
					position, tokenIndex, depth = position133, tokenIndex133, depth133
					if buffer[position] != rune('C') {
						goto l131
					}
					position++
				}
			l133:
				{
					position135, tokenIndex135, depth135 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l136
					}
					position++
					goto l135
				l136:
					position, tokenIndex, depth = position135, tokenIndex135, depth135
					if buffer[position] != rune('R') {
						goto l131
					}
					position++
				}
			l135:
				{
					position137, tokenIndex137, depth137 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l138
					}
					position++
					goto l137
				l138:
					position, tokenIndex, depth = position137, tokenIndex137, depth137
					if buffer[position] != rune('E') {
						goto l131
					}
					position++
				}
			l137:
				{
					position139, tokenIndex139, depth139 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l140
					}
					position++
					goto l139
				l140:
					position, tokenIndex, depth = position139, tokenIndex139, depth139
					if buffer[position] != rune('A') {
						goto l131
					}
					position++
				}
			l139:
				{
					position141, tokenIndex141, depth141 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l142
					}
					position++
					goto l141
				l142:
					position, tokenIndex, depth = position141, tokenIndex141, depth141
					if buffer[position] != rune('T') {
						goto l131
					}
					position++
				}
			l141:
				{
					position143, tokenIndex143, depth143 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l144
					}
					position++
					goto l143
				l144:
					position, tokenIndex, depth = position143, tokenIndex143, depth143
					if buffer[position] != rune('E') {
						goto l131
					}
					position++
				}
			l143:
				if !_rules[rulesp]() {
					goto l131
				}
				{
					position145, tokenIndex145, depth145 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l146
					}
					position++
					goto l145
				l146:
					position, tokenIndex, depth = position145, tokenIndex145, depth145
					if buffer[position] != rune('S') {
						goto l131
					}
					position++
				}
			l145:
				{
					position147, tokenIndex147, depth147 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l148
					}
					position++
					goto l147
				l148:
					position, tokenIndex, depth = position147, tokenIndex147, depth147
					if buffer[position] != rune('T') {
						goto l131
					}
					position++
				}
			l147:
				{
					position149, tokenIndex149, depth149 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l150
					}
					position++
					goto l149
				l150:
					position, tokenIndex, depth = position149, tokenIndex149, depth149
					if buffer[position] != rune('R') {
						goto l131
					}
					position++
				}
			l149:
				{
					position151, tokenIndex151, depth151 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l152
					}
					position++
					goto l151
				l152:
					position, tokenIndex, depth = position151, tokenIndex151, depth151
					if buffer[position] != rune('E') {
						goto l131
					}
					position++
				}
			l151:
				{
					position153, tokenIndex153, depth153 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l154
					}
					position++
					goto l153
				l154:
					position, tokenIndex, depth = position153, tokenIndex153, depth153
					if buffer[position] != rune('A') {
						goto l131
					}
					position++
				}
			l153:
				{
					position155, tokenIndex155, depth155 := position, tokenIndex, depth
					if buffer[position] != rune('m') {
						goto l156
					}
					position++
					goto l155
				l156:
					position, tokenIndex, depth = position155, tokenIndex155, depth155
					if buffer[position] != rune('M') {
						goto l131
					}
					position++
				}
			l155:
				if !_rules[rulesp]() {
					goto l131
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l131
				}
				if !_rules[rulesp]() {
					goto l131
				}
				{
					position157, tokenIndex157, depth157 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l158
					}
					position++
					goto l157
				l158:
					position, tokenIndex, depth = position157, tokenIndex157, depth157
					if buffer[position] != rune('A') {
						goto l131
					}
					position++
				}
			l157:
				{
					position159, tokenIndex159, depth159 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l160
					}
					position++
					goto l159
				l160:
					position, tokenIndex, depth = position159, tokenIndex159, depth159
					if buffer[position] != rune('S') {
						goto l131
					}
					position++
				}
			l159:
				if !_rules[rulesp]() {
					goto l131
				}
				if !_rules[ruleSelectUnionStmt]() {
					goto l131
				}
				if !_rules[ruleAction5]() {
					goto l131
				}
				depth--
				add(ruleCreateStreamAsSelectUnionStmt, position132)
			}
			return true
		l131:
			position, tokenIndex, depth = position131, tokenIndex131, depth131
			return false
		},
		/* 12 CreateSourceStmt <- <(('c' / 'C') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('t' / 'T') ('e' / 'E') PausedOpt sp (('s' / 'S') ('o' / 'O') ('u' / 'U') ('r' / 'R') ('c' / 'C') ('e' / 'E')) sp StreamIdentifier sp (('t' / 'T') ('y' / 'Y') ('p' / 'P') ('e' / 'E')) sp SourceSinkType SourceSinkSpecs Action6)> */
		func() bool {
			position161, tokenIndex161, depth161 := position, tokenIndex, depth
			{
				position162 := position
				depth++
				{
					position163, tokenIndex163, depth163 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l164
					}
					position++
					goto l163
				l164:
					position, tokenIndex, depth = position163, tokenIndex163, depth163
					if buffer[position] != rune('C') {
						goto l161
					}
					position++
				}
			l163:
				{
					position165, tokenIndex165, depth165 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l166
					}
					position++
					goto l165
				l166:
					position, tokenIndex, depth = position165, tokenIndex165, depth165
					if buffer[position] != rune('R') {
						goto l161
					}
					position++
				}
			l165:
				{
					position167, tokenIndex167, depth167 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l168
					}
					position++
					goto l167
				l168:
					position, tokenIndex, depth = position167, tokenIndex167, depth167
					if buffer[position] != rune('E') {
						goto l161
					}
					position++
				}
			l167:
				{
					position169, tokenIndex169, depth169 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l170
					}
					position++
					goto l169
				l170:
					position, tokenIndex, depth = position169, tokenIndex169, depth169
					if buffer[position] != rune('A') {
						goto l161
					}
					position++
				}
			l169:
				{
					position171, tokenIndex171, depth171 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l172
					}
					position++
					goto l171
				l172:
					position, tokenIndex, depth = position171, tokenIndex171, depth171
					if buffer[position] != rune('T') {
						goto l161
					}
					position++
				}
			l171:
				{
					position173, tokenIndex173, depth173 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l174
					}
					position++
					goto l173
				l174:
					position, tokenIndex, depth = position173, tokenIndex173, depth173
					if buffer[position] != rune('E') {
						goto l161
					}
					position++
				}
			l173:
				if !_rules[rulePausedOpt]() {
					goto l161
				}
				if !_rules[rulesp]() {
					goto l161
				}
				{
					position175, tokenIndex175, depth175 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l176
					}
					position++
					goto l175
				l176:
					position, tokenIndex, depth = position175, tokenIndex175, depth175
					if buffer[position] != rune('S') {
						goto l161
					}
					position++
				}
			l175:
				{
					position177, tokenIndex177, depth177 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l178
					}
					position++
					goto l177
				l178:
					position, tokenIndex, depth = position177, tokenIndex177, depth177
					if buffer[position] != rune('O') {
						goto l161
					}
					position++
				}
			l177:
				{
					position179, tokenIndex179, depth179 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l180
					}
					position++
					goto l179
				l180:
					position, tokenIndex, depth = position179, tokenIndex179, depth179
					if buffer[position] != rune('U') {
						goto l161
					}
					position++
				}
			l179:
				{
					position181, tokenIndex181, depth181 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l182
					}
					position++
					goto l181
				l182:
					position, tokenIndex, depth = position181, tokenIndex181, depth181
					if buffer[position] != rune('R') {
						goto l161
					}
					position++
				}
			l181:
				{
					position183, tokenIndex183, depth183 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l184
					}
					position++
					goto l183
				l184:
					position, tokenIndex, depth = position183, tokenIndex183, depth183
					if buffer[position] != rune('C') {
						goto l161
					}
					position++
				}
			l183:
				{
					position185, tokenIndex185, depth185 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l186
					}
					position++
					goto l185
				l186:
					position, tokenIndex, depth = position185, tokenIndex185, depth185
					if buffer[position] != rune('E') {
						goto l161
					}
					position++
				}
			l185:
				if !_rules[rulesp]() {
					goto l161
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l161
				}
				if !_rules[rulesp]() {
					goto l161
				}
				{
					position187, tokenIndex187, depth187 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l188
					}
					position++
					goto l187
				l188:
					position, tokenIndex, depth = position187, tokenIndex187, depth187
					if buffer[position] != rune('T') {
						goto l161
					}
					position++
				}
			l187:
				{
					position189, tokenIndex189, depth189 := position, tokenIndex, depth
					if buffer[position] != rune('y') {
						goto l190
					}
					position++
					goto l189
				l190:
					position, tokenIndex, depth = position189, tokenIndex189, depth189
					if buffer[position] != rune('Y') {
						goto l161
					}
					position++
				}
			l189:
				{
					position191, tokenIndex191, depth191 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l192
					}
					position++
					goto l191
				l192:
					position, tokenIndex, depth = position191, tokenIndex191, depth191
					if buffer[position] != rune('P') {
						goto l161
					}
					position++
				}
			l191:
				{
					position193, tokenIndex193, depth193 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l194
					}
					position++
					goto l193
				l194:
					position, tokenIndex, depth = position193, tokenIndex193, depth193
					if buffer[position] != rune('E') {
						goto l161
					}
					position++
				}
			l193:
				if !_rules[rulesp]() {
					goto l161
				}
				if !_rules[ruleSourceSinkType]() {
					goto l161
				}
				if !_rules[ruleSourceSinkSpecs]() {
					goto l161
				}
				if !_rules[ruleAction6]() {
					goto l161
				}
				depth--
				add(ruleCreateSourceStmt, position162)
			}
			return true
		l161:
			position, tokenIndex, depth = position161, tokenIndex161, depth161
			return false
		},
		/* 13 CreateSinkStmt <- <(('c' / 'C') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('t' / 'T') ('e' / 'E') sp (('s' / 'S') ('i' / 'I') ('n' / 'N') ('k' / 'K')) sp StreamIdentifier sp (('t' / 'T') ('y' / 'Y') ('p' / 'P') ('e' / 'E')) sp SourceSinkType SourceSinkSpecs Action7)> */
		func() bool {
			position195, tokenIndex195, depth195 := position, tokenIndex, depth
			{
				position196 := position
				depth++
				{
					position197, tokenIndex197, depth197 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l198
					}
					position++
					goto l197
				l198:
					position, tokenIndex, depth = position197, tokenIndex197, depth197
					if buffer[position] != rune('C') {
						goto l195
					}
					position++
				}
			l197:
				{
					position199, tokenIndex199, depth199 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l200
					}
					position++
					goto l199
				l200:
					position, tokenIndex, depth = position199, tokenIndex199, depth199
					if buffer[position] != rune('R') {
						goto l195
					}
					position++
				}
			l199:
				{
					position201, tokenIndex201, depth201 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l202
					}
					position++
					goto l201
				l202:
					position, tokenIndex, depth = position201, tokenIndex201, depth201
					if buffer[position] != rune('E') {
						goto l195
					}
					position++
				}
			l201:
				{
					position203, tokenIndex203, depth203 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l204
					}
					position++
					goto l203
				l204:
					position, tokenIndex, depth = position203, tokenIndex203, depth203
					if buffer[position] != rune('A') {
						goto l195
					}
					position++
				}
			l203:
				{
					position205, tokenIndex205, depth205 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l206
					}
					position++
					goto l205
				l206:
					position, tokenIndex, depth = position205, tokenIndex205, depth205
					if buffer[position] != rune('T') {
						goto l195
					}
					position++
				}
			l205:
				{
					position207, tokenIndex207, depth207 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l208
					}
					position++
					goto l207
				l208:
					position, tokenIndex, depth = position207, tokenIndex207, depth207
					if buffer[position] != rune('E') {
						goto l195
					}
					position++
				}
			l207:
				if !_rules[rulesp]() {
					goto l195
				}
				{
					position209, tokenIndex209, depth209 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l210
					}
					position++
					goto l209
				l210:
					position, tokenIndex, depth = position209, tokenIndex209, depth209
					if buffer[position] != rune('S') {
						goto l195
					}
					position++
				}
			l209:
				{
					position211, tokenIndex211, depth211 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l212
					}
					position++
					goto l211
				l212:
					position, tokenIndex, depth = position211, tokenIndex211, depth211
					if buffer[position] != rune('I') {
						goto l195
					}
					position++
				}
			l211:
				{
					position213, tokenIndex213, depth213 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l214
					}
					position++
					goto l213
				l214:
					position, tokenIndex, depth = position213, tokenIndex213, depth213
					if buffer[position] != rune('N') {
						goto l195
					}
					position++
				}
			l213:
				{
					position215, tokenIndex215, depth215 := position, tokenIndex, depth
					if buffer[position] != rune('k') {
						goto l216
					}
					position++
					goto l215
				l216:
					position, tokenIndex, depth = position215, tokenIndex215, depth215
					if buffer[position] != rune('K') {
						goto l195
					}
					position++
				}
			l215:
				if !_rules[rulesp]() {
					goto l195
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l195
				}
				if !_rules[rulesp]() {
					goto l195
				}
				{
					position217, tokenIndex217, depth217 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l218
					}
					position++
					goto l217
				l218:
					position, tokenIndex, depth = position217, tokenIndex217, depth217
					if buffer[position] != rune('T') {
						goto l195
					}
					position++
				}
			l217:
				{
					position219, tokenIndex219, depth219 := position, tokenIndex, depth
					if buffer[position] != rune('y') {
						goto l220
					}
					position++
					goto l219
				l220:
					position, tokenIndex, depth = position219, tokenIndex219, depth219
					if buffer[position] != rune('Y') {
						goto l195
					}
					position++
				}
			l219:
				{
					position221, tokenIndex221, depth221 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l222
					}
					position++
					goto l221
				l222:
					position, tokenIndex, depth = position221, tokenIndex221, depth221
					if buffer[position] != rune('P') {
						goto l195
					}
					position++
				}
			l221:
				{
					position223, tokenIndex223, depth223 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l224
					}
					position++
					goto l223
				l224:
					position, tokenIndex, depth = position223, tokenIndex223, depth223
					if buffer[position] != rune('E') {
						goto l195
					}
					position++
				}
			l223:
				if !_rules[rulesp]() {
					goto l195
				}
				if !_rules[ruleSourceSinkType]() {
					goto l195
				}
				if !_rules[ruleSourceSinkSpecs]() {
					goto l195
				}
				if !_rules[ruleAction7]() {
					goto l195
				}
				depth--
				add(ruleCreateSinkStmt, position196)
			}
			return true
		l195:
			position, tokenIndex, depth = position195, tokenIndex195, depth195
			return false
		},
		/* 14 CreateStateStmt <- <(('c' / 'C') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('t' / 'T') ('e' / 'E') sp (('s' / 'S') ('t' / 'T') ('a' / 'A') ('t' / 'T') ('e' / 'E')) sp StreamIdentifier sp (('t' / 'T') ('y' / 'Y') ('p' / 'P') ('e' / 'E')) sp SourceSinkType SourceSinkSpecs Action8)> */
		func() bool {
			position225, tokenIndex225, depth225 := position, tokenIndex, depth
			{
				position226 := position
				depth++
				{
					position227, tokenIndex227, depth227 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l228
					}
					position++
					goto l227
				l228:
					position, tokenIndex, depth = position227, tokenIndex227, depth227
					if buffer[position] != rune('C') {
						goto l225
					}
					position++
				}
			l227:
				{
					position229, tokenIndex229, depth229 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l230
					}
					position++
					goto l229
				l230:
					position, tokenIndex, depth = position229, tokenIndex229, depth229
					if buffer[position] != rune('R') {
						goto l225
					}
					position++
				}
			l229:
				{
					position231, tokenIndex231, depth231 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l232
					}
					position++
					goto l231
				l232:
					position, tokenIndex, depth = position231, tokenIndex231, depth231
					if buffer[position] != rune('E') {
						goto l225
					}
					position++
				}
			l231:
				{
					position233, tokenIndex233, depth233 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l234
					}
					position++
					goto l233
				l234:
					position, tokenIndex, depth = position233, tokenIndex233, depth233
					if buffer[position] != rune('A') {
						goto l225
					}
					position++
				}
			l233:
				{
					position235, tokenIndex235, depth235 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l236
					}
					position++
					goto l235
				l236:
					position, tokenIndex, depth = position235, tokenIndex235, depth235
					if buffer[position] != rune('T') {
						goto l225
					}
					position++
				}
			l235:
				{
					position237, tokenIndex237, depth237 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l238
					}
					position++
					goto l237
				l238:
					position, tokenIndex, depth = position237, tokenIndex237, depth237
					if buffer[position] != rune('E') {
						goto l225
					}
					position++
				}
			l237:
				if !_rules[rulesp]() {
					goto l225
				}
				{
					position239, tokenIndex239, depth239 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l240
					}
					position++
					goto l239
				l240:
					position, tokenIndex, depth = position239, tokenIndex239, depth239
					if buffer[position] != rune('S') {
						goto l225
					}
					position++
				}
			l239:
				{
					position241, tokenIndex241, depth241 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l242
					}
					position++
					goto l241
				l242:
					position, tokenIndex, depth = position241, tokenIndex241, depth241
					if buffer[position] != rune('T') {
						goto l225
					}
					position++
				}
			l241:
				{
					position243, tokenIndex243, depth243 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l244
					}
					position++
					goto l243
				l244:
					position, tokenIndex, depth = position243, tokenIndex243, depth243
					if buffer[position] != rune('A') {
						goto l225
					}
					position++
				}
			l243:
				{
					position245, tokenIndex245, depth245 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l246
					}
					position++
					goto l245
				l246:
					position, tokenIndex, depth = position245, tokenIndex245, depth245
					if buffer[position] != rune('T') {
						goto l225
					}
					position++
				}
			l245:
				{
					position247, tokenIndex247, depth247 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l248
					}
					position++
					goto l247
				l248:
					position, tokenIndex, depth = position247, tokenIndex247, depth247
					if buffer[position] != rune('E') {
						goto l225
					}
					position++
				}
			l247:
				if !_rules[rulesp]() {
					goto l225
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l225
				}
				if !_rules[rulesp]() {
					goto l225
				}
				{
					position249, tokenIndex249, depth249 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l250
					}
					position++
					goto l249
				l250:
					position, tokenIndex, depth = position249, tokenIndex249, depth249
					if buffer[position] != rune('T') {
						goto l225
					}
					position++
				}
			l249:
				{
					position251, tokenIndex251, depth251 := position, tokenIndex, depth
					if buffer[position] != rune('y') {
						goto l252
					}
					position++
					goto l251
				l252:
					position, tokenIndex, depth = position251, tokenIndex251, depth251
					if buffer[position] != rune('Y') {
						goto l225
					}
					position++
				}
			l251:
				{
					position253, tokenIndex253, depth253 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l254
					}
					position++
					goto l253
				l254:
					position, tokenIndex, depth = position253, tokenIndex253, depth253
					if buffer[position] != rune('P') {
						goto l225
					}
					position++
				}
			l253:
				{
					position255, tokenIndex255, depth255 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l256
					}
					position++
					goto l255
				l256:
					position, tokenIndex, depth = position255, tokenIndex255, depth255
					if buffer[position] != rune('E') {
						goto l225
					}
					position++
				}
			l255:
				if !_rules[rulesp]() {
					goto l225
				}
				if !_rules[ruleSourceSinkType]() {
					goto l225
				}
				if !_rules[ruleSourceSinkSpecs]() {
					goto l225
				}
				if !_rules[ruleAction8]() {
					goto l225
				}
				depth--
				add(ruleCreateStateStmt, position226)
			}
			return true
		l225:
			position, tokenIndex, depth = position225, tokenIndex225, depth225
			return false
		},
		/* 15 UpdateStateStmt <- <(('u' / 'U') ('p' / 'P') ('d' / 'D') ('a' / 'A') ('t' / 'T') ('e' / 'E') sp (('s' / 'S') ('t' / 'T') ('a' / 'A') ('t' / 'T') ('e' / 'E')) sp StreamIdentifier UpdateSourceSinkSpecs Action9)> */
		func() bool {
			position257, tokenIndex257, depth257 := position, tokenIndex, depth
			{
				position258 := position
				depth++
				{
					position259, tokenIndex259, depth259 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l260
					}
					position++
					goto l259
				l260:
					position, tokenIndex, depth = position259, tokenIndex259, depth259
					if buffer[position] != rune('U') {
						goto l257
					}
					position++
				}
			l259:
				{
					position261, tokenIndex261, depth261 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l262
					}
					position++
					goto l261
				l262:
					position, tokenIndex, depth = position261, tokenIndex261, depth261
					if buffer[position] != rune('P') {
						goto l257
					}
					position++
				}
			l261:
				{
					position263, tokenIndex263, depth263 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l264
					}
					position++
					goto l263
				l264:
					position, tokenIndex, depth = position263, tokenIndex263, depth263
					if buffer[position] != rune('D') {
						goto l257
					}
					position++
				}
			l263:
				{
					position265, tokenIndex265, depth265 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l266
					}
					position++
					goto l265
				l266:
					position, tokenIndex, depth = position265, tokenIndex265, depth265
					if buffer[position] != rune('A') {
						goto l257
					}
					position++
				}
			l265:
				{
					position267, tokenIndex267, depth267 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l268
					}
					position++
					goto l267
				l268:
					position, tokenIndex, depth = position267, tokenIndex267, depth267
					if buffer[position] != rune('T') {
						goto l257
					}
					position++
				}
			l267:
				{
					position269, tokenIndex269, depth269 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l270
					}
					position++
					goto l269
				l270:
					position, tokenIndex, depth = position269, tokenIndex269, depth269
					if buffer[position] != rune('E') {
						goto l257
					}
					position++
				}
			l269:
				if !_rules[rulesp]() {
					goto l257
				}
				{
					position271, tokenIndex271, depth271 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l272
					}
					position++
					goto l271
				l272:
					position, tokenIndex, depth = position271, tokenIndex271, depth271
					if buffer[position] != rune('S') {
						goto l257
					}
					position++
				}
			l271:
				{
					position273, tokenIndex273, depth273 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l274
					}
					position++
					goto l273
				l274:
					position, tokenIndex, depth = position273, tokenIndex273, depth273
					if buffer[position] != rune('T') {
						goto l257
					}
					position++
				}
			l273:
				{
					position275, tokenIndex275, depth275 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l276
					}
					position++
					goto l275
				l276:
					position, tokenIndex, depth = position275, tokenIndex275, depth275
					if buffer[position] != rune('A') {
						goto l257
					}
					position++
				}
			l275:
				{
					position277, tokenIndex277, depth277 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l278
					}
					position++
					goto l277
				l278:
					position, tokenIndex, depth = position277, tokenIndex277, depth277
					if buffer[position] != rune('T') {
						goto l257
					}
					position++
				}
			l277:
				{
					position279, tokenIndex279, depth279 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l280
					}
					position++
					goto l279
				l280:
					position, tokenIndex, depth = position279, tokenIndex279, depth279
					if buffer[position] != rune('E') {
						goto l257
					}
					position++
				}
			l279:
				if !_rules[rulesp]() {
					goto l257
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l257
				}
				if !_rules[ruleUpdateSourceSinkSpecs]() {
					goto l257
				}
				if !_rules[ruleAction9]() {
					goto l257
				}
				depth--
				add(ruleUpdateStateStmt, position258)
			}
			return true
		l257:
			position, tokenIndex, depth = position257, tokenIndex257, depth257
			return false
		},
		/* 16 UpdateSourceStmt <- <(('u' / 'U') ('p' / 'P') ('d' / 'D') ('a' / 'A') ('t' / 'T') ('e' / 'E') sp (('s' / 'S') ('o' / 'O') ('u' / 'U') ('r' / 'R') ('c' / 'C') ('e' / 'E')) sp StreamIdentifier UpdateSourceSinkSpecs Action10)> */
		func() bool {
			position281, tokenIndex281, depth281 := position, tokenIndex, depth
			{
				position282 := position
				depth++
				{
					position283, tokenIndex283, depth283 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l284
					}
					position++
					goto l283
				l284:
					position, tokenIndex, depth = position283, tokenIndex283, depth283
					if buffer[position] != rune('U') {
						goto l281
					}
					position++
				}
			l283:
				{
					position285, tokenIndex285, depth285 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l286
					}
					position++
					goto l285
				l286:
					position, tokenIndex, depth = position285, tokenIndex285, depth285
					if buffer[position] != rune('P') {
						goto l281
					}
					position++
				}
			l285:
				{
					position287, tokenIndex287, depth287 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l288
					}
					position++
					goto l287
				l288:
					position, tokenIndex, depth = position287, tokenIndex287, depth287
					if buffer[position] != rune('D') {
						goto l281
					}
					position++
				}
			l287:
				{
					position289, tokenIndex289, depth289 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l290
					}
					position++
					goto l289
				l290:
					position, tokenIndex, depth = position289, tokenIndex289, depth289
					if buffer[position] != rune('A') {
						goto l281
					}
					position++
				}
			l289:
				{
					position291, tokenIndex291, depth291 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l292
					}
					position++
					goto l291
				l292:
					position, tokenIndex, depth = position291, tokenIndex291, depth291
					if buffer[position] != rune('T') {
						goto l281
					}
					position++
				}
			l291:
				{
					position293, tokenIndex293, depth293 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l294
					}
					position++
					goto l293
				l294:
					position, tokenIndex, depth = position293, tokenIndex293, depth293
					if buffer[position] != rune('E') {
						goto l281
					}
					position++
				}
			l293:
				if !_rules[rulesp]() {
					goto l281
				}
				{
					position295, tokenIndex295, depth295 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l296
					}
					position++
					goto l295
				l296:
					position, tokenIndex, depth = position295, tokenIndex295, depth295
					if buffer[position] != rune('S') {
						goto l281
					}
					position++
				}
			l295:
				{
					position297, tokenIndex297, depth297 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l298
					}
					position++
					goto l297
				l298:
					position, tokenIndex, depth = position297, tokenIndex297, depth297
					if buffer[position] != rune('O') {
						goto l281
					}
					position++
				}
			l297:
				{
					position299, tokenIndex299, depth299 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l300
					}
					position++
					goto l299
				l300:
					position, tokenIndex, depth = position299, tokenIndex299, depth299
					if buffer[position] != rune('U') {
						goto l281
					}
					position++
				}
			l299:
				{
					position301, tokenIndex301, depth301 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l302
					}
					position++
					goto l301
				l302:
					position, tokenIndex, depth = position301, tokenIndex301, depth301
					if buffer[position] != rune('R') {
						goto l281
					}
					position++
				}
			l301:
				{
					position303, tokenIndex303, depth303 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l304
					}
					position++
					goto l303
				l304:
					position, tokenIndex, depth = position303, tokenIndex303, depth303
					if buffer[position] != rune('C') {
						goto l281
					}
					position++
				}
			l303:
				{
					position305, tokenIndex305, depth305 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l306
					}
					position++
					goto l305
				l306:
					position, tokenIndex, depth = position305, tokenIndex305, depth305
					if buffer[position] != rune('E') {
						goto l281
					}
					position++
				}
			l305:
				if !_rules[rulesp]() {
					goto l281
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l281
				}
				if !_rules[ruleUpdateSourceSinkSpecs]() {
					goto l281
				}
				if !_rules[ruleAction10]() {
					goto l281
				}
				depth--
				add(ruleUpdateSourceStmt, position282)
			}
			return true
		l281:
			position, tokenIndex, depth = position281, tokenIndex281, depth281
			return false
		},
		/* 17 UpdateSinkStmt <- <(('u' / 'U') ('p' / 'P') ('d' / 'D') ('a' / 'A') ('t' / 'T') ('e' / 'E') sp (('s' / 'S') ('i' / 'I') ('n' / 'N') ('k' / 'K')) sp StreamIdentifier UpdateSourceSinkSpecs Action11)> */
		func() bool {
			position307, tokenIndex307, depth307 := position, tokenIndex, depth
			{
				position308 := position
				depth++
				{
					position309, tokenIndex309, depth309 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l310
					}
					position++
					goto l309
				l310:
					position, tokenIndex, depth = position309, tokenIndex309, depth309
					if buffer[position] != rune('U') {
						goto l307
					}
					position++
				}
			l309:
				{
					position311, tokenIndex311, depth311 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l312
					}
					position++
					goto l311
				l312:
					position, tokenIndex, depth = position311, tokenIndex311, depth311
					if buffer[position] != rune('P') {
						goto l307
					}
					position++
				}
			l311:
				{
					position313, tokenIndex313, depth313 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l314
					}
					position++
					goto l313
				l314:
					position, tokenIndex, depth = position313, tokenIndex313, depth313
					if buffer[position] != rune('D') {
						goto l307
					}
					position++
				}
			l313:
				{
					position315, tokenIndex315, depth315 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l316
					}
					position++
					goto l315
				l316:
					position, tokenIndex, depth = position315, tokenIndex315, depth315
					if buffer[position] != rune('A') {
						goto l307
					}
					position++
				}
			l315:
				{
					position317, tokenIndex317, depth317 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l318
					}
					position++
					goto l317
				l318:
					position, tokenIndex, depth = position317, tokenIndex317, depth317
					if buffer[position] != rune('T') {
						goto l307
					}
					position++
				}
			l317:
				{
					position319, tokenIndex319, depth319 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l320
					}
					position++
					goto l319
				l320:
					position, tokenIndex, depth = position319, tokenIndex319, depth319
					if buffer[position] != rune('E') {
						goto l307
					}
					position++
				}
			l319:
				if !_rules[rulesp]() {
					goto l307
				}
				{
					position321, tokenIndex321, depth321 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l322
					}
					position++
					goto l321
				l322:
					position, tokenIndex, depth = position321, tokenIndex321, depth321
					if buffer[position] != rune('S') {
						goto l307
					}
					position++
				}
			l321:
				{
					position323, tokenIndex323, depth323 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l324
					}
					position++
					goto l323
				l324:
					position, tokenIndex, depth = position323, tokenIndex323, depth323
					if buffer[position] != rune('I') {
						goto l307
					}
					position++
				}
			l323:
				{
					position325, tokenIndex325, depth325 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l326
					}
					position++
					goto l325
				l326:
					position, tokenIndex, depth = position325, tokenIndex325, depth325
					if buffer[position] != rune('N') {
						goto l307
					}
					position++
				}
			l325:
				{
					position327, tokenIndex327, depth327 := position, tokenIndex, depth
					if buffer[position] != rune('k') {
						goto l328
					}
					position++
					goto l327
				l328:
					position, tokenIndex, depth = position327, tokenIndex327, depth327
					if buffer[position] != rune('K') {
						goto l307
					}
					position++
				}
			l327:
				if !_rules[rulesp]() {
					goto l307
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l307
				}
				if !_rules[ruleUpdateSourceSinkSpecs]() {
					goto l307
				}
				if !_rules[ruleAction11]() {
					goto l307
				}
				depth--
				add(ruleUpdateSinkStmt, position308)
			}
			return true
		l307:
			position, tokenIndex, depth = position307, tokenIndex307, depth307
			return false
		},
		/* 18 InsertIntoSelectStmt <- <(('i' / 'I') ('n' / 'N') ('s' / 'S') ('e' / 'E') ('r' / 'R') ('t' / 'T') sp (('i' / 'I') ('n' / 'N') ('t' / 'T') ('o' / 'O')) sp StreamIdentifier sp SelectStmt Action12)> */
		func() bool {
			position329, tokenIndex329, depth329 := position, tokenIndex, depth
			{
				position330 := position
				depth++
				{
					position331, tokenIndex331, depth331 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l332
					}
					position++
					goto l331
				l332:
					position, tokenIndex, depth = position331, tokenIndex331, depth331
					if buffer[position] != rune('I') {
						goto l329
					}
					position++
				}
			l331:
				{
					position333, tokenIndex333, depth333 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l334
					}
					position++
					goto l333
				l334:
					position, tokenIndex, depth = position333, tokenIndex333, depth333
					if buffer[position] != rune('N') {
						goto l329
					}
					position++
				}
			l333:
				{
					position335, tokenIndex335, depth335 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l336
					}
					position++
					goto l335
				l336:
					position, tokenIndex, depth = position335, tokenIndex335, depth335
					if buffer[position] != rune('S') {
						goto l329
					}
					position++
				}
			l335:
				{
					position337, tokenIndex337, depth337 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l338
					}
					position++
					goto l337
				l338:
					position, tokenIndex, depth = position337, tokenIndex337, depth337
					if buffer[position] != rune('E') {
						goto l329
					}
					position++
				}
			l337:
				{
					position339, tokenIndex339, depth339 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l340
					}
					position++
					goto l339
				l340:
					position, tokenIndex, depth = position339, tokenIndex339, depth339
					if buffer[position] != rune('R') {
						goto l329
					}
					position++
				}
			l339:
				{
					position341, tokenIndex341, depth341 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l342
					}
					position++
					goto l341
				l342:
					position, tokenIndex, depth = position341, tokenIndex341, depth341
					if buffer[position] != rune('T') {
						goto l329
					}
					position++
				}
			l341:
				if !_rules[rulesp]() {
					goto l329
				}
				{
					position343, tokenIndex343, depth343 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l344
					}
					position++
					goto l343
				l344:
					position, tokenIndex, depth = position343, tokenIndex343, depth343
					if buffer[position] != rune('I') {
						goto l329
					}
					position++
				}
			l343:
				{
					position345, tokenIndex345, depth345 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l346
					}
					position++
					goto l345
				l346:
					position, tokenIndex, depth = position345, tokenIndex345, depth345
					if buffer[position] != rune('N') {
						goto l329
					}
					position++
				}
			l345:
				{
					position347, tokenIndex347, depth347 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l348
					}
					position++
					goto l347
				l348:
					position, tokenIndex, depth = position347, tokenIndex347, depth347
					if buffer[position] != rune('T') {
						goto l329
					}
					position++
				}
			l347:
				{
					position349, tokenIndex349, depth349 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l350
					}
					position++
					goto l349
				l350:
					position, tokenIndex, depth = position349, tokenIndex349, depth349
					if buffer[position] != rune('O') {
						goto l329
					}
					position++
				}
			l349:
				if !_rules[rulesp]() {
					goto l329
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l329
				}
				if !_rules[rulesp]() {
					goto l329
				}
				if !_rules[ruleSelectStmt]() {
					goto l329
				}
				if !_rules[ruleAction12]() {
					goto l329
				}
				depth--
				add(ruleInsertIntoSelectStmt, position330)
			}
			return true
		l329:
			position, tokenIndex, depth = position329, tokenIndex329, depth329
			return false
		},
		/* 19 InsertIntoFromStmt <- <(('i' / 'I') ('n' / 'N') ('s' / 'S') ('e' / 'E') ('r' / 'R') ('t' / 'T') sp (('i' / 'I') ('n' / 'N') ('t' / 'T') ('o' / 'O')) sp StreamIdentifier sp (('f' / 'F') ('r' / 'R') ('o' / 'O') ('m' / 'M')) sp StreamIdentifier Action13)> */
		func() bool {
			position351, tokenIndex351, depth351 := position, tokenIndex, depth
			{
				position352 := position
				depth++
				{
					position353, tokenIndex353, depth353 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l354
					}
					position++
					goto l353
				l354:
					position, tokenIndex, depth = position353, tokenIndex353, depth353
					if buffer[position] != rune('I') {
						goto l351
					}
					position++
				}
			l353:
				{
					position355, tokenIndex355, depth355 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l356
					}
					position++
					goto l355
				l356:
					position, tokenIndex, depth = position355, tokenIndex355, depth355
					if buffer[position] != rune('N') {
						goto l351
					}
					position++
				}
			l355:
				{
					position357, tokenIndex357, depth357 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l358
					}
					position++
					goto l357
				l358:
					position, tokenIndex, depth = position357, tokenIndex357, depth357
					if buffer[position] != rune('S') {
						goto l351
					}
					position++
				}
			l357:
				{
					position359, tokenIndex359, depth359 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l360
					}
					position++
					goto l359
				l360:
					position, tokenIndex, depth = position359, tokenIndex359, depth359
					if buffer[position] != rune('E') {
						goto l351
					}
					position++
				}
			l359:
				{
					position361, tokenIndex361, depth361 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l362
					}
					position++
					goto l361
				l362:
					position, tokenIndex, depth = position361, tokenIndex361, depth361
					if buffer[position] != rune('R') {
						goto l351
					}
					position++
				}
			l361:
				{
					position363, tokenIndex363, depth363 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l364
					}
					position++
					goto l363
				l364:
					position, tokenIndex, depth = position363, tokenIndex363, depth363
					if buffer[position] != rune('T') {
						goto l351
					}
					position++
				}
			l363:
				if !_rules[rulesp]() {
					goto l351
				}
				{
					position365, tokenIndex365, depth365 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l366
					}
					position++
					goto l365
				l366:
					position, tokenIndex, depth = position365, tokenIndex365, depth365
					if buffer[position] != rune('I') {
						goto l351
					}
					position++
				}
			l365:
				{
					position367, tokenIndex367, depth367 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l368
					}
					position++
					goto l367
				l368:
					position, tokenIndex, depth = position367, tokenIndex367, depth367
					if buffer[position] != rune('N') {
						goto l351
					}
					position++
				}
			l367:
				{
					position369, tokenIndex369, depth369 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l370
					}
					position++
					goto l369
				l370:
					position, tokenIndex, depth = position369, tokenIndex369, depth369
					if buffer[position] != rune('T') {
						goto l351
					}
					position++
				}
			l369:
				{
					position371, tokenIndex371, depth371 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l372
					}
					position++
					goto l371
				l372:
					position, tokenIndex, depth = position371, tokenIndex371, depth371
					if buffer[position] != rune('O') {
						goto l351
					}
					position++
				}
			l371:
				if !_rules[rulesp]() {
					goto l351
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l351
				}
				if !_rules[rulesp]() {
					goto l351
				}
				{
					position373, tokenIndex373, depth373 := position, tokenIndex, depth
					if buffer[position] != rune('f') {
						goto l374
					}
					position++
					goto l373
				l374:
					position, tokenIndex, depth = position373, tokenIndex373, depth373
					if buffer[position] != rune('F') {
						goto l351
					}
					position++
				}
			l373:
				{
					position375, tokenIndex375, depth375 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l376
					}
					position++
					goto l375
				l376:
					position, tokenIndex, depth = position375, tokenIndex375, depth375
					if buffer[position] != rune('R') {
						goto l351
					}
					position++
				}
			l375:
				{
					position377, tokenIndex377, depth377 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l378
					}
					position++
					goto l377
				l378:
					position, tokenIndex, depth = position377, tokenIndex377, depth377
					if buffer[position] != rune('O') {
						goto l351
					}
					position++
				}
			l377:
				{
					position379, tokenIndex379, depth379 := position, tokenIndex, depth
					if buffer[position] != rune('m') {
						goto l380
					}
					position++
					goto l379
				l380:
					position, tokenIndex, depth = position379, tokenIndex379, depth379
					if buffer[position] != rune('M') {
						goto l351
					}
					position++
				}
			l379:
				if !_rules[rulesp]() {
					goto l351
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l351
				}
				if !_rules[ruleAction13]() {
					goto l351
				}
				depth--
				add(ruleInsertIntoFromStmt, position352)
			}
			return true
		l351:
			position, tokenIndex, depth = position351, tokenIndex351, depth351
			return false
		},
		/* 20 PauseSourceStmt <- <(('p' / 'P') ('a' / 'A') ('u' / 'U') ('s' / 'S') ('e' / 'E') sp (('s' / 'S') ('o' / 'O') ('u' / 'U') ('r' / 'R') ('c' / 'C') ('e' / 'E')) sp StreamIdentifier Action14)> */
		func() bool {
			position381, tokenIndex381, depth381 := position, tokenIndex, depth
			{
				position382 := position
				depth++
				{
					position383, tokenIndex383, depth383 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l384
					}
					position++
					goto l383
				l384:
					position, tokenIndex, depth = position383, tokenIndex383, depth383
					if buffer[position] != rune('P') {
						goto l381
					}
					position++
				}
			l383:
				{
					position385, tokenIndex385, depth385 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l386
					}
					position++
					goto l385
				l386:
					position, tokenIndex, depth = position385, tokenIndex385, depth385
					if buffer[position] != rune('A') {
						goto l381
					}
					position++
				}
			l385:
				{
					position387, tokenIndex387, depth387 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l388
					}
					position++
					goto l387
				l388:
					position, tokenIndex, depth = position387, tokenIndex387, depth387
					if buffer[position] != rune('U') {
						goto l381
					}
					position++
				}
			l387:
				{
					position389, tokenIndex389, depth389 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l390
					}
					position++
					goto l389
				l390:
					position, tokenIndex, depth = position389, tokenIndex389, depth389
					if buffer[position] != rune('S') {
						goto l381
					}
					position++
				}
			l389:
				{
					position391, tokenIndex391, depth391 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l392
					}
					position++
					goto l391
				l392:
					position, tokenIndex, depth = position391, tokenIndex391, depth391
					if buffer[position] != rune('E') {
						goto l381
					}
					position++
				}
			l391:
				if !_rules[rulesp]() {
					goto l381
				}
				{
					position393, tokenIndex393, depth393 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l394
					}
					position++
					goto l393
				l394:
					position, tokenIndex, depth = position393, tokenIndex393, depth393
					if buffer[position] != rune('S') {
						goto l381
					}
					position++
				}
			l393:
				{
					position395, tokenIndex395, depth395 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l396
					}
					position++
					goto l395
				l396:
					position, tokenIndex, depth = position395, tokenIndex395, depth395
					if buffer[position] != rune('O') {
						goto l381
					}
					position++
				}
			l395:
				{
					position397, tokenIndex397, depth397 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l398
					}
					position++
					goto l397
				l398:
					position, tokenIndex, depth = position397, tokenIndex397, depth397
					if buffer[position] != rune('U') {
						goto l381
					}
					position++
				}
			l397:
				{
					position399, tokenIndex399, depth399 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l400
					}
					position++
					goto l399
				l400:
					position, tokenIndex, depth = position399, tokenIndex399, depth399
					if buffer[position] != rune('R') {
						goto l381
					}
					position++
				}
			l399:
				{
					position401, tokenIndex401, depth401 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l402
					}
					position++
					goto l401
				l402:
					position, tokenIndex, depth = position401, tokenIndex401, depth401
					if buffer[position] != rune('C') {
						goto l381
					}
					position++
				}
			l401:
				{
					position403, tokenIndex403, depth403 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l404
					}
					position++
					goto l403
				l404:
					position, tokenIndex, depth = position403, tokenIndex403, depth403
					if buffer[position] != rune('E') {
						goto l381
					}
					position++
				}
			l403:
				if !_rules[rulesp]() {
					goto l381
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l381
				}
				if !_rules[ruleAction14]() {
					goto l381
				}
				depth--
				add(rulePauseSourceStmt, position382)
			}
			return true
		l381:
			position, tokenIndex, depth = position381, tokenIndex381, depth381
			return false
		},
		/* 21 ResumeSourceStmt <- <(('r' / 'R') ('e' / 'E') ('s' / 'S') ('u' / 'U') ('m' / 'M') ('e' / 'E') sp (('s' / 'S') ('o' / 'O') ('u' / 'U') ('r' / 'R') ('c' / 'C') ('e' / 'E')) sp StreamIdentifier Action15)> */
		func() bool {
			position405, tokenIndex405, depth405 := position, tokenIndex, depth
			{
				position406 := position
				depth++
				{
					position407, tokenIndex407, depth407 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l408
					}
					position++
					goto l407
				l408:
					position, tokenIndex, depth = position407, tokenIndex407, depth407
					if buffer[position] != rune('R') {
						goto l405
					}
					position++
				}
			l407:
				{
					position409, tokenIndex409, depth409 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l410
					}
					position++
					goto l409
				l410:
					position, tokenIndex, depth = position409, tokenIndex409, depth409
					if buffer[position] != rune('E') {
						goto l405
					}
					position++
				}
			l409:
				{
					position411, tokenIndex411, depth411 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l412
					}
					position++
					goto l411
				l412:
					position, tokenIndex, depth = position411, tokenIndex411, depth411
					if buffer[position] != rune('S') {
						goto l405
					}
					position++
				}
			l411:
				{
					position413, tokenIndex413, depth413 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l414
					}
					position++
					goto l413
				l414:
					position, tokenIndex, depth = position413, tokenIndex413, depth413
					if buffer[position] != rune('U') {
						goto l405
					}
					position++
				}
			l413:
				{
					position415, tokenIndex415, depth415 := position, tokenIndex, depth
					if buffer[position] != rune('m') {
						goto l416
					}
					position++
					goto l415
				l416:
					position, tokenIndex, depth = position415, tokenIndex415, depth415
					if buffer[position] != rune('M') {
						goto l405
					}
					position++
				}
			l415:
				{
					position417, tokenIndex417, depth417 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l418
					}
					position++
					goto l417
				l418:
					position, tokenIndex, depth = position417, tokenIndex417, depth417
					if buffer[position] != rune('E') {
						goto l405
					}
					position++
				}
			l417:
				if !_rules[rulesp]() {
					goto l405
				}
				{
					position419, tokenIndex419, depth419 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l420
					}
					position++
					goto l419
				l420:
					position, tokenIndex, depth = position419, tokenIndex419, depth419
					if buffer[position] != rune('S') {
						goto l405
					}
					position++
				}
			l419:
				{
					position421, tokenIndex421, depth421 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l422
					}
					position++
					goto l421
				l422:
					position, tokenIndex, depth = position421, tokenIndex421, depth421
					if buffer[position] != rune('O') {
						goto l405
					}
					position++
				}
			l421:
				{
					position423, tokenIndex423, depth423 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l424
					}
					position++
					goto l423
				l424:
					position, tokenIndex, depth = position423, tokenIndex423, depth423
					if buffer[position] != rune('U') {
						goto l405
					}
					position++
				}
			l423:
				{
					position425, tokenIndex425, depth425 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l426
					}
					position++
					goto l425
				l426:
					position, tokenIndex, depth = position425, tokenIndex425, depth425
					if buffer[position] != rune('R') {
						goto l405
					}
					position++
				}
			l425:
				{
					position427, tokenIndex427, depth427 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l428
					}
					position++
					goto l427
				l428:
					position, tokenIndex, depth = position427, tokenIndex427, depth427
					if buffer[position] != rune('C') {
						goto l405
					}
					position++
				}
			l427:
				{
					position429, tokenIndex429, depth429 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l430
					}
					position++
					goto l429
				l430:
					position, tokenIndex, depth = position429, tokenIndex429, depth429
					if buffer[position] != rune('E') {
						goto l405
					}
					position++
				}
			l429:
				if !_rules[rulesp]() {
					goto l405
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l405
				}
				if !_rules[ruleAction15]() {
					goto l405
				}
				depth--
				add(ruleResumeSourceStmt, position406)
			}
			return true
		l405:
			position, tokenIndex, depth = position405, tokenIndex405, depth405
			return false
		},
		/* 22 RewindSourceStmt <- <(('r' / 'R') ('e' / 'E') ('w' / 'W') ('i' / 'I') ('n' / 'N') ('d' / 'D') sp (('s' / 'S') ('o' / 'O') ('u' / 'U') ('r' / 'R') ('c' / 'C') ('e' / 'E')) sp StreamIdentifier Action16)> */
		func() bool {
			position431, tokenIndex431, depth431 := position, tokenIndex, depth
			{
				position432 := position
				depth++
				{
					position433, tokenIndex433, depth433 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l434
					}
					position++
					goto l433
				l434:
					position, tokenIndex, depth = position433, tokenIndex433, depth433
					if buffer[position] != rune('R') {
						goto l431
					}
					position++
				}
			l433:
				{
					position435, tokenIndex435, depth435 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l436
					}
					position++
					goto l435
				l436:
					position, tokenIndex, depth = position435, tokenIndex435, depth435
					if buffer[position] != rune('E') {
						goto l431
					}
					position++
				}
			l435:
				{
					position437, tokenIndex437, depth437 := position, tokenIndex, depth
					if buffer[position] != rune('w') {
						goto l438
					}
					position++
					goto l437
				l438:
					position, tokenIndex, depth = position437, tokenIndex437, depth437
					if buffer[position] != rune('W') {
						goto l431
					}
					position++
				}
			l437:
				{
					position439, tokenIndex439, depth439 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l440
					}
					position++
					goto l439
				l440:
					position, tokenIndex, depth = position439, tokenIndex439, depth439
					if buffer[position] != rune('I') {
						goto l431
					}
					position++
				}
			l439:
				{
					position441, tokenIndex441, depth441 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l442
					}
					position++
					goto l441
				l442:
					position, tokenIndex, depth = position441, tokenIndex441, depth441
					if buffer[position] != rune('N') {
						goto l431
					}
					position++
				}
			l441:
				{
					position443, tokenIndex443, depth443 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l444
					}
					position++
					goto l443
				l444:
					position, tokenIndex, depth = position443, tokenIndex443, depth443
					if buffer[position] != rune('D') {
						goto l431
					}
					position++
				}
			l443:
				if !_rules[rulesp]() {
					goto l431
				}
				{
					position445, tokenIndex445, depth445 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l446
					}
					position++
					goto l445
				l446:
					position, tokenIndex, depth = position445, tokenIndex445, depth445
					if buffer[position] != rune('S') {
						goto l431
					}
					position++
				}
			l445:
				{
					position447, tokenIndex447, depth447 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l448
					}
					position++
					goto l447
				l448:
					position, tokenIndex, depth = position447, tokenIndex447, depth447
					if buffer[position] != rune('O') {
						goto l431
					}
					position++
				}
			l447:
				{
					position449, tokenIndex449, depth449 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l450
					}
					position++
					goto l449
				l450:
					position, tokenIndex, depth = position449, tokenIndex449, depth449
					if buffer[position] != rune('U') {
						goto l431
					}
					position++
				}
			l449:
				{
					position451, tokenIndex451, depth451 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l452
					}
					position++
					goto l451
				l452:
					position, tokenIndex, depth = position451, tokenIndex451, depth451
					if buffer[position] != rune('R') {
						goto l431
					}
					position++
				}
			l451:
				{
					position453, tokenIndex453, depth453 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l454
					}
					position++
					goto l453
				l454:
					position, tokenIndex, depth = position453, tokenIndex453, depth453
					if buffer[position] != rune('C') {
						goto l431
					}
					position++
				}
			l453:
				{
					position455, tokenIndex455, depth455 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l456
					}
					position++
					goto l455
				l456:
					position, tokenIndex, depth = position455, tokenIndex455, depth455
					if buffer[position] != rune('E') {
						goto l431
					}
					position++
				}
			l455:
				if !_rules[rulesp]() {
					goto l431
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l431
				}
				if !_rules[ruleAction16]() {
					goto l431
				}
				depth--
				add(ruleRewindSourceStmt, position432)
			}
			return true
		l431:
			position, tokenIndex, depth = position431, tokenIndex431, depth431
			return false
		},
		/* 23 DropSourceStmt <- <(('d' / 'D') ('r' / 'R') ('o' / 'O') ('p' / 'P') sp (('s' / 'S') ('o' / 'O') ('u' / 'U') ('r' / 'R') ('c' / 'C') ('e' / 'E')) sp StreamIdentifier Action17)> */
		func() bool {
			position457, tokenIndex457, depth457 := position, tokenIndex, depth
			{
				position458 := position
				depth++
				{
					position459, tokenIndex459, depth459 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l460
					}
					position++
					goto l459
				l460:
					position, tokenIndex, depth = position459, tokenIndex459, depth459
					if buffer[position] != rune('D') {
						goto l457
					}
					position++
				}
			l459:
				{
					position461, tokenIndex461, depth461 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l462
					}
					position++
					goto l461
				l462:
					position, tokenIndex, depth = position461, tokenIndex461, depth461
					if buffer[position] != rune('R') {
						goto l457
					}
					position++
				}
			l461:
				{
					position463, tokenIndex463, depth463 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l464
					}
					position++
					goto l463
				l464:
					position, tokenIndex, depth = position463, tokenIndex463, depth463
					if buffer[position] != rune('O') {
						goto l457
					}
					position++
				}
			l463:
				{
					position465, tokenIndex465, depth465 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l466
					}
					position++
					goto l465
				l466:
					position, tokenIndex, depth = position465, tokenIndex465, depth465
					if buffer[position] != rune('P') {
						goto l457
					}
					position++
				}
			l465:
				if !_rules[rulesp]() {
					goto l457
				}
				{
					position467, tokenIndex467, depth467 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l468
					}
					position++
					goto l467
				l468:
					position, tokenIndex, depth = position467, tokenIndex467, depth467
					if buffer[position] != rune('S') {
						goto l457
					}
					position++
				}
			l467:
				{
					position469, tokenIndex469, depth469 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l470
					}
					position++
					goto l469
				l470:
					position, tokenIndex, depth = position469, tokenIndex469, depth469
					if buffer[position] != rune('O') {
						goto l457
					}
					position++
				}
			l469:
				{
					position471, tokenIndex471, depth471 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l472
					}
					position++
					goto l471
				l472:
					position, tokenIndex, depth = position471, tokenIndex471, depth471
					if buffer[position] != rune('U') {
						goto l457
					}
					position++
				}
			l471:
				{
					position473, tokenIndex473, depth473 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l474
					}
					position++
					goto l473
				l474:
					position, tokenIndex, depth = position473, tokenIndex473, depth473
					if buffer[position] != rune('R') {
						goto l457
					}
					position++
				}
			l473:
				{
					position475, tokenIndex475, depth475 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l476
					}
					position++
					goto l475
				l476:
					position, tokenIndex, depth = position475, tokenIndex475, depth475
					if buffer[position] != rune('C') {
						goto l457
					}
					position++
				}
			l475:
				{
					position477, tokenIndex477, depth477 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l478
					}
					position++
					goto l477
				l478:
					position, tokenIndex, depth = position477, tokenIndex477, depth477
					if buffer[position] != rune('E') {
						goto l457
					}
					position++
				}
			l477:
				if !_rules[rulesp]() {
					goto l457
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l457
				}
				if !_rules[ruleAction17]() {
					goto l457
				}
				depth--
				add(ruleDropSourceStmt, position458)
			}
			return true
		l457:
			position, tokenIndex, depth = position457, tokenIndex457, depth457
			return false
		},
		/* 24 DropStreamStmt <- <(('d' / 'D') ('r' / 'R') ('o' / 'O') ('p' / 'P') sp (('s' / 'S') ('t' / 'T') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('m' / 'M')) sp StreamIdentifier Action18)> */
		func() bool {
			position479, tokenIndex479, depth479 := position, tokenIndex, depth
			{
				position480 := position
				depth++
				{
					position481, tokenIndex481, depth481 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l482
					}
					position++
					goto l481
				l482:
					position, tokenIndex, depth = position481, tokenIndex481, depth481
					if buffer[position] != rune('D') {
						goto l479
					}
					position++
				}
			l481:
				{
					position483, tokenIndex483, depth483 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l484
					}
					position++
					goto l483
				l484:
					position, tokenIndex, depth = position483, tokenIndex483, depth483
					if buffer[position] != rune('R') {
						goto l479
					}
					position++
				}
			l483:
				{
					position485, tokenIndex485, depth485 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l486
					}
					position++
					goto l485
				l486:
					position, tokenIndex, depth = position485, tokenIndex485, depth485
					if buffer[position] != rune('O') {
						goto l479
					}
					position++
				}
			l485:
				{
					position487, tokenIndex487, depth487 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l488
					}
					position++
					goto l487
				l488:
					position, tokenIndex, depth = position487, tokenIndex487, depth487
					if buffer[position] != rune('P') {
						goto l479
					}
					position++
				}
			l487:
				if !_rules[rulesp]() {
					goto l479
				}
				{
					position489, tokenIndex489, depth489 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l490
					}
					position++
					goto l489
				l490:
					position, tokenIndex, depth = position489, tokenIndex489, depth489
					if buffer[position] != rune('S') {
						goto l479
					}
					position++
				}
			l489:
				{
					position491, tokenIndex491, depth491 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l492
					}
					position++
					goto l491
				l492:
					position, tokenIndex, depth = position491, tokenIndex491, depth491
					if buffer[position] != rune('T') {
						goto l479
					}
					position++
				}
			l491:
				{
					position493, tokenIndex493, depth493 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l494
					}
					position++
					goto l493
				l494:
					position, tokenIndex, depth = position493, tokenIndex493, depth493
					if buffer[position] != rune('R') {
						goto l479
					}
					position++
				}
			l493:
				{
					position495, tokenIndex495, depth495 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l496
					}
					position++
					goto l495
				l496:
					position, tokenIndex, depth = position495, tokenIndex495, depth495
					if buffer[position] != rune('E') {
						goto l479
					}
					position++
				}
			l495:
				{
					position497, tokenIndex497, depth497 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l498
					}
					position++
					goto l497
				l498:
					position, tokenIndex, depth = position497, tokenIndex497, depth497
					if buffer[position] != rune('A') {
						goto l479
					}
					position++
				}
			l497:
				{
					position499, tokenIndex499, depth499 := position, tokenIndex, depth
					if buffer[position] != rune('m') {
						goto l500
					}
					position++
					goto l499
				l500:
					position, tokenIndex, depth = position499, tokenIndex499, depth499
					if buffer[position] != rune('M') {
						goto l479
					}
					position++
				}
			l499:
				if !_rules[rulesp]() {
					goto l479
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l479
				}
				if !_rules[ruleAction18]() {
					goto l479
				}
				depth--
				add(ruleDropStreamStmt, position480)
			}
			return true
		l479:
			position, tokenIndex, depth = position479, tokenIndex479, depth479
			return false
		},
		/* 25 DropSinkStmt <- <(('d' / 'D') ('r' / 'R') ('o' / 'O') ('p' / 'P') sp (('s' / 'S') ('i' / 'I') ('n' / 'N') ('k' / 'K')) sp StreamIdentifier Action19)> */
		func() bool {
			position501, tokenIndex501, depth501 := position, tokenIndex, depth
			{
				position502 := position
				depth++
				{
					position503, tokenIndex503, depth503 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l504
					}
					position++
					goto l503
				l504:
					position, tokenIndex, depth = position503, tokenIndex503, depth503
					if buffer[position] != rune('D') {
						goto l501
					}
					position++
				}
			l503:
				{
					position505, tokenIndex505, depth505 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l506
					}
					position++
					goto l505
				l506:
					position, tokenIndex, depth = position505, tokenIndex505, depth505
					if buffer[position] != rune('R') {
						goto l501
					}
					position++
				}
			l505:
				{
					position507, tokenIndex507, depth507 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l508
					}
					position++
					goto l507
				l508:
					position, tokenIndex, depth = position507, tokenIndex507, depth507
					if buffer[position] != rune('O') {
						goto l501
					}
					position++
				}
			l507:
				{
					position509, tokenIndex509, depth509 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l510
					}
					position++
					goto l509
				l510:
					position, tokenIndex, depth = position509, tokenIndex509, depth509
					if buffer[position] != rune('P') {
						goto l501
					}
					position++
				}
			l509:
				if !_rules[rulesp]() {
					goto l501
				}
				{
					position511, tokenIndex511, depth511 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l512
					}
					position++
					goto l511
				l512:
					position, tokenIndex, depth = position511, tokenIndex511, depth511
					if buffer[position] != rune('S') {
						goto l501
					}
					position++
				}
			l511:
				{
					position513, tokenIndex513, depth513 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l514
					}
					position++
					goto l513
				l514:
					position, tokenIndex, depth = position513, tokenIndex513, depth513
					if buffer[position] != rune('I') {
						goto l501
					}
					position++
				}
			l513:
				{
					position515, tokenIndex515, depth515 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l516
					}
					position++
					goto l515
				l516:
					position, tokenIndex, depth = position515, tokenIndex515, depth515
					if buffer[position] != rune('N') {
						goto l501
					}
					position++
				}
			l515:
				{
					position517, tokenIndex517, depth517 := position, tokenIndex, depth
					if buffer[position] != rune('k') {
						goto l518
					}
					position++
					goto l517
				l518:
					position, tokenIndex, depth = position517, tokenIndex517, depth517
					if buffer[position] != rune('K') {
						goto l501
					}
					position++
				}
			l517:
				if !_rules[rulesp]() {
					goto l501
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l501
				}
				if !_rules[ruleAction19]() {
					goto l501
				}
				depth--
				add(ruleDropSinkStmt, position502)
			}
			return true
		l501:
			position, tokenIndex, depth = position501, tokenIndex501, depth501
			return false
		},
		/* 26 DropStateStmt <- <(('d' / 'D') ('r' / 'R') ('o' / 'O') ('p' / 'P') sp (('s' / 'S') ('t' / 'T') ('a' / 'A') ('t' / 'T') ('e' / 'E')) sp StreamIdentifier Action20)> */
		func() bool {
			position519, tokenIndex519, depth519 := position, tokenIndex, depth
			{
				position520 := position
				depth++
				{
					position521, tokenIndex521, depth521 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l522
					}
					position++
					goto l521
				l522:
					position, tokenIndex, depth = position521, tokenIndex521, depth521
					if buffer[position] != rune('D') {
						goto l519
					}
					position++
				}
			l521:
				{
					position523, tokenIndex523, depth523 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l524
					}
					position++
					goto l523
				l524:
					position, tokenIndex, depth = position523, tokenIndex523, depth523
					if buffer[position] != rune('R') {
						goto l519
					}
					position++
				}
			l523:
				{
					position525, tokenIndex525, depth525 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l526
					}
					position++
					goto l525
				l526:
					position, tokenIndex, depth = position525, tokenIndex525, depth525
					if buffer[position] != rune('O') {
						goto l519
					}
					position++
				}
			l525:
				{
					position527, tokenIndex527, depth527 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l528
					}
					position++
					goto l527
				l528:
					position, tokenIndex, depth = position527, tokenIndex527, depth527
					if buffer[position] != rune('P') {
						goto l519
					}
					position++
				}
			l527:
				if !_rules[rulesp]() {
					goto l519
				}
				{
					position529, tokenIndex529, depth529 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l530
					}
					position++
					goto l529
				l530:
					position, tokenIndex, depth = position529, tokenIndex529, depth529
					if buffer[position] != rune('S') {
						goto l519
					}
					position++
				}
			l529:
				{
					position531, tokenIndex531, depth531 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l532
					}
					position++
					goto l531
				l532:
					position, tokenIndex, depth = position531, tokenIndex531, depth531
					if buffer[position] != rune('T') {
						goto l519
					}
					position++
				}
			l531:
				{
					position533, tokenIndex533, depth533 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l534
					}
					position++
					goto l533
				l534:
					position, tokenIndex, depth = position533, tokenIndex533, depth533
					if buffer[position] != rune('A') {
						goto l519
					}
					position++
				}
			l533:
				{
					position535, tokenIndex535, depth535 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l536
					}
					position++
					goto l535
				l536:
					position, tokenIndex, depth = position535, tokenIndex535, depth535
					if buffer[position] != rune('T') {
						goto l519
					}
					position++
				}
			l535:
				{
					position537, tokenIndex537, depth537 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l538
					}
					position++
					goto l537
				l538:
					position, tokenIndex, depth = position537, tokenIndex537, depth537
					if buffer[position] != rune('E') {
						goto l519
					}
					position++
				}
			l537:
				if !_rules[rulesp]() {
					goto l519
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l519
				}
				if !_rules[ruleAction20]() {
					goto l519
				}
				depth--
				add(ruleDropStateStmt, position520)
			}
			return true
		l519:
			position, tokenIndex, depth = position519, tokenIndex519, depth519
			return false
		},
		/* 27 LoadStateStmt <- <(('l' / 'L') ('o' / 'O') ('a' / 'A') ('d' / 'D') sp (('s' / 'S') ('t' / 'T') ('a' / 'A') ('t' / 'T') ('e' / 'E')) sp StreamIdentifier sp (('t' / 'T') ('y' / 'Y') ('p' / 'P') ('e' / 'E')) sp SourceSinkType SetOptSpecs Action21)> */
		func() bool {
			position539, tokenIndex539, depth539 := position, tokenIndex, depth
			{
				position540 := position
				depth++
				{
					position541, tokenIndex541, depth541 := position, tokenIndex, depth
					if buffer[position] != rune('l') {
						goto l542
					}
					position++
					goto l541
				l542:
					position, tokenIndex, depth = position541, tokenIndex541, depth541
					if buffer[position] != rune('L') {
						goto l539
					}
					position++
				}
			l541:
				{
					position543, tokenIndex543, depth543 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l544
					}
					position++
					goto l543
				l544:
					position, tokenIndex, depth = position543, tokenIndex543, depth543
					if buffer[position] != rune('O') {
						goto l539
					}
					position++
				}
			l543:
				{
					position545, tokenIndex545, depth545 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l546
					}
					position++
					goto l545
				l546:
					position, tokenIndex, depth = position545, tokenIndex545, depth545
					if buffer[position] != rune('A') {
						goto l539
					}
					position++
				}
			l545:
				{
					position547, tokenIndex547, depth547 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l548
					}
					position++
					goto l547
				l548:
					position, tokenIndex, depth = position547, tokenIndex547, depth547
					if buffer[position] != rune('D') {
						goto l539
					}
					position++
				}
			l547:
				if !_rules[rulesp]() {
					goto l539
				}
				{
					position549, tokenIndex549, depth549 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l550
					}
					position++
					goto l549
				l550:
					position, tokenIndex, depth = position549, tokenIndex549, depth549
					if buffer[position] != rune('S') {
						goto l539
					}
					position++
				}
			l549:
				{
					position551, tokenIndex551, depth551 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l552
					}
					position++
					goto l551
				l552:
					position, tokenIndex, depth = position551, tokenIndex551, depth551
					if buffer[position] != rune('T') {
						goto l539
					}
					position++
				}
			l551:
				{
					position553, tokenIndex553, depth553 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l554
					}
					position++
					goto l553
				l554:
					position, tokenIndex, depth = position553, tokenIndex553, depth553
					if buffer[position] != rune('A') {
						goto l539
					}
					position++
				}
			l553:
				{
					position555, tokenIndex555, depth555 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l556
					}
					position++
					goto l555
				l556:
					position, tokenIndex, depth = position555, tokenIndex555, depth555
					if buffer[position] != rune('T') {
						goto l539
					}
					position++
				}
			l555:
				{
					position557, tokenIndex557, depth557 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l558
					}
					position++
					goto l557
				l558:
					position, tokenIndex, depth = position557, tokenIndex557, depth557
					if buffer[position] != rune('E') {
						goto l539
					}
					position++
				}
			l557:
				if !_rules[rulesp]() {
					goto l539
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l539
				}
				if !_rules[rulesp]() {
					goto l539
				}
				{
					position559, tokenIndex559, depth559 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l560
					}
					position++
					goto l559
				l560:
					position, tokenIndex, depth = position559, tokenIndex559, depth559
					if buffer[position] != rune('T') {
						goto l539
					}
					position++
				}
			l559:
				{
					position561, tokenIndex561, depth561 := position, tokenIndex, depth
					if buffer[position] != rune('y') {
						goto l562
					}
					position++
					goto l561
				l562:
					position, tokenIndex, depth = position561, tokenIndex561, depth561
					if buffer[position] != rune('Y') {
						goto l539
					}
					position++
				}
			l561:
				{
					position563, tokenIndex563, depth563 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l564
					}
					position++
					goto l563
				l564:
					position, tokenIndex, depth = position563, tokenIndex563, depth563
					if buffer[position] != rune('P') {
						goto l539
					}
					position++
				}
			l563:
				{
					position565, tokenIndex565, depth565 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l566
					}
					position++
					goto l565
				l566:
					position, tokenIndex, depth = position565, tokenIndex565, depth565
					if buffer[position] != rune('E') {
						goto l539
					}
					position++
				}
			l565:
				if !_rules[rulesp]() {
					goto l539
				}
				if !_rules[ruleSourceSinkType]() {
					goto l539
				}
				if !_rules[ruleSetOptSpecs]() {
					goto l539
				}
				if !_rules[ruleAction21]() {
					goto l539
				}
				depth--
				add(ruleLoadStateStmt, position540)
			}
			return true
		l539:
			position, tokenIndex, depth = position539, tokenIndex539, depth539
			return false
		},
		/* 28 LoadStateOrCreateStmt <- <(LoadStateStmt sp (('o' / 'O') ('r' / 'R')) sp (('c' / 'C') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('t' / 'T') ('e' / 'E')) sp (('i' / 'I') ('f' / 'F')) sp (('n' / 'N') ('o' / 'O') ('t' / 'T')) sp (('e' / 'E') ('x' / 'X') ('i' / 'I') ('s' / 'S') ('t' / 'T') ('s' / 'S')) SourceSinkSpecs Action22)> */
		func() bool {
			position567, tokenIndex567, depth567 := position, tokenIndex, depth
			{
				position568 := position
				depth++
				if !_rules[ruleLoadStateStmt]() {
					goto l567
				}
				if !_rules[rulesp]() {
					goto l567
				}
				{
					position569, tokenIndex569, depth569 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l570
					}
					position++
					goto l569
				l570:
					position, tokenIndex, depth = position569, tokenIndex569, depth569
					if buffer[position] != rune('O') {
						goto l567
					}
					position++
				}
			l569:
				{
					position571, tokenIndex571, depth571 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l572
					}
					position++
					goto l571
				l572:
					position, tokenIndex, depth = position571, tokenIndex571, depth571
					if buffer[position] != rune('R') {
						goto l567
					}
					position++
				}
			l571:
				if !_rules[rulesp]() {
					goto l567
				}
				{
					position573, tokenIndex573, depth573 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l574
					}
					position++
					goto l573
				l574:
					position, tokenIndex, depth = position573, tokenIndex573, depth573
					if buffer[position] != rune('C') {
						goto l567
					}
					position++
				}
			l573:
				{
					position575, tokenIndex575, depth575 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l576
					}
					position++
					goto l575
				l576:
					position, tokenIndex, depth = position575, tokenIndex575, depth575
					if buffer[position] != rune('R') {
						goto l567
					}
					position++
				}
			l575:
				{
					position577, tokenIndex577, depth577 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l578
					}
					position++
					goto l577
				l578:
					position, tokenIndex, depth = position577, tokenIndex577, depth577
					if buffer[position] != rune('E') {
						goto l567
					}
					position++
				}
			l577:
				{
					position579, tokenIndex579, depth579 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l580
					}
					position++
					goto l579
				l580:
					position, tokenIndex, depth = position579, tokenIndex579, depth579
					if buffer[position] != rune('A') {
						goto l567
					}
					position++
				}
			l579:
				{
					position581, tokenIndex581, depth581 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l582
					}
					position++
					goto l581
				l582:
					position, tokenIndex, depth = position581, tokenIndex581, depth581
					if buffer[position] != rune('T') {
						goto l567
					}
					position++
				}
			l581:
				{
					position583, tokenIndex583, depth583 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l584
					}
					position++
					goto l583
				l584:
					position, tokenIndex, depth = position583, tokenIndex583, depth583
					if buffer[position] != rune('E') {
						goto l567
					}
					position++
				}
			l583:
				if !_rules[rulesp]() {
					goto l567
				}
				{
					position585, tokenIndex585, depth585 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l586
					}
					position++
					goto l585
				l586:
					position, tokenIndex, depth = position585, tokenIndex585, depth585
					if buffer[position] != rune('I') {
						goto l567
					}
					position++
				}
			l585:
				{
					position587, tokenIndex587, depth587 := position, tokenIndex, depth
					if buffer[position] != rune('f') {
						goto l588
					}
					position++
					goto l587
				l588:
					position, tokenIndex, depth = position587, tokenIndex587, depth587
					if buffer[position] != rune('F') {
						goto l567
					}
					position++
				}
			l587:
				if !_rules[rulesp]() {
					goto l567
				}
				{
					position589, tokenIndex589, depth589 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l590
					}
					position++
					goto l589
				l590:
					position, tokenIndex, depth = position589, tokenIndex589, depth589
					if buffer[position] != rune('N') {
						goto l567
					}
					position++
				}
			l589:
				{
					position591, tokenIndex591, depth591 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l592
					}
					position++
					goto l591
				l592:
					position, tokenIndex, depth = position591, tokenIndex591, depth591
					if buffer[position] != rune('O') {
						goto l567
					}
					position++
				}
			l591:
				{
					position593, tokenIndex593, depth593 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l594
					}
					position++
					goto l593
				l594:
					position, tokenIndex, depth = position593, tokenIndex593, depth593
					if buffer[position] != rune('T') {
						goto l567
					}
					position++
				}
			l593:
				if !_rules[rulesp]() {
					goto l567
				}
				{
					position595, tokenIndex595, depth595 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l596
					}
					position++
					goto l595
				l596:
					position, tokenIndex, depth = position595, tokenIndex595, depth595
					if buffer[position] != rune('E') {
						goto l567
					}
					position++
				}
			l595:
				{
					position597, tokenIndex597, depth597 := position, tokenIndex, depth
					if buffer[position] != rune('x') {
						goto l598
					}
					position++
					goto l597
				l598:
					position, tokenIndex, depth = position597, tokenIndex597, depth597
					if buffer[position] != rune('X') {
						goto l567
					}
					position++
				}
			l597:
				{
					position599, tokenIndex599, depth599 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l600
					}
					position++
					goto l599
				l600:
					position, tokenIndex, depth = position599, tokenIndex599, depth599
					if buffer[position] != rune('I') {
						goto l567
					}
					position++
				}
			l599:
				{
					position601, tokenIndex601, depth601 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l602
					}
					position++
					goto l601
				l602:
					position, tokenIndex, depth = position601, tokenIndex601, depth601
					if buffer[position] != rune('S') {
						goto l567
					}
					position++
				}
			l601:
				{
					position603, tokenIndex603, depth603 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l604
					}
					position++
					goto l603
				l604:
					position, tokenIndex, depth = position603, tokenIndex603, depth603
					if buffer[position] != rune('T') {
						goto l567
					}
					position++
				}
			l603:
				{
					position605, tokenIndex605, depth605 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l606
					}
					position++
					goto l605
				l606:
					position, tokenIndex, depth = position605, tokenIndex605, depth605
					if buffer[position] != rune('S') {
						goto l567
					}
					position++
				}
			l605:
				if !_rules[ruleSourceSinkSpecs]() {
					goto l567
				}
				if !_rules[ruleAction22]() {
					goto l567
				}
				depth--
				add(ruleLoadStateOrCreateStmt, position568)
			}
			return true
		l567:
			position, tokenIndex, depth = position567, tokenIndex567, depth567
			return false
		},
		/* 29 SaveStateStmt <- <(('s' / 'S') ('a' / 'A') ('v' / 'V') ('e' / 'E') sp (('s' / 'S') ('t' / 'T') ('a' / 'A') ('t' / 'T') ('e' / 'E')) sp StreamIdentifier Action23)> */
		func() bool {
			position607, tokenIndex607, depth607 := position, tokenIndex, depth
			{
				position608 := position
				depth++
				{
					position609, tokenIndex609, depth609 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l610
					}
					position++
					goto l609
				l610:
					position, tokenIndex, depth = position609, tokenIndex609, depth609
					if buffer[position] != rune('S') {
						goto l607
					}
					position++
				}
			l609:
				{
					position611, tokenIndex611, depth611 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l612
					}
					position++
					goto l611
				l612:
					position, tokenIndex, depth = position611, tokenIndex611, depth611
					if buffer[position] != rune('A') {
						goto l607
					}
					position++
				}
			l611:
				{
					position613, tokenIndex613, depth613 := position, tokenIndex, depth
					if buffer[position] != rune('v') {
						goto l614
					}
					position++
					goto l613
				l614:
					position, tokenIndex, depth = position613, tokenIndex613, depth613
					if buffer[position] != rune('V') {
						goto l607
					}
					position++
				}
			l613:
				{
					position615, tokenIndex615, depth615 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l616
					}
					position++
					goto l615
				l616:
					position, tokenIndex, depth = position615, tokenIndex615, depth615
					if buffer[position] != rune('E') {
						goto l607
					}
					position++
				}
			l615:
				if !_rules[rulesp]() {
					goto l607
				}
				{
					position617, tokenIndex617, depth617 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l618
					}
					position++
					goto l617
				l618:
					position, tokenIndex, depth = position617, tokenIndex617, depth617
					if buffer[position] != rune('S') {
						goto l607
					}
					position++
				}
			l617:
				{
					position619, tokenIndex619, depth619 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l620
					}
					position++
					goto l619
				l620:
					position, tokenIndex, depth = position619, tokenIndex619, depth619
					if buffer[position] != rune('T') {
						goto l607
					}
					position++
				}
			l619:
				{
					position621, tokenIndex621, depth621 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l622
					}
					position++
					goto l621
				l622:
					position, tokenIndex, depth = position621, tokenIndex621, depth621
					if buffer[position] != rune('A') {
						goto l607
					}
					position++
				}
			l621:
				{
					position623, tokenIndex623, depth623 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l624
					}
					position++
					goto l623
				l624:
					position, tokenIndex, depth = position623, tokenIndex623, depth623
					if buffer[position] != rune('T') {
						goto l607
					}
					position++
				}
			l623:
				{
					position625, tokenIndex625, depth625 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l626
					}
					position++
					goto l625
				l626:
					position, tokenIndex, depth = position625, tokenIndex625, depth625
					if buffer[position] != rune('E') {
						goto l607
					}
					position++
				}
			l625:
				if !_rules[rulesp]() {
					goto l607
				}
				if !_rules[ruleStreamIdentifier]() {
					goto l607
				}
				if !_rules[ruleAction23]() {
					goto l607
				}
				depth--
				add(ruleSaveStateStmt, position608)
			}
			return true
		l607:
			position, tokenIndex, depth = position607, tokenIndex607, depth607
			return false
		},
		/* 30 EvalStmt <- <(('e' / 'E') ('v' / 'V') ('a' / 'A') ('l' / 'L') sp Expression <(sp (('o' / 'O') ('n' / 'N')) sp MapExpr)?> Action24)> */
		func() bool {
			position627, tokenIndex627, depth627 := position, tokenIndex, depth
			{
				position628 := position
				depth++
				{
					position629, tokenIndex629, depth629 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l630
					}
					position++
					goto l629
				l630:
					position, tokenIndex, depth = position629, tokenIndex629, depth629
					if buffer[position] != rune('E') {
						goto l627
					}
					position++
				}
			l629:
				{
					position631, tokenIndex631, depth631 := position, tokenIndex, depth
					if buffer[position] != rune('v') {
						goto l632
					}
					position++
					goto l631
				l632:
					position, tokenIndex, depth = position631, tokenIndex631, depth631
					if buffer[position] != rune('V') {
						goto l627
					}
					position++
				}
			l631:
				{
					position633, tokenIndex633, depth633 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l634
					}
					position++
					goto l633
				l634:
					position, tokenIndex, depth = position633, tokenIndex633, depth633
					if buffer[position] != rune('A') {
						goto l627
					}
					position++
				}
			l633:
				{
					position635, tokenIndex635, depth635 := position, tokenIndex, depth
					if buffer[position] != rune('l') {
						goto l636
					}
					position++
					goto l635
				l636:
					position, tokenIndex, depth = position635, tokenIndex635, depth635
					if buffer[position] != rune('L') {
						goto l627
					}
					position++
				}
			l635:
				if !_rules[rulesp]() {
					goto l627
				}
				if !_rules[ruleExpression]() {
					goto l627
				}
				{
					position637 := position
					depth++
					{
						position638, tokenIndex638, depth638 := position, tokenIndex, depth
						if !_rules[rulesp]() {
							goto l638
						}
						{
							position640, tokenIndex640, depth640 := position, tokenIndex, depth
							if buffer[position] != rune('o') {
								goto l641
							}
							position++
							goto l640
						l641:
							position, tokenIndex, depth = position640, tokenIndex640, depth640
							if buffer[position] != rune('O') {
								goto l638
							}
							position++
						}
					l640:
						{
							position642, tokenIndex642, depth642 := position, tokenIndex, depth
							if buffer[position] != rune('n') {
								goto l643
							}
							position++
							goto l642
						l643:
							position, tokenIndex, depth = position642, tokenIndex642, depth642
							if buffer[position] != rune('N') {
								goto l638
							}
							position++
						}
					l642:
						if !_rules[rulesp]() {
							goto l638
						}
						if !_rules[ruleMapExpr]() {
							goto l638
						}
						goto l639
					l638:
						position, tokenIndex, depth = position638, tokenIndex638, depth638
					}
				l639:
					depth--
					add(rulePegText, position637)
				}
				if !_rules[ruleAction24]() {
					goto l627
				}
				depth--
				add(ruleEvalStmt, position628)
			}
			return true
		l627:
			position, tokenIndex, depth = position627, tokenIndex627, depth627
			return false
		},
		/* 31 Emitter <- <(sp (ISTREAM / DSTREAM / RSTREAM) EmitterOptions Action25)> */
		func() bool {
			position644, tokenIndex644, depth644 := position, tokenIndex, depth
			{
				position645 := position
				depth++
				if !_rules[rulesp]() {
					goto l644
				}
				{
					position646, tokenIndex646, depth646 := position, tokenIndex, depth
					if !_rules[ruleISTREAM]() {
						goto l647
					}
					goto l646
				l647:
					position, tokenIndex, depth = position646, tokenIndex646, depth646
					if !_rules[ruleDSTREAM]() {
						goto l648
					}
					goto l646
				l648:
					position, tokenIndex, depth = position646, tokenIndex646, depth646
					if !_rules[ruleRSTREAM]() {
						goto l644
					}
				}
			l646:
				if !_rules[ruleEmitterOptions]() {
					goto l644
				}
				if !_rules[ruleAction25]() {
					goto l644
				}
				depth--
				add(ruleEmitter, position645)
			}
			return true
		l644:
			position, tokenIndex, depth = position644, tokenIndex644, depth644
			return false
		},
		/* 32 EmitterOptions <- <(<(spOpt '[' spOpt EmitterOptionCombinations spOpt ']')?> Action26)> */
		func() bool {
			position649, tokenIndex649, depth649 := position, tokenIndex, depth
			{
				position650 := position
				depth++
				{
					position651 := position
					depth++
					{
						position652, tokenIndex652, depth652 := position, tokenIndex, depth
						if !_rules[rulespOpt]() {
							goto l652
						}
						if buffer[position] != rune('[') {
							goto l652
						}
						position++
						if !_rules[rulespOpt]() {
							goto l652
						}
						if !_rules[ruleEmitterOptionCombinations]() {
							goto l652
						}
						if !_rules[rulespOpt]() {
							goto l652
						}
						if buffer[position] != rune(']') {
							goto l652
						}
						position++
						goto l653
					l652:
						position, tokenIndex, depth = position652, tokenIndex652, depth652
					}
				l653:
					depth--
					add(rulePegText, position651)
				}
				if !_rules[ruleAction26]() {
					goto l649
				}
				depth--
				add(ruleEmitterOptions, position650)
			}
			return true
		l649:
			position, tokenIndex, depth = position649, tokenIndex649, depth649
			return false
		},
		/* 33 EmitterOptionCombinations <- <(EmitterLimit / (EmitterSample sp EmitterLimit) / EmitterSample)> */
		func() bool {
			position654, tokenIndex654, depth654 := position, tokenIndex, depth
			{
				position655 := position
				depth++
				{
					position656, tokenIndex656, depth656 := position, tokenIndex, depth
					if !_rules[ruleEmitterLimit]() {
						goto l657
					}
					goto l656
				l657:
					position, tokenIndex, depth = position656, tokenIndex656, depth656
					if !_rules[ruleEmitterSample]() {
						goto l658
					}
					if !_rules[rulesp]() {
						goto l658
					}
					if !_rules[ruleEmitterLimit]() {
						goto l658
					}
					goto l656
				l658:
					position, tokenIndex, depth = position656, tokenIndex656, depth656
					if !_rules[ruleEmitterSample]() {
						goto l654
					}
				}
			l656:
				depth--
				add(ruleEmitterOptionCombinations, position655)
			}
			return true
		l654:
			position, tokenIndex, depth = position654, tokenIndex654, depth654
			return false
		},
		/* 34 EmitterLimit <- <(('l' / 'L') ('i' / 'I') ('m' / 'M') ('i' / 'I') ('t' / 'T') sp NumericLiteral Action27)> */
		func() bool {
			position659, tokenIndex659, depth659 := position, tokenIndex, depth
			{
				position660 := position
				depth++
				{
					position661, tokenIndex661, depth661 := position, tokenIndex, depth
					if buffer[position] != rune('l') {
						goto l662
					}
					position++
					goto l661
				l662:
					position, tokenIndex, depth = position661, tokenIndex661, depth661
					if buffer[position] != rune('L') {
						goto l659
					}
					position++
				}
			l661:
				{
					position663, tokenIndex663, depth663 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l664
					}
					position++
					goto l663
				l664:
					position, tokenIndex, depth = position663, tokenIndex663, depth663
					if buffer[position] != rune('I') {
						goto l659
					}
					position++
				}
			l663:
				{
					position665, tokenIndex665, depth665 := position, tokenIndex, depth
					if buffer[position] != rune('m') {
						goto l666
					}
					position++
					goto l665
				l666:
					position, tokenIndex, depth = position665, tokenIndex665, depth665
					if buffer[position] != rune('M') {
						goto l659
					}
					position++
				}
			l665:
				{
					position667, tokenIndex667, depth667 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l668
					}
					position++
					goto l667
				l668:
					position, tokenIndex, depth = position667, tokenIndex667, depth667
					if buffer[position] != rune('I') {
						goto l659
					}
					position++
				}
			l667:
				{
					position669, tokenIndex669, depth669 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l670
					}
					position++
					goto l669
				l670:
					position, tokenIndex, depth = position669, tokenIndex669, depth669
					if buffer[position] != rune('T') {
						goto l659
					}
					position++
				}
			l669:
				if !_rules[rulesp]() {
					goto l659
				}
				if !_rules[ruleNumericLiteral]() {
					goto l659
				}
				if !_rules[ruleAction27]() {
					goto l659
				}
				depth--
				add(ruleEmitterLimit, position660)
			}
			return true
		l659:
			position, tokenIndex, depth = position659, tokenIndex659, depth659
			return false
		},
		/* 35 EmitterSample <- <(CountBasedSampling / RandomizedSampling / TimeBasedSampling)> */
		func() bool {
			position671, tokenIndex671, depth671 := position, tokenIndex, depth
			{
				position672 := position
				depth++
				{
					position673, tokenIndex673, depth673 := position, tokenIndex, depth
					if !_rules[ruleCountBasedSampling]() {
						goto l674
					}
					goto l673
				l674:
					position, tokenIndex, depth = position673, tokenIndex673, depth673
					if !_rules[ruleRandomizedSampling]() {
						goto l675
					}
					goto l673
				l675:
					position, tokenIndex, depth = position673, tokenIndex673, depth673
					if !_rules[ruleTimeBasedSampling]() {
						goto l671
					}
				}
			l673:
				depth--
				add(ruleEmitterSample, position672)
			}
			return true
		l671:
			position, tokenIndex, depth = position671, tokenIndex671, depth671
			return false
		},
		/* 36 CountBasedSampling <- <(('e' / 'E') ('v' / 'V') ('e' / 'E') ('r' / 'R') ('y' / 'Y') sp NumericLiteral spOpt '-'? spOpt ((('s' / 'S') ('t' / 'T')) / (('n' / 'N') ('d' / 'D')) / (('r' / 'R') ('d' / 'D')) / (('t' / 'T') ('h' / 'H'))) sp (('t' / 'T') ('u' / 'U') ('p' / 'P') ('l' / 'L') ('e' / 'E')) Action28)> */
		func() bool {
			position676, tokenIndex676, depth676 := position, tokenIndex, depth
			{
				position677 := position
				depth++
				{
					position678, tokenIndex678, depth678 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l679
					}
					position++
					goto l678
				l679:
					position, tokenIndex, depth = position678, tokenIndex678, depth678
					if buffer[position] != rune('E') {
						goto l676
					}
					position++
				}
			l678:
				{
					position680, tokenIndex680, depth680 := position, tokenIndex, depth
					if buffer[position] != rune('v') {
						goto l681
					}
					position++
					goto l680
				l681:
					position, tokenIndex, depth = position680, tokenIndex680, depth680
					if buffer[position] != rune('V') {
						goto l676
					}
					position++
				}
			l680:
				{
					position682, tokenIndex682, depth682 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l683
					}
					position++
					goto l682
				l683:
					position, tokenIndex, depth = position682, tokenIndex682, depth682
					if buffer[position] != rune('E') {
						goto l676
					}
					position++
				}
			l682:
				{
					position684, tokenIndex684, depth684 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l685
					}
					position++
					goto l684
				l685:
					position, tokenIndex, depth = position684, tokenIndex684, depth684
					if buffer[position] != rune('R') {
						goto l676
					}
					position++
				}
			l684:
				{
					position686, tokenIndex686, depth686 := position, tokenIndex, depth
					if buffer[position] != rune('y') {
						goto l687
					}
					position++
					goto l686
				l687:
					position, tokenIndex, depth = position686, tokenIndex686, depth686
					if buffer[position] != rune('Y') {
						goto l676
					}
					position++
				}
			l686:
				if !_rules[rulesp]() {
					goto l676
				}
				if !_rules[ruleNumericLiteral]() {
					goto l676
				}
				if !_rules[rulespOpt]() {
					goto l676
				}
				{
					position688, tokenIndex688, depth688 := position, tokenIndex, depth
					if buffer[position] != rune('-') {
						goto l688
					}
					position++
					goto l689
				l688:
					position, tokenIndex, depth = position688, tokenIndex688, depth688
				}
			l689:
				if !_rules[rulespOpt]() {
					goto l676
				}
				{
					position690, tokenIndex690, depth690 := position, tokenIndex, depth
					{
						position692, tokenIndex692, depth692 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l693
						}
						position++
						goto l692
					l693:
						position, tokenIndex, depth = position692, tokenIndex692, depth692
						if buffer[position] != rune('S') {
							goto l691
						}
						position++
					}
				l692:
					{
						position694, tokenIndex694, depth694 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l695
						}
						position++
						goto l694
					l695:
						position, tokenIndex, depth = position694, tokenIndex694, depth694
						if buffer[position] != rune('T') {
							goto l691
						}
						position++
					}
				l694:
					goto l690
				l691:
					position, tokenIndex, depth = position690, tokenIndex690, depth690
					{
						position697, tokenIndex697, depth697 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l698
						}
						position++
						goto l697
					l698:
						position, tokenIndex, depth = position697, tokenIndex697, depth697
						if buffer[position] != rune('N') {
							goto l696
						}
						position++
					}
				l697:
					{
						position699, tokenIndex699, depth699 := position, tokenIndex, depth
						if buffer[position] != rune('d') {
							goto l700
						}
						position++
						goto l699
					l700:
						position, tokenIndex, depth = position699, tokenIndex699, depth699
						if buffer[position] != rune('D') {
							goto l696
						}
						position++
					}
				l699:
					goto l690
				l696:
					position, tokenIndex, depth = position690, tokenIndex690, depth690
					{
						position702, tokenIndex702, depth702 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l703
						}
						position++
						goto l702
					l703:
						position, tokenIndex, depth = position702, tokenIndex702, depth702
						if buffer[position] != rune('R') {
							goto l701
						}
						position++
					}
				l702:
					{
						position704, tokenIndex704, depth704 := position, tokenIndex, depth
						if buffer[position] != rune('d') {
							goto l705
						}
						position++
						goto l704
					l705:
						position, tokenIndex, depth = position704, tokenIndex704, depth704
						if buffer[position] != rune('D') {
							goto l701
						}
						position++
					}
				l704:
					goto l690
				l701:
					position, tokenIndex, depth = position690, tokenIndex690, depth690
					{
						position706, tokenIndex706, depth706 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l707
						}
						position++
						goto l706
					l707:
						position, tokenIndex, depth = position706, tokenIndex706, depth706
						if buffer[position] != rune('T') {
							goto l676
						}
						position++
					}
				l706:
					{
						position708, tokenIndex708, depth708 := position, tokenIndex, depth
						if buffer[position] != rune('h') {
							goto l709
						}
						position++
						goto l708
					l709:
						position, tokenIndex, depth = position708, tokenIndex708, depth708
						if buffer[position] != rune('H') {
							goto l676
						}
						position++
					}
				l708:
				}
			l690:
				if !_rules[rulesp]() {
					goto l676
				}
				{
					position710, tokenIndex710, depth710 := position, tokenIndex, depth
					if buffer[position] != rune('t') {
						goto l711
					}
					position++
					goto l710
				l711:
					position, tokenIndex, depth = position710, tokenIndex710, depth710
					if buffer[position] != rune('T') {
						goto l676
					}
					position++
				}
			l710:
				{
					position712, tokenIndex712, depth712 := position, tokenIndex, depth
					if buffer[position] != rune('u') {
						goto l713
					}
					position++
					goto l712
				l713:
					position, tokenIndex, depth = position712, tokenIndex712, depth712
					if buffer[position] != rune('U') {
						goto l676
					}
					position++
				}
			l712:
				{
					position714, tokenIndex714, depth714 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l715
					}
					position++
					goto l714
				l715:
					position, tokenIndex, depth = position714, tokenIndex714, depth714
					if buffer[position] != rune('P') {
						goto l676
					}
					position++
				}
			l714:
				{
					position716, tokenIndex716, depth716 := position, tokenIndex, depth
					if buffer[position] != rune('l') {
						goto l717
					}
					position++
					goto l716
				l717:
					position, tokenIndex, depth = position716, tokenIndex716, depth716
					if buffer[position] != rune('L') {
						goto l676
					}
					position++
				}
			l716:
				{
					position718, tokenIndex718, depth718 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l719
					}
					position++
					goto l718
				l719:
					position, tokenIndex, depth = position718, tokenIndex718, depth718
					if buffer[position] != rune('E') {
						goto l676
					}
					position++
				}
			l718:
				if !_rules[ruleAction28]() {
					goto l676
				}
				depth--
				add(ruleCountBasedSampling, position677)
			}
			return true
		l676:
			position, tokenIndex, depth = position676, tokenIndex676, depth676
			return false
		},
		/* 37 RandomizedSampling <- <(('s' / 'S') ('a' / 'A') ('m' / 'M') ('p' / 'P') ('l' / 'L') ('e' / 'E') sp (FloatLiteral / NumericLiteral) spOpt '%' Action29)> */
		func() bool {
			position720, tokenIndex720, depth720 := position, tokenIndex, depth
			{
				position721 := position
				depth++
				{
					position722, tokenIndex722, depth722 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l723
					}
					position++
					goto l722
				l723:
					position, tokenIndex, depth = position722, tokenIndex722, depth722
					if buffer[position] != rune('S') {
						goto l720
					}
					position++
				}
			l722:
				{
					position724, tokenIndex724, depth724 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l725
					}
					position++
					goto l724
				l725:
					position, tokenIndex, depth = position724, tokenIndex724, depth724
					if buffer[position] != rune('A') {
						goto l720
					}
					position++
				}
			l724:
				{
					position726, tokenIndex726, depth726 := position, tokenIndex, depth
					if buffer[position] != rune('m') {
						goto l727
					}
					position++
					goto l726
				l727:
					position, tokenIndex, depth = position726, tokenIndex726, depth726
					if buffer[position] != rune('M') {
						goto l720
					}
					position++
				}
			l726:
				{
					position728, tokenIndex728, depth728 := position, tokenIndex, depth
					if buffer[position] != rune('p') {
						goto l729
					}
					position++
					goto l728
				l729:
					position, tokenIndex, depth = position728, tokenIndex728, depth728
					if buffer[position] != rune('P') {
						goto l720
					}
					position++
				}
			l728:
				{
					position730, tokenIndex730, depth730 := position, tokenIndex, depth
					if buffer[position] != rune('l') {
						goto l731
					}
					position++
					goto l730
				l731:
					position, tokenIndex, depth = position730, tokenIndex730, depth730
					if buffer[position] != rune('L') {
						goto l720
					}
					position++
				}
			l730:
				{
					position732, tokenIndex732, depth732 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l733
					}
					position++
					goto l732
				l733:
					position, tokenIndex, depth = position732, tokenIndex732, depth732
					if buffer[position] != rune('E') {
						goto l720
					}
					position++
				}
			l732:
				if !_rules[rulesp]() {
					goto l720
				}
				{
					position734, tokenIndex734, depth734 := position, tokenIndex, depth
					if !_rules[ruleFloatLiteral]() {
						goto l735
					}
					goto l734
				l735:
					position, tokenIndex, depth = position734, tokenIndex734, depth734
					if !_rules[ruleNumericLiteral]() {
						goto l720
					}
				}
			l734:
				if !_rules[rulespOpt]() {
					goto l720
				}
				if buffer[position] != rune('%') {
					goto l720
				}
				position++
				if !_rules[ruleAction29]() {
					goto l720
				}
				depth--
				add(ruleRandomizedSampling, position721)
			}
			return true
		l720:
			position, tokenIndex, depth = position720, tokenIndex720, depth720
			return false
		},
		/* 38 TimeBasedSampling <- <(TimeBasedSamplingSeconds / TimeBasedSamplingMilliseconds)> */
		func() bool {
			position736, tokenIndex736, depth736 := position, tokenIndex, depth
			{
				position737 := position
				depth++
				{
					position738, tokenIndex738, depth738 := position, tokenIndex, depth
					if !_rules[ruleTimeBasedSamplingSeconds]() {
						goto l739
					}
					goto l738
				l739:
					position, tokenIndex, depth = position738, tokenIndex738, depth738
					if !_rules[ruleTimeBasedSamplingMilliseconds]() {
						goto l736
					}
				}
			l738:
				depth--
				add(ruleTimeBasedSampling, position737)
			}
			return true
		l736:
			position, tokenIndex, depth = position736, tokenIndex736, depth736
			return false
		},
		/* 39 TimeBasedSamplingSeconds <- <(('e' / 'E') ('v' / 'V') ('e' / 'E') ('r' / 'R') ('y' / 'Y') sp (FloatLiteral / NumericLiteral) sp (('s' / 'S') ('e' / 'E') ('c' / 'C') ('o' / 'O') ('n' / 'N') ('d' / 'D') ('s' / 'S')) Action30)> */
		func() bool {
			position740, tokenIndex740, depth740 := position, tokenIndex, depth
			{
				position741 := position
				depth++
				{
					position742, tokenIndex742, depth742 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l743
					}
					position++
					goto l742
				l743:
					position, tokenIndex, depth = position742, tokenIndex742, depth742
					if buffer[position] != rune('E') {
						goto l740
					}
					position++
				}
			l742:
				{
					position744, tokenIndex744, depth744 := position, tokenIndex, depth
					if buffer[position] != rune('v') {
						goto l745
					}
					position++
					goto l744
				l745:
					position, tokenIndex, depth = position744, tokenIndex744, depth744
					if buffer[position] != rune('V') {
						goto l740
					}
					position++
				}
			l744:
				{
					position746, tokenIndex746, depth746 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l747
					}
					position++
					goto l746
				l747:
					position, tokenIndex, depth = position746, tokenIndex746, depth746
					if buffer[position] != rune('E') {
						goto l740
					}
					position++
				}
			l746:
				{
					position748, tokenIndex748, depth748 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l749
					}
					position++
					goto l748
				l749:
					position, tokenIndex, depth = position748, tokenIndex748, depth748
					if buffer[position] != rune('R') {
						goto l740
					}
					position++
				}
			l748:
				{
					position750, tokenIndex750, depth750 := position, tokenIndex, depth
					if buffer[position] != rune('y') {
						goto l751
					}
					position++
					goto l750
				l751:
					position, tokenIndex, depth = position750, tokenIndex750, depth750
					if buffer[position] != rune('Y') {
						goto l740
					}
					position++
				}
			l750:
				if !_rules[rulesp]() {
					goto l740
				}
				{
					position752, tokenIndex752, depth752 := position, tokenIndex, depth
					if !_rules[ruleFloatLiteral]() {
						goto l753
					}
					goto l752
				l753:
					position, tokenIndex, depth = position752, tokenIndex752, depth752
					if !_rules[ruleNumericLiteral]() {
						goto l740
					}
				}
			l752:
				if !_rules[rulesp]() {
					goto l740
				}
				{
					position754, tokenIndex754, depth754 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l755
					}
					position++
					goto l754
				l755:
					position, tokenIndex, depth = position754, tokenIndex754, depth754
					if buffer[position] != rune('S') {
						goto l740
					}
					position++
				}
			l754:
				{
					position756, tokenIndex756, depth756 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l757
					}
					position++
					goto l756
				l757:
					position, tokenIndex, depth = position756, tokenIndex756, depth756
					if buffer[position] != rune('E') {
						goto l740
					}
					position++
				}
			l756:
				{
					position758, tokenIndex758, depth758 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l759
					}
					position++
					goto l758
				l759:
					position, tokenIndex, depth = position758, tokenIndex758, depth758
					if buffer[position] != rune('C') {
						goto l740
					}
					position++
				}
			l758:
				{
					position760, tokenIndex760, depth760 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l761
					}
					position++
					goto l760
				l761:
					position, tokenIndex, depth = position760, tokenIndex760, depth760
					if buffer[position] != rune('O') {
						goto l740
					}
					position++
				}
			l760:
				{
					position762, tokenIndex762, depth762 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l763
					}
					position++
					goto l762
				l763:
					position, tokenIndex, depth = position762, tokenIndex762, depth762
					if buffer[position] != rune('N') {
						goto l740
					}
					position++
				}
			l762:
				{
					position764, tokenIndex764, depth764 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l765
					}
					position++
					goto l764
				l765:
					position, tokenIndex, depth = position764, tokenIndex764, depth764
					if buffer[position] != rune('D') {
						goto l740
					}
					position++
				}
			l764:
				{
					position766, tokenIndex766, depth766 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l767
					}
					position++
					goto l766
				l767:
					position, tokenIndex, depth = position766, tokenIndex766, depth766
					if buffer[position] != rune('S') {
						goto l740
					}
					position++
				}
			l766:
				if !_rules[ruleAction30]() {
					goto l740
				}
				depth--
				add(ruleTimeBasedSamplingSeconds, position741)
			}
			return true
		l740:
			position, tokenIndex, depth = position740, tokenIndex740, depth740
			return false
		},
		/* 40 TimeBasedSamplingMilliseconds <- <(('e' / 'E') ('v' / 'V') ('e' / 'E') ('r' / 'R') ('y' / 'Y') sp (FloatLiteral / NumericLiteral) sp (('m' / 'M') ('i' / 'I') ('l' / 'L') ('l' / 'L') ('i' / 'I') ('s' / 'S') ('e' / 'E') ('c' / 'C') ('o' / 'O') ('n' / 'N') ('d' / 'D') ('s' / 'S')) Action31)> */
		func() bool {
			position768, tokenIndex768, depth768 := position, tokenIndex, depth
			{
				position769 := position
				depth++
				{
					position770, tokenIndex770, depth770 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l771
					}
					position++
					goto l770
				l771:
					position, tokenIndex, depth = position770, tokenIndex770, depth770
					if buffer[position] != rune('E') {
						goto l768
					}
					position++
				}
			l770:
				{
					position772, tokenIndex772, depth772 := position, tokenIndex, depth
					if buffer[position] != rune('v') {
						goto l773
					}
					position++
					goto l772
				l773:
					position, tokenIndex, depth = position772, tokenIndex772, depth772
					if buffer[position] != rune('V') {
						goto l768
					}
					position++
				}
			l772:
				{
					position774, tokenIndex774, depth774 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l775
					}
					position++
					goto l774
				l775:
					position, tokenIndex, depth = position774, tokenIndex774, depth774
					if buffer[position] != rune('E') {
						goto l768
					}
					position++
				}
			l774:
				{
					position776, tokenIndex776, depth776 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l777
					}
					position++
					goto l776
				l777:
					position, tokenIndex, depth = position776, tokenIndex776, depth776
					if buffer[position] != rune('R') {
						goto l768
					}
					position++
				}
			l776:
				{
					position778, tokenIndex778, depth778 := position, tokenIndex, depth
					if buffer[position] != rune('y') {
						goto l779
					}
					position++
					goto l778
				l779:
					position, tokenIndex, depth = position778, tokenIndex778, depth778
					if buffer[position] != rune('Y') {
						goto l768
					}
					position++
				}
			l778:
				if !_rules[rulesp]() {
					goto l768
				}
				{
					position780, tokenIndex780, depth780 := position, tokenIndex, depth
					if !_rules[ruleFloatLiteral]() {
						goto l781
					}
					goto l780
				l781:
					position, tokenIndex, depth = position780, tokenIndex780, depth780
					if !_rules[ruleNumericLiteral]() {
						goto l768
					}
				}
			l780:
				if !_rules[rulesp]() {
					goto l768
				}
				{
					position782, tokenIndex782, depth782 := position, tokenIndex, depth
					if buffer[position] != rune('m') {
						goto l783
					}
					position++
					goto l782
				l783:
					position, tokenIndex, depth = position782, tokenIndex782, depth782
					if buffer[position] != rune('M') {
						goto l768
					}
					position++
				}
			l782:
				{
					position784, tokenIndex784, depth784 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l785
					}
					position++
					goto l784
				l785:
					position, tokenIndex, depth = position784, tokenIndex784, depth784
					if buffer[position] != rune('I') {
						goto l768
					}
					position++
				}
			l784:
				{
					position786, tokenIndex786, depth786 := position, tokenIndex, depth
					if buffer[position] != rune('l') {
						goto l787
					}
					position++
					goto l786
				l787:
					position, tokenIndex, depth = position786, tokenIndex786, depth786
					if buffer[position] != rune('L') {
						goto l768
					}
					position++
				}
			l786:
				{
					position788, tokenIndex788, depth788 := position, tokenIndex, depth
					if buffer[position] != rune('l') {
						goto l789
					}
					position++
					goto l788
				l789:
					position, tokenIndex, depth = position788, tokenIndex788, depth788
					if buffer[position] != rune('L') {
						goto l768
					}
					position++
				}
			l788:
				{
					position790, tokenIndex790, depth790 := position, tokenIndex, depth
					if buffer[position] != rune('i') {
						goto l791
					}
					position++
					goto l790
				l791:
					position, tokenIndex, depth = position790, tokenIndex790, depth790
					if buffer[position] != rune('I') {
						goto l768
					}
					position++
				}
			l790:
				{
					position792, tokenIndex792, depth792 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l793
					}
					position++
					goto l792
				l793:
					position, tokenIndex, depth = position792, tokenIndex792, depth792
					if buffer[position] != rune('S') {
						goto l768
					}
					position++
				}
			l792:
				{
					position794, tokenIndex794, depth794 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l795
					}
					position++
					goto l794
				l795:
					position, tokenIndex, depth = position794, tokenIndex794, depth794
					if buffer[position] != rune('E') {
						goto l768
					}
					position++
				}
			l794:
				{
					position796, tokenIndex796, depth796 := position, tokenIndex, depth
					if buffer[position] != rune('c') {
						goto l797
					}
					position++
					goto l796
				l797:
					position, tokenIndex, depth = position796, tokenIndex796, depth796
					if buffer[position] != rune('C') {
						goto l768
					}
					position++
				}
			l796:
				{
					position798, tokenIndex798, depth798 := position, tokenIndex, depth
					if buffer[position] != rune('o') {
						goto l799
					}
					position++
					goto l798
				l799:
					position, tokenIndex, depth = position798, tokenIndex798, depth798
					if buffer[position] != rune('O') {
						goto l768
					}
					position++
				}
			l798:
				{
					position800, tokenIndex800, depth800 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l801
					}
					position++
					goto l800
				l801:
					position, tokenIndex, depth = position800, tokenIndex800, depth800
					if buffer[position] != rune('N') {
						goto l768
					}
					position++
				}
			l800:
				{
					position802, tokenIndex802, depth802 := position, tokenIndex, depth
					if buffer[position] != rune('d') {
						goto l803
					}
					position++
					goto l802
				l803:
					position, tokenIndex, depth = position802, tokenIndex802, depth802
					if buffer[position] != rune('D') {
						goto l768
					}
					position++
				}
			l802:
				{
					position804, tokenIndex804, depth804 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l805
					}
					position++
					goto l804
				l805:
					position, tokenIndex, depth = position804, tokenIndex804, depth804
					if buffer[position] != rune('S') {
						goto l768
					}
					position++
				}
			l804:
				if !_rules[ruleAction31]() {
					goto l768
				}
				depth--
				add(ruleTimeBasedSamplingMilliseconds, position769)
			}
			return true
		l768:
			position, tokenIndex, depth = position768, tokenIndex768, depth768
			return false
		},
		/* 41 Projections <- <(<(sp Projection (spOpt ',' spOpt Projection)*)> Action32)> */
		func() bool {
			position806, tokenIndex806, depth806 := position, tokenIndex, depth
			{
				position807 := position
				depth++
				{
					position808 := position
					depth++
					if !_rules[rulesp]() {
						goto l806
					}
					if !_rules[ruleProjection]() {
						goto l806
					}
				l809:
					{
						position810, tokenIndex810, depth810 := position, tokenIndex, depth
						if !_rules[rulespOpt]() {
							goto l810
						}
						if buffer[position] != rune(',') {
							goto l810
						}
						position++
						if !_rules[rulespOpt]() {
							goto l810
						}
						if !_rules[ruleProjection]() {
							goto l810
						}
						goto l809
					l810:
						position, tokenIndex, depth = position810, tokenIndex810, depth810
					}
					depth--
					add(rulePegText, position808)
				}
				if !_rules[ruleAction32]() {
					goto l806
				}
				depth--
				add(ruleProjections, position807)
			}
			return true
		l806:
			position, tokenIndex, depth = position806, tokenIndex806, depth806
			return false
		},
		/* 42 Projection <- <(AliasExpression / ExpressionOrWildcard)> */
		func() bool {
			position811, tokenIndex811, depth811 := position, tokenIndex, depth
			{
				position812 := position
				depth++
				{
					position813, tokenIndex813, depth813 := position, tokenIndex, depth
					if !_rules[ruleAliasExpression]() {
						goto l814
					}
					goto l813
				l814:
					position, tokenIndex, depth = position813, tokenIndex813, depth813
					if !_rules[ruleExpressionOrWildcard]() {
						goto l811
					}
				}
			l813:
				depth--
				add(ruleProjection, position812)
			}
			return true
		l811:
			position, tokenIndex, depth = position811, tokenIndex811, depth811
			return false
		},
		/* 43 AliasExpression <- <(ExpressionOrWildcard sp (('a' / 'A') ('s' / 'S')) sp TargetIdentifier Action33)> */
		func() bool {
			position815, tokenIndex815, depth815 := position, tokenIndex, depth
			{
				position816 := position
				depth++
				if !_rules[ruleExpressionOrWildcard]() {
					goto l815
				}
				if !_rules[rulesp]() {
					goto l815
				}
				{
					position817, tokenIndex817, depth817 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l818
					}
					position++
					goto l817
				l818:
					position, tokenIndex, depth = position817, tokenIndex817, depth817
					if buffer[position] != rune('A') {
						goto l815
					}
					position++
				}
			l817:
				{
					position819, tokenIndex819, depth819 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l820
					}
					position++
					goto l819
				l820:
					position, tokenIndex, depth = position819, tokenIndex819, depth819
					if buffer[position] != rune('S') {
						goto l815
					}
					position++
				}
			l819:
				if !_rules[rulesp]() {
					goto l815
				}
				if !_rules[ruleTargetIdentifier]() {
					goto l815
				}
				if !_rules[ruleAction33]() {
					goto l815
				}
				depth--
				add(ruleAliasExpression, position816)
			}
			return true
		l815:
			position, tokenIndex, depth = position815, tokenIndex815, depth815
			return false
		},
		/* 44 WindowedFrom <- <(<(sp (('f' / 'F') ('r' / 'R') ('o' / 'O') ('m' / 'M')) sp Relations)?> Action34)> */
		func() bool {
			position821, tokenIndex821, depth821 := position, tokenIndex, depth
			{
				position822 := position
				depth++
				{
					position823 := position
					depth++
					{
						position824, tokenIndex824, depth824 := position, tokenIndex, depth
						if !_rules[rulesp]() {
							goto l824
						}
						{
							position826, tokenIndex826, depth826 := position, tokenIndex, depth
							if buffer[position] != rune('f') {
								goto l827
							}
							position++
							goto l826
						l827:
							position, tokenIndex, depth = position826, tokenIndex826, depth826
							if buffer[position] != rune('F') {
								goto l824
							}
							position++
						}
					l826:
						{
							position828, tokenIndex828, depth828 := position, tokenIndex, depth
							if buffer[position] != rune('r') {
								goto l829
							}
							position++
							goto l828
						l829:
							position, tokenIndex, depth = position828, tokenIndex828, depth828
							if buffer[position] != rune('R') {
								goto l824
							}
							position++
						}
					l828:
						{
							position830, tokenIndex830, depth830 := position, tokenIndex, depth
							if buffer[position] != rune('o') {
								goto l831
							}
							position++
							goto l830
						l831:
							position, tokenIndex, depth = position830, tokenIndex830, depth830
							if buffer[position] != rune('O') {
								goto l824
							}
							position++
						}
					l830:
						{
							position832, tokenIndex832, depth832 := position, tokenIndex, depth
							if buffer[position] != rune('m') {
								goto l833
							}
							position++
							goto l832
						l833:
							position, tokenIndex, depth = position832, tokenIndex832, depth832
							if buffer[position] != rune('M') {
								goto l824
							}
							position++
						}
					l832:
						if !_rules[rulesp]() {
							goto l824
						}
						if !_rules[ruleRelations]() {
							goto l824
						}
						goto l825
					l824:
						position, tokenIndex, depth = position824, tokenIndex824, depth824
					}
				l825:
					depth--
					add(rulePegText, position823)
				}
				if !_rules[ruleAction34]() {
					goto l821
				}
				depth--
				add(ruleWindowedFrom, position822)
			}
			return true
		l821:
			position, tokenIndex, depth = position821, tokenIndex821, depth821
			return false
		},
		/* 45 Interval <- <(TimeInterval / TuplesInterval)> */
		func() bool {
			position834, tokenIndex834, depth834 := position, tokenIndex, depth
			{
				position835 := position
				depth++
				{
					position836, tokenIndex836, depth836 := position, tokenIndex, depth
					if !_rules[ruleTimeInterval]() {
						goto l837
					}
					goto l836
				l837:
					position, tokenIndex, depth = position836, tokenIndex836, depth836
					if !_rules[ruleTuplesInterval]() {
						goto l834
					}
				}
			l836:
				depth--
				add(ruleInterval, position835)
			}
			return true
		l834:
			position, tokenIndex, depth = position834, tokenIndex834, depth834
			return false
		},
		/* 46 TimeInterval <- <((FloatLiteral / NumericLiteral) sp (SECONDS / MILLISECONDS) Action35)> */
		func() bool {
			position838, tokenIndex838, depth838 := position, tokenIndex, depth
			{
				position839 := position
				depth++
				{
					position840, tokenIndex840, depth840 := position, tokenIndex, depth
					if !_rules[ruleFloatLiteral]() {
						goto l841
					}
					goto l840
				l841:
					position, tokenIndex, depth = position840, tokenIndex840, depth840
					if !_rules[ruleNumericLiteral]() {
						goto l838
					}
				}
			l840:
				if !_rules[rulesp]() {
					goto l838
				}
				{
					position842, tokenIndex842, depth842 := position, tokenIndex, depth
					if !_rules[ruleSECONDS]() {
						goto l843
					}
					goto l842
				l843:
					position, tokenIndex, depth = position842, tokenIndex842, depth842
					if !_rules[ruleMILLISECONDS]() {
						goto l838
					}
				}
			l842:
				if !_rules[ruleAction35]() {
					goto l838
				}
				depth--
				add(ruleTimeInterval, position839)
			}
			return true
		l838:
			position, tokenIndex, depth = position838, tokenIndex838, depth838
			return false
		},
		/* 47 TuplesInterval <- <(NumericLiteral sp TUPLES Action36)> */
		func() bool {
			position844, tokenIndex844, depth844 := position, tokenIndex, depth
			{
				position845 := position
				depth++
				if !_rules[ruleNumericLiteral]() {
					goto l844
				}
				if !_rules[rulesp]() {
					goto l844
				}
				if !_rules[ruleTUPLES]() {
					goto l844
				}
				if !_rules[ruleAction36]() {
					goto l844
				}
				depth--
				add(ruleTuplesInterval, position845)
			}
			return true
		l844:
			position, tokenIndex, depth = position844, tokenIndex844, depth844
			return false
		},
		/* 48 Relations <- <(RelationLike (spOpt ',' spOpt RelationLike)*)> */
		func() bool {
			position846, tokenIndex846, depth846 := position, tokenIndex, depth
			{
				position847 := position
				depth++
				if !_rules[ruleRelationLike]() {
					goto l846
				}
			l848:
				{
					position849, tokenIndex849, depth849 := position, tokenIndex, depth
					if !_rules[rulespOpt]() {
						goto l849
					}
					if buffer[position] != rune(',') {
						goto l849
					}
					position++
					if !_rules[rulespOpt]() {
						goto l849
					}
					if !_rules[ruleRelationLike]() {
						goto l849
					}
					goto l848
				l849:
					position, tokenIndex, depth = position849, tokenIndex849, depth849
				}
				depth--
				add(ruleRelations, position847)
			}
			return true
		l846:
			position, tokenIndex, depth = position846, tokenIndex846, depth846
			return false
		},
		/* 49 Filter <- <(<(sp (('w' / 'W') ('h' / 'H') ('e' / 'E') ('r' / 'R') ('e' / 'E')) sp Expression)?> Action37)> */
		func() bool {
			position850, tokenIndex850, depth850 := position, tokenIndex, depth
			{
				position851 := position
				depth++
				{
					position852 := position
					depth++
					{
						position853, tokenIndex853, depth853 := position, tokenIndex, depth
						if !_rules[rulesp]() {
							goto l853
						}
						{
							position855, tokenIndex855, depth855 := position, tokenIndex, depth
							if buffer[position] != rune('w') {
								goto l856
							}
							position++
							goto l855
						l856:
							position, tokenIndex, depth = position855, tokenIndex855, depth855
							if buffer[position] != rune('W') {
								goto l853
							}
							position++
						}
					l855:
						{
							position857, tokenIndex857, depth857 := position, tokenIndex, depth
							if buffer[position] != rune('h') {
								goto l858
							}
							position++
							goto l857
						l858:
							position, tokenIndex, depth = position857, tokenIndex857, depth857
							if buffer[position] != rune('H') {
								goto l853
							}
							position++
						}
					l857:
						{
							position859, tokenIndex859, depth859 := position, tokenIndex, depth
							if buffer[position] != rune('e') {
								goto l860
							}
							position++
							goto l859
						l860:
							position, tokenIndex, depth = position859, tokenIndex859, depth859
							if buffer[position] != rune('E') {
								goto l853
							}
							position++
						}
					l859:
						{
							position861, tokenIndex861, depth861 := position, tokenIndex, depth
							if buffer[position] != rune('r') {
								goto l862
							}
							position++
							goto l861
						l862:
							position, tokenIndex, depth = position861, tokenIndex861, depth861
							if buffer[position] != rune('R') {
								goto l853
							}
							position++
						}
					l861:
						{
							position863, tokenIndex863, depth863 := position, tokenIndex, depth
							if buffer[position] != rune('e') {
								goto l864
							}
							position++
							goto l863
						l864:
							position, tokenIndex, depth = position863, tokenIndex863, depth863
							if buffer[position] != rune('E') {
								goto l853
							}
							position++
						}
					l863:
						if !_rules[rulesp]() {
							goto l853
						}
						if !_rules[ruleExpression]() {
							goto l853
						}
						goto l854
					l853:
						position, tokenIndex, depth = position853, tokenIndex853, depth853
					}
				l854:
					depth--
					add(rulePegText, position852)
				}
				if !_rules[ruleAction37]() {
					goto l850
				}
				depth--
				add(ruleFilter, position851)
			}
			return true
		l850:
			position, tokenIndex, depth = position850, tokenIndex850, depth850
			return false
		},
		/* 50 Grouping <- <(<(sp (('g' / 'G') ('r' / 'R') ('o' / 'O') ('u' / 'U') ('p' / 'P')) sp (('b' / 'B') ('y' / 'Y')) sp GroupList)?> Action38)> */
		func() bool {
			position865, tokenIndex865, depth865 := position, tokenIndex, depth
			{
				position866 := position
				depth++
				{
					position867 := position
					depth++
					{
						position868, tokenIndex868, depth868 := position, tokenIndex, depth
						if !_rules[rulesp]() {
							goto l868
						}
						{
							position870, tokenIndex870, depth870 := position, tokenIndex, depth
							if buffer[position] != rune('g') {
								goto l871
							}
							position++
							goto l870
						l871:
							position, tokenIndex, depth = position870, tokenIndex870, depth870
							if buffer[position] != rune('G') {
								goto l868
							}
							position++
						}
					l870:
						{
							position872, tokenIndex872, depth872 := position, tokenIndex, depth
							if buffer[position] != rune('r') {
								goto l873
							}
							position++
							goto l872
						l873:
							position, tokenIndex, depth = position872, tokenIndex872, depth872
							if buffer[position] != rune('R') {
								goto l868
							}
							position++
						}
					l872:
						{
							position874, tokenIndex874, depth874 := position, tokenIndex, depth
							if buffer[position] != rune('o') {
								goto l875
							}
							position++
							goto l874
						l875:
							position, tokenIndex, depth = position874, tokenIndex874, depth874
							if buffer[position] != rune('O') {
								goto l868
							}
							position++
						}
					l874:
						{
							position876, tokenIndex876, depth876 := position, tokenIndex, depth
							if buffer[position] != rune('u') {
								goto l877
							}
							position++
							goto l876
						l877:
							position, tokenIndex, depth = position876, tokenIndex876, depth876
							if buffer[position] != rune('U') {
								goto l868
							}
							position++
						}
					l876:
						{
							position878, tokenIndex878, depth878 := position, tokenIndex, depth
							if buffer[position] != rune('p') {
								goto l879
							}
							position++
							goto l878
						l879:
							position, tokenIndex, depth = position878, tokenIndex878, depth878
							if buffer[position] != rune('P') {
								goto l868
							}
							position++
						}
					l878:
						if !_rules[rulesp]() {
							goto l868
						}
						{
							position880, tokenIndex880, depth880 := position, tokenIndex, depth
							if buffer[position] != rune('b') {
								goto l881
							}
							position++
							goto l880
						l881:
							position, tokenIndex, depth = position880, tokenIndex880, depth880
							if buffer[position] != rune('B') {
								goto l868
							}
							position++
						}
					l880:
						{
							position882, tokenIndex882, depth882 := position, tokenIndex, depth
							if buffer[position] != rune('y') {
								goto l883
							}
							position++
							goto l882
						l883:
							position, tokenIndex, depth = position882, tokenIndex882, depth882
							if buffer[position] != rune('Y') {
								goto l868
							}
							position++
						}
					l882:
						if !_rules[rulesp]() {
							goto l868
						}
						if !_rules[ruleGroupList]() {
							goto l868
						}
						goto l869
					l868:
						position, tokenIndex, depth = position868, tokenIndex868, depth868
					}
				l869:
					depth--
					add(rulePegText, position867)
				}
				if !_rules[ruleAction38]() {
					goto l865
				}
				depth--
				add(ruleGrouping, position866)
			}
			return true
		l865:
			position, tokenIndex, depth = position865, tokenIndex865, depth865
			return false
		},
		/* 51 GroupList <- <(Expression (spOpt ',' spOpt Expression)*)> */
		func() bool {
			position884, tokenIndex884, depth884 := position, tokenIndex, depth
			{
				position885 := position
				depth++
				if !_rules[ruleExpression]() {
					goto l884
				}
			l886:
				{
					position887, tokenIndex887, depth887 := position, tokenIndex, depth
					if !_rules[rulespOpt]() {
						goto l887
					}
					if buffer[position] != rune(',') {
						goto l887
					}
					position++
					if !_rules[rulespOpt]() {
						goto l887
					}
					if !_rules[ruleExpression]() {
						goto l887
					}
					goto l886
				l887:
					position, tokenIndex, depth = position887, tokenIndex887, depth887
				}
				depth--
				add(ruleGroupList, position885)
			}
			return true
		l884:
			position, tokenIndex, depth = position884, tokenIndex884, depth884
			return false
		},
		/* 52 Having <- <(<(sp (('h' / 'H') ('a' / 'A') ('v' / 'V') ('i' / 'I') ('n' / 'N') ('g' / 'G')) sp Expression)?> Action39)> */
		func() bool {
			position888, tokenIndex888, depth888 := position, tokenIndex, depth
			{
				position889 := position
				depth++
				{
					position890 := position
					depth++
					{
						position891, tokenIndex891, depth891 := position, tokenIndex, depth
						if !_rules[rulesp]() {
							goto l891
						}
						{
							position893, tokenIndex893, depth893 := position, tokenIndex, depth
							if buffer[position] != rune('h') {
								goto l894
							}
							position++
							goto l893
						l894:
							position, tokenIndex, depth = position893, tokenIndex893, depth893
							if buffer[position] != rune('H') {
								goto l891
							}
							position++
						}
					l893:
						{
							position895, tokenIndex895, depth895 := position, tokenIndex, depth
							if buffer[position] != rune('a') {
								goto l896
							}
							position++
							goto l895
						l896:
							position, tokenIndex, depth = position895, tokenIndex895, depth895
							if buffer[position] != rune('A') {
								goto l891
							}
							position++
						}
					l895:
						{
							position897, tokenIndex897, depth897 := position, tokenIndex, depth
							if buffer[position] != rune('v') {
								goto l898
							}
							position++
							goto l897
						l898:
							position, tokenIndex, depth = position897, tokenIndex897, depth897
							if buffer[position] != rune('V') {
								goto l891
							}
							position++
						}
					l897:
						{
							position899, tokenIndex899, depth899 := position, tokenIndex, depth
							if buffer[position] != rune('i') {
								goto l900
							}
							position++
							goto l899
						l900:
							position, tokenIndex, depth = position899, tokenIndex899, depth899
							if buffer[position] != rune('I') {
								goto l891
							}
							position++
						}
					l899:
						{
							position901, tokenIndex901, depth901 := position, tokenIndex, depth
							if buffer[position] != rune('n') {
								goto l902
							}
							position++
							goto l901
						l902:
							position, tokenIndex, depth = position901, tokenIndex901, depth901
							if buffer[position] != rune('N') {
								goto l891
							}
							position++
						}
					l901:
						{
							position903, tokenIndex903, depth903 := position, tokenIndex, depth
							if buffer[position] != rune('g') {
								goto l904
							}
							position++
							goto l903
						l904:
							position, tokenIndex, depth = position903, tokenIndex903, depth903
							if buffer[position] != rune('G') {
								goto l891
							}
							position++
						}
					l903:
						if !_rules[rulesp]() {
							goto l891
						}
						if !_rules[ruleExpression]() {
							goto l891
						}
						goto l892
					l891:
						position, tokenIndex, depth = position891, tokenIndex891, depth891
					}
				l892:
					depth--
					add(rulePegText, position890)
				}
				if !_rules[ruleAction39]() {
					goto l888
				}
				depth--
				add(ruleHaving, position889)
			}
			return true
		l888:
			position, tokenIndex, depth = position888, tokenIndex888, depth888
			return false
		},
		/* 53 RelationLike <- <(AliasedStreamWindow / (StreamWindow Action40))> */
		func() bool {
			position905, tokenIndex905, depth905 := position, tokenIndex, depth
			{
				position906 := position
				depth++
				{
					position907, tokenIndex907, depth907 := position, tokenIndex, depth
					if !_rules[ruleAliasedStreamWindow]() {
						goto l908
					}
					goto l907
				l908:
					position, tokenIndex, depth = position907, tokenIndex907, depth907
					if !_rules[ruleStreamWindow]() {
						goto l905
					}
					if !_rules[ruleAction40]() {
						goto l905
					}
				}
			l907:
				depth--
				add(ruleRelationLike, position906)
			}
			return true
		l905:
			position, tokenIndex, depth = position905, tokenIndex905, depth905
			return false
		},
		/* 54 AliasedStreamWindow <- <(StreamWindow sp (('a' / 'A') ('s' / 'S')) sp Identifier Action41)> */
		func() bool {
			position909, tokenIndex909, depth909 := position, tokenIndex, depth
			{
				position910 := position
				depth++
				if !_rules[ruleStreamWindow]() {
					goto l909
				}
				if !_rules[rulesp]() {
					goto l909
				}
				{
					position911, tokenIndex911, depth911 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l912
					}
					position++
					goto l911
				l912:
					position, tokenIndex, depth = position911, tokenIndex911, depth911
					if buffer[position] != rune('A') {
						goto l909
					}
					position++
				}
			l911:
				{
					position913, tokenIndex913, depth913 := position, tokenIndex, depth
					if buffer[position] != rune('s') {
						goto l914
					}
					position++
					goto l913
				l914:
					position, tokenIndex, depth = position913, tokenIndex913, depth913
					if buffer[position] != rune('S') {
						goto l909
					}
					position++
				}
			l913:
				if !_rules[rulesp]() {
					goto l909
				}
				if !_rules[ruleIdentifier]() {
					goto l909
				}
				if !_rules[ruleAction41]() {
					goto l909
				}
				depth--
				add(ruleAliasedStreamWindow, position910)
			}
			return true
		l909:
			position, tokenIndex, depth = position909, tokenIndex909, depth909
			return false
		},
		/* 55 StreamWindow <- <(StreamLike spOpt '[' spOpt (('r' / 'R') ('a' / 'A') ('n' / 'N') ('g' / 'G') ('e' / 'E')) sp Interval spOpt ']' Action42)> */
		func() bool {
			position915, tokenIndex915, depth915 := position, tokenIndex, depth
			{
				position916 := position
				depth++
				if !_rules[ruleStreamLike]() {
					goto l915
				}
				if !_rules[rulespOpt]() {
					goto l915
				}
				if buffer[position] != rune('[') {
					goto l915
				}
				position++
				if !_rules[rulespOpt]() {
					goto l915
				}
				{
					position917, tokenIndex917, depth917 := position, tokenIndex, depth
					if buffer[position] != rune('r') {
						goto l918
					}
					position++
					goto l917
				l918:
					position, tokenIndex, depth = position917, tokenIndex917, depth917
					if buffer[position] != rune('R') {
						goto l915
					}
					position++
				}
			l917:
				{
					position919, tokenIndex919, depth919 := position, tokenIndex, depth
					if buffer[position] != rune('a') {
						goto l920
					}
					position++
					goto l919
				l920:
					position, tokenIndex, depth = position919, tokenIndex919, depth919
					if buffer[position] != rune('A') {
						goto l915
					}
					position++
				}
			l919:
				{
					position921, tokenIndex921, depth921 := position, tokenIndex, depth
					if buffer[position] != rune('n') {
						goto l922
					}
					position++
					goto l921
				l922:
					position, tokenIndex, depth = position921, tokenIndex921, depth921
					if buffer[position] != rune('N') {
						goto l915
					}
					position++
				}
			l921:
				{
					position923, tokenIndex923, depth923 := position, tokenIndex, depth
					if buffer[position] != rune('g') {
						goto l924
					}
					position++
					goto l923
				l924:
					position, tokenIndex, depth = position923, tokenIndex923, depth923
					if buffer[position] != rune('G') {
						goto l915
					}
					position++
				}
			l923:
				{
					position925, tokenIndex925, depth925 := position, tokenIndex, depth
					if buffer[position] != rune('e') {
						goto l926
					}
					position++
					goto l925
				l926:
					position, tokenIndex, depth = position925, tokenIndex925, depth925
					if buffer[position] != rune('E') {
						goto l915
					}
					position++
				}
			l925:
				if !_rules[rulesp]() {
					goto l915
				}
				if !_rules[ruleInterval]() {
					goto l915
				}
				if !_rules[rulespOpt]() {
					goto l915
				}
				if buffer[position] != rune(']') {
					goto l915
				}
				position++
				if !_rules[ruleAction42]() {
					goto l915
				}
				depth--
				add(ruleStreamWindow, position916)
			}
			return true
		l915:
			position, tokenIndex, depth = position915, tokenIndex915, depth915
			return false
		},
		/* 56 StreamLike <- <(UDSFFuncApp / Stream)> */
		func() bool {
			position927, tokenIndex927, depth927 := position, tokenIndex, depth
			{
				position928 := position
				depth++
				{
					position929, tokenIndex929, depth929 := position, tokenIndex, depth
					if !_rules[ruleUDSFFuncApp]() {
						goto l930
					}
					goto l929
				l930:
					position, tokenIndex, depth = position929, tokenIndex929, depth929
					if !_rules[ruleStream]() {
						goto l927
					}
				}
			l929:
				depth--
				add(ruleStreamLike, position928)
			}
			return true
		l927:
			position, tokenIndex, depth = position927, tokenIndex927, depth927
			return false
		},
		/* 57 UDSFFuncApp <- <(FuncAppWithoutOrderBy Action43)> */
		func() bool {
			position931, tokenIndex931, depth931 := position, tokenIndex, depth
			{
				position932 := position
				depth++
				if !_rules[ruleFuncAppWithoutOrderBy]() {
					goto l931
				}
				if !_rules[ruleAction43]() {
					goto l931
				}
				depth--
				add(ruleUDSFFuncApp, position932)
			}
			return true
		l931:
			position, tokenIndex, depth = position931, tokenIndex931, depth931
			return false
		},
		/* 58 SourceSinkSpecs <- <(<(sp (('w' / 'W') ('i' / 'I') ('t' / 'T') ('h' / 'H')) sp SourceSinkParam (spOpt ',' spOpt SourceSinkParam)*)?> Action44)> */
		func() bool {
			position933, tokenIndex933, depth933 := position, tokenIndex, depth
			{
				position934 := position
				depth++
				{
					position935 := position
					depth++
					{
						position936, tokenIndex936, depth936 := position, tokenIndex, depth
						if !_rules[rulesp]() {
							goto l936
						}
						{
							position938, tokenIndex938, depth938 := position, tokenIndex, depth
							if buffer[position] != rune('w') {
								goto l939
							}
							position++
							goto l938
						l939:
							position, tokenIndex, depth = position938, tokenIndex938, depth938
							if buffer[position] != rune('W') {
								goto l936
							}
							position++
						}
					l938:
						{
							position940, tokenIndex940, depth940 := position, tokenIndex, depth
							if buffer[position] != rune('i') {
								goto l941
							}
							position++
							goto l940
						l941:
							position, tokenIndex, depth = position940, tokenIndex940, depth940
							if buffer[position] != rune('I') {
								goto l936
							}
							position++
						}
					l940:
						{
							position942, tokenIndex942, depth942 := position, tokenIndex, depth
							if buffer[position] != rune('t') {
								goto l943
							}
							position++
							goto l942
						l943:
							position, tokenIndex, depth = position942, tokenIndex942, depth942
							if buffer[position] != rune('T') {
								goto l936
							}
							position++
						}
					l942:
						{
							position944, tokenIndex944, depth944 := position, tokenIndex, depth
							if buffer[position] != rune('h') {
								goto l945
							}
							position++
							goto l944
						l945:
							position, tokenIndex, depth = position944, tokenIndex944, depth944
							if buffer[position] != rune('H') {
								goto l936
							}
							position++
						}
					l944:
						if !_rules[rulesp]() {
							goto l936
						}
						if !_rules[ruleSourceSinkParam]() {
							goto l936
						}
					l946:
						{
							position947, tokenIndex947, depth947 := position, tokenIndex, depth
							if !_rules[rulespOpt]() {
								goto l947
							}
							if buffer[position] != rune(',') {
								goto l947
							}
							position++
							if !_rules[rulespOpt]() {
								goto l947
							}
							if !_rules[ruleSourceSinkParam]() {
								goto l947
							}
							goto l946
						l947:
							position, tokenIndex, depth = position947, tokenIndex947, depth947
						}
						goto l937
					l936:
						position, tokenIndex, depth = position936, tokenIndex936, depth936
					}
				l937:
					depth--
					add(rulePegText, position935)
				}
				if !_rules[ruleAction44]() {
					goto l933
				}
				depth--
				add(ruleSourceSinkSpecs, position934)
			}
			return true
		l933:
			position, tokenIndex, depth = position933, tokenIndex933, depth933
			return false
		},
		/* 59 UpdateSourceSinkSpecs <- <(<(sp (('s' / 'S') ('e' / 'E') ('t' / 'T')) sp SourceSinkParam (spOpt ',' spOpt SourceSinkParam)*)> Action45)> */
		func() bool {
			position948, tokenIndex948, depth948 := position, tokenIndex, depth
			{
				position949 := position
				depth++
				{
					position950 := position
					depth++
					if !_rules[rulesp]() {
						goto l948
					}
					{
						position951, tokenIndex951, depth951 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l952
						}
						position++
						goto l951
					l952:
						position, tokenIndex, depth = position951, tokenIndex951, depth951
						if buffer[position] != rune('S') {
							goto l948
						}
						position++
					}
				l951:
					{
						position953, tokenIndex953, depth953 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l954
						}
						position++
						goto l953
					l954:
						position, tokenIndex, depth = position953, tokenIndex953, depth953
						if buffer[position] != rune('E') {
							goto l948
						}
						position++
					}
				l953:
					{
						position955, tokenIndex955, depth955 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l956
						}
						position++
						goto l955
					l956:
						position, tokenIndex, depth = position955, tokenIndex955, depth955
						if buffer[position] != rune('T') {
							goto l948
						}
						position++
					}
				l955:
					if !_rules[rulesp]() {
						goto l948
					}
					if !_rules[ruleSourceSinkParam]() {
						goto l948
					}
				l957:
					{
						position958, tokenIndex958, depth958 := position, tokenIndex, depth
						if !_rules[rulespOpt]() {
							goto l958
						}
						if buffer[position] != rune(',') {
							goto l958
						}
						position++
						if !_rules[rulespOpt]() {
							goto l958
						}
						if !_rules[ruleSourceSinkParam]() {
							goto l958
						}
						goto l957
					l958:
						position, tokenIndex, depth = position958, tokenIndex958, depth958
					}
					depth--
					add(rulePegText, position950)
				}
				if !_rules[ruleAction45]() {
					goto l948
				}
				depth--
				add(ruleUpdateSourceSinkSpecs, position949)
			}
			return true
		l948:
			position, tokenIndex, depth = position948, tokenIndex948, depth948
			return false
		},
		/* 60 SetOptSpecs <- <(<(sp (('s' / 'S') ('e' / 'E') ('t' / 'T')) sp SourceSinkParam (spOpt ',' spOpt SourceSinkParam)*)?> Action46)> */
		func() bool {
			position959, tokenIndex959, depth959 := position, tokenIndex, depth
			{
				position960 := position
				depth++
				{
					position961 := position
					depth++
					{
						position962, tokenIndex962, depth962 := position, tokenIndex, depth
						if !_rules[rulesp]() {
							goto l962
						}
						{
							position964, tokenIndex964, depth964 := position, tokenIndex, depth
							if buffer[position] != rune('s') {
								goto l965
							}
							position++
							goto l964
						l965:
							position, tokenIndex, depth = position964, tokenIndex964, depth964
							if buffer[position] != rune('S') {
								goto l962
							}
							position++
						}
					l964:
						{
							position966, tokenIndex966, depth966 := position, tokenIndex, depth
							if buffer[position] != rune('e') {
								goto l967
							}
							position++
							goto l966
						l967:
							position, tokenIndex, depth = position966, tokenIndex966, depth966
							if buffer[position] != rune('E') {
								goto l962
							}
							position++
						}
					l966:
						{
							position968, tokenIndex968, depth968 := position, tokenIndex, depth
							if buffer[position] != rune('t') {
								goto l969
							}
							position++
							goto l968
						l969:
							position, tokenIndex, depth = position968, tokenIndex968, depth968
							if buffer[position] != rune('T') {
								goto l962
							}
							position++
						}
					l968:
						if !_rules[rulesp]() {
							goto l962
						}
						if !_rules[ruleSourceSinkParam]() {
							goto l962
						}
					l970:
						{
							position971, tokenIndex971, depth971 := position, tokenIndex, depth
							if !_rules[rulespOpt]() {
								goto l971
							}
							if buffer[position] != rune(',') {
								goto l971
							}
							position++
							if !_rules[rulespOpt]() {
								goto l971
							}
							if !_rules[ruleSourceSinkParam]() {
								goto l971
							}
							goto l970
						l971:
							position, tokenIndex, depth = position971, tokenIndex971, depth971
						}
						goto l963
					l962:
						position, tokenIndex, depth = position962, tokenIndex962, depth962
					}
				l963:
					depth--
					add(rulePegText, position961)
				}
				if !_rules[ruleAction46]() {
					goto l959
				}
				depth--
				add(ruleSetOptSpecs, position960)
			}
			return true
		l959:
			position, tokenIndex, depth = position959, tokenIndex959, depth959
			return false
		},
		/* 61 SourceSinkParam <- <(SourceSinkParamKey '=' SourceSinkParamVal Action47)> */
		func() bool {
			position972, tokenIndex972, depth972 := position, tokenIndex, depth
			{
				position973 := position
				depth++
				if !_rules[ruleSourceSinkParamKey]() {
					goto l972
				}
				if buffer[position] != rune('=') {
					goto l972
				}
				position++
				if !_rules[ruleSourceSinkParamVal]() {
					goto l972
				}
				if !_rules[ruleAction47]() {
					goto l972
				}
				depth--
				add(ruleSourceSinkParam, position973)
			}
			return true
		l972:
			position, tokenIndex, depth = position972, tokenIndex972, depth972
			return false
		},
		/* 62 SourceSinkParamVal <- <(ParamLiteral / ParamArrayExpr)> */
		func() bool {
			position974, tokenIndex974, depth974 := position, tokenIndex, depth
			{
				position975 := position
				depth++
				{
					position976, tokenIndex976, depth976 := position, tokenIndex, depth
					if !_rules[ruleParamLiteral]() {
						goto l977
					}
					goto l976
				l977:
					position, tokenIndex, depth = position976, tokenIndex976, depth976
					if !_rules[ruleParamArrayExpr]() {
						goto l974
					}
				}
			l976:
				depth--
				add(ruleSourceSinkParamVal, position975)
			}
			return true
		l974:
			position, tokenIndex, depth = position974, tokenIndex974, depth974
			return false
		},
		/* 63 ParamLiteral <- <(BooleanLiteral / Literal)> */
		func() bool {
			position978, tokenIndex978, depth978 := position, tokenIndex, depth
			{
				position979 := position
				depth++
				{
					position980, tokenIndex980, depth980 := position, tokenIndex, depth
					if !_rules[ruleBooleanLiteral]() {
						goto l981
					}
					goto l980
				l981:
					position, tokenIndex, depth = position980, tokenIndex980, depth980
					if !_rules[ruleLiteral]() {
						goto l978
					}
				}
			l980:
				depth--
				add(ruleParamLiteral, position979)
			}
			return true
		l978:
			position, tokenIndex, depth = position978, tokenIndex978, depth978
			return false
		},
		/* 64 ParamArrayExpr <- <(<('[' spOpt (ParamLiteral (',' spOpt ParamLiteral)*)? spOpt ','? spOpt ']')> Action48)> */
		func() bool {
			position982, tokenIndex982, depth982 := position, tokenIndex, depth
			{
				position983 := position
				depth++
				{
					position984 := position
					depth++
					if buffer[position] != rune('[') {
						goto l982
					}
					position++
					if !_rules[rulespOpt]() {
						goto l982
					}
					{
						position985, tokenIndex985, depth985 := position, tokenIndex, depth
						if !_rules[ruleParamLiteral]() {
							goto l985
						}
					l987:
						{
							position988, tokenIndex988, depth988 := position, tokenIndex, depth
							if buffer[position] != rune(',') {
								goto l988
							}
							position++
							if !_rules[rulespOpt]() {
								goto l988
							}
							if !_rules[ruleParamLiteral]() {
								goto l988
							}
							goto l987
						l988:
							position, tokenIndex, depth = position988, tokenIndex988, depth988
						}
						goto l986
					l985:
						position, tokenIndex, depth = position985, tokenIndex985, depth985
					}
				l986:
					if !_rules[rulespOpt]() {
						goto l982
					}
					{
						position989, tokenIndex989, depth989 := position, tokenIndex, depth
						if buffer[position] != rune(',') {
							goto l989
						}
						position++
						goto l990
					l989:
						position, tokenIndex, depth = position989, tokenIndex989, depth989
					}
				l990:
					if !_rules[rulespOpt]() {
						goto l982
					}
					if buffer[position] != rune(']') {
						goto l982
					}
					position++
					depth--
					add(rulePegText, position984)
				}
				if !_rules[ruleAction48]() {
					goto l982
				}
				depth--
				add(ruleParamArrayExpr, position983)
			}
			return true
		l982:
			position, tokenIndex, depth = position982, tokenIndex982, depth982
			return false
		},
		/* 65 PausedOpt <- <(<(sp (Paused / Unpaused))?> Action49)> */
		func() bool {
			position991, tokenIndex991, depth991 := position, tokenIndex, depth
			{
				position992 := position
				depth++
				{
					position993 := position
					depth++
					{
						position994, tokenIndex994, depth994 := position, tokenIndex, depth
						if !_rules[rulesp]() {
							goto l994
						}
						{
							position996, tokenIndex996, depth996 := position, tokenIndex, depth
							if !_rules[rulePaused]() {
								goto l997
							}
							goto l996
						l997:
							position, tokenIndex, depth = position996, tokenIndex996, depth996
							if !_rules[ruleUnpaused]() {
								goto l994
							}
						}
					l996:
						goto l995
					l994:
						position, tokenIndex, depth = position994, tokenIndex994, depth994
					}
				l995:
					depth--
					add(rulePegText, position993)
				}
				if !_rules[ruleAction49]() {
					goto l991
				}
				depth--
				add(rulePausedOpt, position992)
			}
			return true
		l991:
			position, tokenIndex, depth = position991, tokenIndex991, depth991
			return false
		},
		/* 66 ExpressionOrWildcard <- <(Wildcard / Expression)> */
		func() bool {
			position998, tokenIndex998, depth998 := position, tokenIndex, depth
			{
				position999 := position
				depth++
				{
					position1000, tokenIndex1000, depth1000 := position, tokenIndex, depth
					if !_rules[ruleWildcard]() {
						goto l1001
					}
					goto l1000
				l1001:
					position, tokenIndex, depth = position1000, tokenIndex1000, depth1000
					if !_rules[ruleExpression]() {
						goto l998
					}
				}
			l1000:
				depth--
				add(ruleExpressionOrWildcard, position999)
			}
			return true
		l998:
			position, tokenIndex, depth = position998, tokenIndex998, depth998
			return false
		},
		/* 67 Expression <- <orExpr> */
		func() bool {
			position1002, tokenIndex1002, depth1002 := position, tokenIndex, depth
			{
				position1003 := position
				depth++
				if !_rules[ruleorExpr]() {
					goto l1002
				}
				depth--
				add(ruleExpression, position1003)
			}
			return true
		l1002:
			position, tokenIndex, depth = position1002, tokenIndex1002, depth1002
			return false
		},
		/* 68 orExpr <- <(<(andExpr (sp Or sp andExpr)*)> Action50)> */
		func() bool {
			position1004, tokenIndex1004, depth1004 := position, tokenIndex, depth
			{
				position1005 := position
				depth++
				{
					position1006 := position
					depth++
					if !_rules[ruleandExpr]() {
						goto l1004
					}
				l1007:
					{
						position1008, tokenIndex1008, depth1008 := position, tokenIndex, depth
						if !_rules[rulesp]() {
							goto l1008
						}
						if !_rules[ruleOr]() {
							goto l1008
						}
						if !_rules[rulesp]() {
							goto l1008
						}
						if !_rules[ruleandExpr]() {
							goto l1008
						}
						goto l1007
					l1008:
						position, tokenIndex, depth = position1008, tokenIndex1008, depth1008
					}
					depth--
					add(rulePegText, position1006)
				}
				if !_rules[ruleAction50]() {
					goto l1004
				}
				depth--
				add(ruleorExpr, position1005)
			}
			return true
		l1004:
			position, tokenIndex, depth = position1004, tokenIndex1004, depth1004
			return false
		},
		/* 69 andExpr <- <(<(notExpr (sp And sp notExpr)*)> Action51)> */
		func() bool {
			position1009, tokenIndex1009, depth1009 := position, tokenIndex, depth
			{
				position1010 := position
				depth++
				{
					position1011 := position
					depth++
					if !_rules[rulenotExpr]() {
						goto l1009
					}
				l1012:
					{
						position1013, tokenIndex1013, depth1013 := position, tokenIndex, depth
						if !_rules[rulesp]() {
							goto l1013
						}
						if !_rules[ruleAnd]() {
							goto l1013
						}
						if !_rules[rulesp]() {
							goto l1013
						}
						if !_rules[rulenotExpr]() {
							goto l1013
						}
						goto l1012
					l1013:
						position, tokenIndex, depth = position1013, tokenIndex1013, depth1013
					}
					depth--
					add(rulePegText, position1011)
				}
				if !_rules[ruleAction51]() {
					goto l1009
				}
				depth--
				add(ruleandExpr, position1010)
			}
			return true
		l1009:
			position, tokenIndex, depth = position1009, tokenIndex1009, depth1009
			return false
		},
		/* 70 notExpr <- <(<((Not sp)? comparisonExpr)> Action52)> */
		func() bool {
			position1014, tokenIndex1014, depth1014 := position, tokenIndex, depth
			{
				position1015 := position
				depth++
				{
					position1016 := position
					depth++
					{
						position1017, tokenIndex1017, depth1017 := position, tokenIndex, depth
						if !_rules[ruleNot]() {
							goto l1017
						}
						if !_rules[rulesp]() {
							goto l1017
						}
						goto l1018
					l1017:
						position, tokenIndex, depth = position1017, tokenIndex1017, depth1017
					}
				l1018:
					if !_rules[rulecomparisonExpr]() {
						goto l1014
					}
					depth--
					add(rulePegText, position1016)
				}
				if !_rules[ruleAction52]() {
					goto l1014
				}
				depth--
				add(rulenotExpr, position1015)
			}
			return true
		l1014:
			position, tokenIndex, depth = position1014, tokenIndex1014, depth1014
			return false
		},
		/* 71 comparisonExpr <- <(<(otherOpExpr (spOpt ComparisonOp spOpt otherOpExpr)?)> Action53)> */
		func() bool {
			position1019, tokenIndex1019, depth1019 := position, tokenIndex, depth
			{
				position1020 := position
				depth++
				{
					position1021 := position
					depth++
					if !_rules[ruleotherOpExpr]() {
						goto l1019
					}
					{
						position1022, tokenIndex1022, depth1022 := position, tokenIndex, depth
						if !_rules[rulespOpt]() {
							goto l1022
						}
						if !_rules[ruleComparisonOp]() {
							goto l1022
						}
						if !_rules[rulespOpt]() {
							goto l1022
						}
						if !_rules[ruleotherOpExpr]() {
							goto l1022
						}
						goto l1023
					l1022:
						position, tokenIndex, depth = position1022, tokenIndex1022, depth1022
					}
				l1023:
					depth--
					add(rulePegText, position1021)
				}
				if !_rules[ruleAction53]() {
					goto l1019
				}
				depth--
				add(rulecomparisonExpr, position1020)
			}
			return true
		l1019:
			position, tokenIndex, depth = position1019, tokenIndex1019, depth1019
			return false
		},
		/* 72 otherOpExpr <- <(<(isExpr (spOpt OtherOp spOpt isExpr)*)> Action54)> */
		func() bool {
			position1024, tokenIndex1024, depth1024 := position, tokenIndex, depth
			{
				position1025 := position
				depth++
				{
					position1026 := position
					depth++
					if !_rules[ruleisExpr]() {
						goto l1024
					}
				l1027:
					{
						position1028, tokenIndex1028, depth1028 := position, tokenIndex, depth
						if !_rules[rulespOpt]() {
							goto l1028
						}
						if !_rules[ruleOtherOp]() {
							goto l1028
						}
						if !_rules[rulespOpt]() {
							goto l1028
						}
						if !_rules[ruleisExpr]() {
							goto l1028
						}
						goto l1027
					l1028:
						position, tokenIndex, depth = position1028, tokenIndex1028, depth1028
					}
					depth--
					add(rulePegText, position1026)
				}
				if !_rules[ruleAction54]() {
					goto l1024
				}
				depth--
				add(ruleotherOpExpr, position1025)
			}
			return true
		l1024:
			position, tokenIndex, depth = position1024, tokenIndex1024, depth1024
			return false
		},
		/* 73 isExpr <- <(<(termExpr (sp IsOp sp NullLiteral)?)> Action55)> */
		func() bool {
			position1029, tokenIndex1029, depth1029 := position, tokenIndex, depth
			{
				position1030 := position
				depth++
				{
					position1031 := position
					depth++
					if !_rules[ruletermExpr]() {
						goto l1029
					}
					{
						position1032, tokenIndex1032, depth1032 := position, tokenIndex, depth
						if !_rules[rulesp]() {
							goto l1032
						}
						if !_rules[ruleIsOp]() {
							goto l1032
						}
						if !_rules[rulesp]() {
							goto l1032
						}
						if !_rules[ruleNullLiteral]() {
							goto l1032
						}
						goto l1033
					l1032:
						position, tokenIndex, depth = position1032, tokenIndex1032, depth1032
					}
				l1033:
					depth--
					add(rulePegText, position1031)
				}
				if !_rules[ruleAction55]() {
					goto l1029
				}
				depth--
				add(ruleisExpr, position1030)
			}
			return true
		l1029:
			position, tokenIndex, depth = position1029, tokenIndex1029, depth1029
			return false
		},
		/* 74 termExpr <- <(<(productExpr (spOpt PlusMinusOp spOpt productExpr)*)> Action56)> */
		func() bool {
			position1034, tokenIndex1034, depth1034 := position, tokenIndex, depth
			{
				position1035 := position
				depth++
				{
					position1036 := position
					depth++
					if !_rules[ruleproductExpr]() {
						goto l1034
					}
				l1037:
					{
						position1038, tokenIndex1038, depth1038 := position, tokenIndex, depth
						if !_rules[rulespOpt]() {
							goto l1038
						}
						if !_rules[rulePlusMinusOp]() {
							goto l1038
						}
						if !_rules[rulespOpt]() {
							goto l1038
						}
						if !_rules[ruleproductExpr]() {
							goto l1038
						}
						goto l1037
					l1038:
						position, tokenIndex, depth = position1038, tokenIndex1038, depth1038
					}
					depth--
					add(rulePegText, position1036)
				}
				if !_rules[ruleAction56]() {
					goto l1034
				}
				depth--
				add(ruletermExpr, position1035)
			}
			return true
		l1034:
			position, tokenIndex, depth = position1034, tokenIndex1034, depth1034
			return false
		},
		/* 75 productExpr <- <(<(minusExpr (spOpt MultDivOp spOpt minusExpr)*)> Action57)> */
		func() bool {
			position1039, tokenIndex1039, depth1039 := position, tokenIndex, depth
			{
				position1040 := position
				depth++
				{
					position1041 := position
					depth++
					if !_rules[ruleminusExpr]() {
						goto l1039
					}
				l1042:
					{
						position1043, tokenIndex1043, depth1043 := position, tokenIndex, depth
						if !_rules[rulespOpt]() {
							goto l1043
						}
						if !_rules[ruleMultDivOp]() {
							goto l1043
						}
						if !_rules[rulespOpt]() {
							goto l1043
						}
						if !_rules[ruleminusExpr]() {
							goto l1043
						}
						goto l1042
					l1043:
						position, tokenIndex, depth = position1043, tokenIndex1043, depth1043
					}
					depth--
					add(rulePegText, position1041)
				}
				if !_rules[ruleAction57]() {
					goto l1039
				}
				depth--
				add(ruleproductExpr, position1040)
			}
			return true
		l1039:
			position, tokenIndex, depth = position1039, tokenIndex1039, depth1039
			return false
		},
		/* 76 minusExpr <- <(<((UnaryMinus spOpt)? castExpr)> Action58)> */
		func() bool {
			position1044, tokenIndex1044, depth1044 := position, tokenIndex, depth
			{
				position1045 := position
				depth++
				{
					position1046 := position
					depth++
					{
						position1047, tokenIndex1047, depth1047 := position, tokenIndex, depth
						if !_rules[ruleUnaryMinus]() {
							goto l1047
						}
						if !_rules[rulespOpt]() {
							goto l1047
						}
						goto l1048
					l1047:
						position, tokenIndex, depth = position1047, tokenIndex1047, depth1047
					}
				l1048:
					if !_rules[rulecastExpr]() {
						goto l1044
					}
					depth--
					add(rulePegText, position1046)
				}
				if !_rules[ruleAction58]() {
					goto l1044
				}
				depth--
				add(ruleminusExpr, position1045)
			}
			return true
		l1044:
			position, tokenIndex, depth = position1044, tokenIndex1044, depth1044
			return false
		},
		/* 77 castExpr <- <(<(baseExpr (spOpt (':' ':') spOpt Type)?)> Action59)> */
		func() bool {
			position1049, tokenIndex1049, depth1049 := position, tokenIndex, depth
			{
				position1050 := position
				depth++
				{
					position1051 := position
					depth++
					if !_rules[rulebaseExpr]() {
						goto l1049
					}
					{
						position1052, tokenIndex1052, depth1052 := position, tokenIndex, depth
						if !_rules[rulespOpt]() {
							goto l1052
						}
						if buffer[position] != rune(':') {
							goto l1052
						}
						position++
						if buffer[position] != rune(':') {
							goto l1052
						}
						position++
						if !_rules[rulespOpt]() {
							goto l1052
						}
						if !_rules[ruleType]() {
							goto l1052
						}
						goto l1053
					l1052:
						position, tokenIndex, depth = position1052, tokenIndex1052, depth1052
					}
				l1053:
					depth--
					add(rulePegText, position1051)
				}
				if !_rules[ruleAction59]() {
					goto l1049
				}
				depth--
				add(rulecastExpr, position1050)
			}
			return true
		l1049:
			position, tokenIndex, depth = position1049, tokenIndex1049, depth1049
			return false
		},
		/* 78 baseExpr <- <(('(' spOpt Expression spOpt ')') / MapExpr / BooleanLiteral / NullLiteral / RowMeta / FuncTypeCast / FuncApp / RowValue / ArrayExpr / Literal)> */
		func() bool {
			position1054, tokenIndex1054, depth1054 := position, tokenIndex, depth
			{
				position1055 := position
				depth++
				{
					position1056, tokenIndex1056, depth1056 := position, tokenIndex, depth
					if buffer[position] != rune('(') {
						goto l1057
					}
					position++
					if !_rules[rulespOpt]() {
						goto l1057
					}
					if !_rules[ruleExpression]() {
						goto l1057
					}
					if !_rules[rulespOpt]() {
						goto l1057
					}
					if buffer[position] != rune(')') {
						goto l1057
					}
					position++
					goto l1056
				l1057:
					position, tokenIndex, depth = position1056, tokenIndex1056, depth1056
					if !_rules[ruleMapExpr]() {
						goto l1058
					}
					goto l1056
				l1058:
					position, tokenIndex, depth = position1056, tokenIndex1056, depth1056
					if !_rules[ruleBooleanLiteral]() {
						goto l1059
					}
					goto l1056
				l1059:
					position, tokenIndex, depth = position1056, tokenIndex1056, depth1056
					if !_rules[ruleNullLiteral]() {
						goto l1060
					}
					goto l1056
				l1060:
					position, tokenIndex, depth = position1056, tokenIndex1056, depth1056
					if !_rules[ruleRowMeta]() {
						goto l1061
					}
					goto l1056
				l1061:
					position, tokenIndex, depth = position1056, tokenIndex1056, depth1056
					if !_rules[ruleFuncTypeCast]() {
						goto l1062
					}
					goto l1056
				l1062:
					position, tokenIndex, depth = position1056, tokenIndex1056, depth1056
					if !_rules[ruleFuncApp]() {
						goto l1063
					}
					goto l1056
				l1063:
					position, tokenIndex, depth = position1056, tokenIndex1056, depth1056
					if !_rules[ruleRowValue]() {
						goto l1064
					}
					goto l1056
				l1064:
					position, tokenIndex, depth = position1056, tokenIndex1056, depth1056
					if !_rules[ruleArrayExpr]() {
						goto l1065
					}
					goto l1056
				l1065:
					position, tokenIndex, depth = position1056, tokenIndex1056, depth1056
					if !_rules[ruleLiteral]() {
						goto l1054
					}
				}
			l1056:
				depth--
				add(rulebaseExpr, position1055)
			}
			return true
		l1054:
			position, tokenIndex, depth = position1054, tokenIndex1054, depth1054
			return false
		},
		/* 79 FuncTypeCast <- <(<(('c' / 'C') ('a' / 'A') ('s' / 'S') ('t' / 'T') spOpt '(' spOpt Expression sp (('a' / 'A') ('s' / 'S')) sp Type spOpt ')')> Action60)> */
		func() bool {
			position1066, tokenIndex1066, depth1066 := position, tokenIndex, depth
			{
				position1067 := position
				depth++
				{
					position1068 := position
					depth++
					{
						position1069, tokenIndex1069, depth1069 := position, tokenIndex, depth
						if buffer[position] != rune('c') {
							goto l1070
						}
						position++
						goto l1069
					l1070:
						position, tokenIndex, depth = position1069, tokenIndex1069, depth1069
						if buffer[position] != rune('C') {
							goto l1066
						}
						position++
					}
				l1069:
					{
						position1071, tokenIndex1071, depth1071 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1072
						}
						position++
						goto l1071
					l1072:
						position, tokenIndex, depth = position1071, tokenIndex1071, depth1071
						if buffer[position] != rune('A') {
							goto l1066
						}
						position++
					}
				l1071:
					{
						position1073, tokenIndex1073, depth1073 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1074
						}
						position++
						goto l1073
					l1074:
						position, tokenIndex, depth = position1073, tokenIndex1073, depth1073
						if buffer[position] != rune('S') {
							goto l1066
						}
						position++
					}
				l1073:
					{
						position1075, tokenIndex1075, depth1075 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1076
						}
						position++
						goto l1075
					l1076:
						position, tokenIndex, depth = position1075, tokenIndex1075, depth1075
						if buffer[position] != rune('T') {
							goto l1066
						}
						position++
					}
				l1075:
					if !_rules[rulespOpt]() {
						goto l1066
					}
					if buffer[position] != rune('(') {
						goto l1066
					}
					position++
					if !_rules[rulespOpt]() {
						goto l1066
					}
					if !_rules[ruleExpression]() {
						goto l1066
					}
					if !_rules[rulesp]() {
						goto l1066
					}
					{
						position1077, tokenIndex1077, depth1077 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1078
						}
						position++
						goto l1077
					l1078:
						position, tokenIndex, depth = position1077, tokenIndex1077, depth1077
						if buffer[position] != rune('A') {
							goto l1066
						}
						position++
					}
				l1077:
					{
						position1079, tokenIndex1079, depth1079 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1080
						}
						position++
						goto l1079
					l1080:
						position, tokenIndex, depth = position1079, tokenIndex1079, depth1079
						if buffer[position] != rune('S') {
							goto l1066
						}
						position++
					}
				l1079:
					if !_rules[rulesp]() {
						goto l1066
					}
					if !_rules[ruleType]() {
						goto l1066
					}
					if !_rules[rulespOpt]() {
						goto l1066
					}
					if buffer[position] != rune(')') {
						goto l1066
					}
					position++
					depth--
					add(rulePegText, position1068)
				}
				if !_rules[ruleAction60]() {
					goto l1066
				}
				depth--
				add(ruleFuncTypeCast, position1067)
			}
			return true
		l1066:
			position, tokenIndex, depth = position1066, tokenIndex1066, depth1066
			return false
		},
		/* 80 FuncApp <- <(FuncAppWithOrderBy / FuncAppWithoutOrderBy)> */
		func() bool {
			position1081, tokenIndex1081, depth1081 := position, tokenIndex, depth
			{
				position1082 := position
				depth++
				{
					position1083, tokenIndex1083, depth1083 := position, tokenIndex, depth
					if !_rules[ruleFuncAppWithOrderBy]() {
						goto l1084
					}
					goto l1083
				l1084:
					position, tokenIndex, depth = position1083, tokenIndex1083, depth1083
					if !_rules[ruleFuncAppWithoutOrderBy]() {
						goto l1081
					}
				}
			l1083:
				depth--
				add(ruleFuncApp, position1082)
			}
			return true
		l1081:
			position, tokenIndex, depth = position1081, tokenIndex1081, depth1081
			return false
		},
		/* 81 FuncAppWithOrderBy <- <(Function spOpt '(' spOpt FuncParams sp ParamsOrder spOpt ')' Action61)> */
		func() bool {
			position1085, tokenIndex1085, depth1085 := position, tokenIndex, depth
			{
				position1086 := position
				depth++
				if !_rules[ruleFunction]() {
					goto l1085
				}
				if !_rules[rulespOpt]() {
					goto l1085
				}
				if buffer[position] != rune('(') {
					goto l1085
				}
				position++
				if !_rules[rulespOpt]() {
					goto l1085
				}
				if !_rules[ruleFuncParams]() {
					goto l1085
				}
				if !_rules[rulesp]() {
					goto l1085
				}
				if !_rules[ruleParamsOrder]() {
					goto l1085
				}
				if !_rules[rulespOpt]() {
					goto l1085
				}
				if buffer[position] != rune(')') {
					goto l1085
				}
				position++
				if !_rules[ruleAction61]() {
					goto l1085
				}
				depth--
				add(ruleFuncAppWithOrderBy, position1086)
			}
			return true
		l1085:
			position, tokenIndex, depth = position1085, tokenIndex1085, depth1085
			return false
		},
		/* 82 FuncAppWithoutOrderBy <- <(Function spOpt '(' spOpt FuncParams <spOpt> ')' Action62)> */
		func() bool {
			position1087, tokenIndex1087, depth1087 := position, tokenIndex, depth
			{
				position1088 := position
				depth++
				if !_rules[ruleFunction]() {
					goto l1087
				}
				if !_rules[rulespOpt]() {
					goto l1087
				}
				if buffer[position] != rune('(') {
					goto l1087
				}
				position++
				if !_rules[rulespOpt]() {
					goto l1087
				}
				if !_rules[ruleFuncParams]() {
					goto l1087
				}
				{
					position1089 := position
					depth++
					if !_rules[rulespOpt]() {
						goto l1087
					}
					depth--
					add(rulePegText, position1089)
				}
				if buffer[position] != rune(')') {
					goto l1087
				}
				position++
				if !_rules[ruleAction62]() {
					goto l1087
				}
				depth--
				add(ruleFuncAppWithoutOrderBy, position1088)
			}
			return true
		l1087:
			position, tokenIndex, depth = position1087, tokenIndex1087, depth1087
			return false
		},
		/* 83 FuncParams <- <(<(ExpressionOrWildcard (spOpt ',' spOpt ExpressionOrWildcard)*)?> Action63)> */
		func() bool {
			position1090, tokenIndex1090, depth1090 := position, tokenIndex, depth
			{
				position1091 := position
				depth++
				{
					position1092 := position
					depth++
					{
						position1093, tokenIndex1093, depth1093 := position, tokenIndex, depth
						if !_rules[ruleExpressionOrWildcard]() {
							goto l1093
						}
					l1095:
						{
							position1096, tokenIndex1096, depth1096 := position, tokenIndex, depth
							if !_rules[rulespOpt]() {
								goto l1096
							}
							if buffer[position] != rune(',') {
								goto l1096
							}
							position++
							if !_rules[rulespOpt]() {
								goto l1096
							}
							if !_rules[ruleExpressionOrWildcard]() {
								goto l1096
							}
							goto l1095
						l1096:
							position, tokenIndex, depth = position1096, tokenIndex1096, depth1096
						}
						goto l1094
					l1093:
						position, tokenIndex, depth = position1093, tokenIndex1093, depth1093
					}
				l1094:
					depth--
					add(rulePegText, position1092)
				}
				if !_rules[ruleAction63]() {
					goto l1090
				}
				depth--
				add(ruleFuncParams, position1091)
			}
			return true
		l1090:
			position, tokenIndex, depth = position1090, tokenIndex1090, depth1090
			return false
		},
		/* 84 ParamsOrder <- <(<(('o' / 'O') ('r' / 'R') ('d' / 'D') ('e' / 'E') ('r' / 'R') sp (('b' / 'B') ('y' / 'Y')) sp SortedExpression (spOpt ',' spOpt SortedExpression)*)> Action64)> */
		func() bool {
			position1097, tokenIndex1097, depth1097 := position, tokenIndex, depth
			{
				position1098 := position
				depth++
				{
					position1099 := position
					depth++
					{
						position1100, tokenIndex1100, depth1100 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l1101
						}
						position++
						goto l1100
					l1101:
						position, tokenIndex, depth = position1100, tokenIndex1100, depth1100
						if buffer[position] != rune('O') {
							goto l1097
						}
						position++
					}
				l1100:
					{
						position1102, tokenIndex1102, depth1102 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l1103
						}
						position++
						goto l1102
					l1103:
						position, tokenIndex, depth = position1102, tokenIndex1102, depth1102
						if buffer[position] != rune('R') {
							goto l1097
						}
						position++
					}
				l1102:
					{
						position1104, tokenIndex1104, depth1104 := position, tokenIndex, depth
						if buffer[position] != rune('d') {
							goto l1105
						}
						position++
						goto l1104
					l1105:
						position, tokenIndex, depth = position1104, tokenIndex1104, depth1104
						if buffer[position] != rune('D') {
							goto l1097
						}
						position++
					}
				l1104:
					{
						position1106, tokenIndex1106, depth1106 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l1107
						}
						position++
						goto l1106
					l1107:
						position, tokenIndex, depth = position1106, tokenIndex1106, depth1106
						if buffer[position] != rune('E') {
							goto l1097
						}
						position++
					}
				l1106:
					{
						position1108, tokenIndex1108, depth1108 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l1109
						}
						position++
						goto l1108
					l1109:
						position, tokenIndex, depth = position1108, tokenIndex1108, depth1108
						if buffer[position] != rune('R') {
							goto l1097
						}
						position++
					}
				l1108:
					if !_rules[rulesp]() {
						goto l1097
					}
					{
						position1110, tokenIndex1110, depth1110 := position, tokenIndex, depth
						if buffer[position] != rune('b') {
							goto l1111
						}
						position++
						goto l1110
					l1111:
						position, tokenIndex, depth = position1110, tokenIndex1110, depth1110
						if buffer[position] != rune('B') {
							goto l1097
						}
						position++
					}
				l1110:
					{
						position1112, tokenIndex1112, depth1112 := position, tokenIndex, depth
						if buffer[position] != rune('y') {
							goto l1113
						}
						position++
						goto l1112
					l1113:
						position, tokenIndex, depth = position1112, tokenIndex1112, depth1112
						if buffer[position] != rune('Y') {
							goto l1097
						}
						position++
					}
				l1112:
					if !_rules[rulesp]() {
						goto l1097
					}
					if !_rules[ruleSortedExpression]() {
						goto l1097
					}
				l1114:
					{
						position1115, tokenIndex1115, depth1115 := position, tokenIndex, depth
						if !_rules[rulespOpt]() {
							goto l1115
						}
						if buffer[position] != rune(',') {
							goto l1115
						}
						position++
						if !_rules[rulespOpt]() {
							goto l1115
						}
						if !_rules[ruleSortedExpression]() {
							goto l1115
						}
						goto l1114
					l1115:
						position, tokenIndex, depth = position1115, tokenIndex1115, depth1115
					}
					depth--
					add(rulePegText, position1099)
				}
				if !_rules[ruleAction64]() {
					goto l1097
				}
				depth--
				add(ruleParamsOrder, position1098)
			}
			return true
		l1097:
			position, tokenIndex, depth = position1097, tokenIndex1097, depth1097
			return false
		},
		/* 85 SortedExpression <- <(Expression OrderDirectionOpt Action65)> */
		func() bool {
			position1116, tokenIndex1116, depth1116 := position, tokenIndex, depth
			{
				position1117 := position
				depth++
				if !_rules[ruleExpression]() {
					goto l1116
				}
				if !_rules[ruleOrderDirectionOpt]() {
					goto l1116
				}
				if !_rules[ruleAction65]() {
					goto l1116
				}
				depth--
				add(ruleSortedExpression, position1117)
			}
			return true
		l1116:
			position, tokenIndex, depth = position1116, tokenIndex1116, depth1116
			return false
		},
		/* 86 OrderDirectionOpt <- <(<(sp (Ascending / Descending))?> Action66)> */
		func() bool {
			position1118, tokenIndex1118, depth1118 := position, tokenIndex, depth
			{
				position1119 := position
				depth++
				{
					position1120 := position
					depth++
					{
						position1121, tokenIndex1121, depth1121 := position, tokenIndex, depth
						if !_rules[rulesp]() {
							goto l1121
						}
						{
							position1123, tokenIndex1123, depth1123 := position, tokenIndex, depth
							if !_rules[ruleAscending]() {
								goto l1124
							}
							goto l1123
						l1124:
							position, tokenIndex, depth = position1123, tokenIndex1123, depth1123
							if !_rules[ruleDescending]() {
								goto l1121
							}
						}
					l1123:
						goto l1122
					l1121:
						position, tokenIndex, depth = position1121, tokenIndex1121, depth1121
					}
				l1122:
					depth--
					add(rulePegText, position1120)
				}
				if !_rules[ruleAction66]() {
					goto l1118
				}
				depth--
				add(ruleOrderDirectionOpt, position1119)
			}
			return true
		l1118:
			position, tokenIndex, depth = position1118, tokenIndex1118, depth1118
			return false
		},
		/* 87 ArrayExpr <- <(<('[' spOpt (ExpressionOrWildcard (spOpt ',' spOpt ExpressionOrWildcard)*)? spOpt ','? spOpt ']')> Action67)> */
		func() bool {
			position1125, tokenIndex1125, depth1125 := position, tokenIndex, depth
			{
				position1126 := position
				depth++
				{
					position1127 := position
					depth++
					if buffer[position] != rune('[') {
						goto l1125
					}
					position++
					if !_rules[rulespOpt]() {
						goto l1125
					}
					{
						position1128, tokenIndex1128, depth1128 := position, tokenIndex, depth
						if !_rules[ruleExpressionOrWildcard]() {
							goto l1128
						}
					l1130:
						{
							position1131, tokenIndex1131, depth1131 := position, tokenIndex, depth
							if !_rules[rulespOpt]() {
								goto l1131
							}
							if buffer[position] != rune(',') {
								goto l1131
							}
							position++
							if !_rules[rulespOpt]() {
								goto l1131
							}
							if !_rules[ruleExpressionOrWildcard]() {
								goto l1131
							}
							goto l1130
						l1131:
							position, tokenIndex, depth = position1131, tokenIndex1131, depth1131
						}
						goto l1129
					l1128:
						position, tokenIndex, depth = position1128, tokenIndex1128, depth1128
					}
				l1129:
					if !_rules[rulespOpt]() {
						goto l1125
					}
					{
						position1132, tokenIndex1132, depth1132 := position, tokenIndex, depth
						if buffer[position] != rune(',') {
							goto l1132
						}
						position++
						goto l1133
					l1132:
						position, tokenIndex, depth = position1132, tokenIndex1132, depth1132
					}
				l1133:
					if !_rules[rulespOpt]() {
						goto l1125
					}
					if buffer[position] != rune(']') {
						goto l1125
					}
					position++
					depth--
					add(rulePegText, position1127)
				}
				if !_rules[ruleAction67]() {
					goto l1125
				}
				depth--
				add(ruleArrayExpr, position1126)
			}
			return true
		l1125:
			position, tokenIndex, depth = position1125, tokenIndex1125, depth1125
			return false
		},
		/* 88 MapExpr <- <(<('{' spOpt (KeyValuePair (spOpt ',' spOpt KeyValuePair)*)? spOpt '}')> Action68)> */
		func() bool {
			position1134, tokenIndex1134, depth1134 := position, tokenIndex, depth
			{
				position1135 := position
				depth++
				{
					position1136 := position
					depth++
					if buffer[position] != rune('{') {
						goto l1134
					}
					position++
					if !_rules[rulespOpt]() {
						goto l1134
					}
					{
						position1137, tokenIndex1137, depth1137 := position, tokenIndex, depth
						if !_rules[ruleKeyValuePair]() {
							goto l1137
						}
					l1139:
						{
							position1140, tokenIndex1140, depth1140 := position, tokenIndex, depth
							if !_rules[rulespOpt]() {
								goto l1140
							}
							if buffer[position] != rune(',') {
								goto l1140
							}
							position++
							if !_rules[rulespOpt]() {
								goto l1140
							}
							if !_rules[ruleKeyValuePair]() {
								goto l1140
							}
							goto l1139
						l1140:
							position, tokenIndex, depth = position1140, tokenIndex1140, depth1140
						}
						goto l1138
					l1137:
						position, tokenIndex, depth = position1137, tokenIndex1137, depth1137
					}
				l1138:
					if !_rules[rulespOpt]() {
						goto l1134
					}
					if buffer[position] != rune('}') {
						goto l1134
					}
					position++
					depth--
					add(rulePegText, position1136)
				}
				if !_rules[ruleAction68]() {
					goto l1134
				}
				depth--
				add(ruleMapExpr, position1135)
			}
			return true
		l1134:
			position, tokenIndex, depth = position1134, tokenIndex1134, depth1134
			return false
		},
		/* 89 KeyValuePair <- <(<(StringLiteral spOpt ':' spOpt ExpressionOrWildcard)> Action69)> */
		func() bool {
			position1141, tokenIndex1141, depth1141 := position, tokenIndex, depth
			{
				position1142 := position
				depth++
				{
					position1143 := position
					depth++
					if !_rules[ruleStringLiteral]() {
						goto l1141
					}
					if !_rules[rulespOpt]() {
						goto l1141
					}
					if buffer[position] != rune(':') {
						goto l1141
					}
					position++
					if !_rules[rulespOpt]() {
						goto l1141
					}
					if !_rules[ruleExpressionOrWildcard]() {
						goto l1141
					}
					depth--
					add(rulePegText, position1143)
				}
				if !_rules[ruleAction69]() {
					goto l1141
				}
				depth--
				add(ruleKeyValuePair, position1142)
			}
			return true
		l1141:
			position, tokenIndex, depth = position1141, tokenIndex1141, depth1141
			return false
		},
		/* 90 Literal <- <(FloatLiteral / NumericLiteral / StringLiteral)> */
		func() bool {
			position1144, tokenIndex1144, depth1144 := position, tokenIndex, depth
			{
				position1145 := position
				depth++
				{
					position1146, tokenIndex1146, depth1146 := position, tokenIndex, depth
					if !_rules[ruleFloatLiteral]() {
						goto l1147
					}
					goto l1146
				l1147:
					position, tokenIndex, depth = position1146, tokenIndex1146, depth1146
					if !_rules[ruleNumericLiteral]() {
						goto l1148
					}
					goto l1146
				l1148:
					position, tokenIndex, depth = position1146, tokenIndex1146, depth1146
					if !_rules[ruleStringLiteral]() {
						goto l1144
					}
				}
			l1146:
				depth--
				add(ruleLiteral, position1145)
			}
			return true
		l1144:
			position, tokenIndex, depth = position1144, tokenIndex1144, depth1144
			return false
		},
		/* 91 ComparisonOp <- <(Equal / NotEqual / LessOrEqual / Less / GreaterOrEqual / Greater / NotEqual)> */
		func() bool {
			position1149, tokenIndex1149, depth1149 := position, tokenIndex, depth
			{
				position1150 := position
				depth++
				{
					position1151, tokenIndex1151, depth1151 := position, tokenIndex, depth
					if !_rules[ruleEqual]() {
						goto l1152
					}
					goto l1151
				l1152:
					position, tokenIndex, depth = position1151, tokenIndex1151, depth1151
					if !_rules[ruleNotEqual]() {
						goto l1153
					}
					goto l1151
				l1153:
					position, tokenIndex, depth = position1151, tokenIndex1151, depth1151
					if !_rules[ruleLessOrEqual]() {
						goto l1154
					}
					goto l1151
				l1154:
					position, tokenIndex, depth = position1151, tokenIndex1151, depth1151
					if !_rules[ruleLess]() {
						goto l1155
					}
					goto l1151
				l1155:
					position, tokenIndex, depth = position1151, tokenIndex1151, depth1151
					if !_rules[ruleGreaterOrEqual]() {
						goto l1156
					}
					goto l1151
				l1156:
					position, tokenIndex, depth = position1151, tokenIndex1151, depth1151
					if !_rules[ruleGreater]() {
						goto l1157
					}
					goto l1151
				l1157:
					position, tokenIndex, depth = position1151, tokenIndex1151, depth1151
					if !_rules[ruleNotEqual]() {
						goto l1149
					}
				}
			l1151:
				depth--
				add(ruleComparisonOp, position1150)
			}
			return true
		l1149:
			position, tokenIndex, depth = position1149, tokenIndex1149, depth1149
			return false
		},
		/* 92 OtherOp <- <Concat> */
		func() bool {
			position1158, tokenIndex1158, depth1158 := position, tokenIndex, depth
			{
				position1159 := position
				depth++
				if !_rules[ruleConcat]() {
					goto l1158
				}
				depth--
				add(ruleOtherOp, position1159)
			}
			return true
		l1158:
			position, tokenIndex, depth = position1158, tokenIndex1158, depth1158
			return false
		},
		/* 93 IsOp <- <(IsNot / Is)> */
		func() bool {
			position1160, tokenIndex1160, depth1160 := position, tokenIndex, depth
			{
				position1161 := position
				depth++
				{
					position1162, tokenIndex1162, depth1162 := position, tokenIndex, depth
					if !_rules[ruleIsNot]() {
						goto l1163
					}
					goto l1162
				l1163:
					position, tokenIndex, depth = position1162, tokenIndex1162, depth1162
					if !_rules[ruleIs]() {
						goto l1160
					}
				}
			l1162:
				depth--
				add(ruleIsOp, position1161)
			}
			return true
		l1160:
			position, tokenIndex, depth = position1160, tokenIndex1160, depth1160
			return false
		},
		/* 94 PlusMinusOp <- <(Plus / Minus)> */
		func() bool {
			position1164, tokenIndex1164, depth1164 := position, tokenIndex, depth
			{
				position1165 := position
				depth++
				{
					position1166, tokenIndex1166, depth1166 := position, tokenIndex, depth
					if !_rules[rulePlus]() {
						goto l1167
					}
					goto l1166
				l1167:
					position, tokenIndex, depth = position1166, tokenIndex1166, depth1166
					if !_rules[ruleMinus]() {
						goto l1164
					}
				}
			l1166:
				depth--
				add(rulePlusMinusOp, position1165)
			}
			return true
		l1164:
			position, tokenIndex, depth = position1164, tokenIndex1164, depth1164
			return false
		},
		/* 95 MultDivOp <- <(Multiply / Divide / Modulo)> */
		func() bool {
			position1168, tokenIndex1168, depth1168 := position, tokenIndex, depth
			{
				position1169 := position
				depth++
				{
					position1170, tokenIndex1170, depth1170 := position, tokenIndex, depth
					if !_rules[ruleMultiply]() {
						goto l1171
					}
					goto l1170
				l1171:
					position, tokenIndex, depth = position1170, tokenIndex1170, depth1170
					if !_rules[ruleDivide]() {
						goto l1172
					}
					goto l1170
				l1172:
					position, tokenIndex, depth = position1170, tokenIndex1170, depth1170
					if !_rules[ruleModulo]() {
						goto l1168
					}
				}
			l1170:
				depth--
				add(ruleMultDivOp, position1169)
			}
			return true
		l1168:
			position, tokenIndex, depth = position1168, tokenIndex1168, depth1168
			return false
		},
		/* 96 Stream <- <(<ident> Action70)> */
		func() bool {
			position1173, tokenIndex1173, depth1173 := position, tokenIndex, depth
			{
				position1174 := position
				depth++
				{
					position1175 := position
					depth++
					if !_rules[ruleident]() {
						goto l1173
					}
					depth--
					add(rulePegText, position1175)
				}
				if !_rules[ruleAction70]() {
					goto l1173
				}
				depth--
				add(ruleStream, position1174)
			}
			return true
		l1173:
			position, tokenIndex, depth = position1173, tokenIndex1173, depth1173
			return false
		},
		/* 97 RowMeta <- <RowTimestamp> */
		func() bool {
			position1176, tokenIndex1176, depth1176 := position, tokenIndex, depth
			{
				position1177 := position
				depth++
				if !_rules[ruleRowTimestamp]() {
					goto l1176
				}
				depth--
				add(ruleRowMeta, position1177)
			}
			return true
		l1176:
			position, tokenIndex, depth = position1176, tokenIndex1176, depth1176
			return false
		},
		/* 98 RowTimestamp <- <(<((ident ':')? ('t' 's' '(' ')'))> Action71)> */
		func() bool {
			position1178, tokenIndex1178, depth1178 := position, tokenIndex, depth
			{
				position1179 := position
				depth++
				{
					position1180 := position
					depth++
					{
						position1181, tokenIndex1181, depth1181 := position, tokenIndex, depth
						if !_rules[ruleident]() {
							goto l1181
						}
						if buffer[position] != rune(':') {
							goto l1181
						}
						position++
						goto l1182
					l1181:
						position, tokenIndex, depth = position1181, tokenIndex1181, depth1181
					}
				l1182:
					if buffer[position] != rune('t') {
						goto l1178
					}
					position++
					if buffer[position] != rune('s') {
						goto l1178
					}
					position++
					if buffer[position] != rune('(') {
						goto l1178
					}
					position++
					if buffer[position] != rune(')') {
						goto l1178
					}
					position++
					depth--
					add(rulePegText, position1180)
				}
				if !_rules[ruleAction71]() {
					goto l1178
				}
				depth--
				add(ruleRowTimestamp, position1179)
			}
			return true
		l1178:
			position, tokenIndex, depth = position1178, tokenIndex1178, depth1178
			return false
		},
		/* 99 RowValue <- <(<((ident ':' !':')? jsonGetPath)> Action72)> */
		func() bool {
			position1183, tokenIndex1183, depth1183 := position, tokenIndex, depth
			{
				position1184 := position
				depth++
				{
					position1185 := position
					depth++
					{
						position1186, tokenIndex1186, depth1186 := position, tokenIndex, depth
						if !_rules[ruleident]() {
							goto l1186
						}
						if buffer[position] != rune(':') {
							goto l1186
						}
						position++
						{
							position1188, tokenIndex1188, depth1188 := position, tokenIndex, depth
							if buffer[position] != rune(':') {
								goto l1188
							}
							position++
							goto l1186
						l1188:
							position, tokenIndex, depth = position1188, tokenIndex1188, depth1188
						}
						goto l1187
					l1186:
						position, tokenIndex, depth = position1186, tokenIndex1186, depth1186
					}
				l1187:
					if !_rules[rulejsonGetPath]() {
						goto l1183
					}
					depth--
					add(rulePegText, position1185)
				}
				if !_rules[ruleAction72]() {
					goto l1183
				}
				depth--
				add(ruleRowValue, position1184)
			}
			return true
		l1183:
			position, tokenIndex, depth = position1183, tokenIndex1183, depth1183
			return false
		},
		/* 100 NumericLiteral <- <(<('-'? [0-9]+)> Action73)> */
		func() bool {
			position1189, tokenIndex1189, depth1189 := position, tokenIndex, depth
			{
				position1190 := position
				depth++
				{
					position1191 := position
					depth++
					{
						position1192, tokenIndex1192, depth1192 := position, tokenIndex, depth
						if buffer[position] != rune('-') {
							goto l1192
						}
						position++
						goto l1193
					l1192:
						position, tokenIndex, depth = position1192, tokenIndex1192, depth1192
					}
				l1193:
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l1189
					}
					position++
				l1194:
					{
						position1195, tokenIndex1195, depth1195 := position, tokenIndex, depth
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1195
						}
						position++
						goto l1194
					l1195:
						position, tokenIndex, depth = position1195, tokenIndex1195, depth1195
					}
					depth--
					add(rulePegText, position1191)
				}
				if !_rules[ruleAction73]() {
					goto l1189
				}
				depth--
				add(ruleNumericLiteral, position1190)
			}
			return true
		l1189:
			position, tokenIndex, depth = position1189, tokenIndex1189, depth1189
			return false
		},
		/* 101 FloatLiteral <- <(<('-'? [0-9]+ '.' [0-9]+)> Action74)> */
		func() bool {
			position1196, tokenIndex1196, depth1196 := position, tokenIndex, depth
			{
				position1197 := position
				depth++
				{
					position1198 := position
					depth++
					{
						position1199, tokenIndex1199, depth1199 := position, tokenIndex, depth
						if buffer[position] != rune('-') {
							goto l1199
						}
						position++
						goto l1200
					l1199:
						position, tokenIndex, depth = position1199, tokenIndex1199, depth1199
					}
				l1200:
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l1196
					}
					position++
				l1201:
					{
						position1202, tokenIndex1202, depth1202 := position, tokenIndex, depth
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1202
						}
						position++
						goto l1201
					l1202:
						position, tokenIndex, depth = position1202, tokenIndex1202, depth1202
					}
					if buffer[position] != rune('.') {
						goto l1196
					}
					position++
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l1196
					}
					position++
				l1203:
					{
						position1204, tokenIndex1204, depth1204 := position, tokenIndex, depth
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1204
						}
						position++
						goto l1203
					l1204:
						position, tokenIndex, depth = position1204, tokenIndex1204, depth1204
					}
					depth--
					add(rulePegText, position1198)
				}
				if !_rules[ruleAction74]() {
					goto l1196
				}
				depth--
				add(ruleFloatLiteral, position1197)
			}
			return true
		l1196:
			position, tokenIndex, depth = position1196, tokenIndex1196, depth1196
			return false
		},
		/* 102 Function <- <(<ident> Action75)> */
		func() bool {
			position1205, tokenIndex1205, depth1205 := position, tokenIndex, depth
			{
				position1206 := position
				depth++
				{
					position1207 := position
					depth++
					if !_rules[ruleident]() {
						goto l1205
					}
					depth--
					add(rulePegText, position1207)
				}
				if !_rules[ruleAction75]() {
					goto l1205
				}
				depth--
				add(ruleFunction, position1206)
			}
			return true
		l1205:
			position, tokenIndex, depth = position1205, tokenIndex1205, depth1205
			return false
		},
		/* 103 NullLiteral <- <(<(('n' / 'N') ('u' / 'U') ('l' / 'L') ('l' / 'L'))> Action76)> */
		func() bool {
			position1208, tokenIndex1208, depth1208 := position, tokenIndex, depth
			{
				position1209 := position
				depth++
				{
					position1210 := position
					depth++
					{
						position1211, tokenIndex1211, depth1211 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l1212
						}
						position++
						goto l1211
					l1212:
						position, tokenIndex, depth = position1211, tokenIndex1211, depth1211
						if buffer[position] != rune('N') {
							goto l1208
						}
						position++
					}
				l1211:
					{
						position1213, tokenIndex1213, depth1213 := position, tokenIndex, depth
						if buffer[position] != rune('u') {
							goto l1214
						}
						position++
						goto l1213
					l1214:
						position, tokenIndex, depth = position1213, tokenIndex1213, depth1213
						if buffer[position] != rune('U') {
							goto l1208
						}
						position++
					}
				l1213:
					{
						position1215, tokenIndex1215, depth1215 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l1216
						}
						position++
						goto l1215
					l1216:
						position, tokenIndex, depth = position1215, tokenIndex1215, depth1215
						if buffer[position] != rune('L') {
							goto l1208
						}
						position++
					}
				l1215:
					{
						position1217, tokenIndex1217, depth1217 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l1218
						}
						position++
						goto l1217
					l1218:
						position, tokenIndex, depth = position1217, tokenIndex1217, depth1217
						if buffer[position] != rune('L') {
							goto l1208
						}
						position++
					}
				l1217:
					depth--
					add(rulePegText, position1210)
				}
				if !_rules[ruleAction76]() {
					goto l1208
				}
				depth--
				add(ruleNullLiteral, position1209)
			}
			return true
		l1208:
			position, tokenIndex, depth = position1208, tokenIndex1208, depth1208
			return false
		},
		/* 104 BooleanLiteral <- <(TRUE / FALSE)> */
		func() bool {
			position1219, tokenIndex1219, depth1219 := position, tokenIndex, depth
			{
				position1220 := position
				depth++
				{
					position1221, tokenIndex1221, depth1221 := position, tokenIndex, depth
					if !_rules[ruleTRUE]() {
						goto l1222
					}
					goto l1221
				l1222:
					position, tokenIndex, depth = position1221, tokenIndex1221, depth1221
					if !_rules[ruleFALSE]() {
						goto l1219
					}
				}
			l1221:
				depth--
				add(ruleBooleanLiteral, position1220)
			}
			return true
		l1219:
			position, tokenIndex, depth = position1219, tokenIndex1219, depth1219
			return false
		},
		/* 105 TRUE <- <(<(('t' / 'T') ('r' / 'R') ('u' / 'U') ('e' / 'E'))> Action77)> */
		func() bool {
			position1223, tokenIndex1223, depth1223 := position, tokenIndex, depth
			{
				position1224 := position
				depth++
				{
					position1225 := position
					depth++
					{
						position1226, tokenIndex1226, depth1226 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1227
						}
						position++
						goto l1226
					l1227:
						position, tokenIndex, depth = position1226, tokenIndex1226, depth1226
						if buffer[position] != rune('T') {
							goto l1223
						}
						position++
					}
				l1226:
					{
						position1228, tokenIndex1228, depth1228 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l1229
						}
						position++
						goto l1228
					l1229:
						position, tokenIndex, depth = position1228, tokenIndex1228, depth1228
						if buffer[position] != rune('R') {
							goto l1223
						}
						position++
					}
				l1228:
					{
						position1230, tokenIndex1230, depth1230 := position, tokenIndex, depth
						if buffer[position] != rune('u') {
							goto l1231
						}
						position++
						goto l1230
					l1231:
						position, tokenIndex, depth = position1230, tokenIndex1230, depth1230
						if buffer[position] != rune('U') {
							goto l1223
						}
						position++
					}
				l1230:
					{
						position1232, tokenIndex1232, depth1232 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l1233
						}
						position++
						goto l1232
					l1233:
						position, tokenIndex, depth = position1232, tokenIndex1232, depth1232
						if buffer[position] != rune('E') {
							goto l1223
						}
						position++
					}
				l1232:
					depth--
					add(rulePegText, position1225)
				}
				if !_rules[ruleAction77]() {
					goto l1223
				}
				depth--
				add(ruleTRUE, position1224)
			}
			return true
		l1223:
			position, tokenIndex, depth = position1223, tokenIndex1223, depth1223
			return false
		},
		/* 106 FALSE <- <(<(('f' / 'F') ('a' / 'A') ('l' / 'L') ('s' / 'S') ('e' / 'E'))> Action78)> */
		func() bool {
			position1234, tokenIndex1234, depth1234 := position, tokenIndex, depth
			{
				position1235 := position
				depth++
				{
					position1236 := position
					depth++
					{
						position1237, tokenIndex1237, depth1237 := position, tokenIndex, depth
						if buffer[position] != rune('f') {
							goto l1238
						}
						position++
						goto l1237
					l1238:
						position, tokenIndex, depth = position1237, tokenIndex1237, depth1237
						if buffer[position] != rune('F') {
							goto l1234
						}
						position++
					}
				l1237:
					{
						position1239, tokenIndex1239, depth1239 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1240
						}
						position++
						goto l1239
					l1240:
						position, tokenIndex, depth = position1239, tokenIndex1239, depth1239
						if buffer[position] != rune('A') {
							goto l1234
						}
						position++
					}
				l1239:
					{
						position1241, tokenIndex1241, depth1241 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l1242
						}
						position++
						goto l1241
					l1242:
						position, tokenIndex, depth = position1241, tokenIndex1241, depth1241
						if buffer[position] != rune('L') {
							goto l1234
						}
						position++
					}
				l1241:
					{
						position1243, tokenIndex1243, depth1243 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1244
						}
						position++
						goto l1243
					l1244:
						position, tokenIndex, depth = position1243, tokenIndex1243, depth1243
						if buffer[position] != rune('S') {
							goto l1234
						}
						position++
					}
				l1243:
					{
						position1245, tokenIndex1245, depth1245 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l1246
						}
						position++
						goto l1245
					l1246:
						position, tokenIndex, depth = position1245, tokenIndex1245, depth1245
						if buffer[position] != rune('E') {
							goto l1234
						}
						position++
					}
				l1245:
					depth--
					add(rulePegText, position1236)
				}
				if !_rules[ruleAction78]() {
					goto l1234
				}
				depth--
				add(ruleFALSE, position1235)
			}
			return true
		l1234:
			position, tokenIndex, depth = position1234, tokenIndex1234, depth1234
			return false
		},
		/* 107 Wildcard <- <(<((ident ':' !':')? '*')> Action79)> */
		func() bool {
			position1247, tokenIndex1247, depth1247 := position, tokenIndex, depth
			{
				position1248 := position
				depth++
				{
					position1249 := position
					depth++
					{
						position1250, tokenIndex1250, depth1250 := position, tokenIndex, depth
						if !_rules[ruleident]() {
							goto l1250
						}
						if buffer[position] != rune(':') {
							goto l1250
						}
						position++
						{
							position1252, tokenIndex1252, depth1252 := position, tokenIndex, depth
							if buffer[position] != rune(':') {
								goto l1252
							}
							position++
							goto l1250
						l1252:
							position, tokenIndex, depth = position1252, tokenIndex1252, depth1252
						}
						goto l1251
					l1250:
						position, tokenIndex, depth = position1250, tokenIndex1250, depth1250
					}
				l1251:
					if buffer[position] != rune('*') {
						goto l1247
					}
					position++
					depth--
					add(rulePegText, position1249)
				}
				if !_rules[ruleAction79]() {
					goto l1247
				}
				depth--
				add(ruleWildcard, position1248)
			}
			return true
		l1247:
			position, tokenIndex, depth = position1247, tokenIndex1247, depth1247
			return false
		},
		/* 108 StringLiteral <- <(<('\'' (('\'' '\'') / (!'\'' .))* '\'')> Action80)> */
		func() bool {
			position1253, tokenIndex1253, depth1253 := position, tokenIndex, depth
			{
				position1254 := position
				depth++
				{
					position1255 := position
					depth++
					if buffer[position] != rune('\'') {
						goto l1253
					}
					position++
				l1256:
					{
						position1257, tokenIndex1257, depth1257 := position, tokenIndex, depth
						{
							position1258, tokenIndex1258, depth1258 := position, tokenIndex, depth
							if buffer[position] != rune('\'') {
								goto l1259
							}
							position++
							if buffer[position] != rune('\'') {
								goto l1259
							}
							position++
							goto l1258
						l1259:
							position, tokenIndex, depth = position1258, tokenIndex1258, depth1258
							{
								position1260, tokenIndex1260, depth1260 := position, tokenIndex, depth
								if buffer[position] != rune('\'') {
									goto l1260
								}
								position++
								goto l1257
							l1260:
								position, tokenIndex, depth = position1260, tokenIndex1260, depth1260
							}
							if !matchDot() {
								goto l1257
							}
						}
					l1258:
						goto l1256
					l1257:
						position, tokenIndex, depth = position1257, tokenIndex1257, depth1257
					}
					if buffer[position] != rune('\'') {
						goto l1253
					}
					position++
					depth--
					add(rulePegText, position1255)
				}
				if !_rules[ruleAction80]() {
					goto l1253
				}
				depth--
				add(ruleStringLiteral, position1254)
			}
			return true
		l1253:
			position, tokenIndex, depth = position1253, tokenIndex1253, depth1253
			return false
		},
		/* 109 ISTREAM <- <(<(('i' / 'I') ('s' / 'S') ('t' / 'T') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('m' / 'M'))> Action81)> */
		func() bool {
			position1261, tokenIndex1261, depth1261 := position, tokenIndex, depth
			{
				position1262 := position
				depth++
				{
					position1263 := position
					depth++
					{
						position1264, tokenIndex1264, depth1264 := position, tokenIndex, depth
						if buffer[position] != rune('i') {
							goto l1265
						}
						position++
						goto l1264
					l1265:
						position, tokenIndex, depth = position1264, tokenIndex1264, depth1264
						if buffer[position] != rune('I') {
							goto l1261
						}
						position++
					}
				l1264:
					{
						position1266, tokenIndex1266, depth1266 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1267
						}
						position++
						goto l1266
					l1267:
						position, tokenIndex, depth = position1266, tokenIndex1266, depth1266
						if buffer[position] != rune('S') {
							goto l1261
						}
						position++
					}
				l1266:
					{
						position1268, tokenIndex1268, depth1268 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1269
						}
						position++
						goto l1268
					l1269:
						position, tokenIndex, depth = position1268, tokenIndex1268, depth1268
						if buffer[position] != rune('T') {
							goto l1261
						}
						position++
					}
				l1268:
					{
						position1270, tokenIndex1270, depth1270 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l1271
						}
						position++
						goto l1270
					l1271:
						position, tokenIndex, depth = position1270, tokenIndex1270, depth1270
						if buffer[position] != rune('R') {
							goto l1261
						}
						position++
					}
				l1270:
					{
						position1272, tokenIndex1272, depth1272 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l1273
						}
						position++
						goto l1272
					l1273:
						position, tokenIndex, depth = position1272, tokenIndex1272, depth1272
						if buffer[position] != rune('E') {
							goto l1261
						}
						position++
					}
				l1272:
					{
						position1274, tokenIndex1274, depth1274 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1275
						}
						position++
						goto l1274
					l1275:
						position, tokenIndex, depth = position1274, tokenIndex1274, depth1274
						if buffer[position] != rune('A') {
							goto l1261
						}
						position++
					}
				l1274:
					{
						position1276, tokenIndex1276, depth1276 := position, tokenIndex, depth
						if buffer[position] != rune('m') {
							goto l1277
						}
						position++
						goto l1276
					l1277:
						position, tokenIndex, depth = position1276, tokenIndex1276, depth1276
						if buffer[position] != rune('M') {
							goto l1261
						}
						position++
					}
				l1276:
					depth--
					add(rulePegText, position1263)
				}
				if !_rules[ruleAction81]() {
					goto l1261
				}
				depth--
				add(ruleISTREAM, position1262)
			}
			return true
		l1261:
			position, tokenIndex, depth = position1261, tokenIndex1261, depth1261
			return false
		},
		/* 110 DSTREAM <- <(<(('d' / 'D') ('s' / 'S') ('t' / 'T') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('m' / 'M'))> Action82)> */
		func() bool {
			position1278, tokenIndex1278, depth1278 := position, tokenIndex, depth
			{
				position1279 := position
				depth++
				{
					position1280 := position
					depth++
					{
						position1281, tokenIndex1281, depth1281 := position, tokenIndex, depth
						if buffer[position] != rune('d') {
							goto l1282
						}
						position++
						goto l1281
					l1282:
						position, tokenIndex, depth = position1281, tokenIndex1281, depth1281
						if buffer[position] != rune('D') {
							goto l1278
						}
						position++
					}
				l1281:
					{
						position1283, tokenIndex1283, depth1283 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1284
						}
						position++
						goto l1283
					l1284:
						position, tokenIndex, depth = position1283, tokenIndex1283, depth1283
						if buffer[position] != rune('S') {
							goto l1278
						}
						position++
					}
				l1283:
					{
						position1285, tokenIndex1285, depth1285 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1286
						}
						position++
						goto l1285
					l1286:
						position, tokenIndex, depth = position1285, tokenIndex1285, depth1285
						if buffer[position] != rune('T') {
							goto l1278
						}
						position++
					}
				l1285:
					{
						position1287, tokenIndex1287, depth1287 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l1288
						}
						position++
						goto l1287
					l1288:
						position, tokenIndex, depth = position1287, tokenIndex1287, depth1287
						if buffer[position] != rune('R') {
							goto l1278
						}
						position++
					}
				l1287:
					{
						position1289, tokenIndex1289, depth1289 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l1290
						}
						position++
						goto l1289
					l1290:
						position, tokenIndex, depth = position1289, tokenIndex1289, depth1289
						if buffer[position] != rune('E') {
							goto l1278
						}
						position++
					}
				l1289:
					{
						position1291, tokenIndex1291, depth1291 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1292
						}
						position++
						goto l1291
					l1292:
						position, tokenIndex, depth = position1291, tokenIndex1291, depth1291
						if buffer[position] != rune('A') {
							goto l1278
						}
						position++
					}
				l1291:
					{
						position1293, tokenIndex1293, depth1293 := position, tokenIndex, depth
						if buffer[position] != rune('m') {
							goto l1294
						}
						position++
						goto l1293
					l1294:
						position, tokenIndex, depth = position1293, tokenIndex1293, depth1293
						if buffer[position] != rune('M') {
							goto l1278
						}
						position++
					}
				l1293:
					depth--
					add(rulePegText, position1280)
				}
				if !_rules[ruleAction82]() {
					goto l1278
				}
				depth--
				add(ruleDSTREAM, position1279)
			}
			return true
		l1278:
			position, tokenIndex, depth = position1278, tokenIndex1278, depth1278
			return false
		},
		/* 111 RSTREAM <- <(<(('r' / 'R') ('s' / 'S') ('t' / 'T') ('r' / 'R') ('e' / 'E') ('a' / 'A') ('m' / 'M'))> Action83)> */
		func() bool {
			position1295, tokenIndex1295, depth1295 := position, tokenIndex, depth
			{
				position1296 := position
				depth++
				{
					position1297 := position
					depth++
					{
						position1298, tokenIndex1298, depth1298 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l1299
						}
						position++
						goto l1298
					l1299:
						position, tokenIndex, depth = position1298, tokenIndex1298, depth1298
						if buffer[position] != rune('R') {
							goto l1295
						}
						position++
					}
				l1298:
					{
						position1300, tokenIndex1300, depth1300 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1301
						}
						position++
						goto l1300
					l1301:
						position, tokenIndex, depth = position1300, tokenIndex1300, depth1300
						if buffer[position] != rune('S') {
							goto l1295
						}
						position++
					}
				l1300:
					{
						position1302, tokenIndex1302, depth1302 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1303
						}
						position++
						goto l1302
					l1303:
						position, tokenIndex, depth = position1302, tokenIndex1302, depth1302
						if buffer[position] != rune('T') {
							goto l1295
						}
						position++
					}
				l1302:
					{
						position1304, tokenIndex1304, depth1304 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l1305
						}
						position++
						goto l1304
					l1305:
						position, tokenIndex, depth = position1304, tokenIndex1304, depth1304
						if buffer[position] != rune('R') {
							goto l1295
						}
						position++
					}
				l1304:
					{
						position1306, tokenIndex1306, depth1306 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l1307
						}
						position++
						goto l1306
					l1307:
						position, tokenIndex, depth = position1306, tokenIndex1306, depth1306
						if buffer[position] != rune('E') {
							goto l1295
						}
						position++
					}
				l1306:
					{
						position1308, tokenIndex1308, depth1308 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1309
						}
						position++
						goto l1308
					l1309:
						position, tokenIndex, depth = position1308, tokenIndex1308, depth1308
						if buffer[position] != rune('A') {
							goto l1295
						}
						position++
					}
				l1308:
					{
						position1310, tokenIndex1310, depth1310 := position, tokenIndex, depth
						if buffer[position] != rune('m') {
							goto l1311
						}
						position++
						goto l1310
					l1311:
						position, tokenIndex, depth = position1310, tokenIndex1310, depth1310
						if buffer[position] != rune('M') {
							goto l1295
						}
						position++
					}
				l1310:
					depth--
					add(rulePegText, position1297)
				}
				if !_rules[ruleAction83]() {
					goto l1295
				}
				depth--
				add(ruleRSTREAM, position1296)
			}
			return true
		l1295:
			position, tokenIndex, depth = position1295, tokenIndex1295, depth1295
			return false
		},
		/* 112 TUPLES <- <(<(('t' / 'T') ('u' / 'U') ('p' / 'P') ('l' / 'L') ('e' / 'E') ('s' / 'S'))> Action84)> */
		func() bool {
			position1312, tokenIndex1312, depth1312 := position, tokenIndex, depth
			{
				position1313 := position
				depth++
				{
					position1314 := position
					depth++
					{
						position1315, tokenIndex1315, depth1315 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1316
						}
						position++
						goto l1315
					l1316:
						position, tokenIndex, depth = position1315, tokenIndex1315, depth1315
						if buffer[position] != rune('T') {
							goto l1312
						}
						position++
					}
				l1315:
					{
						position1317, tokenIndex1317, depth1317 := position, tokenIndex, depth
						if buffer[position] != rune('u') {
							goto l1318
						}
						position++
						goto l1317
					l1318:
						position, tokenIndex, depth = position1317, tokenIndex1317, depth1317
						if buffer[position] != rune('U') {
							goto l1312
						}
						position++
					}
				l1317:
					{
						position1319, tokenIndex1319, depth1319 := position, tokenIndex, depth
						if buffer[position] != rune('p') {
							goto l1320
						}
						position++
						goto l1319
					l1320:
						position, tokenIndex, depth = position1319, tokenIndex1319, depth1319
						if buffer[position] != rune('P') {
							goto l1312
						}
						position++
					}
				l1319:
					{
						position1321, tokenIndex1321, depth1321 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l1322
						}
						position++
						goto l1321
					l1322:
						position, tokenIndex, depth = position1321, tokenIndex1321, depth1321
						if buffer[position] != rune('L') {
							goto l1312
						}
						position++
					}
				l1321:
					{
						position1323, tokenIndex1323, depth1323 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l1324
						}
						position++
						goto l1323
					l1324:
						position, tokenIndex, depth = position1323, tokenIndex1323, depth1323
						if buffer[position] != rune('E') {
							goto l1312
						}
						position++
					}
				l1323:
					{
						position1325, tokenIndex1325, depth1325 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1326
						}
						position++
						goto l1325
					l1326:
						position, tokenIndex, depth = position1325, tokenIndex1325, depth1325
						if buffer[position] != rune('S') {
							goto l1312
						}
						position++
					}
				l1325:
					depth--
					add(rulePegText, position1314)
				}
				if !_rules[ruleAction84]() {
					goto l1312
				}
				depth--
				add(ruleTUPLES, position1313)
			}
			return true
		l1312:
			position, tokenIndex, depth = position1312, tokenIndex1312, depth1312
			return false
		},
		/* 113 SECONDS <- <(<(('s' / 'S') ('e' / 'E') ('c' / 'C') ('o' / 'O') ('n' / 'N') ('d' / 'D') ('s' / 'S'))> Action85)> */
		func() bool {
			position1327, tokenIndex1327, depth1327 := position, tokenIndex, depth
			{
				position1328 := position
				depth++
				{
					position1329 := position
					depth++
					{
						position1330, tokenIndex1330, depth1330 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1331
						}
						position++
						goto l1330
					l1331:
						position, tokenIndex, depth = position1330, tokenIndex1330, depth1330
						if buffer[position] != rune('S') {
							goto l1327
						}
						position++
					}
				l1330:
					{
						position1332, tokenIndex1332, depth1332 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l1333
						}
						position++
						goto l1332
					l1333:
						position, tokenIndex, depth = position1332, tokenIndex1332, depth1332
						if buffer[position] != rune('E') {
							goto l1327
						}
						position++
					}
				l1332:
					{
						position1334, tokenIndex1334, depth1334 := position, tokenIndex, depth
						if buffer[position] != rune('c') {
							goto l1335
						}
						position++
						goto l1334
					l1335:
						position, tokenIndex, depth = position1334, tokenIndex1334, depth1334
						if buffer[position] != rune('C') {
							goto l1327
						}
						position++
					}
				l1334:
					{
						position1336, tokenIndex1336, depth1336 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l1337
						}
						position++
						goto l1336
					l1337:
						position, tokenIndex, depth = position1336, tokenIndex1336, depth1336
						if buffer[position] != rune('O') {
							goto l1327
						}
						position++
					}
				l1336:
					{
						position1338, tokenIndex1338, depth1338 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l1339
						}
						position++
						goto l1338
					l1339:
						position, tokenIndex, depth = position1338, tokenIndex1338, depth1338
						if buffer[position] != rune('N') {
							goto l1327
						}
						position++
					}
				l1338:
					{
						position1340, tokenIndex1340, depth1340 := position, tokenIndex, depth
						if buffer[position] != rune('d') {
							goto l1341
						}
						position++
						goto l1340
					l1341:
						position, tokenIndex, depth = position1340, tokenIndex1340, depth1340
						if buffer[position] != rune('D') {
							goto l1327
						}
						position++
					}
				l1340:
					{
						position1342, tokenIndex1342, depth1342 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1343
						}
						position++
						goto l1342
					l1343:
						position, tokenIndex, depth = position1342, tokenIndex1342, depth1342
						if buffer[position] != rune('S') {
							goto l1327
						}
						position++
					}
				l1342:
					depth--
					add(rulePegText, position1329)
				}
				if !_rules[ruleAction85]() {
					goto l1327
				}
				depth--
				add(ruleSECONDS, position1328)
			}
			return true
		l1327:
			position, tokenIndex, depth = position1327, tokenIndex1327, depth1327
			return false
		},
		/* 114 MILLISECONDS <- <(<(('m' / 'M') ('i' / 'I') ('l' / 'L') ('l' / 'L') ('i' / 'I') ('s' / 'S') ('e' / 'E') ('c' / 'C') ('o' / 'O') ('n' / 'N') ('d' / 'D') ('s' / 'S'))> Action86)> */
		func() bool {
			position1344, tokenIndex1344, depth1344 := position, tokenIndex, depth
			{
				position1345 := position
				depth++
				{
					position1346 := position
					depth++
					{
						position1347, tokenIndex1347, depth1347 := position, tokenIndex, depth
						if buffer[position] != rune('m') {
							goto l1348
						}
						position++
						goto l1347
					l1348:
						position, tokenIndex, depth = position1347, tokenIndex1347, depth1347
						if buffer[position] != rune('M') {
							goto l1344
						}
						position++
					}
				l1347:
					{
						position1349, tokenIndex1349, depth1349 := position, tokenIndex, depth
						if buffer[position] != rune('i') {
							goto l1350
						}
						position++
						goto l1349
					l1350:
						position, tokenIndex, depth = position1349, tokenIndex1349, depth1349
						if buffer[position] != rune('I') {
							goto l1344
						}
						position++
					}
				l1349:
					{
						position1351, tokenIndex1351, depth1351 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l1352
						}
						position++
						goto l1351
					l1352:
						position, tokenIndex, depth = position1351, tokenIndex1351, depth1351
						if buffer[position] != rune('L') {
							goto l1344
						}
						position++
					}
				l1351:
					{
						position1353, tokenIndex1353, depth1353 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l1354
						}
						position++
						goto l1353
					l1354:
						position, tokenIndex, depth = position1353, tokenIndex1353, depth1353
						if buffer[position] != rune('L') {
							goto l1344
						}
						position++
					}
				l1353:
					{
						position1355, tokenIndex1355, depth1355 := position, tokenIndex, depth
						if buffer[position] != rune('i') {
							goto l1356
						}
						position++
						goto l1355
					l1356:
						position, tokenIndex, depth = position1355, tokenIndex1355, depth1355
						if buffer[position] != rune('I') {
							goto l1344
						}
						position++
					}
				l1355:
					{
						position1357, tokenIndex1357, depth1357 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1358
						}
						position++
						goto l1357
					l1358:
						position, tokenIndex, depth = position1357, tokenIndex1357, depth1357
						if buffer[position] != rune('S') {
							goto l1344
						}
						position++
					}
				l1357:
					{
						position1359, tokenIndex1359, depth1359 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l1360
						}
						position++
						goto l1359
					l1360:
						position, tokenIndex, depth = position1359, tokenIndex1359, depth1359
						if buffer[position] != rune('E') {
							goto l1344
						}
						position++
					}
				l1359:
					{
						position1361, tokenIndex1361, depth1361 := position, tokenIndex, depth
						if buffer[position] != rune('c') {
							goto l1362
						}
						position++
						goto l1361
					l1362:
						position, tokenIndex, depth = position1361, tokenIndex1361, depth1361
						if buffer[position] != rune('C') {
							goto l1344
						}
						position++
					}
				l1361:
					{
						position1363, tokenIndex1363, depth1363 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l1364
						}
						position++
						goto l1363
					l1364:
						position, tokenIndex, depth = position1363, tokenIndex1363, depth1363
						if buffer[position] != rune('O') {
							goto l1344
						}
						position++
					}
				l1363:
					{
						position1365, tokenIndex1365, depth1365 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l1366
						}
						position++
						goto l1365
					l1366:
						position, tokenIndex, depth = position1365, tokenIndex1365, depth1365
						if buffer[position] != rune('N') {
							goto l1344
						}
						position++
					}
				l1365:
					{
						position1367, tokenIndex1367, depth1367 := position, tokenIndex, depth
						if buffer[position] != rune('d') {
							goto l1368
						}
						position++
						goto l1367
					l1368:
						position, tokenIndex, depth = position1367, tokenIndex1367, depth1367
						if buffer[position] != rune('D') {
							goto l1344
						}
						position++
					}
				l1367:
					{
						position1369, tokenIndex1369, depth1369 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1370
						}
						position++
						goto l1369
					l1370:
						position, tokenIndex, depth = position1369, tokenIndex1369, depth1369
						if buffer[position] != rune('S') {
							goto l1344
						}
						position++
					}
				l1369:
					depth--
					add(rulePegText, position1346)
				}
				if !_rules[ruleAction86]() {
					goto l1344
				}
				depth--
				add(ruleMILLISECONDS, position1345)
			}
			return true
		l1344:
			position, tokenIndex, depth = position1344, tokenIndex1344, depth1344
			return false
		},
		/* 115 StreamIdentifier <- <(<ident> Action87)> */
		func() bool {
			position1371, tokenIndex1371, depth1371 := position, tokenIndex, depth
			{
				position1372 := position
				depth++
				{
					position1373 := position
					depth++
					if !_rules[ruleident]() {
						goto l1371
					}
					depth--
					add(rulePegText, position1373)
				}
				if !_rules[ruleAction87]() {
					goto l1371
				}
				depth--
				add(ruleStreamIdentifier, position1372)
			}
			return true
		l1371:
			position, tokenIndex, depth = position1371, tokenIndex1371, depth1371
			return false
		},
		/* 116 SourceSinkType <- <(<ident> Action88)> */
		func() bool {
			position1374, tokenIndex1374, depth1374 := position, tokenIndex, depth
			{
				position1375 := position
				depth++
				{
					position1376 := position
					depth++
					if !_rules[ruleident]() {
						goto l1374
					}
					depth--
					add(rulePegText, position1376)
				}
				if !_rules[ruleAction88]() {
					goto l1374
				}
				depth--
				add(ruleSourceSinkType, position1375)
			}
			return true
		l1374:
			position, tokenIndex, depth = position1374, tokenIndex1374, depth1374
			return false
		},
		/* 117 SourceSinkParamKey <- <(<ident> Action89)> */
		func() bool {
			position1377, tokenIndex1377, depth1377 := position, tokenIndex, depth
			{
				position1378 := position
				depth++
				{
					position1379 := position
					depth++
					if !_rules[ruleident]() {
						goto l1377
					}
					depth--
					add(rulePegText, position1379)
				}
				if !_rules[ruleAction89]() {
					goto l1377
				}
				depth--
				add(ruleSourceSinkParamKey, position1378)
			}
			return true
		l1377:
			position, tokenIndex, depth = position1377, tokenIndex1377, depth1377
			return false
		},
		/* 118 Paused <- <(<(('p' / 'P') ('a' / 'A') ('u' / 'U') ('s' / 'S') ('e' / 'E') ('d' / 'D'))> Action90)> */
		func() bool {
			position1380, tokenIndex1380, depth1380 := position, tokenIndex, depth
			{
				position1381 := position
				depth++
				{
					position1382 := position
					depth++
					{
						position1383, tokenIndex1383, depth1383 := position, tokenIndex, depth
						if buffer[position] != rune('p') {
							goto l1384
						}
						position++
						goto l1383
					l1384:
						position, tokenIndex, depth = position1383, tokenIndex1383, depth1383
						if buffer[position] != rune('P') {
							goto l1380
						}
						position++
					}
				l1383:
					{
						position1385, tokenIndex1385, depth1385 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1386
						}
						position++
						goto l1385
					l1386:
						position, tokenIndex, depth = position1385, tokenIndex1385, depth1385
						if buffer[position] != rune('A') {
							goto l1380
						}
						position++
					}
				l1385:
					{
						position1387, tokenIndex1387, depth1387 := position, tokenIndex, depth
						if buffer[position] != rune('u') {
							goto l1388
						}
						position++
						goto l1387
					l1388:
						position, tokenIndex, depth = position1387, tokenIndex1387, depth1387
						if buffer[position] != rune('U') {
							goto l1380
						}
						position++
					}
				l1387:
					{
						position1389, tokenIndex1389, depth1389 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1390
						}
						position++
						goto l1389
					l1390:
						position, tokenIndex, depth = position1389, tokenIndex1389, depth1389
						if buffer[position] != rune('S') {
							goto l1380
						}
						position++
					}
				l1389:
					{
						position1391, tokenIndex1391, depth1391 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l1392
						}
						position++
						goto l1391
					l1392:
						position, tokenIndex, depth = position1391, tokenIndex1391, depth1391
						if buffer[position] != rune('E') {
							goto l1380
						}
						position++
					}
				l1391:
					{
						position1393, tokenIndex1393, depth1393 := position, tokenIndex, depth
						if buffer[position] != rune('d') {
							goto l1394
						}
						position++
						goto l1393
					l1394:
						position, tokenIndex, depth = position1393, tokenIndex1393, depth1393
						if buffer[position] != rune('D') {
							goto l1380
						}
						position++
					}
				l1393:
					depth--
					add(rulePegText, position1382)
				}
				if !_rules[ruleAction90]() {
					goto l1380
				}
				depth--
				add(rulePaused, position1381)
			}
			return true
		l1380:
			position, tokenIndex, depth = position1380, tokenIndex1380, depth1380
			return false
		},
		/* 119 Unpaused <- <(<(('u' / 'U') ('n' / 'N') ('p' / 'P') ('a' / 'A') ('u' / 'U') ('s' / 'S') ('e' / 'E') ('d' / 'D'))> Action91)> */
		func() bool {
			position1395, tokenIndex1395, depth1395 := position, tokenIndex, depth
			{
				position1396 := position
				depth++
				{
					position1397 := position
					depth++
					{
						position1398, tokenIndex1398, depth1398 := position, tokenIndex, depth
						if buffer[position] != rune('u') {
							goto l1399
						}
						position++
						goto l1398
					l1399:
						position, tokenIndex, depth = position1398, tokenIndex1398, depth1398
						if buffer[position] != rune('U') {
							goto l1395
						}
						position++
					}
				l1398:
					{
						position1400, tokenIndex1400, depth1400 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l1401
						}
						position++
						goto l1400
					l1401:
						position, tokenIndex, depth = position1400, tokenIndex1400, depth1400
						if buffer[position] != rune('N') {
							goto l1395
						}
						position++
					}
				l1400:
					{
						position1402, tokenIndex1402, depth1402 := position, tokenIndex, depth
						if buffer[position] != rune('p') {
							goto l1403
						}
						position++
						goto l1402
					l1403:
						position, tokenIndex, depth = position1402, tokenIndex1402, depth1402
						if buffer[position] != rune('P') {
							goto l1395
						}
						position++
					}
				l1402:
					{
						position1404, tokenIndex1404, depth1404 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1405
						}
						position++
						goto l1404
					l1405:
						position, tokenIndex, depth = position1404, tokenIndex1404, depth1404
						if buffer[position] != rune('A') {
							goto l1395
						}
						position++
					}
				l1404:
					{
						position1406, tokenIndex1406, depth1406 := position, tokenIndex, depth
						if buffer[position] != rune('u') {
							goto l1407
						}
						position++
						goto l1406
					l1407:
						position, tokenIndex, depth = position1406, tokenIndex1406, depth1406
						if buffer[position] != rune('U') {
							goto l1395
						}
						position++
					}
				l1406:
					{
						position1408, tokenIndex1408, depth1408 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1409
						}
						position++
						goto l1408
					l1409:
						position, tokenIndex, depth = position1408, tokenIndex1408, depth1408
						if buffer[position] != rune('S') {
							goto l1395
						}
						position++
					}
				l1408:
					{
						position1410, tokenIndex1410, depth1410 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l1411
						}
						position++
						goto l1410
					l1411:
						position, tokenIndex, depth = position1410, tokenIndex1410, depth1410
						if buffer[position] != rune('E') {
							goto l1395
						}
						position++
					}
				l1410:
					{
						position1412, tokenIndex1412, depth1412 := position, tokenIndex, depth
						if buffer[position] != rune('d') {
							goto l1413
						}
						position++
						goto l1412
					l1413:
						position, tokenIndex, depth = position1412, tokenIndex1412, depth1412
						if buffer[position] != rune('D') {
							goto l1395
						}
						position++
					}
				l1412:
					depth--
					add(rulePegText, position1397)
				}
				if !_rules[ruleAction91]() {
					goto l1395
				}
				depth--
				add(ruleUnpaused, position1396)
			}
			return true
		l1395:
			position, tokenIndex, depth = position1395, tokenIndex1395, depth1395
			return false
		},
		/* 120 Ascending <- <(<(('a' / 'A') ('s' / 'S') ('c' / 'C'))> Action92)> */
		func() bool {
			position1414, tokenIndex1414, depth1414 := position, tokenIndex, depth
			{
				position1415 := position
				depth++
				{
					position1416 := position
					depth++
					{
						position1417, tokenIndex1417, depth1417 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1418
						}
						position++
						goto l1417
					l1418:
						position, tokenIndex, depth = position1417, tokenIndex1417, depth1417
						if buffer[position] != rune('A') {
							goto l1414
						}
						position++
					}
				l1417:
					{
						position1419, tokenIndex1419, depth1419 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1420
						}
						position++
						goto l1419
					l1420:
						position, tokenIndex, depth = position1419, tokenIndex1419, depth1419
						if buffer[position] != rune('S') {
							goto l1414
						}
						position++
					}
				l1419:
					{
						position1421, tokenIndex1421, depth1421 := position, tokenIndex, depth
						if buffer[position] != rune('c') {
							goto l1422
						}
						position++
						goto l1421
					l1422:
						position, tokenIndex, depth = position1421, tokenIndex1421, depth1421
						if buffer[position] != rune('C') {
							goto l1414
						}
						position++
					}
				l1421:
					depth--
					add(rulePegText, position1416)
				}
				if !_rules[ruleAction92]() {
					goto l1414
				}
				depth--
				add(ruleAscending, position1415)
			}
			return true
		l1414:
			position, tokenIndex, depth = position1414, tokenIndex1414, depth1414
			return false
		},
		/* 121 Descending <- <(<(('d' / 'D') ('e' / 'E') ('s' / 'S') ('c' / 'C'))> Action93)> */
		func() bool {
			position1423, tokenIndex1423, depth1423 := position, tokenIndex, depth
			{
				position1424 := position
				depth++
				{
					position1425 := position
					depth++
					{
						position1426, tokenIndex1426, depth1426 := position, tokenIndex, depth
						if buffer[position] != rune('d') {
							goto l1427
						}
						position++
						goto l1426
					l1427:
						position, tokenIndex, depth = position1426, tokenIndex1426, depth1426
						if buffer[position] != rune('D') {
							goto l1423
						}
						position++
					}
				l1426:
					{
						position1428, tokenIndex1428, depth1428 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l1429
						}
						position++
						goto l1428
					l1429:
						position, tokenIndex, depth = position1428, tokenIndex1428, depth1428
						if buffer[position] != rune('E') {
							goto l1423
						}
						position++
					}
				l1428:
					{
						position1430, tokenIndex1430, depth1430 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1431
						}
						position++
						goto l1430
					l1431:
						position, tokenIndex, depth = position1430, tokenIndex1430, depth1430
						if buffer[position] != rune('S') {
							goto l1423
						}
						position++
					}
				l1430:
					{
						position1432, tokenIndex1432, depth1432 := position, tokenIndex, depth
						if buffer[position] != rune('c') {
							goto l1433
						}
						position++
						goto l1432
					l1433:
						position, tokenIndex, depth = position1432, tokenIndex1432, depth1432
						if buffer[position] != rune('C') {
							goto l1423
						}
						position++
					}
				l1432:
					depth--
					add(rulePegText, position1425)
				}
				if !_rules[ruleAction93]() {
					goto l1423
				}
				depth--
				add(ruleDescending, position1424)
			}
			return true
		l1423:
			position, tokenIndex, depth = position1423, tokenIndex1423, depth1423
			return false
		},
		/* 122 Type <- <(Bool / Int / Float / String / Blob / Timestamp / Array / Map)> */
		func() bool {
			position1434, tokenIndex1434, depth1434 := position, tokenIndex, depth
			{
				position1435 := position
				depth++
				{
					position1436, tokenIndex1436, depth1436 := position, tokenIndex, depth
					if !_rules[ruleBool]() {
						goto l1437
					}
					goto l1436
				l1437:
					position, tokenIndex, depth = position1436, tokenIndex1436, depth1436
					if !_rules[ruleInt]() {
						goto l1438
					}
					goto l1436
				l1438:
					position, tokenIndex, depth = position1436, tokenIndex1436, depth1436
					if !_rules[ruleFloat]() {
						goto l1439
					}
					goto l1436
				l1439:
					position, tokenIndex, depth = position1436, tokenIndex1436, depth1436
					if !_rules[ruleString]() {
						goto l1440
					}
					goto l1436
				l1440:
					position, tokenIndex, depth = position1436, tokenIndex1436, depth1436
					if !_rules[ruleBlob]() {
						goto l1441
					}
					goto l1436
				l1441:
					position, tokenIndex, depth = position1436, tokenIndex1436, depth1436
					if !_rules[ruleTimestamp]() {
						goto l1442
					}
					goto l1436
				l1442:
					position, tokenIndex, depth = position1436, tokenIndex1436, depth1436
					if !_rules[ruleArray]() {
						goto l1443
					}
					goto l1436
				l1443:
					position, tokenIndex, depth = position1436, tokenIndex1436, depth1436
					if !_rules[ruleMap]() {
						goto l1434
					}
				}
			l1436:
				depth--
				add(ruleType, position1435)
			}
			return true
		l1434:
			position, tokenIndex, depth = position1434, tokenIndex1434, depth1434
			return false
		},
		/* 123 Bool <- <(<(('b' / 'B') ('o' / 'O') ('o' / 'O') ('l' / 'L'))> Action94)> */
		func() bool {
			position1444, tokenIndex1444, depth1444 := position, tokenIndex, depth
			{
				position1445 := position
				depth++
				{
					position1446 := position
					depth++
					{
						position1447, tokenIndex1447, depth1447 := position, tokenIndex, depth
						if buffer[position] != rune('b') {
							goto l1448
						}
						position++
						goto l1447
					l1448:
						position, tokenIndex, depth = position1447, tokenIndex1447, depth1447
						if buffer[position] != rune('B') {
							goto l1444
						}
						position++
					}
				l1447:
					{
						position1449, tokenIndex1449, depth1449 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l1450
						}
						position++
						goto l1449
					l1450:
						position, tokenIndex, depth = position1449, tokenIndex1449, depth1449
						if buffer[position] != rune('O') {
							goto l1444
						}
						position++
					}
				l1449:
					{
						position1451, tokenIndex1451, depth1451 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l1452
						}
						position++
						goto l1451
					l1452:
						position, tokenIndex, depth = position1451, tokenIndex1451, depth1451
						if buffer[position] != rune('O') {
							goto l1444
						}
						position++
					}
				l1451:
					{
						position1453, tokenIndex1453, depth1453 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l1454
						}
						position++
						goto l1453
					l1454:
						position, tokenIndex, depth = position1453, tokenIndex1453, depth1453
						if buffer[position] != rune('L') {
							goto l1444
						}
						position++
					}
				l1453:
					depth--
					add(rulePegText, position1446)
				}
				if !_rules[ruleAction94]() {
					goto l1444
				}
				depth--
				add(ruleBool, position1445)
			}
			return true
		l1444:
			position, tokenIndex, depth = position1444, tokenIndex1444, depth1444
			return false
		},
		/* 124 Int <- <(<(('i' / 'I') ('n' / 'N') ('t' / 'T'))> Action95)> */
		func() bool {
			position1455, tokenIndex1455, depth1455 := position, tokenIndex, depth
			{
				position1456 := position
				depth++
				{
					position1457 := position
					depth++
					{
						position1458, tokenIndex1458, depth1458 := position, tokenIndex, depth
						if buffer[position] != rune('i') {
							goto l1459
						}
						position++
						goto l1458
					l1459:
						position, tokenIndex, depth = position1458, tokenIndex1458, depth1458
						if buffer[position] != rune('I') {
							goto l1455
						}
						position++
					}
				l1458:
					{
						position1460, tokenIndex1460, depth1460 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l1461
						}
						position++
						goto l1460
					l1461:
						position, tokenIndex, depth = position1460, tokenIndex1460, depth1460
						if buffer[position] != rune('N') {
							goto l1455
						}
						position++
					}
				l1460:
					{
						position1462, tokenIndex1462, depth1462 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1463
						}
						position++
						goto l1462
					l1463:
						position, tokenIndex, depth = position1462, tokenIndex1462, depth1462
						if buffer[position] != rune('T') {
							goto l1455
						}
						position++
					}
				l1462:
					depth--
					add(rulePegText, position1457)
				}
				if !_rules[ruleAction95]() {
					goto l1455
				}
				depth--
				add(ruleInt, position1456)
			}
			return true
		l1455:
			position, tokenIndex, depth = position1455, tokenIndex1455, depth1455
			return false
		},
		/* 125 Float <- <(<(('f' / 'F') ('l' / 'L') ('o' / 'O') ('a' / 'A') ('t' / 'T'))> Action96)> */
		func() bool {
			position1464, tokenIndex1464, depth1464 := position, tokenIndex, depth
			{
				position1465 := position
				depth++
				{
					position1466 := position
					depth++
					{
						position1467, tokenIndex1467, depth1467 := position, tokenIndex, depth
						if buffer[position] != rune('f') {
							goto l1468
						}
						position++
						goto l1467
					l1468:
						position, tokenIndex, depth = position1467, tokenIndex1467, depth1467
						if buffer[position] != rune('F') {
							goto l1464
						}
						position++
					}
				l1467:
					{
						position1469, tokenIndex1469, depth1469 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l1470
						}
						position++
						goto l1469
					l1470:
						position, tokenIndex, depth = position1469, tokenIndex1469, depth1469
						if buffer[position] != rune('L') {
							goto l1464
						}
						position++
					}
				l1469:
					{
						position1471, tokenIndex1471, depth1471 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l1472
						}
						position++
						goto l1471
					l1472:
						position, tokenIndex, depth = position1471, tokenIndex1471, depth1471
						if buffer[position] != rune('O') {
							goto l1464
						}
						position++
					}
				l1471:
					{
						position1473, tokenIndex1473, depth1473 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1474
						}
						position++
						goto l1473
					l1474:
						position, tokenIndex, depth = position1473, tokenIndex1473, depth1473
						if buffer[position] != rune('A') {
							goto l1464
						}
						position++
					}
				l1473:
					{
						position1475, tokenIndex1475, depth1475 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1476
						}
						position++
						goto l1475
					l1476:
						position, tokenIndex, depth = position1475, tokenIndex1475, depth1475
						if buffer[position] != rune('T') {
							goto l1464
						}
						position++
					}
				l1475:
					depth--
					add(rulePegText, position1466)
				}
				if !_rules[ruleAction96]() {
					goto l1464
				}
				depth--
				add(ruleFloat, position1465)
			}
			return true
		l1464:
			position, tokenIndex, depth = position1464, tokenIndex1464, depth1464
			return false
		},
		/* 126 String <- <(<(('s' / 'S') ('t' / 'T') ('r' / 'R') ('i' / 'I') ('n' / 'N') ('g' / 'G'))> Action97)> */
		func() bool {
			position1477, tokenIndex1477, depth1477 := position, tokenIndex, depth
			{
				position1478 := position
				depth++
				{
					position1479 := position
					depth++
					{
						position1480, tokenIndex1480, depth1480 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1481
						}
						position++
						goto l1480
					l1481:
						position, tokenIndex, depth = position1480, tokenIndex1480, depth1480
						if buffer[position] != rune('S') {
							goto l1477
						}
						position++
					}
				l1480:
					{
						position1482, tokenIndex1482, depth1482 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1483
						}
						position++
						goto l1482
					l1483:
						position, tokenIndex, depth = position1482, tokenIndex1482, depth1482
						if buffer[position] != rune('T') {
							goto l1477
						}
						position++
					}
				l1482:
					{
						position1484, tokenIndex1484, depth1484 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l1485
						}
						position++
						goto l1484
					l1485:
						position, tokenIndex, depth = position1484, tokenIndex1484, depth1484
						if buffer[position] != rune('R') {
							goto l1477
						}
						position++
					}
				l1484:
					{
						position1486, tokenIndex1486, depth1486 := position, tokenIndex, depth
						if buffer[position] != rune('i') {
							goto l1487
						}
						position++
						goto l1486
					l1487:
						position, tokenIndex, depth = position1486, tokenIndex1486, depth1486
						if buffer[position] != rune('I') {
							goto l1477
						}
						position++
					}
				l1486:
					{
						position1488, tokenIndex1488, depth1488 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l1489
						}
						position++
						goto l1488
					l1489:
						position, tokenIndex, depth = position1488, tokenIndex1488, depth1488
						if buffer[position] != rune('N') {
							goto l1477
						}
						position++
					}
				l1488:
					{
						position1490, tokenIndex1490, depth1490 := position, tokenIndex, depth
						if buffer[position] != rune('g') {
							goto l1491
						}
						position++
						goto l1490
					l1491:
						position, tokenIndex, depth = position1490, tokenIndex1490, depth1490
						if buffer[position] != rune('G') {
							goto l1477
						}
						position++
					}
				l1490:
					depth--
					add(rulePegText, position1479)
				}
				if !_rules[ruleAction97]() {
					goto l1477
				}
				depth--
				add(ruleString, position1478)
			}
			return true
		l1477:
			position, tokenIndex, depth = position1477, tokenIndex1477, depth1477
			return false
		},
		/* 127 Blob <- <(<(('b' / 'B') ('l' / 'L') ('o' / 'O') ('b' / 'B'))> Action98)> */
		func() bool {
			position1492, tokenIndex1492, depth1492 := position, tokenIndex, depth
			{
				position1493 := position
				depth++
				{
					position1494 := position
					depth++
					{
						position1495, tokenIndex1495, depth1495 := position, tokenIndex, depth
						if buffer[position] != rune('b') {
							goto l1496
						}
						position++
						goto l1495
					l1496:
						position, tokenIndex, depth = position1495, tokenIndex1495, depth1495
						if buffer[position] != rune('B') {
							goto l1492
						}
						position++
					}
				l1495:
					{
						position1497, tokenIndex1497, depth1497 := position, tokenIndex, depth
						if buffer[position] != rune('l') {
							goto l1498
						}
						position++
						goto l1497
					l1498:
						position, tokenIndex, depth = position1497, tokenIndex1497, depth1497
						if buffer[position] != rune('L') {
							goto l1492
						}
						position++
					}
				l1497:
					{
						position1499, tokenIndex1499, depth1499 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l1500
						}
						position++
						goto l1499
					l1500:
						position, tokenIndex, depth = position1499, tokenIndex1499, depth1499
						if buffer[position] != rune('O') {
							goto l1492
						}
						position++
					}
				l1499:
					{
						position1501, tokenIndex1501, depth1501 := position, tokenIndex, depth
						if buffer[position] != rune('b') {
							goto l1502
						}
						position++
						goto l1501
					l1502:
						position, tokenIndex, depth = position1501, tokenIndex1501, depth1501
						if buffer[position] != rune('B') {
							goto l1492
						}
						position++
					}
				l1501:
					depth--
					add(rulePegText, position1494)
				}
				if !_rules[ruleAction98]() {
					goto l1492
				}
				depth--
				add(ruleBlob, position1493)
			}
			return true
		l1492:
			position, tokenIndex, depth = position1492, tokenIndex1492, depth1492
			return false
		},
		/* 128 Timestamp <- <(<(('t' / 'T') ('i' / 'I') ('m' / 'M') ('e' / 'E') ('s' / 'S') ('t' / 'T') ('a' / 'A') ('m' / 'M') ('p' / 'P'))> Action99)> */
		func() bool {
			position1503, tokenIndex1503, depth1503 := position, tokenIndex, depth
			{
				position1504 := position
				depth++
				{
					position1505 := position
					depth++
					{
						position1506, tokenIndex1506, depth1506 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1507
						}
						position++
						goto l1506
					l1507:
						position, tokenIndex, depth = position1506, tokenIndex1506, depth1506
						if buffer[position] != rune('T') {
							goto l1503
						}
						position++
					}
				l1506:
					{
						position1508, tokenIndex1508, depth1508 := position, tokenIndex, depth
						if buffer[position] != rune('i') {
							goto l1509
						}
						position++
						goto l1508
					l1509:
						position, tokenIndex, depth = position1508, tokenIndex1508, depth1508
						if buffer[position] != rune('I') {
							goto l1503
						}
						position++
					}
				l1508:
					{
						position1510, tokenIndex1510, depth1510 := position, tokenIndex, depth
						if buffer[position] != rune('m') {
							goto l1511
						}
						position++
						goto l1510
					l1511:
						position, tokenIndex, depth = position1510, tokenIndex1510, depth1510
						if buffer[position] != rune('M') {
							goto l1503
						}
						position++
					}
				l1510:
					{
						position1512, tokenIndex1512, depth1512 := position, tokenIndex, depth
						if buffer[position] != rune('e') {
							goto l1513
						}
						position++
						goto l1512
					l1513:
						position, tokenIndex, depth = position1512, tokenIndex1512, depth1512
						if buffer[position] != rune('E') {
							goto l1503
						}
						position++
					}
				l1512:
					{
						position1514, tokenIndex1514, depth1514 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1515
						}
						position++
						goto l1514
					l1515:
						position, tokenIndex, depth = position1514, tokenIndex1514, depth1514
						if buffer[position] != rune('S') {
							goto l1503
						}
						position++
					}
				l1514:
					{
						position1516, tokenIndex1516, depth1516 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1517
						}
						position++
						goto l1516
					l1517:
						position, tokenIndex, depth = position1516, tokenIndex1516, depth1516
						if buffer[position] != rune('T') {
							goto l1503
						}
						position++
					}
				l1516:
					{
						position1518, tokenIndex1518, depth1518 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1519
						}
						position++
						goto l1518
					l1519:
						position, tokenIndex, depth = position1518, tokenIndex1518, depth1518
						if buffer[position] != rune('A') {
							goto l1503
						}
						position++
					}
				l1518:
					{
						position1520, tokenIndex1520, depth1520 := position, tokenIndex, depth
						if buffer[position] != rune('m') {
							goto l1521
						}
						position++
						goto l1520
					l1521:
						position, tokenIndex, depth = position1520, tokenIndex1520, depth1520
						if buffer[position] != rune('M') {
							goto l1503
						}
						position++
					}
				l1520:
					{
						position1522, tokenIndex1522, depth1522 := position, tokenIndex, depth
						if buffer[position] != rune('p') {
							goto l1523
						}
						position++
						goto l1522
					l1523:
						position, tokenIndex, depth = position1522, tokenIndex1522, depth1522
						if buffer[position] != rune('P') {
							goto l1503
						}
						position++
					}
				l1522:
					depth--
					add(rulePegText, position1505)
				}
				if !_rules[ruleAction99]() {
					goto l1503
				}
				depth--
				add(ruleTimestamp, position1504)
			}
			return true
		l1503:
			position, tokenIndex, depth = position1503, tokenIndex1503, depth1503
			return false
		},
		/* 129 Array <- <(<(('a' / 'A') ('r' / 'R') ('r' / 'R') ('a' / 'A') ('y' / 'Y'))> Action100)> */
		func() bool {
			position1524, tokenIndex1524, depth1524 := position, tokenIndex, depth
			{
				position1525 := position
				depth++
				{
					position1526 := position
					depth++
					{
						position1527, tokenIndex1527, depth1527 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1528
						}
						position++
						goto l1527
					l1528:
						position, tokenIndex, depth = position1527, tokenIndex1527, depth1527
						if buffer[position] != rune('A') {
							goto l1524
						}
						position++
					}
				l1527:
					{
						position1529, tokenIndex1529, depth1529 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l1530
						}
						position++
						goto l1529
					l1530:
						position, tokenIndex, depth = position1529, tokenIndex1529, depth1529
						if buffer[position] != rune('R') {
							goto l1524
						}
						position++
					}
				l1529:
					{
						position1531, tokenIndex1531, depth1531 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l1532
						}
						position++
						goto l1531
					l1532:
						position, tokenIndex, depth = position1531, tokenIndex1531, depth1531
						if buffer[position] != rune('R') {
							goto l1524
						}
						position++
					}
				l1531:
					{
						position1533, tokenIndex1533, depth1533 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1534
						}
						position++
						goto l1533
					l1534:
						position, tokenIndex, depth = position1533, tokenIndex1533, depth1533
						if buffer[position] != rune('A') {
							goto l1524
						}
						position++
					}
				l1533:
					{
						position1535, tokenIndex1535, depth1535 := position, tokenIndex, depth
						if buffer[position] != rune('y') {
							goto l1536
						}
						position++
						goto l1535
					l1536:
						position, tokenIndex, depth = position1535, tokenIndex1535, depth1535
						if buffer[position] != rune('Y') {
							goto l1524
						}
						position++
					}
				l1535:
					depth--
					add(rulePegText, position1526)
				}
				if !_rules[ruleAction100]() {
					goto l1524
				}
				depth--
				add(ruleArray, position1525)
			}
			return true
		l1524:
			position, tokenIndex, depth = position1524, tokenIndex1524, depth1524
			return false
		},
		/* 130 Map <- <(<(('m' / 'M') ('a' / 'A') ('p' / 'P'))> Action101)> */
		func() bool {
			position1537, tokenIndex1537, depth1537 := position, tokenIndex, depth
			{
				position1538 := position
				depth++
				{
					position1539 := position
					depth++
					{
						position1540, tokenIndex1540, depth1540 := position, tokenIndex, depth
						if buffer[position] != rune('m') {
							goto l1541
						}
						position++
						goto l1540
					l1541:
						position, tokenIndex, depth = position1540, tokenIndex1540, depth1540
						if buffer[position] != rune('M') {
							goto l1537
						}
						position++
					}
				l1540:
					{
						position1542, tokenIndex1542, depth1542 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1543
						}
						position++
						goto l1542
					l1543:
						position, tokenIndex, depth = position1542, tokenIndex1542, depth1542
						if buffer[position] != rune('A') {
							goto l1537
						}
						position++
					}
				l1542:
					{
						position1544, tokenIndex1544, depth1544 := position, tokenIndex, depth
						if buffer[position] != rune('p') {
							goto l1545
						}
						position++
						goto l1544
					l1545:
						position, tokenIndex, depth = position1544, tokenIndex1544, depth1544
						if buffer[position] != rune('P') {
							goto l1537
						}
						position++
					}
				l1544:
					depth--
					add(rulePegText, position1539)
				}
				if !_rules[ruleAction101]() {
					goto l1537
				}
				depth--
				add(ruleMap, position1538)
			}
			return true
		l1537:
			position, tokenIndex, depth = position1537, tokenIndex1537, depth1537
			return false
		},
		/* 131 Or <- <(<(('o' / 'O') ('r' / 'R'))> Action102)> */
		func() bool {
			position1546, tokenIndex1546, depth1546 := position, tokenIndex, depth
			{
				position1547 := position
				depth++
				{
					position1548 := position
					depth++
					{
						position1549, tokenIndex1549, depth1549 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l1550
						}
						position++
						goto l1549
					l1550:
						position, tokenIndex, depth = position1549, tokenIndex1549, depth1549
						if buffer[position] != rune('O') {
							goto l1546
						}
						position++
					}
				l1549:
					{
						position1551, tokenIndex1551, depth1551 := position, tokenIndex, depth
						if buffer[position] != rune('r') {
							goto l1552
						}
						position++
						goto l1551
					l1552:
						position, tokenIndex, depth = position1551, tokenIndex1551, depth1551
						if buffer[position] != rune('R') {
							goto l1546
						}
						position++
					}
				l1551:
					depth--
					add(rulePegText, position1548)
				}
				if !_rules[ruleAction102]() {
					goto l1546
				}
				depth--
				add(ruleOr, position1547)
			}
			return true
		l1546:
			position, tokenIndex, depth = position1546, tokenIndex1546, depth1546
			return false
		},
		/* 132 And <- <(<(('a' / 'A') ('n' / 'N') ('d' / 'D'))> Action103)> */
		func() bool {
			position1553, tokenIndex1553, depth1553 := position, tokenIndex, depth
			{
				position1554 := position
				depth++
				{
					position1555 := position
					depth++
					{
						position1556, tokenIndex1556, depth1556 := position, tokenIndex, depth
						if buffer[position] != rune('a') {
							goto l1557
						}
						position++
						goto l1556
					l1557:
						position, tokenIndex, depth = position1556, tokenIndex1556, depth1556
						if buffer[position] != rune('A') {
							goto l1553
						}
						position++
					}
				l1556:
					{
						position1558, tokenIndex1558, depth1558 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l1559
						}
						position++
						goto l1558
					l1559:
						position, tokenIndex, depth = position1558, tokenIndex1558, depth1558
						if buffer[position] != rune('N') {
							goto l1553
						}
						position++
					}
				l1558:
					{
						position1560, tokenIndex1560, depth1560 := position, tokenIndex, depth
						if buffer[position] != rune('d') {
							goto l1561
						}
						position++
						goto l1560
					l1561:
						position, tokenIndex, depth = position1560, tokenIndex1560, depth1560
						if buffer[position] != rune('D') {
							goto l1553
						}
						position++
					}
				l1560:
					depth--
					add(rulePegText, position1555)
				}
				if !_rules[ruleAction103]() {
					goto l1553
				}
				depth--
				add(ruleAnd, position1554)
			}
			return true
		l1553:
			position, tokenIndex, depth = position1553, tokenIndex1553, depth1553
			return false
		},
		/* 133 Not <- <(<(('n' / 'N') ('o' / 'O') ('t' / 'T'))> Action104)> */
		func() bool {
			position1562, tokenIndex1562, depth1562 := position, tokenIndex, depth
			{
				position1563 := position
				depth++
				{
					position1564 := position
					depth++
					{
						position1565, tokenIndex1565, depth1565 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l1566
						}
						position++
						goto l1565
					l1566:
						position, tokenIndex, depth = position1565, tokenIndex1565, depth1565
						if buffer[position] != rune('N') {
							goto l1562
						}
						position++
					}
				l1565:
					{
						position1567, tokenIndex1567, depth1567 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l1568
						}
						position++
						goto l1567
					l1568:
						position, tokenIndex, depth = position1567, tokenIndex1567, depth1567
						if buffer[position] != rune('O') {
							goto l1562
						}
						position++
					}
				l1567:
					{
						position1569, tokenIndex1569, depth1569 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1570
						}
						position++
						goto l1569
					l1570:
						position, tokenIndex, depth = position1569, tokenIndex1569, depth1569
						if buffer[position] != rune('T') {
							goto l1562
						}
						position++
					}
				l1569:
					depth--
					add(rulePegText, position1564)
				}
				if !_rules[ruleAction104]() {
					goto l1562
				}
				depth--
				add(ruleNot, position1563)
			}
			return true
		l1562:
			position, tokenIndex, depth = position1562, tokenIndex1562, depth1562
			return false
		},
		/* 134 Equal <- <(<'='> Action105)> */
		func() bool {
			position1571, tokenIndex1571, depth1571 := position, tokenIndex, depth
			{
				position1572 := position
				depth++
				{
					position1573 := position
					depth++
					if buffer[position] != rune('=') {
						goto l1571
					}
					position++
					depth--
					add(rulePegText, position1573)
				}
				if !_rules[ruleAction105]() {
					goto l1571
				}
				depth--
				add(ruleEqual, position1572)
			}
			return true
		l1571:
			position, tokenIndex, depth = position1571, tokenIndex1571, depth1571
			return false
		},
		/* 135 Less <- <(<'<'> Action106)> */
		func() bool {
			position1574, tokenIndex1574, depth1574 := position, tokenIndex, depth
			{
				position1575 := position
				depth++
				{
					position1576 := position
					depth++
					if buffer[position] != rune('<') {
						goto l1574
					}
					position++
					depth--
					add(rulePegText, position1576)
				}
				if !_rules[ruleAction106]() {
					goto l1574
				}
				depth--
				add(ruleLess, position1575)
			}
			return true
		l1574:
			position, tokenIndex, depth = position1574, tokenIndex1574, depth1574
			return false
		},
		/* 136 LessOrEqual <- <(<('<' '=')> Action107)> */
		func() bool {
			position1577, tokenIndex1577, depth1577 := position, tokenIndex, depth
			{
				position1578 := position
				depth++
				{
					position1579 := position
					depth++
					if buffer[position] != rune('<') {
						goto l1577
					}
					position++
					if buffer[position] != rune('=') {
						goto l1577
					}
					position++
					depth--
					add(rulePegText, position1579)
				}
				if !_rules[ruleAction107]() {
					goto l1577
				}
				depth--
				add(ruleLessOrEqual, position1578)
			}
			return true
		l1577:
			position, tokenIndex, depth = position1577, tokenIndex1577, depth1577
			return false
		},
		/* 137 Greater <- <(<'>'> Action108)> */
		func() bool {
			position1580, tokenIndex1580, depth1580 := position, tokenIndex, depth
			{
				position1581 := position
				depth++
				{
					position1582 := position
					depth++
					if buffer[position] != rune('>') {
						goto l1580
					}
					position++
					depth--
					add(rulePegText, position1582)
				}
				if !_rules[ruleAction108]() {
					goto l1580
				}
				depth--
				add(ruleGreater, position1581)
			}
			return true
		l1580:
			position, tokenIndex, depth = position1580, tokenIndex1580, depth1580
			return false
		},
		/* 138 GreaterOrEqual <- <(<('>' '=')> Action109)> */
		func() bool {
			position1583, tokenIndex1583, depth1583 := position, tokenIndex, depth
			{
				position1584 := position
				depth++
				{
					position1585 := position
					depth++
					if buffer[position] != rune('>') {
						goto l1583
					}
					position++
					if buffer[position] != rune('=') {
						goto l1583
					}
					position++
					depth--
					add(rulePegText, position1585)
				}
				if !_rules[ruleAction109]() {
					goto l1583
				}
				depth--
				add(ruleGreaterOrEqual, position1584)
			}
			return true
		l1583:
			position, tokenIndex, depth = position1583, tokenIndex1583, depth1583
			return false
		},
		/* 139 NotEqual <- <(<(('!' '=') / ('<' '>'))> Action110)> */
		func() bool {
			position1586, tokenIndex1586, depth1586 := position, tokenIndex, depth
			{
				position1587 := position
				depth++
				{
					position1588 := position
					depth++
					{
						position1589, tokenIndex1589, depth1589 := position, tokenIndex, depth
						if buffer[position] != rune('!') {
							goto l1590
						}
						position++
						if buffer[position] != rune('=') {
							goto l1590
						}
						position++
						goto l1589
					l1590:
						position, tokenIndex, depth = position1589, tokenIndex1589, depth1589
						if buffer[position] != rune('<') {
							goto l1586
						}
						position++
						if buffer[position] != rune('>') {
							goto l1586
						}
						position++
					}
				l1589:
					depth--
					add(rulePegText, position1588)
				}
				if !_rules[ruleAction110]() {
					goto l1586
				}
				depth--
				add(ruleNotEqual, position1587)
			}
			return true
		l1586:
			position, tokenIndex, depth = position1586, tokenIndex1586, depth1586
			return false
		},
		/* 140 Concat <- <(<('|' '|')> Action111)> */
		func() bool {
			position1591, tokenIndex1591, depth1591 := position, tokenIndex, depth
			{
				position1592 := position
				depth++
				{
					position1593 := position
					depth++
					if buffer[position] != rune('|') {
						goto l1591
					}
					position++
					if buffer[position] != rune('|') {
						goto l1591
					}
					position++
					depth--
					add(rulePegText, position1593)
				}
				if !_rules[ruleAction111]() {
					goto l1591
				}
				depth--
				add(ruleConcat, position1592)
			}
			return true
		l1591:
			position, tokenIndex, depth = position1591, tokenIndex1591, depth1591
			return false
		},
		/* 141 Is <- <(<(('i' / 'I') ('s' / 'S'))> Action112)> */
		func() bool {
			position1594, tokenIndex1594, depth1594 := position, tokenIndex, depth
			{
				position1595 := position
				depth++
				{
					position1596 := position
					depth++
					{
						position1597, tokenIndex1597, depth1597 := position, tokenIndex, depth
						if buffer[position] != rune('i') {
							goto l1598
						}
						position++
						goto l1597
					l1598:
						position, tokenIndex, depth = position1597, tokenIndex1597, depth1597
						if buffer[position] != rune('I') {
							goto l1594
						}
						position++
					}
				l1597:
					{
						position1599, tokenIndex1599, depth1599 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1600
						}
						position++
						goto l1599
					l1600:
						position, tokenIndex, depth = position1599, tokenIndex1599, depth1599
						if buffer[position] != rune('S') {
							goto l1594
						}
						position++
					}
				l1599:
					depth--
					add(rulePegText, position1596)
				}
				if !_rules[ruleAction112]() {
					goto l1594
				}
				depth--
				add(ruleIs, position1595)
			}
			return true
		l1594:
			position, tokenIndex, depth = position1594, tokenIndex1594, depth1594
			return false
		},
		/* 142 IsNot <- <(<(('i' / 'I') ('s' / 'S') sp (('n' / 'N') ('o' / 'O') ('t' / 'T')))> Action113)> */
		func() bool {
			position1601, tokenIndex1601, depth1601 := position, tokenIndex, depth
			{
				position1602 := position
				depth++
				{
					position1603 := position
					depth++
					{
						position1604, tokenIndex1604, depth1604 := position, tokenIndex, depth
						if buffer[position] != rune('i') {
							goto l1605
						}
						position++
						goto l1604
					l1605:
						position, tokenIndex, depth = position1604, tokenIndex1604, depth1604
						if buffer[position] != rune('I') {
							goto l1601
						}
						position++
					}
				l1604:
					{
						position1606, tokenIndex1606, depth1606 := position, tokenIndex, depth
						if buffer[position] != rune('s') {
							goto l1607
						}
						position++
						goto l1606
					l1607:
						position, tokenIndex, depth = position1606, tokenIndex1606, depth1606
						if buffer[position] != rune('S') {
							goto l1601
						}
						position++
					}
				l1606:
					if !_rules[rulesp]() {
						goto l1601
					}
					{
						position1608, tokenIndex1608, depth1608 := position, tokenIndex, depth
						if buffer[position] != rune('n') {
							goto l1609
						}
						position++
						goto l1608
					l1609:
						position, tokenIndex, depth = position1608, tokenIndex1608, depth1608
						if buffer[position] != rune('N') {
							goto l1601
						}
						position++
					}
				l1608:
					{
						position1610, tokenIndex1610, depth1610 := position, tokenIndex, depth
						if buffer[position] != rune('o') {
							goto l1611
						}
						position++
						goto l1610
					l1611:
						position, tokenIndex, depth = position1610, tokenIndex1610, depth1610
						if buffer[position] != rune('O') {
							goto l1601
						}
						position++
					}
				l1610:
					{
						position1612, tokenIndex1612, depth1612 := position, tokenIndex, depth
						if buffer[position] != rune('t') {
							goto l1613
						}
						position++
						goto l1612
					l1613:
						position, tokenIndex, depth = position1612, tokenIndex1612, depth1612
						if buffer[position] != rune('T') {
							goto l1601
						}
						position++
					}
				l1612:
					depth--
					add(rulePegText, position1603)
				}
				if !_rules[ruleAction113]() {
					goto l1601
				}
				depth--
				add(ruleIsNot, position1602)
			}
			return true
		l1601:
			position, tokenIndex, depth = position1601, tokenIndex1601, depth1601
			return false
		},
		/* 143 Plus <- <(<'+'> Action114)> */
		func() bool {
			position1614, tokenIndex1614, depth1614 := position, tokenIndex, depth
			{
				position1615 := position
				depth++
				{
					position1616 := position
					depth++
					if buffer[position] != rune('+') {
						goto l1614
					}
					position++
					depth--
					add(rulePegText, position1616)
				}
				if !_rules[ruleAction114]() {
					goto l1614
				}
				depth--
				add(rulePlus, position1615)
			}
			return true
		l1614:
			position, tokenIndex, depth = position1614, tokenIndex1614, depth1614
			return false
		},
		/* 144 Minus <- <(<'-'> Action115)> */
		func() bool {
			position1617, tokenIndex1617, depth1617 := position, tokenIndex, depth
			{
				position1618 := position
				depth++
				{
					position1619 := position
					depth++
					if buffer[position] != rune('-') {
						goto l1617
					}
					position++
					depth--
					add(rulePegText, position1619)
				}
				if !_rules[ruleAction115]() {
					goto l1617
				}
				depth--
				add(ruleMinus, position1618)
			}
			return true
		l1617:
			position, tokenIndex, depth = position1617, tokenIndex1617, depth1617
			return false
		},
		/* 145 Multiply <- <(<'*'> Action116)> */
		func() bool {
			position1620, tokenIndex1620, depth1620 := position, tokenIndex, depth
			{
				position1621 := position
				depth++
				{
					position1622 := position
					depth++
					if buffer[position] != rune('*') {
						goto l1620
					}
					position++
					depth--
					add(rulePegText, position1622)
				}
				if !_rules[ruleAction116]() {
					goto l1620
				}
				depth--
				add(ruleMultiply, position1621)
			}
			return true
		l1620:
			position, tokenIndex, depth = position1620, tokenIndex1620, depth1620
			return false
		},
		/* 146 Divide <- <(<'/'> Action117)> */
		func() bool {
			position1623, tokenIndex1623, depth1623 := position, tokenIndex, depth
			{
				position1624 := position
				depth++
				{
					position1625 := position
					depth++
					if buffer[position] != rune('/') {
						goto l1623
					}
					position++
					depth--
					add(rulePegText, position1625)
				}
				if !_rules[ruleAction117]() {
					goto l1623
				}
				depth--
				add(ruleDivide, position1624)
			}
			return true
		l1623:
			position, tokenIndex, depth = position1623, tokenIndex1623, depth1623
			return false
		},
		/* 147 Modulo <- <(<'%'> Action118)> */
		func() bool {
			position1626, tokenIndex1626, depth1626 := position, tokenIndex, depth
			{
				position1627 := position
				depth++
				{
					position1628 := position
					depth++
					if buffer[position] != rune('%') {
						goto l1626
					}
					position++
					depth--
					add(rulePegText, position1628)
				}
				if !_rules[ruleAction118]() {
					goto l1626
				}
				depth--
				add(ruleModulo, position1627)
			}
			return true
		l1626:
			position, tokenIndex, depth = position1626, tokenIndex1626, depth1626
			return false
		},
		/* 148 UnaryMinus <- <(<'-'> Action119)> */
		func() bool {
			position1629, tokenIndex1629, depth1629 := position, tokenIndex, depth
			{
				position1630 := position
				depth++
				{
					position1631 := position
					depth++
					if buffer[position] != rune('-') {
						goto l1629
					}
					position++
					depth--
					add(rulePegText, position1631)
				}
				if !_rules[ruleAction119]() {
					goto l1629
				}
				depth--
				add(ruleUnaryMinus, position1630)
			}
			return true
		l1629:
			position, tokenIndex, depth = position1629, tokenIndex1629, depth1629
			return false
		},
		/* 149 Identifier <- <(<ident> Action120)> */
		func() bool {
			position1632, tokenIndex1632, depth1632 := position, tokenIndex, depth
			{
				position1633 := position
				depth++
				{
					position1634 := position
					depth++
					if !_rules[ruleident]() {
						goto l1632
					}
					depth--
					add(rulePegText, position1634)
				}
				if !_rules[ruleAction120]() {
					goto l1632
				}
				depth--
				add(ruleIdentifier, position1633)
			}
			return true
		l1632:
			position, tokenIndex, depth = position1632, tokenIndex1632, depth1632
			return false
		},
		/* 150 TargetIdentifier <- <(<('*' / jsonSetPath)> Action121)> */
		func() bool {
			position1635, tokenIndex1635, depth1635 := position, tokenIndex, depth
			{
				position1636 := position
				depth++
				{
					position1637 := position
					depth++
					{
						position1638, tokenIndex1638, depth1638 := position, tokenIndex, depth
						if buffer[position] != rune('*') {
							goto l1639
						}
						position++
						goto l1638
					l1639:
						position, tokenIndex, depth = position1638, tokenIndex1638, depth1638
						if !_rules[rulejsonSetPath]() {
							goto l1635
						}
					}
				l1638:
					depth--
					add(rulePegText, position1637)
				}
				if !_rules[ruleAction121]() {
					goto l1635
				}
				depth--
				add(ruleTargetIdentifier, position1636)
			}
			return true
		l1635:
			position, tokenIndex, depth = position1635, tokenIndex1635, depth1635
			return false
		},
		/* 151 ident <- <(([a-z] / [A-Z]) ([a-z] / [A-Z] / [0-9] / '_')*)> */
		func() bool {
			position1640, tokenIndex1640, depth1640 := position, tokenIndex, depth
			{
				position1641 := position
				depth++
				{
					position1642, tokenIndex1642, depth1642 := position, tokenIndex, depth
					if c := buffer[position]; c < rune('a') || c > rune('z') {
						goto l1643
					}
					position++
					goto l1642
				l1643:
					position, tokenIndex, depth = position1642, tokenIndex1642, depth1642
					if c := buffer[position]; c < rune('A') || c > rune('Z') {
						goto l1640
					}
					position++
				}
			l1642:
			l1644:
				{
					position1645, tokenIndex1645, depth1645 := position, tokenIndex, depth
					{
						position1646, tokenIndex1646, depth1646 := position, tokenIndex, depth
						if c := buffer[position]; c < rune('a') || c > rune('z') {
							goto l1647
						}
						position++
						goto l1646
					l1647:
						position, tokenIndex, depth = position1646, tokenIndex1646, depth1646
						if c := buffer[position]; c < rune('A') || c > rune('Z') {
							goto l1648
						}
						position++
						goto l1646
					l1648:
						position, tokenIndex, depth = position1646, tokenIndex1646, depth1646
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1649
						}
						position++
						goto l1646
					l1649:
						position, tokenIndex, depth = position1646, tokenIndex1646, depth1646
						if buffer[position] != rune('_') {
							goto l1645
						}
						position++
					}
				l1646:
					goto l1644
				l1645:
					position, tokenIndex, depth = position1645, tokenIndex1645, depth1645
				}
				depth--
				add(ruleident, position1641)
			}
			return true
		l1640:
			position, tokenIndex, depth = position1640, tokenIndex1640, depth1640
			return false
		},
		/* 152 jsonGetPath <- <(jsonPathHead jsonGetPathNonHead*)> */
		func() bool {
			position1650, tokenIndex1650, depth1650 := position, tokenIndex, depth
			{
				position1651 := position
				depth++
				if !_rules[rulejsonPathHead]() {
					goto l1650
				}
			l1652:
				{
					position1653, tokenIndex1653, depth1653 := position, tokenIndex, depth
					if !_rules[rulejsonGetPathNonHead]() {
						goto l1653
					}
					goto l1652
				l1653:
					position, tokenIndex, depth = position1653, tokenIndex1653, depth1653
				}
				depth--
				add(rulejsonGetPath, position1651)
			}
			return true
		l1650:
			position, tokenIndex, depth = position1650, tokenIndex1650, depth1650
			return false
		},
		/* 153 jsonSetPath <- <(jsonPathHead jsonSetPathNonHead*)> */
		func() bool {
			position1654, tokenIndex1654, depth1654 := position, tokenIndex, depth
			{
				position1655 := position
				depth++
				if !_rules[rulejsonPathHead]() {
					goto l1654
				}
			l1656:
				{
					position1657, tokenIndex1657, depth1657 := position, tokenIndex, depth
					if !_rules[rulejsonSetPathNonHead]() {
						goto l1657
					}
					goto l1656
				l1657:
					position, tokenIndex, depth = position1657, tokenIndex1657, depth1657
				}
				depth--
				add(rulejsonSetPath, position1655)
			}
			return true
		l1654:
			position, tokenIndex, depth = position1654, tokenIndex1654, depth1654
			return false
		},
		/* 154 jsonPathHead <- <(jsonMapAccessString / jsonMapAccessBracket)> */
		func() bool {
			position1658, tokenIndex1658, depth1658 := position, tokenIndex, depth
			{
				position1659 := position
				depth++
				{
					position1660, tokenIndex1660, depth1660 := position, tokenIndex, depth
					if !_rules[rulejsonMapAccessString]() {
						goto l1661
					}
					goto l1660
				l1661:
					position, tokenIndex, depth = position1660, tokenIndex1660, depth1660
					if !_rules[rulejsonMapAccessBracket]() {
						goto l1658
					}
				}
			l1660:
				depth--
				add(rulejsonPathHead, position1659)
			}
			return true
		l1658:
			position, tokenIndex, depth = position1658, tokenIndex1658, depth1658
			return false
		},
		/* 155 jsonGetPathNonHead <- <(jsonMapMultipleLevel / jsonMapSingleLevel / jsonArrayFullSlice / jsonArrayPartialSlice / jsonArraySlice / jsonArrayAccess)> */
		func() bool {
			position1662, tokenIndex1662, depth1662 := position, tokenIndex, depth
			{
				position1663 := position
				depth++
				{
					position1664, tokenIndex1664, depth1664 := position, tokenIndex, depth
					if !_rules[rulejsonMapMultipleLevel]() {
						goto l1665
					}
					goto l1664
				l1665:
					position, tokenIndex, depth = position1664, tokenIndex1664, depth1664
					if !_rules[rulejsonMapSingleLevel]() {
						goto l1666
					}
					goto l1664
				l1666:
					position, tokenIndex, depth = position1664, tokenIndex1664, depth1664
					if !_rules[rulejsonArrayFullSlice]() {
						goto l1667
					}
					goto l1664
				l1667:
					position, tokenIndex, depth = position1664, tokenIndex1664, depth1664
					if !_rules[rulejsonArrayPartialSlice]() {
						goto l1668
					}
					goto l1664
				l1668:
					position, tokenIndex, depth = position1664, tokenIndex1664, depth1664
					if !_rules[rulejsonArraySlice]() {
						goto l1669
					}
					goto l1664
				l1669:
					position, tokenIndex, depth = position1664, tokenIndex1664, depth1664
					if !_rules[rulejsonArrayAccess]() {
						goto l1662
					}
				}
			l1664:
				depth--
				add(rulejsonGetPathNonHead, position1663)
			}
			return true
		l1662:
			position, tokenIndex, depth = position1662, tokenIndex1662, depth1662
			return false
		},
		/* 156 jsonSetPathNonHead <- <(jsonMapSingleLevel / jsonNonNegativeArrayAccess)> */
		func() bool {
			position1670, tokenIndex1670, depth1670 := position, tokenIndex, depth
			{
				position1671 := position
				depth++
				{
					position1672, tokenIndex1672, depth1672 := position, tokenIndex, depth
					if !_rules[rulejsonMapSingleLevel]() {
						goto l1673
					}
					goto l1672
				l1673:
					position, tokenIndex, depth = position1672, tokenIndex1672, depth1672
					if !_rules[rulejsonNonNegativeArrayAccess]() {
						goto l1670
					}
				}
			l1672:
				depth--
				add(rulejsonSetPathNonHead, position1671)
			}
			return true
		l1670:
			position, tokenIndex, depth = position1670, tokenIndex1670, depth1670
			return false
		},
		/* 157 jsonMapSingleLevel <- <(('.' jsonMapAccessString) / jsonMapAccessBracket)> */
		func() bool {
			position1674, tokenIndex1674, depth1674 := position, tokenIndex, depth
			{
				position1675 := position
				depth++
				{
					position1676, tokenIndex1676, depth1676 := position, tokenIndex, depth
					if buffer[position] != rune('.') {
						goto l1677
					}
					position++
					if !_rules[rulejsonMapAccessString]() {
						goto l1677
					}
					goto l1676
				l1677:
					position, tokenIndex, depth = position1676, tokenIndex1676, depth1676
					if !_rules[rulejsonMapAccessBracket]() {
						goto l1674
					}
				}
			l1676:
				depth--
				add(rulejsonMapSingleLevel, position1675)
			}
			return true
		l1674:
			position, tokenIndex, depth = position1674, tokenIndex1674, depth1674
			return false
		},
		/* 158 jsonMapMultipleLevel <- <('.' '.' (jsonMapAccessString / jsonMapAccessBracket))> */
		func() bool {
			position1678, tokenIndex1678, depth1678 := position, tokenIndex, depth
			{
				position1679 := position
				depth++
				if buffer[position] != rune('.') {
					goto l1678
				}
				position++
				if buffer[position] != rune('.') {
					goto l1678
				}
				position++
				{
					position1680, tokenIndex1680, depth1680 := position, tokenIndex, depth
					if !_rules[rulejsonMapAccessString]() {
						goto l1681
					}
					goto l1680
				l1681:
					position, tokenIndex, depth = position1680, tokenIndex1680, depth1680
					if !_rules[rulejsonMapAccessBracket]() {
						goto l1678
					}
				}
			l1680:
				depth--
				add(rulejsonMapMultipleLevel, position1679)
			}
			return true
		l1678:
			position, tokenIndex, depth = position1678, tokenIndex1678, depth1678
			return false
		},
		/* 159 jsonMapAccessString <- <<(([a-z] / [A-Z]) ([a-z] / [A-Z] / [0-9] / '_')*)>> */
		func() bool {
			position1682, tokenIndex1682, depth1682 := position, tokenIndex, depth
			{
				position1683 := position
				depth++
				{
					position1684 := position
					depth++
					{
						position1685, tokenIndex1685, depth1685 := position, tokenIndex, depth
						if c := buffer[position]; c < rune('a') || c > rune('z') {
							goto l1686
						}
						position++
						goto l1685
					l1686:
						position, tokenIndex, depth = position1685, tokenIndex1685, depth1685
						if c := buffer[position]; c < rune('A') || c > rune('Z') {
							goto l1682
						}
						position++
					}
				l1685:
				l1687:
					{
						position1688, tokenIndex1688, depth1688 := position, tokenIndex, depth
						{
							position1689, tokenIndex1689, depth1689 := position, tokenIndex, depth
							if c := buffer[position]; c < rune('a') || c > rune('z') {
								goto l1690
							}
							position++
							goto l1689
						l1690:
							position, tokenIndex, depth = position1689, tokenIndex1689, depth1689
							if c := buffer[position]; c < rune('A') || c > rune('Z') {
								goto l1691
							}
							position++
							goto l1689
						l1691:
							position, tokenIndex, depth = position1689, tokenIndex1689, depth1689
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1692
							}
							position++
							goto l1689
						l1692:
							position, tokenIndex, depth = position1689, tokenIndex1689, depth1689
							if buffer[position] != rune('_') {
								goto l1688
							}
							position++
						}
					l1689:
						goto l1687
					l1688:
						position, tokenIndex, depth = position1688, tokenIndex1688, depth1688
					}
					depth--
					add(rulePegText, position1684)
				}
				depth--
				add(rulejsonMapAccessString, position1683)
			}
			return true
		l1682:
			position, tokenIndex, depth = position1682, tokenIndex1682, depth1682
			return false
		},
		/* 160 jsonMapAccessBracket <- <('[' singleQuotedString ']')> */
		func() bool {
			position1693, tokenIndex1693, depth1693 := position, tokenIndex, depth
			{
				position1694 := position
				depth++
				if buffer[position] != rune('[') {
					goto l1693
				}
				position++
				if !_rules[rulesingleQuotedString]() {
					goto l1693
				}
				if buffer[position] != rune(']') {
					goto l1693
				}
				position++
				depth--
				add(rulejsonMapAccessBracket, position1694)
			}
			return true
		l1693:
			position, tokenIndex, depth = position1693, tokenIndex1693, depth1693
			return false
		},
		/* 161 singleQuotedString <- <('\'' <(('\'' '\'') / (!'\'' .))*> '\'')> */
		func() bool {
			position1695, tokenIndex1695, depth1695 := position, tokenIndex, depth
			{
				position1696 := position
				depth++
				if buffer[position] != rune('\'') {
					goto l1695
				}
				position++
				{
					position1697 := position
					depth++
				l1698:
					{
						position1699, tokenIndex1699, depth1699 := position, tokenIndex, depth
						{
							position1700, tokenIndex1700, depth1700 := position, tokenIndex, depth
							if buffer[position] != rune('\'') {
								goto l1701
							}
							position++
							if buffer[position] != rune('\'') {
								goto l1701
							}
							position++
							goto l1700
						l1701:
							position, tokenIndex, depth = position1700, tokenIndex1700, depth1700
							{
								position1702, tokenIndex1702, depth1702 := position, tokenIndex, depth
								if buffer[position] != rune('\'') {
									goto l1702
								}
								position++
								goto l1699
							l1702:
								position, tokenIndex, depth = position1702, tokenIndex1702, depth1702
							}
							if !matchDot() {
								goto l1699
							}
						}
					l1700:
						goto l1698
					l1699:
						position, tokenIndex, depth = position1699, tokenIndex1699, depth1699
					}
					depth--
					add(rulePegText, position1697)
				}
				if buffer[position] != rune('\'') {
					goto l1695
				}
				position++
				depth--
				add(rulesingleQuotedString, position1696)
			}
			return true
		l1695:
			position, tokenIndex, depth = position1695, tokenIndex1695, depth1695
			return false
		},
		/* 162 jsonArrayAccess <- <('[' <('-'? [0-9]+)> ']')> */
		func() bool {
			position1703, tokenIndex1703, depth1703 := position, tokenIndex, depth
			{
				position1704 := position
				depth++
				if buffer[position] != rune('[') {
					goto l1703
				}
				position++
				{
					position1705 := position
					depth++
					{
						position1706, tokenIndex1706, depth1706 := position, tokenIndex, depth
						if buffer[position] != rune('-') {
							goto l1706
						}
						position++
						goto l1707
					l1706:
						position, tokenIndex, depth = position1706, tokenIndex1706, depth1706
					}
				l1707:
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l1703
					}
					position++
				l1708:
					{
						position1709, tokenIndex1709, depth1709 := position, tokenIndex, depth
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1709
						}
						position++
						goto l1708
					l1709:
						position, tokenIndex, depth = position1709, tokenIndex1709, depth1709
					}
					depth--
					add(rulePegText, position1705)
				}
				if buffer[position] != rune(']') {
					goto l1703
				}
				position++
				depth--
				add(rulejsonArrayAccess, position1704)
			}
			return true
		l1703:
			position, tokenIndex, depth = position1703, tokenIndex1703, depth1703
			return false
		},
		/* 163 jsonNonNegativeArrayAccess <- <('[' <[0-9]+> ']')> */
		func() bool {
			position1710, tokenIndex1710, depth1710 := position, tokenIndex, depth
			{
				position1711 := position
				depth++
				if buffer[position] != rune('[') {
					goto l1710
				}
				position++
				{
					position1712 := position
					depth++
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l1710
					}
					position++
				l1713:
					{
						position1714, tokenIndex1714, depth1714 := position, tokenIndex, depth
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1714
						}
						position++
						goto l1713
					l1714:
						position, tokenIndex, depth = position1714, tokenIndex1714, depth1714
					}
					depth--
					add(rulePegText, position1712)
				}
				if buffer[position] != rune(']') {
					goto l1710
				}
				position++
				depth--
				add(rulejsonNonNegativeArrayAccess, position1711)
			}
			return true
		l1710:
			position, tokenIndex, depth = position1710, tokenIndex1710, depth1710
			return false
		},
		/* 164 jsonArraySlice <- <('[' <('-'? [0-9]+ ':' '-'? [0-9]+ (':' '-'? [0-9]+)?)> ']')> */
		func() bool {
			position1715, tokenIndex1715, depth1715 := position, tokenIndex, depth
			{
				position1716 := position
				depth++
				if buffer[position] != rune('[') {
					goto l1715
				}
				position++
				{
					position1717 := position
					depth++
					{
						position1718, tokenIndex1718, depth1718 := position, tokenIndex, depth
						if buffer[position] != rune('-') {
							goto l1718
						}
						position++
						goto l1719
					l1718:
						position, tokenIndex, depth = position1718, tokenIndex1718, depth1718
					}
				l1719:
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l1715
					}
					position++
				l1720:
					{
						position1721, tokenIndex1721, depth1721 := position, tokenIndex, depth
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1721
						}
						position++
						goto l1720
					l1721:
						position, tokenIndex, depth = position1721, tokenIndex1721, depth1721
					}
					if buffer[position] != rune(':') {
						goto l1715
					}
					position++
					{
						position1722, tokenIndex1722, depth1722 := position, tokenIndex, depth
						if buffer[position] != rune('-') {
							goto l1722
						}
						position++
						goto l1723
					l1722:
						position, tokenIndex, depth = position1722, tokenIndex1722, depth1722
					}
				l1723:
					if c := buffer[position]; c < rune('0') || c > rune('9') {
						goto l1715
					}
					position++
				l1724:
					{
						position1725, tokenIndex1725, depth1725 := position, tokenIndex, depth
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1725
						}
						position++
						goto l1724
					l1725:
						position, tokenIndex, depth = position1725, tokenIndex1725, depth1725
					}
					{
						position1726, tokenIndex1726, depth1726 := position, tokenIndex, depth
						if buffer[position] != rune(':') {
							goto l1726
						}
						position++
						{
							position1728, tokenIndex1728, depth1728 := position, tokenIndex, depth
							if buffer[position] != rune('-') {
								goto l1728
							}
							position++
							goto l1729
						l1728:
							position, tokenIndex, depth = position1728, tokenIndex1728, depth1728
						}
					l1729:
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1726
						}
						position++
					l1730:
						{
							position1731, tokenIndex1731, depth1731 := position, tokenIndex, depth
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1731
							}
							position++
							goto l1730
						l1731:
							position, tokenIndex, depth = position1731, tokenIndex1731, depth1731
						}
						goto l1727
					l1726:
						position, tokenIndex, depth = position1726, tokenIndex1726, depth1726
					}
				l1727:
					depth--
					add(rulePegText, position1717)
				}
				if buffer[position] != rune(']') {
					goto l1715
				}
				position++
				depth--
				add(rulejsonArraySlice, position1716)
			}
			return true
		l1715:
			position, tokenIndex, depth = position1715, tokenIndex1715, depth1715
			return false
		},
		/* 165 jsonArrayPartialSlice <- <('[' <((':' '-'? [0-9]+) / ('-'? [0-9]+ ':'))> ']')> */
		func() bool {
			position1732, tokenIndex1732, depth1732 := position, tokenIndex, depth
			{
				position1733 := position
				depth++
				if buffer[position] != rune('[') {
					goto l1732
				}
				position++
				{
					position1734 := position
					depth++
					{
						position1735, tokenIndex1735, depth1735 := position, tokenIndex, depth
						if buffer[position] != rune(':') {
							goto l1736
						}
						position++
						{
							position1737, tokenIndex1737, depth1737 := position, tokenIndex, depth
							if buffer[position] != rune('-') {
								goto l1737
							}
							position++
							goto l1738
						l1737:
							position, tokenIndex, depth = position1737, tokenIndex1737, depth1737
						}
					l1738:
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1736
						}
						position++
					l1739:
						{
							position1740, tokenIndex1740, depth1740 := position, tokenIndex, depth
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1740
							}
							position++
							goto l1739
						l1740:
							position, tokenIndex, depth = position1740, tokenIndex1740, depth1740
						}
						goto l1735
					l1736:
						position, tokenIndex, depth = position1735, tokenIndex1735, depth1735
						{
							position1741, tokenIndex1741, depth1741 := position, tokenIndex, depth
							if buffer[position] != rune('-') {
								goto l1741
							}
							position++
							goto l1742
						l1741:
							position, tokenIndex, depth = position1741, tokenIndex1741, depth1741
						}
					l1742:
						if c := buffer[position]; c < rune('0') || c > rune('9') {
							goto l1732
						}
						position++
					l1743:
						{
							position1744, tokenIndex1744, depth1744 := position, tokenIndex, depth
							if c := buffer[position]; c < rune('0') || c > rune('9') {
								goto l1744
							}
							position++
							goto l1743
						l1744:
							position, tokenIndex, depth = position1744, tokenIndex1744, depth1744
						}
						if buffer[position] != rune(':') {
							goto l1732
						}
						position++
					}
				l1735:
					depth--
					add(rulePegText, position1734)
				}
				if buffer[position] != rune(']') {
					goto l1732
				}
				position++
				depth--
				add(rulejsonArrayPartialSlice, position1733)
			}
			return true
		l1732:
			position, tokenIndex, depth = position1732, tokenIndex1732, depth1732
			return false
		},
		/* 166 jsonArrayFullSlice <- <('[' ':' ']')> */
		func() bool {
			position1745, tokenIndex1745, depth1745 := position, tokenIndex, depth
			{
				position1746 := position
				depth++
				if buffer[position] != rune('[') {
					goto l1745
				}
				position++
				if buffer[position] != rune(':') {
					goto l1745
				}
				position++
				if buffer[position] != rune(']') {
					goto l1745
				}
				position++
				depth--
				add(rulejsonArrayFullSlice, position1746)
			}
			return true
		l1745:
			position, tokenIndex, depth = position1745, tokenIndex1745, depth1745
			return false
		},
		/* 167 spElem <- <(' ' / '\t' / '\n' / '\r' / comment / finalComment)> */
		func() bool {
			position1747, tokenIndex1747, depth1747 := position, tokenIndex, depth
			{
				position1748 := position
				depth++
				{
					position1749, tokenIndex1749, depth1749 := position, tokenIndex, depth
					if buffer[position] != rune(' ') {
						goto l1750
					}
					position++
					goto l1749
				l1750:
					position, tokenIndex, depth = position1749, tokenIndex1749, depth1749
					if buffer[position] != rune('\t') {
						goto l1751
					}
					position++
					goto l1749
				l1751:
					position, tokenIndex, depth = position1749, tokenIndex1749, depth1749
					if buffer[position] != rune('\n') {
						goto l1752
					}
					position++
					goto l1749
				l1752:
					position, tokenIndex, depth = position1749, tokenIndex1749, depth1749
					if buffer[position] != rune('\r') {
						goto l1753
					}
					position++
					goto l1749
				l1753:
					position, tokenIndex, depth = position1749, tokenIndex1749, depth1749
					if !_rules[rulecomment]() {
						goto l1754
					}
					goto l1749
				l1754:
					position, tokenIndex, depth = position1749, tokenIndex1749, depth1749
					if !_rules[rulefinalComment]() {
						goto l1747
					}
				}
			l1749:
				depth--
				add(rulespElem, position1748)
			}
			return true
		l1747:
			position, tokenIndex, depth = position1747, tokenIndex1747, depth1747
			return false
		},
		/* 168 sp <- <spElem+> */
		func() bool {
			position1755, tokenIndex1755, depth1755 := position, tokenIndex, depth
			{
				position1756 := position
				depth++
				if !_rules[rulespElem]() {
					goto l1755
				}
			l1757:
				{
					position1758, tokenIndex1758, depth1758 := position, tokenIndex, depth
					if !_rules[rulespElem]() {
						goto l1758
					}
					goto l1757
				l1758:
					position, tokenIndex, depth = position1758, tokenIndex1758, depth1758
				}
				depth--
				add(rulesp, position1756)
			}
			return true
		l1755:
			position, tokenIndex, depth = position1755, tokenIndex1755, depth1755
			return false
		},
		/* 169 spOpt <- <spElem*> */
		func() bool {
			{
				position1760 := position
				depth++
			l1761:
				{
					position1762, tokenIndex1762, depth1762 := position, tokenIndex, depth
					if !_rules[rulespElem]() {
						goto l1762
					}
					goto l1761
				l1762:
					position, tokenIndex, depth = position1762, tokenIndex1762, depth1762
				}
				depth--
				add(rulespOpt, position1760)
			}
			return true
		},
		/* 170 comment <- <('-' '-' (!('\r' / '\n') .)* ('\r' / '\n'))> */
		func() bool {
			position1763, tokenIndex1763, depth1763 := position, tokenIndex, depth
			{
				position1764 := position
				depth++
				if buffer[position] != rune('-') {
					goto l1763
				}
				position++
				if buffer[position] != rune('-') {
					goto l1763
				}
				position++
			l1765:
				{
					position1766, tokenIndex1766, depth1766 := position, tokenIndex, depth
					{
						position1767, tokenIndex1767, depth1767 := position, tokenIndex, depth
						{
							position1768, tokenIndex1768, depth1768 := position, tokenIndex, depth
							if buffer[position] != rune('\r') {
								goto l1769
							}
							position++
							goto l1768
						l1769:
							position, tokenIndex, depth = position1768, tokenIndex1768, depth1768
							if buffer[position] != rune('\n') {
								goto l1767
							}
							position++
						}
					l1768:
						goto l1766
					l1767:
						position, tokenIndex, depth = position1767, tokenIndex1767, depth1767
					}
					if !matchDot() {
						goto l1766
					}
					goto l1765
				l1766:
					position, tokenIndex, depth = position1766, tokenIndex1766, depth1766
				}
				{
					position1770, tokenIndex1770, depth1770 := position, tokenIndex, depth
					if buffer[position] != rune('\r') {
						goto l1771
					}
					position++
					goto l1770
				l1771:
					position, tokenIndex, depth = position1770, tokenIndex1770, depth1770
					if buffer[position] != rune('\n') {
						goto l1763
					}
					position++
				}
			l1770:
				depth--
				add(rulecomment, position1764)
			}
			return true
		l1763:
			position, tokenIndex, depth = position1763, tokenIndex1763, depth1763
			return false
		},
		/* 171 finalComment <- <('-' '-' (!('\r' / '\n') .)* !.)> */
		func() bool {
			position1772, tokenIndex1772, depth1772 := position, tokenIndex, depth
			{
				position1773 := position
				depth++
				if buffer[position] != rune('-') {
					goto l1772
				}
				position++
				if buffer[position] != rune('-') {
					goto l1772
				}
				position++
			l1774:
				{
					position1775, tokenIndex1775, depth1775 := position, tokenIndex, depth
					{
						position1776, tokenIndex1776, depth1776 := position, tokenIndex, depth
						{
							position1777, tokenIndex1777, depth1777 := position, tokenIndex, depth
							if buffer[position] != rune('\r') {
								goto l1778
							}
							position++
							goto l1777
						l1778:
							position, tokenIndex, depth = position1777, tokenIndex1777, depth1777
							if buffer[position] != rune('\n') {
								goto l1776
							}
							position++
						}
					l1777:
						goto l1775
					l1776:
						position, tokenIndex, depth = position1776, tokenIndex1776, depth1776
					}
					if !matchDot() {
						goto l1775
					}
					goto l1774
				l1775:
					position, tokenIndex, depth = position1775, tokenIndex1775, depth1775
				}
				{
					position1779, tokenIndex1779, depth1779 := position, tokenIndex, depth
					if !matchDot() {
						goto l1779
					}
					goto l1772
				l1779:
					position, tokenIndex, depth = position1779, tokenIndex1779, depth1779
				}
				depth--
				add(rulefinalComment, position1773)
			}
			return true
		l1772:
			position, tokenIndex, depth = position1772, tokenIndex1772, depth1772
			return false
		},
		nil,
		/* 174 Action0 <- <{
		    p.IncludeTrailingWhitespace(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction0, position)
			}
			return true
		},
		/* 175 Action1 <- <{
		    p.IncludeTrailingWhitespace(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction1, position)
			}
			return true
		},
		/* 176 Action2 <- <{
		    p.AssembleSelect()
		}> */
		func() bool {
			{
				add(ruleAction2, position)
			}
			return true
		},
		/* 177 Action3 <- <{
		    p.AssembleSelectUnion(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction3, position)
			}
			return true
		},
		/* 178 Action4 <- <{
		    p.AssembleCreateStreamAsSelect()
		}> */
		func() bool {
			{
				add(ruleAction4, position)
			}
			return true
		},
		/* 179 Action5 <- <{
		    p.AssembleCreateStreamAsSelectUnion()
		}> */
		func() bool {
			{
				add(ruleAction5, position)
			}
			return true
		},
		/* 180 Action6 <- <{
		    p.AssembleCreateSource()
		}> */
		func() bool {
			{
				add(ruleAction6, position)
			}
			return true
		},
		/* 181 Action7 <- <{
		    p.AssembleCreateSink()
		}> */
		func() bool {
			{
				add(ruleAction7, position)
			}
			return true
		},
		/* 182 Action8 <- <{
		    p.AssembleCreateState()
		}> */
		func() bool {
			{
				add(ruleAction8, position)
			}
			return true
		},
		/* 183 Action9 <- <{
		    p.AssembleUpdateState()
		}> */
		func() bool {
			{
				add(ruleAction9, position)
			}
			return true
		},
		/* 184 Action10 <- <{
		    p.AssembleUpdateSource()
		}> */
		func() bool {
			{
				add(ruleAction10, position)
			}
			return true
		},
		/* 185 Action11 <- <{
		    p.AssembleUpdateSink()
		}> */
		func() bool {
			{
				add(ruleAction11, position)
			}
			return true
		},
		/* 186 Action12 <- <{
		    p.AssembleInsertIntoSelect()
		}> */
		func() bool {
			{
				add(ruleAction12, position)
			}
			return true
		},
		/* 187 Action13 <- <{
		    p.AssembleInsertIntoFrom()
		}> */
		func() bool {
			{
				add(ruleAction13, position)
			}
			return true
		},
		/* 188 Action14 <- <{
		    p.AssemblePauseSource()
		}> */
		func() bool {
			{
				add(ruleAction14, position)
			}
			return true
		},
		/* 189 Action15 <- <{
		    p.AssembleResumeSource()
		}> */
		func() bool {
			{
				add(ruleAction15, position)
			}
			return true
		},
		/* 190 Action16 <- <{
		    p.AssembleRewindSource()
		}> */
		func() bool {
			{
				add(ruleAction16, position)
			}
			return true
		},
		/* 191 Action17 <- <{
		    p.AssembleDropSource()
		}> */
		func() bool {
			{
				add(ruleAction17, position)
			}
			return true
		},
		/* 192 Action18 <- <{
		    p.AssembleDropStream()
		}> */
		func() bool {
			{
				add(ruleAction18, position)
			}
			return true
		},
		/* 193 Action19 <- <{
		    p.AssembleDropSink()
		}> */
		func() bool {
			{
				add(ruleAction19, position)
			}
			return true
		},
		/* 194 Action20 <- <{
		    p.AssembleDropState()
		}> */
		func() bool {
			{
				add(ruleAction20, position)
			}
			return true
		},
		/* 195 Action21 <- <{
		    p.AssembleLoadState()
		}> */
		func() bool {
			{
				add(ruleAction21, position)
			}
			return true
		},
		/* 196 Action22 <- <{
		    p.AssembleLoadStateOrCreate()
		}> */
		func() bool {
			{
				add(ruleAction22, position)
			}
			return true
		},
		/* 197 Action23 <- <{
		    p.AssembleSaveState()
		}> */
		func() bool {
			{
				add(ruleAction23, position)
			}
			return true
		},
		/* 198 Action24 <- <{
		    p.AssembleEval(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction24, position)
			}
			return true
		},
		/* 199 Action25 <- <{
		    p.AssembleEmitter()
		}> */
		func() bool {
			{
				add(ruleAction25, position)
			}
			return true
		},
		/* 200 Action26 <- <{
		    p.AssembleEmitterOptions(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction26, position)
			}
			return true
		},
		/* 201 Action27 <- <{
		    p.AssembleEmitterLimit()
		}> */
		func() bool {
			{
				add(ruleAction27, position)
			}
			return true
		},
		/* 202 Action28 <- <{
		    p.AssembleEmitterSampling(CountBasedSampling, 1)
		}> */
		func() bool {
			{
				add(ruleAction28, position)
			}
			return true
		},
		/* 203 Action29 <- <{
		    p.AssembleEmitterSampling(RandomizedSampling, 1)
		}> */
		func() bool {
			{
				add(ruleAction29, position)
			}
			return true
		},
		/* 204 Action30 <- <{
		    p.AssembleEmitterSampling(TimeBasedSampling, 1)
		}> */
		func() bool {
			{
				add(ruleAction30, position)
			}
			return true
		},
		/* 205 Action31 <- <{
		    p.AssembleEmitterSampling(TimeBasedSampling, 0.001)
		}> */
		func() bool {
			{
				add(ruleAction31, position)
			}
			return true
		},
		/* 206 Action32 <- <{
		    p.AssembleProjections(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction32, position)
			}
			return true
		},
		/* 207 Action33 <- <{
		    p.AssembleAlias()
		}> */
		func() bool {
			{
				add(ruleAction33, position)
			}
			return true
		},
		/* 208 Action34 <- <{
		    // This is *always* executed, even if there is no
		    // FROM clause present in the statement.
		    p.AssembleWindowedFrom(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction34, position)
			}
			return true
		},
		/* 209 Action35 <- <{
		    p.AssembleInterval()
		}> */
		func() bool {
			{
				add(ruleAction35, position)
			}
			return true
		},
		/* 210 Action36 <- <{
		    p.AssembleInterval()
		}> */
		func() bool {
			{
				add(ruleAction36, position)
			}
			return true
		},
		/* 211 Action37 <- <{
		    // This is *always* executed, even if there is no
		    // WHERE clause present in the statement.
		    p.AssembleFilter(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction37, position)
			}
			return true
		},
		/* 212 Action38 <- <{
		    // This is *always* executed, even if there is no
		    // GROUP BY clause present in the statement.
		    p.AssembleGrouping(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction38, position)
			}
			return true
		},
		/* 213 Action39 <- <{
		    // This is *always* executed, even if there is no
		    // HAVING clause present in the statement.
		    p.AssembleHaving(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction39, position)
			}
			return true
		},
		/* 214 Action40 <- <{
		    p.EnsureAliasedStreamWindow()
		}> */
		func() bool {
			{
				add(ruleAction40, position)
			}
			return true
		},
		/* 215 Action41 <- <{
		    p.AssembleAliasedStreamWindow()
		}> */
		func() bool {
			{
				add(ruleAction41, position)
			}
			return true
		},
		/* 216 Action42 <- <{
		    p.AssembleStreamWindow()
		}> */
		func() bool {
			{
				add(ruleAction42, position)
			}
			return true
		},
		/* 217 Action43 <- <{
		    p.AssembleUDSFFuncApp()
		}> */
		func() bool {
			{
				add(ruleAction43, position)
			}
			return true
		},
		/* 218 Action44 <- <{
		    p.AssembleSourceSinkSpecs(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction44, position)
			}
			return true
		},
		/* 219 Action45 <- <{
		    p.AssembleSourceSinkSpecs(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction45, position)
			}
			return true
		},
		/* 220 Action46 <- <{
		    p.AssembleSourceSinkSpecs(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction46, position)
			}
			return true
		},
		/* 221 Action47 <- <{
		    p.AssembleSourceSinkParam()
		}> */
		func() bool {
			{
				add(ruleAction47, position)
			}
			return true
		},
		/* 222 Action48 <- <{
		    p.AssembleExpressions(begin, end)
		    p.AssembleArray()
		}> */
		func() bool {
			{
				add(ruleAction48, position)
			}
			return true
		},
		/* 223 Action49 <- <{
		    p.EnsureKeywordPresent(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction49, position)
			}
			return true
		},
		/* 224 Action50 <- <{
		    p.AssembleBinaryOperation(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction50, position)
			}
			return true
		},
		/* 225 Action51 <- <{
		    p.AssembleBinaryOperation(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction51, position)
			}
			return true
		},
		/* 226 Action52 <- <{
		    p.AssembleUnaryPrefixOperation(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction52, position)
			}
			return true
		},
		/* 227 Action53 <- <{
		    p.AssembleBinaryOperation(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction53, position)
			}
			return true
		},
		/* 228 Action54 <- <{
		    p.AssembleBinaryOperation(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction54, position)
			}
			return true
		},
		/* 229 Action55 <- <{
		    p.AssembleBinaryOperation(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction55, position)
			}
			return true
		},
		/* 230 Action56 <- <{
		    p.AssembleBinaryOperation(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction56, position)
			}
			return true
		},
		/* 231 Action57 <- <{
		    p.AssembleBinaryOperation(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction57, position)
			}
			return true
		},
		/* 232 Action58 <- <{
		    p.AssembleUnaryPrefixOperation(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction58, position)
			}
			return true
		},
		/* 233 Action59 <- <{
		    p.AssembleTypeCast(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction59, position)
			}
			return true
		},
		/* 234 Action60 <- <{
		    p.AssembleTypeCast(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction60, position)
			}
			return true
		},
		/* 235 Action61 <- <{
		    p.AssembleFuncApp()
		}> */
		func() bool {
			{
				add(ruleAction61, position)
			}
			return true
		},
		/* 236 Action62 <- <{
		    p.AssembleExpressions(begin, end)
		    p.AssembleFuncApp()
		}> */
		func() bool {
			{
				add(ruleAction62, position)
			}
			return true
		},
		/* 237 Action63 <- <{
		    p.AssembleExpressions(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction63, position)
			}
			return true
		},
		/* 238 Action64 <- <{
		    p.AssembleExpressions(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction64, position)
			}
			return true
		},
		/* 239 Action65 <- <{
		    p.AssembleSortedExpression()
		}> */
		func() bool {
			{
				add(ruleAction65, position)
			}
			return true
		},
		/* 240 Action66 <- <{
		    p.EnsureKeywordPresent(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction66, position)
			}
			return true
		},
		/* 241 Action67 <- <{
		    p.AssembleExpressions(begin, end)
		    p.AssembleArray()
		}> */
		func() bool {
			{
				add(ruleAction67, position)
			}
			return true
		},
		/* 242 Action68 <- <{
		    p.AssembleMap(begin, end)
		}> */
		func() bool {
			{
				add(ruleAction68, position)
			}
			return true
		},
		/* 243 Action69 <- <{
		    p.AssembleKeyValuePair()
		}> */
		func() bool {
			{
				add(ruleAction69, position)
			}
			return true
		},
		/* 244 Action70 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, NewStream(substr))
		}> */
		func() bool {
			{
				add(ruleAction70, position)
			}
			return true
		},
		/* 245 Action71 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, NewRowMeta(substr, TimestampMeta))
		}> */
		func() bool {
			{
				add(ruleAction71, position)
			}
			return true
		},
		/* 246 Action72 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, NewRowValue(substr))
		}> */
		func() bool {
			{
				add(ruleAction72, position)
			}
			return true
		},
		/* 247 Action73 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, NewNumericLiteral(substr))
		}> */
		func() bool {
			{
				add(ruleAction73, position)
			}
			return true
		},
		/* 248 Action74 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, NewFloatLiteral(substr))
		}> */
		func() bool {
			{
				add(ruleAction74, position)
			}
			return true
		},
		/* 249 Action75 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, FuncName(substr))
		}> */
		func() bool {
			{
				add(ruleAction75, position)
			}
			return true
		},
		/* 250 Action76 <- <{
		    p.PushComponent(begin, end, NewNullLiteral())
		}> */
		func() bool {
			{
				add(ruleAction76, position)
			}
			return true
		},
		/* 251 Action77 <- <{
		    p.PushComponent(begin, end, NewBoolLiteral(true))
		}> */
		func() bool {
			{
				add(ruleAction77, position)
			}
			return true
		},
		/* 252 Action78 <- <{
		    p.PushComponent(begin, end, NewBoolLiteral(false))
		}> */
		func() bool {
			{
				add(ruleAction78, position)
			}
			return true
		},
		/* 253 Action79 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, NewWildcard(substr))
		}> */
		func() bool {
			{
				add(ruleAction79, position)
			}
			return true
		},
		/* 254 Action80 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, NewStringLiteral(substr))
		}> */
		func() bool {
			{
				add(ruleAction80, position)
			}
			return true
		},
		/* 255 Action81 <- <{
		    p.PushComponent(begin, end, Istream)
		}> */
		func() bool {
			{
				add(ruleAction81, position)
			}
			return true
		},
		/* 256 Action82 <- <{
		    p.PushComponent(begin, end, Dstream)
		}> */
		func() bool {
			{
				add(ruleAction82, position)
			}
			return true
		},
		/* 257 Action83 <- <{
		    p.PushComponent(begin, end, Rstream)
		}> */
		func() bool {
			{
				add(ruleAction83, position)
			}
			return true
		},
		/* 258 Action84 <- <{
		    p.PushComponent(begin, end, Tuples)
		}> */
		func() bool {
			{
				add(ruleAction84, position)
			}
			return true
		},
		/* 259 Action85 <- <{
		    p.PushComponent(begin, end, Seconds)
		}> */
		func() bool {
			{
				add(ruleAction85, position)
			}
			return true
		},
		/* 260 Action86 <- <{
		    p.PushComponent(begin, end, Milliseconds)
		}> */
		func() bool {
			{
				add(ruleAction86, position)
			}
			return true
		},
		/* 261 Action87 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, StreamIdentifier(substr))
		}> */
		func() bool {
			{
				add(ruleAction87, position)
			}
			return true
		},
		/* 262 Action88 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, SourceSinkType(substr))
		}> */
		func() bool {
			{
				add(ruleAction88, position)
			}
			return true
		},
		/* 263 Action89 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, SourceSinkParamKey(substr))
		}> */
		func() bool {
			{
				add(ruleAction89, position)
			}
			return true
		},
		/* 264 Action90 <- <{
		    p.PushComponent(begin, end, Yes)
		}> */
		func() bool {
			{
				add(ruleAction90, position)
			}
			return true
		},
		/* 265 Action91 <- <{
		    p.PushComponent(begin, end, No)
		}> */
		func() bool {
			{
				add(ruleAction91, position)
			}
			return true
		},
		/* 266 Action92 <- <{
		    p.PushComponent(begin, end, Yes)
		}> */
		func() bool {
			{
				add(ruleAction92, position)
			}
			return true
		},
		/* 267 Action93 <- <{
		    p.PushComponent(begin, end, No)
		}> */
		func() bool {
			{
				add(ruleAction93, position)
			}
			return true
		},
		/* 268 Action94 <- <{
		    p.PushComponent(begin, end, Bool)
		}> */
		func() bool {
			{
				add(ruleAction94, position)
			}
			return true
		},
		/* 269 Action95 <- <{
		    p.PushComponent(begin, end, Int)
		}> */
		func() bool {
			{
				add(ruleAction95, position)
			}
			return true
		},
		/* 270 Action96 <- <{
		    p.PushComponent(begin, end, Float)
		}> */
		func() bool {
			{
				add(ruleAction96, position)
			}
			return true
		},
		/* 271 Action97 <- <{
		    p.PushComponent(begin, end, String)
		}> */
		func() bool {
			{
				add(ruleAction97, position)
			}
			return true
		},
		/* 272 Action98 <- <{
		    p.PushComponent(begin, end, Blob)
		}> */
		func() bool {
			{
				add(ruleAction98, position)
			}
			return true
		},
		/* 273 Action99 <- <{
		    p.PushComponent(begin, end, Timestamp)
		}> */
		func() bool {
			{
				add(ruleAction99, position)
			}
			return true
		},
		/* 274 Action100 <- <{
		    p.PushComponent(begin, end, Array)
		}> */
		func() bool {
			{
				add(ruleAction100, position)
			}
			return true
		},
		/* 275 Action101 <- <{
		    p.PushComponent(begin, end, Map)
		}> */
		func() bool {
			{
				add(ruleAction101, position)
			}
			return true
		},
		/* 276 Action102 <- <{
		    p.PushComponent(begin, end, Or)
		}> */
		func() bool {
			{
				add(ruleAction102, position)
			}
			return true
		},
		/* 277 Action103 <- <{
		    p.PushComponent(begin, end, And)
		}> */
		func() bool {
			{
				add(ruleAction103, position)
			}
			return true
		},
		/* 278 Action104 <- <{
		    p.PushComponent(begin, end, Not)
		}> */
		func() bool {
			{
				add(ruleAction104, position)
			}
			return true
		},
		/* 279 Action105 <- <{
		    p.PushComponent(begin, end, Equal)
		}> */
		func() bool {
			{
				add(ruleAction105, position)
			}
			return true
		},
		/* 280 Action106 <- <{
		    p.PushComponent(begin, end, Less)
		}> */
		func() bool {
			{
				add(ruleAction106, position)
			}
			return true
		},
		/* 281 Action107 <- <{
		    p.PushComponent(begin, end, LessOrEqual)
		}> */
		func() bool {
			{
				add(ruleAction107, position)
			}
			return true
		},
		/* 282 Action108 <- <{
		    p.PushComponent(begin, end, Greater)
		}> */
		func() bool {
			{
				add(ruleAction108, position)
			}
			return true
		},
		/* 283 Action109 <- <{
		    p.PushComponent(begin, end, GreaterOrEqual)
		}> */
		func() bool {
			{
				add(ruleAction109, position)
			}
			return true
		},
		/* 284 Action110 <- <{
		    p.PushComponent(begin, end, NotEqual)
		}> */
		func() bool {
			{
				add(ruleAction110, position)
			}
			return true
		},
		/* 285 Action111 <- <{
		    p.PushComponent(begin, end, Concat)
		}> */
		func() bool {
			{
				add(ruleAction111, position)
			}
			return true
		},
		/* 286 Action112 <- <{
		    p.PushComponent(begin, end, Is)
		}> */
		func() bool {
			{
				add(ruleAction112, position)
			}
			return true
		},
		/* 287 Action113 <- <{
		    p.PushComponent(begin, end, IsNot)
		}> */
		func() bool {
			{
				add(ruleAction113, position)
			}
			return true
		},
		/* 288 Action114 <- <{
		    p.PushComponent(begin, end, Plus)
		}> */
		func() bool {
			{
				add(ruleAction114, position)
			}
			return true
		},
		/* 289 Action115 <- <{
		    p.PushComponent(begin, end, Minus)
		}> */
		func() bool {
			{
				add(ruleAction115, position)
			}
			return true
		},
		/* 290 Action116 <- <{
		    p.PushComponent(begin, end, Multiply)
		}> */
		func() bool {
			{
				add(ruleAction116, position)
			}
			return true
		},
		/* 291 Action117 <- <{
		    p.PushComponent(begin, end, Divide)
		}> */
		func() bool {
			{
				add(ruleAction117, position)
			}
			return true
		},
		/* 292 Action118 <- <{
		    p.PushComponent(begin, end, Modulo)
		}> */
		func() bool {
			{
				add(ruleAction118, position)
			}
			return true
		},
		/* 293 Action119 <- <{
		    p.PushComponent(begin, end, UnaryMinus)
		}> */
		func() bool {
			{
				add(ruleAction119, position)
			}
			return true
		},
		/* 294 Action120 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, Identifier(substr))
		}> */
		func() bool {
			{
				add(ruleAction120, position)
			}
			return true
		},
		/* 295 Action121 <- <{
		    substr := string([]rune(buffer)[begin:end])
		    p.PushComponent(begin, end, Identifier(substr))
		}> */
		func() bool {
			{
				add(ruleAction121, position)
			}
			return true
		},
	}
	p.rules = _rules
}
