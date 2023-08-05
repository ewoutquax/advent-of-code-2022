package puzzle

type Crate string

type Universe struct {
	stack map[int][]Crate
}
