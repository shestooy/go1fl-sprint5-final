package daysteps

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Shestooy/tracker/internal/personaldata"
	"github.com/Shestooy/tracker/internal/spentenergy"
)

const actionInfoTemplate = `Количество шагов: %d.
Дистанция составила %0.2f км.
Вы сожгли %0.2f ккал.
`

type DaySteps struct {
	Steps    int
	Duration time.Duration
	personaldata.Personal
}

func (ds *DaySteps) Parse(data string) error {
	tuples := strings.Split(data, ",")

	if len(tuples) != 2 {
		return fmt.Errorf("bad format of input - expected 2 arg, got %d from %s", len(tuples), data)
	}

	steps, err := strconv.Atoi(tuples[0])
	if err != nil {
		return err
	}
	if steps < 1 {
		return fmt.Errorf("steps must be greater than zero - got %d", steps)
	}

	duration, err := time.ParseDuration(tuples[1])
	if err != nil {
		return err
	}

	if duration.Hours() <= 0 {
		return fmt.Errorf("duration hours must be greater than zero - got %0.2f", duration.Hours())
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

func (ds *DaySteps) ActionInfo() (string, error) {
	calories, err := spentenergy.WalkingSpentCalories(
		ds.Steps,
		ds.Weight,
		ds.Height,
		ds.Duration,
	)

	if err != nil {
		return "", err
	}

	distance := spentenergy.Distance(
		ds.Steps,
		ds.Height,
	)

	result := fmt.Sprintf(
		actionInfoTemplate,
		ds.Steps,
		distance,
		calories,
	)

	return result, nil
}
