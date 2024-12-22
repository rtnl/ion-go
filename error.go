package ion

// #include <ion.h>
import "C"
import "fmt"

func Check(code C.t_ion_result_code) (err error) {
	switch code {
	case C.RESULT_OK:
		return nil

	case C.RESULT_ERROR:
		return fmt.Errorf("error")

	default:
		return fmt.Errorf("unknown")
	}
}
