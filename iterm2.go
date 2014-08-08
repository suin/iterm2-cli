package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"strings"
)

var colorCommands = []cli.Command{
	tabColor("maroon", 128, 0, 0),
	tabColor("dark-red", 139, 0, 0),
	tabColor("brown", 165, 42, 42),
	tabColor("firebrick", 178, 34, 34),
	tabColor("crimson", 220, 20, 60),
	tabColor("tomato", 255, 99, 71),
	tabColor("coral", 255, 127, 80),
	tabColor("indian-red", 205, 92, 92),
	tabColor("light-coral", 240, 128, 128),
	tabColor("dark-salmon", 233, 150, 122),
	tabColor("salmon", 250, 128, 114),
	tabColor("light-salmon", 255, 160, 122),
	tabColor("orange-red", 255, 69, 0),
	tabColor("dark-orange", 255, 140, 0),
	tabColor("gold", 255, 215, 0),
	tabColor("dark-golden-rod", 184, 134, 11),
	tabColor("golden-rod", 218, 165, 32),
	tabColor("pale-golden-rod", 238, 232, 170),
	tabColor("dark-khaki", 189, 183, 107),
	tabColor("khaki", 240, 230, 140),
	tabColor("olive", 128, 128, 0),
	tabColor("yellow-green", 154, 205, 50),
	tabColor("dark-olive-green", 85, 107, 47),
	tabColor("olive-drab", 107, 142, 35),
	tabColor("lawn-green", 124, 252, 0),
	tabColor("chart-reuse", 127, 255, 0),
	tabColor("green-yellow", 173, 255, 47),
	tabColor("dark-green", 0, 100, 0),
	tabColor("forest-green", 34, 139, 34),
	tabColor("lime", 0, 255, 0),
	tabColor("lime-green", 50, 205, 50),
	tabColor("light-green", 144, 238, 144),
	tabColor("pale-green", 152, 251, 152),
	tabColor("dark-sea-green", 143, 188, 143),
	tabColor("medium-spring-green", 0, 250, 154),
	tabColor("spring-green", 0, 255, 127),
	tabColor("sea-green", 46, 139, 87),
	tabColor("medium-aqua-marine", 102, 205, 170),
	tabColor("medium-sea-green", 60, 179, 113),
	tabColor("light-sea-green", 32, 178, 170),
	tabColor("dark-slate-gray", 47, 79, 79),
	tabColor("teal", 0, 128, 128),
	tabColor("dark-cyan", 0, 139, 139),
	tabColor("aqua", 0, 255, 255),
	tabColor("cyan", 0, 255, 255),
	tabColor("light-cyan", 224, 255, 255),
	tabColor("dark-turquoise", 0, 206, 209),
	tabColor("turquoise", 64, 224, 208),
	tabColor("medium-turquoise", 72, 209, 204),
	tabColor("pale-turquoise", 175, 238, 238),
	tabColor("aqua-marine", 127, 255, 212),
	tabColor("powder-blue", 176, 224, 230),
	tabColor("cadet-blue", 95, 158, 160),
	tabColor("steel-blue", 70, 130, 180),
	tabColor("corn-flower-blue", 100, 149, 237),
	tabColor("deep-sky-blue", 0, 191, 255),
	tabColor("dodger-blue", 30, 144, 255),
	tabColor("light-blue", 173, 216, 230),
	tabColor("sky-blue", 135, 206, 235),
	tabColor("light-sky-blue", 135, 206, 250),
	tabColor("navy", 0, 0, 128),
	tabColor("dark-blue", 0, 0, 139),
	tabColor("medium-blue", 0, 0, 205),
	tabColor("royal-blue", 65, 105, 225),
	tabColor("blue-violet", 138, 43, 226),
	tabColor("indigo", 75, 0, 130),
	tabColor("dark-slate-blue", 72, 61, 139),
	tabColor("slate-blue", 106, 90, 205),
	tabColor("medium-slate-blue", 123, 104, 238),
	tabColor("medium-purple", 147, 112, 219),
	tabColor("dark-magenta", 139, 0, 139),
	tabColor("dark-violet", 148, 0, 211),
	tabColor("dark-orchid", 153, 50, 204),
	tabColor("purple", 128, 0, 128),
	tabColor("thistle", 216, 191, 216),
	tabColor("plum", 221, 160, 221),
	tabColor("violet", 238, 130, 238),
	tabColor("magenta-fuchsia", 255, 0, 255),
	tabColor("orchid", 218, 112, 214),
	tabColor("medium-violet-red", 199, 21, 133),
	tabColor("pale-violet-red", 219, 112, 147),
	tabColor("deep-pink", 255, 20, 147),
	tabColor("hot-pink", 255, 105, 180),
	tabColor("light-pink", 255, 182, 193),
	tabColor("pink", 255, 192, 203),
	tabColor("antique-white", 250, 235, 215),
	tabColor("beige", 245, 245, 220),
	tabColor("bisque", 255, 228, 196),
	tabColor("blanched-almond", 255, 235, 205),
	tabColor("wheat", 245, 222, 179),
	tabColor("corn-silk", 255, 248, 220),
	tabColor("lemon-chiffon", 255, 250, 205),
	tabColor("light-golden-rod-yellow", 250, 250, 210),
	tabColor("light-yellow", 255, 255, 224),
	tabColor("saddle-brown", 139, 69, 19),
	tabColor("sienna", 160, 82, 45),
	tabColor("chocolate", 210, 105, 30),
	tabColor("peru", 205, 133, 63),
	tabColor("sandy-brown", 244, 164, 96),
	tabColor("burly-wood", 222, 184, 135),
	tabColor("tan", 210, 180, 140),
	tabColor("rosy-brown", 188, 143, 143),
	tabColor("moccasin", 255, 228, 181),
	tabColor("navajo-white", 255, 222, 173),
	tabColor("peach-puff", 255, 218, 185),
	tabColor("misty-rose", 255, 228, 225),
	tabColor("lavender-blush", 255, 240, 245),
	tabColor("linen", 250, 240, 230),
	tabColor("old-lace", 253, 245, 230),
	tabColor("papaya-whip", 255, 239, 213),
	tabColor("sea-shell", 255, 245, 238),
	tabColor("mint-cream", 245, 255, 250),
	tabColor("slate-gray", 112, 128, 144),
	tabColor("light-slate-gray", 119, 136, 153),
	tabColor("light-steel-blue", 176, 196, 222),
	tabColor("lavender", 230, 230, 250),
	tabColor("floral-white", 255, 250, 240),
	tabColor("alice-blue", 240, 248, 255),
	tabColor("ghost-white", 248, 248, 255),
	tabColor("honeydew", 240, 255, 240),
	tabColor("ivory", 255, 255, 240),
	tabColor("azure", 240, 255, 255),
	tabColor("snow", 255, 250, 250),
	tabColor("black", 0, 0, 0),
	tabColor("dim-gray-dim-grey", 105, 105, 105),
	tabColor("gray-grey", 128, 128, 128),
	tabColor("dark-gray-dark-grey", 169, 169, 169),
	tabColor("silver", 192, 192, 192),
	tabColor("gainsboro", 220, 220, 220),
	tabColor("white-smoke", 245, 245, 245),
	tabColor("white", 255, 255, 255),
	tabColor("pure-red", 255, 0, 0),
	tabColor("pure-orange", 255, 165, 0),
	tabColor("pure-green", 0, 128, 0),
	tabColor("pure-blue", 0, 0, 255),
	tabColor("pure-yellow", 255, 255, 0),
	tabColor("red", 195, 89, 76),
	tabColor("orange", 219, 154, 88),
	tabColor("green", 65, 174, 76),
	tabColor("blue", 92, 155, 204),
	tabColor("yellow", 240, 240, 0),
}

func main() {
	app := cli.NewApp()
	app.Name = "iterm2"
	app.Usage = "control iTerm2 from cli"
	app.Commands = []cli.Command{
		{
			Name:  "tab",
			Usage: "Control tabs",
			Subcommands: []cli.Command{
				{
					Name:        "color",
					ShortName:   "colour",
					Usage:       "Change tab color",
					Subcommands: colorCommands,
				},
				{
					Name:      "title",
					ShortName: "name",
					Usage:     "Change tab title",
					Action: func(c *cli.Context) {
						if len(c.Args()) > 0 {
							tabTitle(strings.Join(c.Args(), " "))
						}
					},
				},
			},
		},
	}
	app.Run(os.Args)
}

func tabColor(color string, r, g, b int) cli.Command {
	return cli.Command{
		Name: color,
		Action: func(c *cli.Context) {
			fmt.Printf("\033]6;1;bg;red;brightness;%d\a", r)
			fmt.Printf("\033]6;1;bg;green;brightness;%d\a", g)
			fmt.Printf("\033]6;1;bg;blue;brightness;%d\a", b)
		},
	}
}

func tabTitle(title string) {
	fmt.Printf("\033]0;%s\007", title)
}
