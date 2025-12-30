package models

type PolicyDecision struct {
    Allowed bool
    Reason  string
}

func NewPolicyDecision(allowed bool, reason string) *PolicyDecision {
	return &PolicyDecision{
		Allowed: allowed,
		Reason: reason,
	}
}