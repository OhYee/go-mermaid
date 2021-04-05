package mermaid

import (
	_ "embed"
	"io/ioutil"
	"math"
	"path"
	"testing"
)

const TESTFOLDER = "testcases"

func TestRender(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		source string
		result string
	}{
		{
			name:   "error",
			source: "0.mmd",
			result: "0.svg",
		},
		{
			name:   "Flowchart",
			source: "1.mmd",
			result: "1.svg",
		},
		{
			name:   "Sequence diagram",
			source: "2.mmd",
			result: "2.svg",
		},
		{
			name:   "Gantt diagram",
			source: "3.mmd",
			result: "3.svg",
		},
		{
			name:   "Class diagram",
			source: "4.mmd",
			result: "4.svg",
		},
		// Waiting for https://github.com/mermaid-js/mermaid/issues/1252
		// {
		// 	name:   "Git graph",
		// 	source: "5.mmd",
		// 	result: "5.svg",
		// },
		{
			name:   "Entity Relationship Diagram",
			source: "6.mmd",
			result: "6.svg",
		},
		{
			name:   "User Journey Diagram",
			source: "7.mmd",
			result: "7.svg",
		},
		{
			name:   "State diagram",
			source: "8.mmd",
			result: "8.svg",
		},
		{
			name:   "Pie chart",
			source: "9.mmd",
			result: "9.svg",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			src, err := ioutil.ReadFile(path.Join(TESTFOLDER, tt.source))
			if err != nil {
				t.Errorf("Can not open source file: %s", err)
			}
			dst, err := ioutil.ReadFile(path.Join(TESTFOLDER, tt.result))
			if err != nil {
				t.Errorf("Can not open target file: %s", err)
			}

			gotResult := Render(string(src))
			la := float64(len(gotResult))
			lb := float64(len(dst))
			d := math.Abs(la-lb) / lb
			if d > 0.1 {
				t.Errorf("%f Render() = %s, want %s", d, gotResult, dst)
			}
		})
	}

}
