package parser

import (
	"reflect"
	"testing"
)

type testCase struct {
	text string
	tree ParseTree
}

var testCases = []testCase{
	{
		text: "",
		tree: ParseTree{},
	},
	{
		text: "A",
		tree: ParseTree{
			Type:     "variable",
			Element:  "a",
			Children: []ParseTree{},
		},
	},
	{
		text: "AND",
		tree: ParseTree{
			Type:     "operator",
			Element:  "and",
			Children: []ParseTree{},
		},
	},
	{
		text: "AND A B",
		tree: ParseTree{
			Type:    "operator",
			Element: "and",
			Children: []ParseTree{
				ParseTree{
					Type:     "variable",
					Element:  "a",
					Children: []ParseTree{},
				},
				ParseTree{
					Type:     "variable",
					Element:  "b",
					Children: []ParseTree{},
				},
			},
		},
	},
	{
		text: "OR A B",
		tree: ParseTree{
			Type:    "operator",
			Element: "or",
			Children: []ParseTree{
				ParseTree{
					Type:     "variable",
					Element:  "a",
					Children: []ParseTree{},
				},
				ParseTree{
					Type:     "variable",
					Element:  "b",
					Children: []ParseTree{},
				},
			},
		},
	},
	{
		text: "AND A (OR B C)",
		tree: ParseTree{
			Type:    "operator",
			Element: "and",
			Children: []ParseTree{
				ParseTree{
					Type:     "variable",
					Element:  "a",
					Children: []ParseTree{},
				},
				ParseTree{
					Type:    "operator",
					Element: "or",
					Children: []ParseTree{
						ParseTree{
							Type:     "variable",
							Element:  "b",
							Children: []ParseTree{},
						},
						ParseTree{
							Type:     "variable",
							Element:  "c",
							Children: []ParseTree{},
						},
					},
				},
			},
		},
	},
	{
		text: "OR A (AND B C)",
		tree: ParseTree{
			Type:    "operator",
			Element: "or",
			Children: []ParseTree{
				ParseTree{
					Type:     "variable",
					Element:  "a",
					Children: []ParseTree{},
				},
				ParseTree{
					Type:    "operator",
					Element: "and",
					Children: []ParseTree{
						ParseTree{
							Type:     "variable",
							Element:  "b",
							Children: []ParseTree{},
						},
						ParseTree{
							Type:     "variable",
							Element:  "c",
							Children: []ParseTree{},
						},
					},
				},
			},
		},
	},
	{
		text: "OR A (AND B (AND (A B)))",
		tree: ParseTree{
			Type:    "operator",
			Element: "or",
			Children: []ParseTree{
				ParseTree{
					Type:     "variable",
					Element:  "a",
					Children: []ParseTree{},
				},
				ParseTree{
					Type:    "operator",
					Element: "and",
					Children: []ParseTree{
						ParseTree{
							Type:     "variable",
							Element:  "b",
							Children: []ParseTree{},
						},
						ParseTree{
							Type:    "operator",
							Element: "and",
							Children: []ParseTree{
								ParseTree{
									Type:     "variable",
									Element:  "a",
									Children: []ParseTree{},
								},
								ParseTree{
									Type:     "variable",
									Element:  "b",
									Children: []ParseTree{},
								},
							},
						},
					},
				},
			},
		},
	},
}

func TestParser(t *testing.T) {
	for _, testCase := range testCases {
		actualTree := Parser(testCase.text)
		if !reflect.DeepEqual(testCase.tree, actualTree) {
			t.Fatalf("\nExpected: %v \nActual: %v", testCase.tree, actualTree)
		}
	}
}
