package goappendfile

import "os"

func appendFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	defer file.Close()

	file.WriteString(message)

	return nil
}
