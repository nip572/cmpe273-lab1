package perimeter
import "testing"
func Test_perimeterRect (t *testing.T) {

	r := Rect{length: 5, width: 3}
	var x Shaper
	x = r 

	Rectperi :=  x.Perimeter()
	if Rectperi != 16 {

		t.Error("expected 16, got ", Rectperi)

	}

}

func Test_perimeterCircle (t *testing.T) {

	c := Circle{r: 5}
	var x Shaper
	x = c 

	CPeri := x.Perimeter()
	if CPeri != 31.41592653589793 {

		t.Error("expected 31.41592653589793, got ", CPeri)

	}

}
