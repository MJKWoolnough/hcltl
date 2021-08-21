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

	"github.com/shakinm/xlsReader/xls"
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
	wb, err := xls.OpenFile(os.Args[1])
	if err != nil {
		return err
	}
	ws, err := wb.GetSheet(0)
	if err != nil {
		return err
	}
	if ws == nil {
		return errors.New("no worksheets")
	}
	header, err := ws.GetRow(0)
	if err != nil {
		return err
	}
	if header == nil {
		return errors.New("no header row")
	}
	accepted := -1
	ended := -1
	user := -1
	logged := -1
	for i := 0; i <= 0x4000; i++ {
		c, err := header.GetCol(i)
		if err != nil {
			return err
		}
		switch strings.ToUpper(c.GetString()) {
		case "ACCEPTED":
			accepted = i
		case "ENDED":
			ended = i
		case "USER":
			user = i
		case "DATE/TIME":
			logged = i
		}
		if accepted != -1 && ended != -1 && user != -1 && logged != -1 {
			break
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
	for i := 1; i < ws.GetNumberRows(); i++ {
		row, err := ws.GetRow(int(i))
		if err != nil {
			return err
		}
		c, err := row.GetCol(user)
		if err != nil {
			return err
		}
		username := strings.TrimSpace(c.GetString())
		if username == "" {
			continue
		}
		c, err = row.GetCol(accepted)
		if err != nil {
			return err
		}
		t := c.GetString()
		if t == "" {
			continue
		}
		startTime := parseTime(t)
		c, err = row.GetCol(ended)
		if err != nil {
			return err
		}
		t = c.GetString()
		if t == "" {
			continue
		}
		endTime := parseTime(t)
		c, err = row.GetCol(logged)
		if err != nil {
			return err
		}
		t = c.GetString()
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
