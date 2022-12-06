package day04

import (
	"strconv"
	"strings"
)

type Strategy = string

type RawAssignment = string
type RawAssignmentPair = string

type AssignmentInterval struct {
	min int
	max int
}

func assignmentIntervalFromRawAssignment(rawAssignment RawAssignment) AssignmentInterval {
	minAndMax := strings.Split(rawAssignment, "-")
	min, _ := strconv.Atoi(minAndMax[0])
	max, _ := strconv.Atoi(minAndMax[1])
	return AssignmentInterval{
		min: min,
		max: max,
	}
}

func (assignmentInterval AssignmentInterval) contains(other AssignmentInterval) bool {
	return assignmentInterval.min <= other.min && other.max <= assignmentInterval.max
}

type AssignmentIntervalPair struct {
	left  AssignmentInterval
	right AssignmentInterval
}

func (assignmentIntervalPair AssignmentIntervalPair) checkForPartialOverlap() bool {
	left := assignmentIntervalPair.left
	right := assignmentIntervalPair.right
	return !(left.max < right.min || right.max < left.min)
}

func (assignmentIntervalPair AssignmentIntervalPair) checkForFullOverlap() bool {
	return assignmentIntervalPair.left.contains(assignmentIntervalPair.right) || assignmentIntervalPair.right.contains(assignmentIntervalPair.left)
}

func assignmentIntervalPairFromRawAssignmentPair(rawAssignmentPair RawAssignmentPair) AssignmentIntervalPair {
	assignmentPairList := strings.Split(rawAssignmentPair, ",")
	return AssignmentIntervalPair{
		left:  assignmentIntervalFromRawAssignment(assignmentPairList[0]),
		right: assignmentIntervalFromRawAssignment(assignmentPairList[1]),
	}
}
