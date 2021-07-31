# Holiday
Holiday verification package

# Installation

```go
go get github.com/mpsdantas/holiday
```

# Usage

## Basic use

```go
package main

import (
	"fmt"
	"time"

	"github.com/mpsdantas/holiday"
	"github.com/mpsdantas/holiday/country"
)

func main() {
	holiday.SetHolidays(country.Brazil)

	today := time.Now()

	if holiday.IsHoliday(today) {
		fmt.Println(today.Format("02/01/2006"), "is a holiday")
		return
	}

	fmt.Println(today.Format("02/01/2006"), "is not a holiday")
}

```

## Adding my holidays

```go
package main

import (
	"fmt"
	"time"

	"github.com/mpsdantas/holiday"
)

func main() {
	holiday.SetHolidays(MyHolidays)

	today := time.Date(2021, 01, 01, 12, 0, 0, 0, time.UTC)

	if holiday.IsHoliday(today) {
		fmt.Println(today.Format("02/01/2006"), "is a holiday")
		return
	}

	fmt.Println(today.Format("02/01/2006"), "is not a holiday")
}

func MyHolidays(year int) map[holiday.Date][]*holiday.Holiday {
	result := holiday.NewCalendarManager()

	result.Add(&holiday.Date{Month: time.January, Day: 1}, &holiday.Holiday{
		Name:    "New Year",
		Details: "Universal New Year's Holiday",
	})

	return result.GetHolidays()
}
```

## Checking Holiday Details

```go
package main

import (
	"fmt"
	"time"

	"github.com/mpsdantas/holiday"
)

func main() {
	holiday.SetHolidays(MyHolidays)

	today := time.Date(2021, 01, 01, 12, 0, 0, 0, time.UTC)

	if holidays := holiday.Holidays(today); len(holidays) > 0 {
		fmt.Println(today.Format("02/01/2006"), "is a holiday")
		fmt.Println("Name:", holidays[0].Name)
		fmt.Println("Details:", holidays[0].Details)
		return
	}

	fmt.Println(today.Format("02/01/2006"), "is not a holiday")
}

func MyHolidays(year int) map[holiday.Date][]*holiday.Holiday {
	result := holiday.NewCalendarManager()

	result.Add(&holiday.Date{Month: time.January, Day: 1}, &holiday.Holiday{
		Name:    "New Year",
		Details: "Universal New Year's Holiday",
	})

	return result.GetHolidays()
}
```

# Details

This pack was inspired by the pack: https://github.com/tobiwild/holidays

The package was created because the package above was not receiving updates or approving new pull requests. In it, for example, we are not able to configure a custom calendar without opening a pull request for the package.
