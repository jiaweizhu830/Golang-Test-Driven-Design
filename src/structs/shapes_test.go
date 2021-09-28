package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		// f means float 64
		// .2 means 2 decimal places
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		// decouple from the concrete type by using interface Shape
		got := shape.Area()

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12.0, 6.0}
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

	// great fit when testing multiple implementations of an interface
	t.Run("table driven tests", func(t *testing.T) {
		// "anonymous struct": slice of struct
		areaTests := []struct {
			name    string
			shape   Shape
			hasArea float64
		}{
			// slice
			// {Rectangle{12, 6}, 72.0},
			// {Circle{10}, 314.1592653589793},
			// {Triangle{12, 6}, 36.0},

			// can also optionally name the struct fields
			{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
			{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
			{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
		}

		for _, tt := range areaTests {
			t.Run(tt.name, func(t *testing.T) {
				got := tt.shape.Area()
				if got != tt.hasArea {
					//%#v print struct
					t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
				}
			})
		}
	})
}

// run specific test case: go test -run TestArea/Rectangle
