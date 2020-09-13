package squeeze

import "log"

type fn func()

func LogIfError(err error, shouldExit bool, cleanup fn)  {
	if err != nil {
		if shouldExit {
			defer cleanup()
			log.Fatal(err)
		} else {
			log.SetPrefix("Error: ")
			log.Println(err)
			log.SetPrefix("")
		}
	}
}