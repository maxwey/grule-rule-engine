package conflict

import "github.com/DataWiseHQ/grule-rule-engine/ast"

type Set interface {
	// AddCandidate adds the provided rule to the conflict set as a possible candidate. Providing a nil rule is a no-op.
	// As this operation is called more frequently than `Resolve()`, where possible, compute should be shifted to be
	// performed at resolution time.
	AddCandidate(rule *ast.RuleEntry)

	// Clear resets the conflict set to its initial zero value. This may be beneficial over creating a new conflict set
	// by allowing the resolution strategy to retain any allocated buffers for future use.
	Clear()

	// Resolve evaluates and returns the selected rule from the conflict set strategy.
	// If the conflict set is empty, the return value may be nil.
	// Resolve must always return a non-nil value if TotalCandidates is greater than 0.
	Resolve() *ast.RuleEntry

	// TotalCandidates returns the number of non-nil candidates that were provided to the conflict set.
	TotalCandidates() int
}
