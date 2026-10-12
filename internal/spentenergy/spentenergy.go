package spentenergy

import (
	"fmt"
	"time"
)

const (
	mInKm                      = 1000 // количество метров в километре.
	minInH                     = 60   // количество минут в часе.
	stepLengthCoefficient      = 0.45 // коэффициент для расчета длины шага на основе роста.
	walkingCaloriesCoefficient = 0.5  // коэффициент для расчета калорий при ходьбе.
)

func WalkingSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	spentCalories, err := RunningSpentCalories(steps, weight, height, duration)
	if err != nil {
		return 0.0, err
	}

	return spentCalories * walkingCaloriesCoefficient, nil
}

func RunningSpentCalories(steps int, weight, height float64, duration time.Duration) (float64, error) {
	if steps <= 0 || weight <= 0 || height <= 0 || duration.Hours() <= 0 {
		return 0.0, fmt.Errorf(
			"the input values must be positive. Input values: [steps %d, weight %0.2f, height %0.2f, hours %0.2f]",
			steps, weight, height, duration.Hours())
	}

	speed := MeanSpeed(steps, height, duration)

	return duration.Minutes() * speed * weight / minInH, nil
}

func MeanSpeed(steps int, height float64, duration time.Duration) float64 {
	if steps < 1 || height <= 0 || duration.Hours() <= 0 {
		return 0
	}

	dist := Distance(steps, height)

	return dist / duration.Hours()
}

func Distance(steps int, height float64) float64 {
	return (float64(steps) * height * stepLengthCoefficient) / mInKm
}
