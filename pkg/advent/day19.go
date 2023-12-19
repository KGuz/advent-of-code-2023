package advent

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Day19 struct {
	/* --- Day 19: Aplenty ---
	The Elves of Gear Island are thankful for your help and send you on your
	way. They even have a hang glider that someone stole from Desert Island;
	since you're already going that direction, it would help them a lot if you
	would use it to get down there and return it to them.

	As you reach the bottom of the relentless avalanche of machine parts, you
	discover that they're already forming a formidable heap. Don't worry,
	though - a group of Elves is already here organizing the parts, and they
	have a system.

	To start, each part is rated in each of four categories:

	- x: Extremely cool looking
	- m: Musical (it makes a noise when you hit it)
	- a: Aerodynamic
	- s: Shiny

	Then, each part is sent through a series of workflows that will ultimately
	accept or reject the part. Each workflow has a name and contains a list of
	rules; each rule specifies a condition and where to send the part if the
	condition is true. The first rule that matches the part being considered is
	applied immediately, and the part moves on to the destination described by
	the rule. (The last rule in each workflow has no condition and always
	applies if reached.)

	Consider the workflow ex{x>10:one,m<20:two,a>30:R,A}. This
	workflow is named ex and contains four rules. If workflow ex were
	considering a specific part, it would perform the following steps in order:

	- Rule "x>10:one": If the part's x is more than 10, send the part to the
	  workflow named one.
	- Rule "m<20:two": Otherwise, if the part's m is less than 20, send the
	  part to the workflow named two.
	- Rule "a>30:R": Otherwise, if the part's a is more than 30, the part is
	  immediately rejected (R).
	- Rule "A": Otherwise, because no other rules matched the part, the part is
	  immediately accepted (A).

	If a part is sent to another workflow, it immediately switches to the start
	of that workflow instead and never returns. If a part is accepted (sent to
	A) or rejected (sent to R), the part immediately stops any further
	processing.

	The system works, but it's not keeping up with the torrent of weird metal
	shapes. The Elves ask if you can help sort a few parts and give you the
	list of workflows and some part ratings (your puzzle input). For example:

	px{a2090:A,rfg}
	pv{a>1716:R,A}
	lnx{m>1548:A,A}
	rfg{s2440:R,A}
	qs{s>3448:A,lnx}
	qkq{x<1416:A,crn}
	crn{x>2662:A,R}
	in{s<1351:px,qqz}
	qqz{s>2770:qs,m<1801:hdj,R}
	gd{a>3333:R,R}
	hdj{m>838:A,pv}

	{x=787,m=2655,a=1222,s=2876}
	{x=1679,m=44,a=2067,s=496}
	{x=2036,m=264,a=79,s=2244}
	{x=2461,m=1339,a=466,s=291}
	{x=2127,m=1623,a=2188,s=1013}

	The workflows are listed first, followed by a blank line, then the ratings
	of the parts the Elves would like you to sort. All parts begin in the
	workflow named in. In this example, the five listed parts go through the
	following workflows:

	- {x=787,m=2655,a=1222,s=2876}: in -> qqz -> qs -> lnx -> A
	- {x=1679,m=44,a=2067,s=496}: in -> px -> rfg -> gd -> R
	- {x=2036,m=264,a=79,s=2244}: in -> qqz -> hdj -> pv -> A
	- {x=2461,m=1339,a=466,s=291}: in -> px -> qkq -> crn -> R
	- {x=2127,m=1623,a=2188,s=1013}: in -> px -> rfg -> A

	Ultimately, three parts are accepted. Adding up the x, m, a, and s rating
	for each of the accepted parts gives 7540 for the part with x=787, 4623 for
	the part with x=2036, and 6951 for the part with x=2127. Adding all of the
	ratings for all of the accepted parts gives the sum total of 19114.

	Sort through all of the parts you've been given; what do you get if you add
	together all of the rating numbers for all of the parts that ultimately get
	accepted? */

	/* --- Part Two ---
	Even with your help, the sorting process still isn't fast enough.

	One of the Elves comes up with a new plan: rather than sort parts
	individually through all of these workflows, maybe you can figure out in
	advance which combinations of ratings will be accepted or rejected.

	Each of the four ratings (x, m, a, s) can have an integer value ranging
	from a minimum of 1 to a maximum of 4000. Of all possible distinct
	combinations of ratings, your job is to figure out which ones will be
	accepted.

	In the above example, there are 167409079868000 distinct combinations of
	ratings that will be accepted.

	Consider only your list of workflows; the list of part ratings that the
	Elves wanted you to sort is no longer relevant. How many distinct
	combinations of ratings will be accepted by the Elves' workflows? */
}

func (d Day19) PartOne(input string) string {
	workflows, parts := d.parse(input)
	sum := 0
	for _, mp := range parts {
		if d.sort(mp, workflows) {
			sum += mp.x + mp.m + mp.a + mp.s
		}
	}
	return strconv.Itoa(sum)
}

func (d Day19) PartTwo(input string) string {
	workflows, _ := d.parse(input)
	answer := d.combinations(workflows)
	return strconv.Itoa(answer)
}

type MachinePart struct {
	x, m, a, s int
}

func (mp MachinePart) accepts(r Rule) bool {
	switch r.cat {
	case 'x':
		if r.op == '<' {
			return mp.x < r.val
		} else {
			return mp.x > r.val
		}
	case 'm':
		if r.op == '<' {
			return mp.m < r.val
		} else {
			return mp.m > r.val
		}
	case 'a':
		if r.op == '<' {
			return mp.a < r.val
		} else {
			return mp.a > r.val
		}
	case 's':
		if r.op == '<' {
			return mp.s < r.val
		} else {
			return mp.s > r.val
		}
	default:
		return true
	}
}

type Rule struct {
	cat    byte
	op     byte
	val    int
	action string
}

func (r Rule) String() string {
	return fmt.Sprintf("{if %c %c %d then %s}", r.cat, r.op, r.val, r.action)
}

func MakeWorkflow(s string) []Rule {
	workflow := make([]Rule, 0)
	for _, str := range strings.Split(s, ",") {
		then := strings.Index(str, ":")
		if then == -1 {
			workflow = append(workflow, Rule{'0', '=', 0, str[:len(str)-1]})
		} else {
			value, action := parse(str[2:then]), str[then+1:]
			category, operation := str[0], str[1]
			workflow = append(workflow, Rule{category, operation, value, action})
		}
	}
	return workflow
}

func (Day19) parse(input string) (map[string][]Rule, []MachinePart) {
	workflows := make(map[string][]Rule)
	parts := make([]MachinePart, 0)

	lines, n := lines(input), 0
	for ; len(lines[n]) > 0; n++ {
		line := lines[n]
		s := strings.Index(line, "{")
		workflows[line[:s]] = MakeWorkflow(line[s+1:])
	}

	re := regexp.MustCompile(`\d+`)
	for n++; n < len(lines); n++ {
		nums := transform(captures(re, lines[n]), parse)
		new := MachinePart{nums[0], nums[1], nums[2], nums[3]}
		parts = append(parts, new)
	}
	return workflows, parts
}

func (Day19) sort(mp MachinePart, workflows map[string][]Rule) bool {
	curr := "in"
	for curr != "R" && curr != "A" {
		w := workflows[curr]
		for n := 0; n < len(w); n++ {
			if mp.accepts(w[n]) {
				curr = w[n].action
				break
			}
		}
	}
	return curr == "A"
}

type State struct {
	start [4]int
	end   [4]int
	node  string
}

func (s State) String() string {
	return fmt.Sprintf("{x:(%d %d) m:(%d %d) a:(%d %d) s:(%d %d)}", s.start[0], s.end[0], s.start[1], s.end[1], s.start[2], s.end[2], s.start[3], s.end[3])
}

func (Day19) combinations(workflows map[string][]Rule) int {
	MIN, MAX := 1, 4000
	queue := []State{{
		start: [4]int{MIN, MIN, MIN, MIN},
		end:   [4]int{MAX, MAX, MAX, MAX},
		node:  "in",
	}}

	combinations := make([]State, 0)

	for len(queue) > 0 {
		state := pop(&queue, len(queue)-1)
		fmt.Printf("Visiting workflow: %v with initial range: %v\n", state.node, state)
		if state.node == "R" {
			fmt.Printf("Rejected range: %v\n", state)
			continue // reject chain
		}
		if state.node == "A" {
			fmt.Printf("Added range: %v\n", state)
			combinations = append(combinations, state)
			continue // accept chain
		}

		for _, rule := range workflows[state.node] {
			new := State{start: state.start, end: state.end, node: rule.action}
			fmt.Printf("Considering rule: %v\n", rule)
			if rule.op == '=' {
				queue = append(queue, new)
				fmt.Printf("\tadded:%v,\n\tthats it\n", new)
				break
			} else {
				idx := 0 // 'x'
				if rule.cat == 'm' {
					idx = 1
				} else if rule.cat == 'a' {
					idx = 2
				} else if rule.cat == 's' {
					idx = 3
				}

				if rule.op == '<' && new.start[idx] < rule.val { // s...v - otherwise whole range is invalid
					if new.end[idx] < rule.val { // s...e...v
						queue = append(queue, new)
						fmt.Printf("\tadded:%v,\n\tthats it\n", new)
						break

					} else { // s...v...e
						new.end[idx] = rule.val - 1
						queue = append(queue, new)
						state.start[idx] = rule.val
						fmt.Printf("\tadded:%v,\n\tremaining:%v\n", new, state)
						continue
					}
				} else if rule.op == '>' && new.end[idx] > rule.val { // v...e - otherwise whole range is invalid
					if new.start[idx] > rule.val { // v...s...e
						queue = append(queue, new)
						fmt.Printf("\tadded:%v,\n\tthats it\n", new)
						break
					} else { // s...v...e
						new.start[idx] = rule.val + 1
						queue = append(queue, new)
						state.end[idx] = rule.val
						fmt.Printf("\tadded:%v,\n\tremaining:%v\n", new, state)
						continue
					}
				}
				fmt.Printf("should not happen")
			}
		}
	}

	sum := 0
	for _, comb := range combinations {
		total := 1
		for n := 0; n < 4; n++ {
			partRange := comb.end[n] - comb.start[n] + 1
			total *= partRange
		}
		sum += total
	}
	fmt.Printf("combs: %v\n", combinations)
	return sum
}
