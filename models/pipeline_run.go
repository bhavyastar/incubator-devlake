package models

import (
	"time"

	"github.com/merico-dev/lake/models/common"
	"gorm.io/datatypes"
)

type Pipeline struct {
	common.Model
	Name           string         `json:"name" gorm:"index"`
	PipelinePlanId uint64         `json:"pipeline_plan_id"`
	Tasks          datatypes.JSON `json:"tasks"`
	TotalTasks     int            `json:"totalTasks"`
	FinishedTasks  int            `json:"finishedTasks"`
	BeganAt        *time.Time     `json:"beganAt"`
	FinishedAt     *time.Time     `json:"finishedAt" gorm:"index"`
	Status         string         `json:"status"`
	Message        string         `json:"message"`
	SpentSeconds   int            `json:"spentSeconds"`
}

// We use a 2D array because the request body must be an array of a set of tasks
// to be executed concurrently, while each set is to be executed sequentially.
type NewPipeline struct {
	Name           string       `json:"name"`
	PipelinePlanId uint64       `json:"pipeline_plan_id"`
	Tasks          [][]*NewTask `json:"tasks"`
	CronConfig     *CronConfig
	Enable         string
}

type CronConfig struct {
	Type   string //weekly, monthly, interval
	Day    int
	Hour   int
	Minute int
}
