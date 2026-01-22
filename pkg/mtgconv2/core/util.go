package core

import (
	"time"
	"strings"
)

func GetDateStr() string {
	return time.Now().Format("2006-01-02")
}

func ParseOutputFormat(raw string) (OutputFormat, error) {
	switch strings.ToLower(raw) {
	case string(OutputDCK):
		return OutputDCK, nil
	case string(OutputJSON):
		return OutputJSON, nil
	default:
		return "", &UnknownOutputFormat{OutputFormat(raw)}
	}

}

// returns the correct BoardType enum from the available options
func ParseBoardType(raw string) (BoardType, error) {
	switch strings.ToLower(raw) {
	case string(BoardMain):
		return BoardMain, nil
	case string(BoardCommander):
		return BoardCommander, nil
	case string(BoardSideboard):
		return BoardSideboard, nil
	default:
		return "", &UnknownBoardType{raw}
	}
}

// returns the correct FinishType enum from the available options
func ParseFinishType(raw string) (FinishType, error) {
	switch strings.ToLower(raw) {
	case string(FinishDefault):
		return FinishDefault, nil
	case string(FinishFoil):
		return FinishFoil, nil
	case string(FinishNonfoil):
		return FinishNonfoil, nil
	case string(FinishEtched):
		return FinishEtched, nil
	default:
		return "", &UnknownFinishType{raw}
	}
}