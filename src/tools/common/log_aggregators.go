package common

import (
	"types"
	"strings"
	"sync"
)

type baseAggregator struct {
	lines []string
}

func (ba *baseAggregator) String() string {
	return strings.Join(ba.lines, "\n")
}

type NopAggregator struct {
	*baseAggregator
}

func NewNopAggregator() *NopAggregator {
	return &NopAggregator{
		baseAggregator: &baseAggregator{},
	}
}

func (na *NopAggregator) ProcessLine(line string) {}

type HeadAggregator struct {
	sync.Mutex
	*baseAggregator
	Want int
}

func NewHeadAggregator(lines int) *HeadAggregator {
	return &HeadAggregator{
		baseAggregator: &baseAggregator{},
		Want:           lines,
	}
}

func (ha *HeadAggregator) ProcessLine(line string) {
	ha.Lock()
	if ha.Want > len(ha.lines) || ha.Want == 0 {
		ha.lines = append(ha.lines, line)
	}
	ha.Unlock()
}

type TailAggregator struct {
	sync.Mutex
	*baseAggregator
	Want int
}

func NewTailAggregator(lines int) *TailAggregator {
	return &TailAggregator{
		baseAggregator: &baseAggregator{},
		Want:           lines,
	}
}

func (ta *TailAggregator) ProcessLine(line string) {
	ta.Lock()
	if ta.Want > len(ta.lines) {
		ta.lines = append(ta.lines, line)
	} else {
		ta.lines = append(ta.lines[1:], line)
	}
	ta.Unlock()
}

type CombinedAggregator []types.LogAggregator

func (ca CombinedAggregator) String() string {
	var combined []string
	for _, c := range ca {
		combined = append(combined, c.String())
	}
	return strings.Join(combined, "\n")
}

func (ca CombinedAggregator) ProcessLine(line string) {
	for _, c := range ca {
		c.ProcessLine(line)
	}
}

type LabeledAggregator struct {
	ag    types.LogAggregator
	Label string
}

func NewLabeledAggregator(label string, source types.LogAggregator) *LabeledAggregator {
	return &LabeledAggregator{
		ag:    source,
		Label: label,
	}
}

func (la *LabeledAggregator) String() string {
	return la.Label + "\n" + la.ag.String()
}

func (la *LabeledAggregator) ProcessLine(line string) {
	la.ag.ProcessLine(line)
}
