package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/extrame/xls"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: %s input", os.Args[0])
	}
	wb, err := xls.Open(os.Args[1], "utf-8")
	if err != nil {
		return err
	}
	ws := wb.GetSheet(0)
	if ws == nil {
		return errors.New("no worksheets")
	}
	header := ws.Row(0)
	if header == nil {
		return errors.New("no header row")
	}
	accepted := -1
	ended := -1
	user := -1
	for i := header.FirstCol(); i < header.LastCol(); i++ {
		switch strings.ToUpper(header.Col(i)) {
		case "ACCEPTED":
			accepted = i
		case "ENDED":
			ended = i
		case "USER":
			user = i
		}
	}
	if accepted == -1 {
		return errors.New("no 'Accepted' column")
	}
	if ended == -1 {
		return errors.New("no 'Ended' column")
	}
	if user == -1 {
		return errors.New("no 'User' column")
	}
	f := os.Stdout
	first := true
	f.WriteString("export const data = [")
	for i := uint16(1); i < ws.MaxRow; i++ {
		row := ws.Row(int(i))
		username := strings.TrimSpace(row.Col(user))
		if username == "" {
			continue
		}
		t := row.Col(accepted)
		if t == "" {
			continue
		}
		startTime := parseTime(t)
		t = row.Col(ended)
		if t == "" {
			continue
		}
		endTime := parseTime(t)
		if startTime == 0 || endTime == 0 || endTime < startTime {
			continue
		}
		if first {
			first = false
		} else {
			f.WriteString(",")
		}
		fmt.Fprintf(f, "[%q,%d,%d]", username, startTime, endTime)
	}
	f.WriteString("];")
	f.Close()
	return nil
}

func parseTime(t string) uint64 {
	p, err := strconv.ParseFloat(t, 64)
	if err != nil {
		return 0
	}
	return uint64(math.Round((p - 25569) * 86400))
}
