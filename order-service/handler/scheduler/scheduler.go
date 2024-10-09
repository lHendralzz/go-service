package scheduler

import (
	"fmt"
	"go-service/usecase"
	"strconv"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/robfig/cron/v3"
)

type Scheduler interface {
	Run()
}

type scheduler struct {
	svc    *usecase.Usecase
	logger *log.Logger
	opt    Option
}

type Option struct {
	CronExpiredOrder     string `env:"CRON_EXPIRED_ORDER_SCHEDULER"`
	ExpiredOrderDuration string `env:"EXPIRED_ORDER_DURATION"`
}

var once = &sync.Once{}

func Init(service *usecase.Usecase, logger *log.Logger, opt Option) Scheduler {
	var r *scheduler
	once.Do(func() {
		r = &scheduler{
			svc:    service,
			logger: logger,
			opt:    opt,
		}
	})
	return r
}

const (
	defaultDuration = 60
)

func (s *scheduler) Run() {
	c := cron.New()

	duration, err := strconv.Atoi(s.opt.ExpiredOrderDuration)
	if err != nil {
		s.logger.Error(fmt.Sprintf("Failed convert expired Order Duration %+v", err))
		duration = defaultDuration
	}

	// Schedule a job to run every minute
	c.AddFunc(s.opt.CronExpiredOrder, func() {
		err := s.svc.Order.ReleaseOrderFromCheckoutStatus(time.Duration(duration) * time.Second)
		if err != nil {
			s.logger.Error(err)
		}
		s.logger.Info("check order expired")
	})
	c.Start()
	s.logger.Info("Success running Cron")
}
