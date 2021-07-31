package holiday

import (
	"time"
)

var (
	m = NewCalendarManager()
)

func NewCalendarManager() *CalendarManager {
	return &CalendarManager{
		f: func(year int) map[Date][]*Holiday {
			return map[Date][]*Holiday{}
		},
		holidays: map[Date][]*Holiday{},
	}
}

type CalendarManager struct {
	f        SetHolidaysFunc
	holidays map[Date][]*Holiday
}

type Holiday struct {
	Name    string
	Details string
}

type Date struct {
	Month time.Month
	Day   int
}

func (c *CalendarManager) ClearHolidays() {
	c.holidays = map[Date][]*Holiday{}
	c.f = func(year int) map[Date][]*Holiday {
		return map[Date][]*Holiday{}
	}
}

func (c *CalendarManager) GetHolidays() map[Date][]*Holiday {
	return c.holidays
}

func (c *CalendarManager) Add(date *Date, holiday *Holiday) {
	if _, ok := c.holidays[*date]; !ok {
		c.holidays = map[Date][]*Holiday{
			*date: {
				holiday,
			},
		}
	}

	c.holidays[*date] = append(c.holidays[*date], holiday)
}

func (c *CalendarManager) AddFromTime(t time.Time, holiday *Holiday) {
	c.Add(&Date{Day: t.Day(), Month: t.Month()}, holiday)
}

func (c *CalendarManager) Holidays(t time.Time) []*Holiday {
	holidays := c.f(t.Year())

	return holidays[Date{
		Day:   t.Day(),
		Month: t.Month(),
	}]
}

func (c *CalendarManager) IsHoliday(t time.Time) bool {
	return len(c.Holidays(t)) > 0
}

type SetHolidaysFunc func(year int) map[Date][]*Holiday

func Add(date *Date, holiday *Holiday) {
	m.Add(date, holiday)
}

func AddFromTime(t time.Time, holiday *Holiday) {
	m.AddFromTime(t, holiday)
}

func SetHolidays(f SetHolidaysFunc) {
	m.f = f
}

func Holidays(t time.Time) []*Holiday {
	return m.Holidays(t)
}

func IsHoliday(t time.Time) bool {
	return m.IsHoliday(t)
}

func ClearHolidays() {
	m.ClearHolidays()
}
