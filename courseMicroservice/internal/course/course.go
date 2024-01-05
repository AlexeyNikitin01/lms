package course

type Lecture struct {
	ID   int
	Text string
}

type Question struct {
	ID     int
	Text   string
	Answer string
}

type Course struct {
	ID        int
	Progress  int
	Name      string
	Lectures  []Lecture // исправить массив на стринг?
	Questions []Question
}
