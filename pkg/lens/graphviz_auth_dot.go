package lens

import (
	"errors"
	"fmt"
	"strings"
)

// GraphvizAuthTreeSpec holds the deterministic diagnostic result produced by EvaluateGraphvizAuthTree.
type GraphvizAuthTreeSpec struct {
	TargetKey string
	PrimaryOp int
	SecondOpN int
	ScoreRate int
	StatusTag string
	DetailMsg string
}

// EvaluateGraphvizAuthTree evaluates XDR diagnostic bounds for targetKey and returns a GraphvizAuthTreeSpec.
func EvaluateGraphvizAuthTree(targetKey string, primaryOp int, secondOpN int, limitUnit int) (GraphvizAuthTreeSpec, error) {
	trimmed := strings.TrimSpace(targetKey)
	if trimmed == "" {
		return GraphvizAuthTreeSpec{}, errors.New("lens: targetKey must not be empty")
	}
	if primaryOp < 0 {
		primaryOp = 0
	}
	if secondOpN < 0 {
		secondOpN = 0
	}
	if limitUnit <= 0 {
		limitUnit = 100
	}

	scoreRate := ((primaryOp + secondOpN) * 10000) / limitUnit
	if scoreRate > 10000 {
		scoreRate = 10000
	}

	statusTag := "nominal"
	if scoreRate >= 8500 {
		statusTag = "critical"
	} else if scoreRate >= 5000 {
		statusTag = "elevated"
	}

	return GraphvizAuthTreeSpec{
		TargetKey: trimmed,
		PrimaryOp: primaryOp,
		SecondOpN: secondOpN,
		ScoreRate: scoreRate,
		StatusTag: statusTag,
		DetailMsg: fmt.Sprintf("%s:%s:%d", trimmed, statusTag, scoreRate),
	}, nil
}
