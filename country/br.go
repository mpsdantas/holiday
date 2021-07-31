package country

import (
	"time"

	"github.com/mpsdantas/holiday"
	"github.com/vjeantet/eastertime"
)

func Brazil(year int) map[holiday.Date][]*holiday.Holiday {
	result := holiday.NewCalendarManager()

	result.Add(&holiday.Date{Month: time.January, Day: 1}, &holiday.Holiday{
		Name: "Ano Novo",
	})
	result.Add(&holiday.Date{Month: time.April, Day: 21}, &holiday.Holiday{
		Name: "Dia de Tiradentes",
	})
	result.Add(&holiday.Date{Month: time.May, Day: 1}, &holiday.Holiday{
		Name: "Dia do Trabalho",
	})
	result.Add(&holiday.Date{Month: time.September, Day: 7}, &holiday.Holiday{
		Name: "Independência do Brasil",
	})
	result.Add(&holiday.Date{Month: time.October, Day: 12}, &holiday.Holiday{
		Name: "Nossa Senhora Aparecida",
	})
	result.Add(&holiday.Date{Month: time.November, Day: 2}, &holiday.Holiday{
		Name: "Dia de Finados",
	})
	result.Add(&holiday.Date{Month: time.November, Day: 15}, &holiday.Holiday{
		Name: "Proclamação da República",
	})
	result.Add(&holiday.Date{Month: time.December, Day: 25}, &holiday.Holiday{
		Name: "Natal",
	})

	easter, _ := eastertime.CatholicByYear(year)

	result.AddFromTime(easter.AddDate(0, 0, 0).UTC(), &holiday.Holiday{
		Name: "Domingo de Páscoa",
	})
	result.AddFromTime(easter.AddDate(0, 0, -2).UTC(), &holiday.Holiday{
		Name: "Sexta-Feira da Paixão",
	})
	result.AddFromTime(easter.AddDate(0, 0, -47).UTC(), &holiday.Holiday{
		Name: "Carnaval",
	})
	result.AddFromTime(easter.AddDate(0, 0, 60).UTC(), &holiday.Holiday{
		Name: "Corpus Christi",
	})

	return result.GetHolidays()
}
