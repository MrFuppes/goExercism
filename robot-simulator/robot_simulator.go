package robot

// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// stage 1 robot
// -----------------------------------------------------------------------------

// Step1Robot - a totally stupid robot which isn't even a type
var Step1Robot struct {
	X, Y int
	Dir
}

// Dir represents a direction the robot is heading to; 0 = N, ... 3 = W
type Dir int

const (
	// N - north 0
	N Dir = iota
	// E - east 1
	E
	// S - south 2
	S
	// W - west 3
	W
)

var dirNames = [4]string{"North", "East", "South", "West"}

// String - a stringer for the Dir type
func (d Dir) String() (s string) {
	return dirNames[d]
}

// Action represents the actions a robot can do;
// type Action int

const (
	// A - advance - 0
	A Action = iota
	// L - turn left - 1
	L
	// R - turn right - 2
	R
)

// Advance advances s1robot
func Advance() {
	switch Step1Robot.Dir {
	case N:
		Step1Robot.Y++
	case E:
		Step1Robot.X++
	case S:
		Step1Robot.Y--
	case W:
		Step1Robot.X--
	}
}

// Left turns s1robot left
func Left() {
	Step1Robot.Dir = (Step1Robot.Dir + 3) % 4
}

// Right turns s1robot right
func Right() {
	Step1Robot.Dir = (Step1Robot.Dir + 1) % 4
}

// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// stage 2 robot
// -----------------------------------------------------------------------------

// Command to command the step 2 robot
type Command byte // valid values are 'R', 'L', 'A'

// Action resulting from Command; to be performed by the robot
type Action rune

// RU room unit
type RU int

// Pos x y in room units
type Pos struct{ Easting, Northing RU }

// Rect rectangle defined by min and max Pos
type Rect struct{ Min, Max Pos }

// validPosInSpace - check if a position p is valid within a rectancle r
func (p Pos) validPosInSpace(r Rect) bool {
	return p.Easting >= r.Min.Easting && p.Easting <= r.Max.Easting &&
		p.Northing >= r.Min.Northing && p.Northing <= r.Max.Northing
}

// Step2Robot - a robot that is bit more capable than the step 1 robot
type Step2Robot struct {
	Dir
	Pos
}

// doStuff within the space s
// since this is a method of the step2robot type, this is conceptually different from the step2 test readme:
// the robot ITSELF anticipates the consequence of its actions (advance...) and only performs valid actions
// so it doesn't fall off the grid.
func (r *Step2Robot) doStuff(a Action, s Rect) {
	switch a {
	case 'R':
		r.Dir = (r.Dir + 1) % 4
	case 'L':
		r.Dir = (r.Dir + 3) % 4
	case 'A':
		pos := r.Pos
		switch r.Dir { // change position based on current facing
		case N:
			pos.Northing++
		case S:
			pos.Northing--
		case E:
			pos.Easting++
		case W:
			pos.Easting--
		}
		if pos.validPosInSpace(s) { // only do the move if the new position is valid
			r.Pos = pos
		}
	}
}

// StartRobot - start a robot
func StartRobot(cmd chan Command, act chan Action) {
	for c := range cmd {
		act <- Action(c)
	}
	close(act)
}

// Room - let robot "r" do stuff in a Room defined by "space"
func Room(space Rect, r Step2Robot, act chan Action, rep chan Step2Robot) {
	for a := range act { // perform all actions on the channel
		r.doStuff(a, space)
	}
	rep <- r // return current state of r to rep channel
}

// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// -----------------------------------------------------------------------------
// stage 3 robot
// -----------------------------------------------------------------------------

// Action3 - an struct to attribute actions to a specific robot by its name
type Action3 struct {
	name   string
	script rune
}

// Step3Robot - the final stage of evolution
type Step3Robot struct {
	Name string
	Step2Robot
}

// StartRobot3 - start a robot3
func StartRobot3(name, script string, action chan Action3, log chan string) {
	if name == "" {
		log <- "A robot without a name"
	}
	for _, a := range script {
		action <- Action3{name, a}
	}

	action <- Action3{name, 0}
}

// Room3 - room that allows more operations
func Room3(extent Rect, robots []Step3Robot, action chan Action3, report chan []Step3Robot, log chan string) {
	defer func() { report <- robots }()

	var (
		activeRobots    = make(map[string]int)
		activePositions = make(map[Pos]int)
		nIdle           int
	)

	for i, r := range robots {
		if _, ok := activeRobots[r.Name]; ok {
			log <- "Duplicate robot names"
			return
		}
		activeRobots[r.Name] = i

		robotPos := r.Step2Robot.Pos

		if !robotPos.validPosInSpace(extent) {
			log <- "A robot placed outside of the room"
		}

		if _, ok := activePositions[robotPos]; ok {
			log <- "Robots placed at the same place"
			return
		}
		activePositions[r.Step2Robot.Pos] = i
	}

	for act := range action {
		i, ok := activeRobots[act.name]
		if !ok {
			log <- "Action by unknown robot"
			return
		}

		robotAction := &robots[i].Step2Robot

		switch act.script {
		case 'R':
			robotAction.Dir = (robotAction.Dir + 1) % 4
		case 'L':
			robotAction.Dir = (robotAction.Dir + 3) % 4
		case 'A':
			robotPos := robotAction.Pos
			switch robotAction.Dir {
			case N:
				robotPos.Northing++
			case S:
				robotPos.Northing--
			case E:
				robotPos.Easting++
			case W:
				robotPos.Easting--
			}
			if !robotPos.validPosInSpace(extent) {
				log <- "A robot attempting to advance into a wall"
				continue
			}

			if _, bump := activePositions[robotPos]; bump {
				log <- "A robot attempting to advance into another robot"
				continue
			}
			delete(activePositions, robotAction.Pos)
			activePositions[robotPos] = i
			robotAction.Pos = robotPos

		case 0:
			nIdle++
			if nIdle == len(robots) {
				return
			}

		default:
			log <- "An undefined command in a script"
			return
		}
	}
}
