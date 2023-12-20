package advent

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
)

type Day20 struct {
	/* --- Day 20: Pulse Propagation ---
	With your help, the Elves manage to find the right parts and fix all of the
	machines. Now, they just need to send the command to boot up the machines
	and get the sand flowing again.

	The machines are far apart and wired together with long cables. The cables
	don't connect to the machines directly, but rather to communication modules
	attached to the machines that perform various initialization tasks and also
	act as communication relays.

	Modules communicate using pulses. Each pulse is either a high pulse or a
	low pulse. When a module sends a pulse, it sends that type of pulse to each
	module in its list of destination modules.

	There are several different types of modules:

	Flip-flop modules (prefix %) are either on or off; they are initially off.
	If a flip-flop module receives a high pulse, it is ignored and nothing
	happens. However, if a flip-flop module receives a low pulse, it flips
	between on and off. If it was off, it turns on and sends a high pulse. If
	it was on, it turns off and sends a low pulse.

	Conjunction modules (prefix &) remember the type of the most recent
	pulse received from each of their connected input modules; they initially
	default to remembering a low pulse for each input. When a pulse is
	received, the conjunction module first updates its memory for that input.
	Then, if it remembers high pulses for all inputs, it sends a low pulse;
	otherwise, it sends a high pulse.

	There is a single broadcast module (named broadcaster). When it receives a
	pulse, it sends the same pulse to all of its destination modules.

	Here at Desert Machine Headquarters, there is a module with a single button
	on it called, aptly, the button module. When you push the button, a single
	low pulse is sent directly to the broadcaster module.

	After pushing the button, you must wait until all pulses have been
	delivered and fully handled before pushing it again. Never push the button
	if modules are still processing pulses.

	Pulses are always processed in the order they are sent. So, if a pulse is
	sent to modules a, b, and c, and then module a processes its pulse and
	sends more pulses, the pulses sent to modules b and c would have to be
	handled first.

	The module configuration (your puzzle input) lists each module. The name of
	the module is preceded by a symbol identifying its type, if any. The name
	is then followed by an arrow and a list of its destination modules. For
	example:

	broadcaster -> a, b, c
	%a -> b
	%b -> c
	%c -> inv
	&inv -> a

	In this module configuration, the broadcaster has three destination modules
	named a, b, and c. Each of these modules is a flip-flop module (as
	indicated by the % prefix). a outputs to b which outputs to c which outputs
	to another module named inv. inv is a conjunction module (as indicated by
	the & prefix) which, because it has only one input, acts like an
	inverter (it sends the opposite of the pulse type it receives); it outputs
	to a.

	By pushing the button once, the following pulses are sent:

	button -low-> broadcaster
	broadcaster -low-> a
	broadcaster -low-> b
	broadcaster -low-> c
	a -high-> b
	b -high-> c
	c -high-> inv
	inv -low-> a
	a -low-> b
	b -low-> c
	c -low-> inv
	inv -high-> a

	After this sequence, the flip-flop modules all end up off, so pushing the
	button again repeats the same sequence.

	Here's a more interesting example:
	broadcaster -> a
	%a -> inv, con
	&inv -> b
	%b -> con
	&con -> output

	This module configuration includes the broadcaster, two flip-flops (named a
	and b), a single-input conjunction module (inv), a multi-input conjunction
	module (con), and an untyped module named output (for testing purposes).
	The multi-input conjunction module con watches the two flip-flop modules
	and, if they're both on, sends a low pulse to the output module.

	Here's what happens if you push the button once:

	button -low-> broadcaster
	broadcaster -low-> a
	a -high-> inv
	a -high-> con
	inv -low-> b
	con -high-> output
	b -high-> con
	con -low-> output

	Both flip-flops turn on and a low pulse is sent to output! However, now
	that both flip-flops are on and con remembers a high pulse from each of its
	two inputs, pushing the button a second time does something different:

	button -low-> broadcaster
	broadcaster -low-> a
	a -low-> inv
	a -low-> con
	inv -high-> b
	con -high-> output

	Flip-flop a turns off! Now, con remembers a low pulse from module a, and so
	it sends only a high pulse to output.

	Push the button a third time:

	button -low-> broadcaster
	broadcaster -low-> a
	a -high-> inv
	a -high-> con
	inv -low-> b
	con -low-> output
	b -low-> con
	con -high-> output

	This time, flip-flop a turns on, then flip-flop b turns off. However,
	before b can turn off, the pulse sent to con is handled first, so it
	briefly remembers all high pulses for its inputs and sends a low pulse to
	output. After that, flip-flop b turns off, which causes con to update its
	state and send a high pulse to output.

	Finally, with a on and b off, push the button a fourth time:

	button -low-> broadcaster
	broadcaster -low-> a
	a -low-> inv
	a -low-> con
	inv -high-> b
	con -high-> output

	This completes the cycle: a turns off, causing con to remember only low
	pulses and restoring all modules to their original states.

	To get the cables warmed up, the Elves have pushed the button 1000 times.
	How many pulses got sent as a result (including the pulses sent by the
	button itself)?

	In the first example, the same thing happens every time the button is
	pushed: 8 low pulses and 4 high pulses are sent. So, after pushing the
	button 1000 times, 8000 low pulses and 4000 high pulses are sent.
	Multiplying these together gives 32000000.

	In the second example, after pushing the button 1000 times, 4250 low pulses
	and 2750 high pulses are sent. Multiplying these together gives 11687500.

	Consult your module configuration; determine the number of low pulses and
	high pulses that would be sent after pushing the button 1000 times, waiting
	for all pulses to be fully handled after each push of the button. What do
	you get if you multiply the total number of low pulses sent by the total
	number of high pulses sent? */

	/* --- Part Two ---
	The final machine responsible for moving the sand down to Island Island has
	a module attached named rx. The machine turns on when a single low pulse is
	sent to rx.

	Reset all modules to their default states. Waiting for all pulses to be
	fully handled after each button press, what is the fewest number of button
	presses required to deliver a single low pulse to the module named rx? */
}

func (d Day20) PartOne(input string) string {
	config, modules := d.parse(input)
	d.init(modules, config)

	low, high := 0, 0
	for n := 0; n < 1000; n++ {
		// fmt.Printf("cycle: %v\n", n+1)
		l, h := d.button(modules)
		low += l
		high += h
	}
	return strconv.Itoa(low * high)
}

func (d Day20) PartTwo(input string) string {
	config, modules := d.parse(input)
	d.init(modules, config)

	n := 0
	for !d.button2(modules) {
		fmt.Printf("iter: %v\r", n)
		n++
	}
	fmt.Println()
	return strconv.Itoa(n)
}

type Mod struct {
	inputs  []string
	outputs []string
	mem     []bool
	typ     int
}

func (m Mod) String() string {
	return fmt.Sprintf("(typ:%v mem:%v)", m.typ, m.mem)
}

const (
	BROADCASTER = iota
	FLIPFLOP
	CONJUNCTION
)

func (m Mod) send() bool {
	switch m.typ {
	case BROADCASTER:
		return false // always send low(false) pulse
	case CONJUNCTION:
		return slices.Contains(m.mem, false) // send low(false) pulse only if remembers high(true) for all inputs
	case FLIPFLOP:
		return m.mem[0] // sends remembered pulse
	default:
		panic("unreachable send")
	}
}

func (m Mod) receive(from string, pulse bool) {
	switch m.typ {
	case CONJUNCTION:
		input := slices.Index(m.inputs, from)
		m.mem[input] = pulse // updates memory for that input with recent pulse
	case FLIPFLOP:
		if pulse == false { // flips between on(true) and off(false) only when receives low(false) pulse
			m.mem[0] = !m.mem[0]
		}
	}
}

func (m Mod) active(pulse bool) bool {
	switch m.typ {
	case FLIPFLOP:
		return !pulse // active only when low(false) pulse
	default: // CONJUNCTION, BROADCASTER
		return true // always active
	}
}

func (Day20) parse(input string) (map[string][]string, map[string]*Mod) {
	re := regexp.MustCompile(`\w+`)

	config, modules := make(map[string][]string), make(map[string]*Mod)
	for _, s := range lines(input) {
		captures := captures(re, s)
		config[captures[0]] = captures[1:]

		switch s[0] {
		case '&':
			modules[captures[0]] = &Mod{typ: CONJUNCTION, outputs: captures[1:]}
		case '%':
			modules[captures[0]] = &Mod{typ: FLIPFLOP, outputs: captures[1:]}
		default:
			modules[captures[0]] = &Mod{typ: BROADCASTER, outputs: captures[1:]}
		}
	}
	return config, modules
}

func (Day20) init(modules map[string]*Mod, config map[string][]string) {
	for name, module := range modules {
		inputs := make([]string, 0)
		for input, outputs := range config {
			if slices.Contains(outputs, name) {
				inputs = append(inputs, input)
			}
		}
		module.inputs = inputs
		if module.typ == CONJUNCTION {
			module.mem = make([]bool, len(inputs))
		} else if module.typ == FLIPFLOP {
			module.mem = []bool{false}
		}
	}
}

func (d Day20) button(modules map[string]*Mod) (int, int) {
	lo, hi := 0, 0
	counter := func(in string, pulse bool, out string) {
		// fmt.Printf("    %s -%v-> %s\n", in, pulse, out)
		if pulse {
			hi++
		} else {
			lo++
		}
	}

	counter("button", false, "broadcaster")
	queue := []string{"broadcaster"}

	for len(queue) > 0 {
		in := pop(&queue, 0)
		imod := modules[in]

		pulse := imod.send()
		for _, out := range imod.outputs {
			counter(in, pulse, out) // count how many low(false) and high(true) pulses were sent

			omod, ok := modules[out]
			if !ok { // sink modules (i.e output, rx)
				continue
			}

			omod.receive(in, pulse) // receive pulse from in, and update memmory
			if omod.active(pulse) { // is module active after update
				queue = append(queue, out)
			}
		}
	}
	return lo, hi
}

func (d Day20) button2(modules map[string]*Mod) bool {
	queue := []string{"broadcaster"}
	rx, rx_low := 0, 0
	for len(queue) > 0 {
		in := pop(&queue, 0)
		imod := modules[in]

		pulse := imod.send()
		for _, out := range imod.outputs {
			omod, ok := modules[out]
			if !ok { // sink modules (i.e output, rx)
				rx += 1
				if pulse == false {
					rx_low += 1
				}
				continue
			}

			omod.receive(in, pulse) // receive pulse from in, and update memmory
			if omod.active(pulse) { // is module active after update
				queue = append(queue, out)
			}
		}
	}
	return rx_low == 1
}
