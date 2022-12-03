package godep3

import "os"

func TestEnv(name string) string {
	return os.Getenv(name)
}

