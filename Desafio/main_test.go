package main

import (
	"testing"
)

func TestDependsAA(t *testing.T) {
	rs := NewRuleSet()
	rs.AddDep("A", "A")
	resultado := rs.IsCoherent()
	if resultado != false {
		t.Errorf("TestDependsAA: esperava false, mas obteve %v", resultado)
	}
}

func TestDependsAB_BA(t *testing.T) {
	rs := NewRuleSet()
	rs.AddDep("A", "B")
	rs.AddDep("B", "A")
	resultado := rs.IsCoherent()
	if resultado != false {
		t.Errorf("TestDependsAB_BA: esperava false mas obteve %v", resultado)
	}
}

func TestExclusiveAB(t *testing.T) {
	rs := NewRuleSet()
	rs.AddDep("A", "B")
	rs.AddConflict("A", "B")
	resultado := rs.IsCoherent()
	if resultado != false {
		t.Errorf("TestExclusiveAB: esperava false, mas obteve %v", resultado)
	}
}

func TestExclusiveAB_BC(t *testing.T) {
	rs := NewRuleSet()
	rs.AddDep("A", "B")
	rs.AddDep("B", "C")
	rs.AddConflict("A", "C")
	resultado := rs.IsCoherent()
	if resultado != false {
		t.Errorf("TestExclusiveAB_BC: esperava false, mas obteve %v", resultado)
	}
}

func TestDeepDeps(t *testing.T) {
	rs := NewRuleSet()
	rs.AddDep("A", "B")
	rs.AddDep("B", "C")
	rs.AddDep("C", "D")
	rs.AddDep("D", "E")
	rs.AddDep("A", "F")
	rs.AddConflict("E", "F")
	resultado := rs.IsCoherent()
	if resultado != false {
		t.Errorf("TestDeepDeps: esperava false, mas obteve %v", resultado)
	}
}
