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

type StringRepo struct {
	strings []string
	ids     map[string]int64
}

func NewStringRepo() StringRepo {
	return StringRepo{
		ids: make(map[string]int64),
	}
}

func (s *StringRepo) GetID(str string) int64 {
	if str == "" {
		return -1
	}
	ustr := strings.ToUpper(str)
	id, ok := s.ids[ustr]
	if !ok {
		id = int64(len(s.ids))
		s.ids[ustr] = id
		s.strings = append(s.strings, str)
	}
	return id
}

func (s *StringRepo) WriteTo(w *os.File) (int64, error) {
	var count int64
	for n, str := range s.strings {
		if n > 0 {
			if _, err := w.WriteString(","); err != nil {
				return count, err
			}
		}
		m, err := fmt.Fprintf(w, "%q", str)
		count += int64(m)
		if err != nil {
			return count, err
		}
	}
	return count, nil
}

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

	users := NewStringRepo()
	alarms := NewStringRepo()
	lines := NewStringRepo()
	reasons := NewStringRepo()

	wb, err := xls.OpenFile(os.Args[1])
	if err != nil {
		return fmt.Errorf("error reading xls file: %w", err)
	}

	ws, err := wb.GetSheet(0)
	if err != nil {
		return fmt.Errorf("error getting first sheet: %w", err)
	} else if ws == nil {
		return errors.New("no worksheets")
	}

	header, err := ws.GetRow(0)
	if err != nil {
		return fmt.Errorf("error getting first row: %w", err)
	} else if header == nil {
		return errors.New("no header row")
	}

	cols := map[string]int{
		"ACCEPTED":    -1,
		"ENDED":       -1,
		"USER":        -1,
		"DATE/TIME":   -1,
		"IN/OUT":      -1,
		"ALARM DESC.": -1,
		"LINE":        -1,
		"CALL REASON": -1,
	}

	var done int

	for i := 0; i <= 0x4000; i++ {
		c, err := header.GetCol(i)
		if err != nil {
			return fmt.Errorf("error reading column %d: %w", i+1, err)
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
		return fmt.Errorf("error creating output file: %w", err)
	}

	if _, err = f.WriteString(start); err != nil {
		return fmt.Errorf("error writing JS 'start': %w", err)
	}

	first := true
	maxRows := ws.GetNumberRows()

	for i := 1; i < maxRows; i++ {
		row, err := ws.GetRow(int(i))
		if err != nil {
			return fmt.Errorf("error reading row %d: %w", i+1, err)
		}

		data := make(map[string]string, len(cols))

		for k, col := range cols {
			cell, err := row.GetCol(col)
			if err != nil {
				return fmt.Errorf("error reading cell %d,%d: %w", col+1, i+1, err)
			}
			d := strings.TrimSpace(cell.GetString())
			data[k] = d
		}

		userID := users.GetID(data["USER"])
		lineID := lines.GetID(data["LINE"])
		reasonID := reasons.GetID(data["CALL REASON"])
		startTime := parseTime(data["ACCEPTED"])
		endTime := parseTime(data["ENDED"])
		logTime := parseTime(data["DATE/TIME"])

		if userID < 0 || lineID < 0 || reasonID < 0 || startTime == 0 || endTime == 0 || logTime == 0 || endTime < startTime || startTime < logTime {
			continue
		}

		if io := strings.ToUpper(data["IN/OUT"]); io == "OUT" {
			logTime = 0
		} else if io == "NONE" {
			continue
		}

		if first {
			first = false
		} else if _, err := f.WriteString(","); err != nil {
			return fmt.Errorf("error writing separator: %w", err)
		}

		if _, err := fmt.Fprintf(f, "[%d,%d,%d,%d,%d,%d", userID, startTime, endTime, logTime, lineID, reasonID); err != nil {
			return fmt.Errorf("error writing row data: %w", err)
		} else if aid := alarms.GetID(data["ALARM DESC."]); aid < 0 {
			if _, err := fmt.Fprint(f, "]"); err != nil {
				return fmt.Errorf("error writing close bracker: %w", err)
			}
		} else if _, err := fmt.Fprintf(f, ",%d]", aid); err != nil {
			return fmt.Errorf("error writing alarm desc: %w", err)
		}
	}

	if _, err := f.WriteString(mid); err != nil {
		return fmt.Errorf("error writing JS 'mid' data: %w", err)
	} else if _, err := users.WriteTo(f); err != nil {
		return fmt.Errorf("error writing user data: %w", err)
	} else if _, err := f.WriteString(mid2); err != nil {
		return fmt.Errorf("error writing JS 'mid2': %w", err)
	} else if _, err := alarms.WriteTo(f); err != nil {
		return fmt.Errorf("error writing alarm data: %w", err)
	} else if _, err := f.WriteString(mid3); err != nil {
		return fmt.Errorf("error writing JS 'mid3': %w", err)
	} else if _, err := lines.WriteTo(f); err != nil {
		return fmt.Errorf("error writing lines data: %w", err)
	} else if _, err := f.WriteString(mid4); err != nil {
		return fmt.Errorf("error writing JS 'mid4': %w", err)
	} else if _, err := reasons.WriteTo(f); err != nil {
		return fmt.Errorf("error writing reasons data: %w", err)
	} else if _, err := f.WriteString(end); err != nil {
		return fmt.Errorf("error writing JS 'end': %w", err)
	}

	return f.Close()
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
