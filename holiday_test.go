package holiday_test

import (
	"testing"
	"time"

	"github.com/mpsdantas/holiday"
	"github.com/mpsdantas/holiday/country"
	assert2 "github.com/stretchr/testify/assert"
)

func TestSetHolidays(t *testing.T) {
	t.Run("simple test", func(t *testing.T) {
		defer holiday.ClearHolidays()

		holiday.SetHolidays(func(year int) map[holiday.Date][]*holiday.Holiday {
			result := holiday.NewCalendarManager()

			result.AddFromTime(time.Now(), &holiday.Holiday{
				Name: "Holiday Test",
			})

			return result.GetHolidays()
		})

		result := holiday.IsHoliday(time.Now())

		assert := assert2.New(t)

		assert.Equal(true, result)
	})

	t.Run("complex test", func(t *testing.T) {
		defer holiday.ClearHolidays()

		holiday.SetHolidays(country.Brazil)

		year := time.Now().Year() + 1

		assert := assert2.New(t)

		for day, details := range country.Brazil(year) {
			result := holiday.IsHoliday(time.Date(year, day.Month, day.Day, 12, 0, 0, 0, time.UTC))

			assert.NotEmpty(details)
			assert.Equal(true, result)
		}
	})
}

func TestHolidays(t *testing.T) {
	t.Run("getting christmas holiday", func(t *testing.T) {
		defer holiday.ClearHolidays()

		holiday.SetHolidays(func(year int) map[holiday.Date][]*holiday.Holiday {
			result := holiday.NewCalendarManager()

			result.Add(&holiday.Date{Month: time.December, Day: 25}, &holiday.Holiday{
				Name: "Christmas holiday",
			})

			return result.GetHolidays()
		})

		h := holiday.Holidays(time.Date(2021, time.December, 25, 12, 0, 0, 0, time.UTC))

		assert := assert2.New(t)

		assert.Equal(h[0].Name, "Christmas holiday")
	})
}
