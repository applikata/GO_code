package src

import (
	"bytes"
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func GetFiles(log_path string, prefix string) []string {
	dt := time.Now()
	n := 2

	if strings.Contains(log_path, "COLLECTOR") {
		n = 3
	}
	if dt.Minute() < 5 {
		n *= 2
	}
	files := make([]string, n)

	for i := 0; i < n; i++ {

		command := "ls -t " + log_path + " | " + "head" + " -n " + strconv.Itoa(i+1) + " | " + "tail " + "-n " + "1"

		cmd := exec.Command("bash", "-c", command) //for logging
		var errb bytes.Buffer
		cmd.Stderr = &errb
		err := cmd.Run()
		logErr := errb.String()

		if len(logErr) > 0 {
			newErr := errors.New(logErr)
			logger("FIND_FILES_OUT: ", newErr)
		}

		if err != nil {
			logger("ERR_CMD2", err)
			continue
		}

		out, err := exec.Command("bash", "-c", command).Output()

		if err != nil {
			logger("ERR_CMD1", err)
			continue
		}

		files[i] = string(out)

	}

	return files

}
