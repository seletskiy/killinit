package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"syscall"

	"github.com/docopt/docopt-go"
	"github.com/kovetskiy/lorg"
)

var version = "1.0"

var usage = `killinit - start specified program and listen for kill signals on fifo.

Tool starts specified command and creates FIFO file which is used to listen for
kill messages for this command.

To send signal text message with signal identifier followed by newline should
be sent to specified FIFO listen file.

Both numeric (e.g. 2) and string signal identifiers (INT) are supported.

Tool will preserve original exit code from specified command, if command
will be killed by unhandled signal, tool will return 127+signal exit code.

Usage:
  killinit -h | --help
  killinit [options] --listen <file> -- <command>...

Options:
  -h --help        Show this help.
  --listen <file>  Create FIFO and continuously listen on kill messages.
  --debug          Print debug messages into stderr.
`

func main() {
	args, err := docopt.ParseArgs(usage, nil, version)
	if err != nil {
		panic(err)
	}

	var (
		listen   = args["--listen"].(string)
		command  = args["<command>"].([]string)
		debug, _ = args["--debug"].(bool)
	)

	var status int

	defer func() { os.Exit(status) }()

	exit := func(code int) {
		status = code

		runtime.Goexit()
	}

	lorg.Exiter = exit

	lorg.SetPrefix("{killinit}")

	if debug {
		lorg.SetLevel(lorg.LevelDebug)
	}

	err = syscall.Mknod(listen, syscall.S_IFIFO|0666, 0)
	if err != nil {
		lorg.Fatalf(
			"unable to create fifo to listen on: %s: %q",
			err,
			listen,
		)
	}

	defer os.Remove(listen)

	cmd := exec.Command(command[0], command[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err = cmd.Start()
	if err != nil {
		lorg.Fatalf(
			"unable to start specified command: %q",
			command,
		)
	}

	lorg.Debugf("command %q started with PID %d", command, cmd.Process.Pid)

	go handle(listen, cmd.Process)

	err = cmd.Wait()
	if err != nil {
		if err, ok := err.(*exec.ExitError); ok {
			if status, ok := err.Sys().(syscall.WaitStatus); ok {
				if status.ExitStatus() >= 0 {
					lorg.Debugf(
						"command %q exited with code %d",
						command,
						status.ExitStatus(),
					)

					exit(status.ExitStatus())
				}

				if status.Signaled() {
					exit(127 + int(status.Signal()))
				}
			}
		}

		lorg.Debugf(
			"command %q finished with error: %s",
			command,
			err,
		)
	} else {
		lorg.Debugf("command %q finished successfully", command)
	}
}

func handle(listen string, process *os.Process) {
	for {
		file, err := os.Open(listen)
		if err != nil {
			lorg.Fatalf(
				"unable to open fifo file: %s: %q",
				err,
				listen,
			)
		}

		reader := bufio.NewReader(file)

		for {
			message, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}

				lorg.Fatalf(
					"unable to read message from fifo: %s",
					err,
				)
			}

			signal, err := parse(message)
			if err != nil {
				lorg.Errorf(
					"unable to process message %q: %s",
					message,
					err,
				)
			}

			lorg.Debugf(
				"sending signal %d (%s) to PID %d",
				signal,
				signal.String(),
				process.Pid,
			)

			err = process.Signal(signal)
			if err != nil {
				lorg.Errorf(
					"unable to send signal %d (%s) to child process %d: %s",
					signal,
					signal.String(),
					process.Pid,
					err,
				)
			}
		}
	}
}

func parse(message string) (syscall.Signal, error) {
	message = strings.TrimSpace(message)

	number, err := strconv.ParseInt(message, 10, 0)
	if err != nil {
		if number, ok := signals[message]; ok {
			return syscall.Signal(number), nil
		}

		return 0, fmt.Errorf(
			"unknown signal specifier: %q",
			message,
		)
	}

	return syscall.Signal(number), nil
}
