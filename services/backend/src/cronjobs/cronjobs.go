package cronjobs

import (
	"github.com/dcssoftware/bafoeg-manager/src/configuration"
	ragvectorprocessor "github.com/dcssoftware/bafoeg-manager/src/cronjobs/rag-vector-processor"
	"github.com/go-co-op/gocron/v2"
	"github.com/go-sqlx/sqlx"
	"github.com/minio/minio-go/v7"
)

type Cronjob struct {
	db *sqlx.DB
	s3 *minio.Client

	cron gocron.Scheduler
}

func NewCronjob(db *sqlx.DB, s3 *minio.Client) (*Cronjob, error) {
	cron, err := gocron.NewScheduler()
	if err != nil {
		return nil, err
	}

	return &Cronjob{
		db:   db,
		s3:   s3,
		cron: cron,
	}, nil
}

func (c *Cronjob) StartCronjobs() {
	c.cron.Start()
}

func (c *Cronjob) RegisterRAGVectorProcessor() error {
	ragprocessor := ragvectorprocessor.NewCronjob(c.db, c.s3)

	_, jobErr := c.cron.NewJob(
		gocron.DurationJob(
			configuration.Conjobs.RagVectorProcessor.Interval,
		),
		gocron.NewTask(ragprocessor.RunCronjob),
		gocron.WithIntervalFromCompletion(),
	)

	if jobErr != nil {
		return jobErr
	}

	return nil
}
