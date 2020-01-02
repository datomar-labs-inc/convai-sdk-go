package byob

// Session is used to keep track of information pertaining to a user
type Session struct {
	Flaggable
	Stack   Stack  `json:"stack" mapstructure:"stack" msgpack:"stack"`
	Version string `json:"ver" mapstructure:"ver" msgpack:"ver"`
}

type Frame struct {
	Module int64 `json:"m" mapstructure:"m" msgpack:"m"`
	Node   int64 `json:"n" mapstructure:"n" msgpack:"n"`
}

type Stack struct {
	Frames []Frame `json:"frames" mapstructure:"frames" msgpack:"frames"`
}

func (s *Stack) Push(frame Frame) {
	s.Frames = append(s.Frames, frame)
}

func (s *Stack) Pop() Frame {
	if len(s.Frames) == 1 {
		panic("cannot be frameless")
	}

	frame := s.GetCurrent()
	s.Frames = s.Frames[:len(s.Frames)-1]
	return frame
}

func (s *Stack) GetCurrent() Frame {
	return s.Frames[len(s.Frames)-1]
}

func (s *Stack) SetCurrent(node int64) {
	frame := s.GetCurrent()
	s.Frames = s.Frames[:len(s.Frames)-1]
	frame.Node = node
	s.Push(frame)
}

func (s *Stack) IsInitialized() bool {
	return len(s.Frames) >= 1
}

func (s *Stack) IsOnMainGraph() bool {
	return len(s.Frames) == 1
}
