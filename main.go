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
	alarms := make([]string, 0)
	alarmIDs := make(map[string]uint64)
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
	cols := map[string]int{
		"ACCEPTED":    -1,
		"ENDED":       -1,
		"USER":        -1,
		"DATE/TIME":   -1,
		"IN/OUT":      -1,
		"ALARM DESC.": -1,
	}
	var done int
	for i := 0; i <= 0x4000; i++ {
		c, err := header.GetCol(i)
		if err != nil {
			return err
		}
		col := strings.ToUpper(c.GetString())
		if n, ok := cols[col]; ok && n == -1 {
			cols[col] = i
			done++
		}
		if len(cols) == done {
			break
		}
	}
	if len(cols) != done {
		return errors.New("cannot find all required headers")
	}
	f, err := os.Create(os.Args[1] + ".html")
	if err != nil {
		return err
	}
	first := true
	f.WriteString(start)
	maxRows := ws.GetNumberRows()
	for i := 1; i < maxRows; i++ {
		row, err := ws.GetRow(int(i))
		if err != nil {
			return err
		}
		data := make(map[string]string, len(cols))
		for k, col := range cols {
			cell, err := row.GetCol(col)
			if err != nil {
				return err
			}
			d := strings.TrimSpace(cell.GetString())
			data[k] = d
		}
		username := data["USER"]
		if username == "" {
			continue
		}
		uname := strings.ToUpper(username)
		userID, ok := userIDs[uname]
		if !ok {
			userID = uint64(len(userIDs))
			userIDs[uname] = userID
			users = append(users, username)
		}
		if io := strings.ToUpper(data["IN/OUT"]); io == "OUT" {
			logTime = 0
		} else if io == "NONE" {
			continue
		}
		startTime := parseTime(data["ACCEPTED"])
		endTime := parseTime(data["ENDED"])
		logTime := parseTime(data["DATE/TIME"])
		if startTime == 0 || endTime == 0 || logTime == 0 || endTime < startTime || startTime < logTime {
			continue
		}
		if first {
			first = false
		} else {
			f.WriteString(",")
		}
		fmt.Fprintf(f, "[%d,%d,%d,%d", userID, startTime, endTime, logTime)
		ad := data["ALARM DESC."]
		if ad != "" {
			adu := strings.ToUpper(ad)
			adID, ok := alarmIDs[adu]
			if !ok {
				adID = uint64(len(alarmIDs))
				alarmIDs[adu] = adID
				alarms = append(alarms, ad)
			}
			fmt.Fprintf(f, ",%d]", adID)
		} else {
			fmt.Fprint(f, "]")
		}
	}
	f.WriteString(mid)
	for n, username := range users {
		if n > 0 {
			f.WriteString(",")
		}
		fmt.Fprintf(f, "%q", username)
	}
	f.WriteString(mid2)
	for n, alarm := range alarms {
		if n > 0 {
			f.WriteString(",")
		}
		fmt.Fprintf(f, "%q", alarm)
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
