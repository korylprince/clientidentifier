package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"howett.net/plist"
)

const plistPath = "/Library/Preferences/ManagedInstalls.plist"

func GetClientIdentifier(filename string) (string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var config map[string]interface{}

	d := plist.NewDecoder(f)
	err = d.Decode(&config)
	if err != nil {
		return "", err
	}

	identifier, ok := config["ClientIdentifier"]
	if !ok {
		return "<empty>", nil
	}

	i, ok := identifier.(string)
	if !ok {
		return "", fmt.Errorf("ClientIdentifier is not a valid string")
	}
	if i == "" {
		return "<empty>", nil
	}

	return i, nil
}

func SetClientIdentifier(filename, identifier string) error {
	fi, err := os.Stat(filename)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filename, os.O_RDWR, fi.Mode())
	if err != nil {
		return err
	}
	defer f.Close()

	var config map[string]interface{}

	d := plist.NewDecoder(f)
	err = d.Decode(&config)
	if err != nil {
		return err
	}

	config["ClientIdentifier"] = identifier

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	e := plist.NewEncoderForFormat(f, plist.AutomaticFormat)

	err = e.Encode(&config)
	if err != nil {
		return err
	}

	l, err := f.Seek(0, io.SeekCurrent)
	if err != nil {
		return err
	}

	err = f.Truncate(l)
	if err != nil {
		return err
	}

	if err := (exec.Command("defaults", "read", plistPath)).Run(); err != nil {
		return err
	}

	return nil
}

func ClearClientIdentifier(filename string) error {
	fi, err := os.Stat(filename)
	if err != nil {
		return err
	}

	f, err := os.OpenFile(filename, os.O_RDWR, fi.Mode())
	if err != nil {
		return err
	}
	defer f.Close()

	var config map[string]interface{}

	d := plist.NewDecoder(f)
	err = d.Decode(&config)
	if err != nil {
		return err
	}

	delete(config, "ClientIdentifier")

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		return err
	}

	e := plist.NewEncoderForFormat(f, plist.AutomaticFormat)

	err = e.Encode(&config)
	if err != nil {
		return err
	}

	l, err := f.Seek(0, io.SeekCurrent)
	if err != nil {
		return err
	}

	err = f.Truncate(l)
	if err != nil {
		return err
	}

	if err := (exec.Command("defaults", "read", plistPath)).Run(); err != nil {
		return err
	}

	return nil
}

func elevate() {
	cmd := exec.Command("sudo", os.Args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error getting root privileges")
		fmt.Printf("%#v\n", err)
	}
}

func help() {
	fmt.Println("Usage: clientidentifier [OPTION...] [IDENTIFIER]")
	fmt.Println("\t-h, --help\tDisplay this help message")
	fmt.Println("\t-c,\tClear the ClientIdentifier")
	fmt.Println("Running this program without any options will display the current ClientIdentifier.")
	fmt.Println("The ClientIdentifier will be changed to IDENTIFIER if given.")
}

func main() {
	if os.Geteuid() != 0 {
		elevate()
		return
	}
	fmt.Println("WARNING: the clientidentifier utility is deprecated. You should run `clientidentifier -c` to clear the ClientIdentifier")

	switch len(os.Args) {
	case 1:
		identifier, err := GetClientIdentifier(plistPath)
		if err != nil {
			fmt.Println("Error reading ClientIdentifier:\n\t", err)
		} else {
			fmt.Println("ClientIdentifier:", identifier)
		}
	case 2:
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			help()
		} else if os.Args[1] == "-c" {
			if err := ClearClientIdentifier(plistPath); err != nil {
				fmt.Println("Error clearing ClientIdentifier:\n\t", err)
				return
			}
		} else {
			identifier := os.Args[1]
			err := SetClientIdentifier(plistPath, identifier)
			if err != nil {
				fmt.Println("Error setting ClientIdentifier:\n\t", err)
				return
			}
			fmt.Println("ClientIdentifier:", identifier)
		}
	default:
		help()
	}
}
