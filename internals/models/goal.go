package models

import (
	"time"
)

type Investment struct {
	Weight  float64 `json:"weight"`
	AssetID int     `json:"asset_id"`
}

type GoalAttribute struct {
	Name                   string       `json:"name"`
	NameWithoutSuffix      string       `json:"name_without_suffix"`
	Nav                    float64      `json:"nav"`
	CreatedAt              time.Time    `json:"created_at"`
	Timeframe              int          `json:"timeframe"`
	Deposited              float64      `json:"deposited"`
	Hidden                 bool         `json:"hidden"`
	Profit                 float64      `json:"profit"`
	Investments            []Investment `json:"investments"`
	PublicLink             *string      `json:"public_link"`
	ParamId                int          `json:"param_id"`
	GoalType               string       `json:"goal_type"`
	TranslatedGoalType     string       `json:"translated_goal_type"`
	Regime                 *string      `json:"regime"`
	Completed              bool         `json:"completed"`
	HasAnyWithdrawals      bool         `json:"has_any_withdrawals"`
	EligibleForDeposits    bool         `json:"eligible_for_deposits"`
	EligibleForInternalMLT bool         `json:"eligible_for_internal_mlt"`
	MonthlyDeposit         float64      `json:"monthly_deposit"`
	SimulatedDeposit       float64      `json:"simulated_deposit"`
	FundsSource            *string      `json:"funds_source"`
	FundsSourceDescription *string      `json:"funds_source_description"`
	NotNetDeposited        float64      `json:"not_net_deposited"`
	Withdrawn              float64      `json:"withdrawn"`
	GroupGoalId            *string      `json:"group_goal_id"`
}

type Goal struct {
	Id         string        `json:"id"`
	Type       string        `json:"type"`
	Attributes GoalAttribute `json:"attributes"`
}

type GoalsResponse struct {
	Data []Goal `json:"data"`
}
