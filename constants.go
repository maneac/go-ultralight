package ultralight

/******************************************************************************
 * WindowFlags
 *****************************************************************************/

// WindowFlag .
type WindowFlag int

//WindowFlags
const (
	WindowBorderless  WindowFlag = 1 << 0
	WindowTitled                 = 1 << 1
	WindowResizable              = 1 << 2
	WindowMaximizable            = 1 << 3
)

/******************************************************************************
 * MessageSource
 *****************************************************************************/

// MessageSource .
type MessageSource int

// MessageSources
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

/******************************************************************************
 * MessageLevel
 *****************************************************************************/

// MessageLevel .
type MessageLevel int

// MessageLevels
const (
	Log MessageLevel = iota + 1
	Warning
	Error
	Debug
	Info
)

/******************************************************************************
 * Cursor
 *****************************************************************************/

// Cursor .
type Cursor int

// Cursors
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

/******************************************************************************
 * BitmapFormat
 *****************************************************************************/

// BitmapFormat .
type BitmapFormat int

// BitmapFormats
const (
	A8 BitmapFormat = iota
	RGBA8
)

/******************************************************************************
 * KeyEventType
 *****************************************************************************/
const (
	KeyDown = iota
	KeyUp
	RawKeyDown
	Char
)

/******************************************************************************
 * MouseEventType
 *****************************************************************************/
const (
	MouseMoved = iota
	MouseDown
	MouseUp
)

/******************************************************************************
 * MouseButton
 *****************************************************************************/
const (
	MouseButtonNone = iota
	Left
	Middle
	Right
)

/******************************************************************************
 * ScrollEventType
 *****************************************************************************/
const (
	ScrollByPixel = iota
	ScrollByPage
)
