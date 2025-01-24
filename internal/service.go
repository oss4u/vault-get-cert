package internal

import (
	"fmt"
	"github.com/go-co-op/gocron/v2"
)

func RunServer(config *Config) error {
	s, err := gocron.NewScheduler()
	if err != nil {
		return fmt.Errorf("failed to create new scheduler: %w", err)
	}

	_, err = s.NewJob(
		gocron.CronJob(
			"1 * * * *",
			false,
		),
		gocron.NewTask(func(cfg *Config) error {
			return GetCertificates(cfg)
		}, config),
	)
	if err != nil {
		return fmt.Errorf("failed to create new job: %w", err)
	}

	s.Start()

	err = s.Shutdown()
	if err != nil {
		return fmt.Errorf("failed to shutdown scheduler: %w", err)
	}
	return nil
}
