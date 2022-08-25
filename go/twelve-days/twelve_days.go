package twelve

func Verse(i int) string {
	type Lyric struct {
		Day  string
		gift string
	}

	Lyrics := map[int]Lyric{
		1:  {"first", "a Partridge in a Pear Tree"},
		2:  {"second", "two Turtle Doves"},
		3:  {"third", "three French Hens"},
		4:  {"fourth", "four Calling Birds"},
		5:  {"fifth", "five Gold Rings"},
		6:  {"sixth", "six Geese-a-Laying"},
		7:  {"seventh", "seven Swans-a-Swimming"},
		8:  {"eighth", "eight Maids-a-Milking"},
		9:  {"ninth", "nine Ladies Dancing"},
		10: {"tenth", "ten Lords-a-Leaping"},
		11: {"eleventh", "eleven Pipers Piping"},
		12: {"twelfth", "twelve Drummers Drumming"},
	}
	out := "On the " + Lyrics[i].Day + " day of Christmas my true love gave to me: "
	for day := i; day >= 2; day-- {
		out += Lyrics[day].gift + ", "
	}

	if i != 1 {
		out += "and "
	}
	out += Lyrics[1].gift + "."
	return out
}

func Song() string {
	out := ""
	for i := 1; i <= 12; i++ {
		out = out + Verse(i)
		if i != 12 {
			out += "\n"
		}
	}
	return out
}
