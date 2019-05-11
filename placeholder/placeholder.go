package placeholder

type placeholders [5]string

var zero = placeholders{
	"███",
	"█ █",
	"█ █",
	"█ █",
	"███",
}

var one = placeholders{
	"██ ",
	" █ ",
	" █ ",
	" █ ",
	"███",
}
var two = placeholders{
	"███",
	"  █",
	"███",
	"█  ",
	"███",
}

var three = placeholders{
	"███",
	"  █",
	"███",
	"  █",
	"███",
}

var four = placeholders{
	"█ █",
	"█ █",
	"███",
	"  █",
	"  █",
}

var five = placeholders{
	"███",
	"█  ",
	"███",
	"  █",
	"███",
}

var six = placeholders{
	"███",
	"█  ",
	"███",
	"█ █",
	"███",
}

var seven = placeholders{
	"███",
	"  █",
	"  █",
	"  █",
	"  █",
}

var eight = placeholders{
	"███",
	"█ █",
	"███",
	"█ █",
	"███",
}

var nine = placeholders{
	"███",
	"█ █",
	"███",
	"  █",
	"███",
}

var Colon = placeholders{
	"   ",
	" ░ ",
	"   ",
	" ░ ",
	"   ",
}

var Digits = [...]placeholders{
	zero, one, two, three, four, five, six, seven, eight, nine,
}
