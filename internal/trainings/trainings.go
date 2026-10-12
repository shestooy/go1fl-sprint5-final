package trainings

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Shestooy/tracker/internal/personaldata"
	"github.com/Shestooy/tracker/internal/spentenergy"
)

const actionInfoTemplate = `Тип тренировки: %s
Длительность: %0.2f ч.
Дистанция: %0.2f км.
Скорость: %0.2f км/ч
Сожгли калорий: %0.2f
`

type Training struct {
	Steps        int
	TrainingType string
	Duration     time.Duration
	personaldata.Personal
}

func (t *Training) Parse(data string) error {
	tuples := strings.Split(data, ",")
	if len(tuples) != 3 {
		return fmt.Errorf("bad format of input - expected 3 arg, got %d from %s", len(tuples), data)
	}

	steps, err := strconv.Atoi(tuples[0])
	if err != nil {
		return err
	}

	if steps < 1 {
		return fmt.Errorf("steps must be greater than zero - got %d", steps)
	}

	duration, err := time.ParseDuration(tuples[2])
	if err != nil {
		return err
	}

	if duration.Hours() <= 0 {
		return fmt.Errorf("duration hours must be greater than zero - got %0.2f", duration.Hours())
	}
	t.Steps = steps
	t.TrainingType = tuples[1]
	t.Duration = duration

	return nil
}

func (t *Training) ActionInfo() (string, error) {
	var (
		calories float64
		err      error
	)

	switch t.TrainingType {
	case "Бег":
		calories, err = spentenergy.RunningSpentCalories(
			t.Steps,
			t.Weight,
			t.Height,
			t.Duration,
		)
	case "Ходьба":
		calories, err = spentenergy.WalkingSpentCalories(
			t.Steps,
			t.Weight,
			t.Height,
			t.Duration,
		)
	default:
		// by the way, this err-message must be like it is now, but it's bad practice for debug scenarios
		return "", errors.New("неизвестный тип тренировки")
	}

	if err != nil {
		return "", err
	}

	dist := spentenergy.Distance(t.Steps, t.Height)
	speed := spentenergy.MeanSpeed(t.Steps, t.Height, t.Duration)

	return fmt.Sprintf(
		actionInfoTemplate,
		t.TrainingType,
		t.Duration.Hours(),
		dist,
		speed,
		calories,
	), nil
}
