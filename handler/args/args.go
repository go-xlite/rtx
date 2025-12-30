package args

import "os"

type args struct {
	// Define argument fields here
}

func newArgs() *args {
	return &args{}
}

func (a *args) HasKey(key string) bool {
	if len(os.Args) <= 1 {
		return false
	}
	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == key {
			return true
		}
	}
	return false
}
func (a *args) GetValue(key string) string {
	if len(os.Args) <= 1 {
		return ""
	}
	for i := 1; i < len(os.Args)-1; i++ {
		if os.Args[i] == key {
			return os.Args[i+1]
		}
	}
	return ""
}

var Args = newArgs()
