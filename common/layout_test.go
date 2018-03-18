package common

import (
	"reflect"
	"testing"
)

func TestLayoutIndices(t *testing.T) {
	layout := NewLayout(0)
	if !reflect.DeepEqual(layout.Indices(), []int{0}) {
		t.Errorf("Indices should be %+v but got %+v", []int{0}, layout.Indices())
	}

	layout = layout.SplitTop(1)
	layout = layout.SplitLeft(2)
	layout = layout.SplitBottom(3)
	layout = layout.SplitRight(4)
	if !reflect.DeepEqual(layout.Indices(), []int{2, 3, 4, 1, 0}) {
		t.Errorf("Indices should be %+v but got %+v", []int{0}, layout.Indices())
	}
}

func TestLayoutSplitCloseReplace(t *testing.T) {
	layout := NewLayout(0)

	layout = layout.SplitTop(1)
	layout = layout.SplitLeft(2)
	layout = layout.SplitBottom(3)
	layout = layout.SplitRight(4)

	expected := LayoutHorizontal{
		Top: LayoutVertical{
			Left: LayoutHorizontal{
				Top: LayoutWindow{Index: 2, Active: false},
				Bottom: LayoutVertical{
					Left:  LayoutWindow{Index: 3, Active: false},
					Right: LayoutWindow{Index: 4, Active: true},
				},
			},
			Right: LayoutWindow{Index: 1, Active: false},
		},
		Bottom: LayoutWindow{Index: 0, Active: false},
	}

	if !reflect.DeepEqual(layout, expected) {
		t.Errorf("layout should be %+v but got %+v", expected, layout)
	}

	w, h := layout.Count()
	if w != 3 {
		t.Errorf("layout width be %d but got %d", 3, w)
	}
	if h != 3 {
		t.Errorf("layout height be %d but got %d", 3, h)
	}

	layout = layout.Close()

	expected = LayoutHorizontal{
		Top: LayoutVertical{
			Left: LayoutHorizontal{
				Top:    LayoutWindow{Index: 2, Active: false},
				Bottom: LayoutWindow{Index: 3, Active: true},
			},
			Right: LayoutWindow{Index: 1, Active: false},
		},
		Bottom: LayoutWindow{Index: 0, Active: false},
	}

	if !reflect.DeepEqual(layout, expected) {
		t.Errorf("layout should be %+v but got %+v", expected, layout)
	}

	w, h = layout.Count()
	if w != 2 {
		t.Errorf("layout width be %d but got %d", 3, w)
	}
	if h != 3 {
		t.Errorf("layout height be %d but got %d", 3, h)
	}

	layout = layout.Replace(5)

	expected = LayoutHorizontal{
		Top: LayoutVertical{
			Left: LayoutHorizontal{
				Top:    LayoutWindow{Index: 2, Active: false},
				Bottom: LayoutWindow{Index: 5, Active: true},
			},
			Right: LayoutWindow{Index: 1, Active: false},
		},
		Bottom: LayoutWindow{Index: 0, Active: false},
	}

	if !reflect.DeepEqual(layout, expected) {
		t.Errorf("layout should be %+v but got %+v", expected, layout)
	}
}