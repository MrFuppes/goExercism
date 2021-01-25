package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"time"
)

// Entry for the ledger
type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

// ledgerCnf - a struct to hold config for ledger output based on locale
type ledgerCnf struct {
	dateFmt, currencySym, thousandsSep, decSeparator, ammountFmtNeg, ammountFmtPos string
}

const (
	headerEN = "Date       | Description               | Change\n"
	headerNL = "Datum      | Omschrijving              | Verandering\n"
)

var (
	errNoCurrecy   = errors.New("empty currency")
	errIvdCurrency = errors.New("invalid currency")
	errNoLocale    = errors.New("empty locale")
	errIvdLocale   = errors.New("invalid locale")
)

// FormatLedger formats entries to string
func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	if currency == "" {
		return "", errNoCurrecy
	}
	if locale == "" {
		return "", errNoLocale
	}

	var (
		ledger, ammount string //, dateFmt, currencySym, thousandsSep, decSeparator, ammountFmtNeg, ammountFmtPos string
		config          ledgerCnf
	)

	switch currency {
	case "USD":
		config.currencySym = "$"
	case "EUR":
		config.currencySym = "€"
	default:
		return "", errIvdCurrency
	}

	switch locale {
	case "en-US":
		ledger += headerEN
		config.dateFmt = "01/02/2006"
		config.ammountFmtNeg, config.ammountFmtPos = "(%s%s)", "%s%s "
		config.thousandsSep, config.decSeparator = ",", "."
	case "nl-NL":
		ledger += headerNL
		config.dateFmt = "02-01-2006"
		config.ammountFmtNeg, config.ammountFmtPos = "%s %s-", "%s %s "
		config.thousandsSep, config.decSeparator = ".", ","
	default:
		return "", errIvdLocale
	}

	if len(entries) == 0 {
		return ledger, nil
	}

	// make a copy so that the input is not modified
	var entriesCopy = make([]Entry, len(entries))
	copy(entriesCopy, entries)

	// sort entries by date or change if date equal
	sort.Slice(entriesCopy,
		func(i, j int) bool {
			if entriesCopy[i].Date == entriesCopy[j].Date {
				return entriesCopy[i].Change < entriesCopy[j].Change
			}
			return entriesCopy[i].Date < entriesCopy[j].Date
		})

	for _, e := range entriesCopy {
		// date to local format
		t, err := time.Parse("2006-01-02", e.Date)
		if err != nil {
			return "", err
		}
		ledger += t.Format(config.dateFmt) + " | "

		// add description
		if len(e.Description) > 25 {
			e.Description = e.Description[:22] + "..."
		}
		ledger += fmt.Sprintf("%-25s | ", e.Description)

		// add ammount
		if e.Change < 0 {
			ammount = fmt.Sprintf(config.ammountFmtNeg, config.currencySym,
				formatCents(e.Change*-1, config.thousandsSep, config.decSeparator))
		} else {
			ammount = fmt.Sprintf(config.ammountFmtPos, config.currencySym,
				formatCents(e.Change, config.thousandsSep, config.decSeparator))
		}
		ledger += fmt.Sprintf("%13s\n", ammount)
	}

	return ledger, nil
}

// formatCents - a helper I stole from the community
func formatCents(cents int, thousandsSep, decSeparator string) string {
	c := cents % 100
	cents -= c
	cents /= 100

	f := fmt.Sprintf("%s%02d", decSeparator, c)

	if cents == 0 {
		return "0" + f
	}

	for cents > 0 {
		c = cents % 1000
		cents -= c
		cents /= 1000

		if cents > 0 {
			f = fmt.Sprintf("%s%03d", thousandsSep, c) + f
		} else {
			f = strconv.Itoa(c) + f
		}
	}

	return f
}
