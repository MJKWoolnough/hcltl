// HCLTL converts a formatted XLS file into a timeline
package main // import "vimagination.zapto.org/hcltl"

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

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
	users := make([]string, 0)
	userIDs := make(map[string]uint64)
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
	logged := -1
	for i := header.FirstCol(); i < header.LastCol(); i++ {
		switch strings.ToUpper(header.Col(i)) {
		case "ACCEPTED":
			accepted = i
		case "ENDED":
			ended = i
		case "USER":
			user = i
		case "DATE/TIME":
			logged = i
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
	if logged == -1 {
		return errors.New("no 'Date/Time' column")
	}
	f, err := os.Create(os.Args[1] + ".html")
	if err != nil {
		return err
	}
	first := true
	f.WriteString(start)
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
		t = row.Col(logged)
		if t == "" {
			continue
		}
		logTime := parseTime(t)
		if startTime == 0 || endTime == 0 || logTime == 0 || endTime < startTime || startTime < logTime {
			continue
		}
		if first {
			first = false
		} else {
			f.WriteString(",")
		}
		uname := strings.ToUpper(username)
		userID, ok := userIDs[uname]
		if !ok {
			userID = uint64(len(userIDs))
			userIDs[uname] = userID
			users = append(users, username)
		}
		fmt.Fprintf(f, "[%d,%d,%d,%d]", userID, startTime, endTime, logTime)
	}
	f.WriteString(mid)
	for n, username := range users {
		if n > 0 {
			f.WriteString(",")
		}
		fmt.Fprintf(f, "%q", username)
	}
	f.WriteString(end)
	f.Close()
	return nil
}

func parseTime(v string) int64 {
	p, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return 0
	}
	u := int64(math.Round((p - 25569) * 86400))
	t := time.Unix(u, 0)
	y, mo, d := t.Date()
	h, mi, s := t.Clock()
	_, offset := time.Date(y, mo, d, h, mi, s, 0, time.Local).Zone()
	return u - int64(offset)
}
