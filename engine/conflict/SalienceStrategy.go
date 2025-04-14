package conflict

import "github.com/DataWiseHQ/grule-rule-engine/ast"

// The SalienceStrategy strategy returns candidate with the highest salience score. If multiple candidates with identical
// salience score are provided, the salience strategy will resolve the first candidate.
type SalienceStrategy struct {
	selectedRule    *ast.RuleEntry
	maxSalience     int
	totalCandidates int
}

var _ Set = &SalienceStrategy{}

func (s *SalienceStrategy) AddCandidate(rule *ast.RuleEntry) {
	if rule == nil {
		return
	}
	s.totalCandidates++
	if s.selectedRule == nil || s.maxSalience < rule.Salience {
		s.selectedRule = rule
		s.maxSalience = rule.Salience
	}
}

func (s *SalienceStrategy) Resolve() *ast.RuleEntry {
	return s.selectedRule
}

func (s *SalienceStrategy) TotalCandidates() int {
	return s.totalCandidates
}

func (s *SalienceStrategy) Clear() {
	*s = SalienceStrategy{}
}
