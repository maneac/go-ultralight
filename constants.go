package ultralight

// MessageSource enumerates the possible sources for a message
type MessageSource int

// The message sources
const (
	XML MessageSource = iota
	JS
	Network
	ConsoleAPI
	Storage
	AppCache
	Rendering
	CSS
	Security
	ContentBlocker
	Other
)

// MessageLevel enumerates the severity levels of a message
type MessageLevel int

// The severity levels
const (
	Log MessageLevel = iota + 1
	Warning
	Error
	Debug
	Info
)

// Cursor enumerates the different types of cursor to show
type Cursor int

// The cursor types
const (
	Pointer Cursor = iota
	Cross
	Hand
	IBeam
	Wait
	Help
	EastResize
	NorthResize
	NorthEastResize
	NorthWestResize
	SouthResize
	SouthEastResize
	SouthWestResize
	WestResize
	NorthSouthResize
	EastWestResize
	NorthEastSouthWestResize
	NorthWestSouthEastResize
	ColumnResize
	RowResize
	MiddlePanning
	EastPanning
	NorthPanning
	NorthEastPanning
	NorthWestPanning
	SouthPanning
	SouthEastPanning
	SouthWestPanning
	WestPanning
	Move
	VerticalText
	Cell
	ContextMenu
	Alias
	Progress
	NoDrop
	Copy
	CursorNone
	NotAllowed
	ZoomIn
	ZoomOut
	Grab
	Grabbing
	Custom
)

// BitmapFormat enumerates the possible formats for a bitmap
type BitmapFormat int

// The bitmap formats
const (
	A8 BitmapFormat = iota
	RGBA8
)

// // KeyEventType .
// type KeyEventType int

// // KeyEventTypes
// const (
// 	KeyDown KeyEventType = iota
// 	KeyUp
// 	RawKeyDown
// 	Char
// )

// // MouseEventType .
// type MouseEventType int

// // MouseEventTypes
// const (
// 	MouseMoved MouseEventType = iota
// 	MouseDown
// 	MouseUp
// )

// // MouseButton .
// type MouseButton int

// // MouseButtons
// const (
// 	MouseButtonNone MouseButton = iota
// 	Left
// 	Middle
// 	Right
// )

// // ScrollEventType .
// type ScrollEventType int

// // ScrollEventTypes
// const (
// 	ScrollByPixel ScrollEventType = iota
// 	ScrollByPage
// )

// WindowFlag enumerates the window features in a bitwise-OR ('|') friendly manner
type WindowFlag int

// Feature flags for the Window instance
const (
	WindowBorderless  WindowFlag = 1 << 0
	WindowTitled      WindowFlag = 1 << 1
	WindowResizable   WindowFlag = 1 << 2
	WindowMaximizable WindowFlag = 1 << 3
)
