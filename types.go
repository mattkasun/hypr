package hypr

type Workspace struct {
	ID              int
	Name            string
	Monitor         string
	Windows         int
	HasFullScreen   bool
	LastWindow      string
	LastWindowTitle string
}

type Client struct {
	Address        string
	Mapped         bool
	Hidden         bool
	At             [2]int
	Size           [2]int
	Workspace      Workspace
	Floating       bool
	Monitor        int
	Class          string
	Title          string
	InitialClass   string
	InitialTitle   string
	PID            int
	XWayland       bool
	Pinned         bool
	FullScreen     bool
	FullScreenMode int
	FakeFullScreen bool
	Grouped        []string
	Tags           []string
	Swallowing     string
	FocusHistoryID int
}

type FloatingClient struct {
	Workspace Workspace
	Floating  bool
}
