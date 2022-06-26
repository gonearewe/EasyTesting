package handlers

import (
	"github.com/gonearewe/EasyTesting/dao"
	"github.com/gonearewe/EasyTesting/models"
	"github.com/rs/zerolog/log"
)

func InitTaskConsumers() {
	const consumerNum = 8
	for i := 0; i < consumerNum; i++ {
		go func(id int) {
			log.Info().Msgf("task consumer %d started", id)
			defer func() {
				if e := recover(); e != nil {
					if err, ok := e.(error); ok {
						log.Fatal().Stack().Err(err).Msgf("task consumer %d panicked", id)
					} else {
						log.Fatal().Stack().Interface("panic value", e).Msgf("task consumer %d panicked", id)
					}
				}
			}()

			for {
				t := <-taskQueue
				dao.SubmitMyAnswers(t.examSessionId, t.mcqs, t.maqs, t.bfqs, t.tfqs, t.crqs, t.cqs)
			}
		}(i)
	}
}

var taskQueue = make(chan task, 1000)

type task struct {
	examSessionId int
	mcqs          []*models.McqAnswer
	maqs          []*models.MaqAnswer
	bfqs          []*models.BfqAnswer
	tfqs          []*models.TfqAnswer
	crqs          []*models.CrqAnswer
	cqs           []*models.CqAnswer
}

func newTask() task {
	return task{
		examSessionId: 0,
		mcqs:          make([]*models.McqAnswer, 0),
		maqs:          make([]*models.MaqAnswer, 0),
		bfqs:          make([]*models.BfqAnswer, 0),
		tfqs:          make([]*models.TfqAnswer, 0),
		crqs:          make([]*models.CrqAnswer, 0),
		cqs:           make([]*models.CqAnswer, 0),
	}
}
