package sleep
import "time"
import "fmt"
import "testing"

func Test_sleep(q *testing.T) {

	t1 := time.Now().Second()
	fmt.Println(t1)
	Sleep(5)
	t2 := time.Now().Second()
	fmt.Println(t2)

	if t2-t1 == 5 {
		q.Log("Pass")
		fmt.Println("Done")


	} else {

		q.Log("Fail")

	}

}

