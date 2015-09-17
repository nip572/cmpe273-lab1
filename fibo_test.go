package fn
import "testing"
func Test_fibo (t *testing.T){
var v int
v = fibo(3)
if v != 2 {
	t.Error("expected 2, got ", v)
}

}