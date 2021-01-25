package school

import "sort"

// Grade - has a number and can have students
type Grade struct {
	number   int
	students []string
}

// School - a list of grades
type School []Grade

// New builds a new school
func New() *School {
	return new(School)
}

// Add adds a student's name to a certain grade
func (s *School) Add(name string, grade int) {
	students := s.Grade(grade)
	if students == nil { // add the grade if it does not exist
		newGrade := Grade{grade, []string{name}}
		*s = append(*s, newGrade)
	} else {
		newStudents := append(students, name)
		sort.Strings(newStudents)
		i := s.getGradeIdx(grade)
		(*s)[i].students = newStudents
	}
}

// updateGrade - a helper needed because the exercise requires Grade() to return a slice instead of a pointer...
func (s *School) getGradeIdx(grade int) int {
	for i, g := range *s {
		if g.number == grade {
			return i
		}
	}
	return -1
}

// Grade returns a sorted list of all students in a certain grade of the school
func (s *School) Grade(i int) []string {
	for _, g := range *s {
		if g.number == i {
			return g.students
		}
	}
	return nil
}

// Enrollment returns a sorted list of all grades in the school
func (s *School) Enrollment() []Grade {
	if len(*s) > 1 { // sort by grade if there is more than one
		sort.Slice(*s, func(i, j int) bool { return (*s)[i].number < (*s)[j].number })
	}
	return *s
}
