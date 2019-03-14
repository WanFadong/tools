package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/wanfadong/tools/service"
	"os"
	"strconv"
	"time"
)

var (
	DefaultDayLayout = "2006-01-02"
	DefaultTimeLayout = "2006-01-02 15:04:05"
	PreciseTimeLayout = "2006-01-02 15:04:05.000000000"
)

func init() {
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {
	app := cli.NewApp()
	app.Name = "one"
	app.Usage = "个人小工具"

	app.Commands = []cli.Command {
		{
			Name: "time",
			Usage: "parse time to timestamp",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "format, f",
					Usage: "time format",
				},
				cli.IntFlag{
					Name: "last",
					Usage: "last days num",
				},
			},
			Action:  CommandTime,
			Subcommands: []cli.Command {
				{
					Name: "last",
					Usage: "return last days timestamp",
					Action: CommandTimeLast,
				},
			},
		},
		{
			Name: "timestamp",
			Usage: "parse timestamp to time",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name: "nanosecond, ns",
					Usage: "is nanosecond",
				},
				cli.BoolFlag{
					Name: "second, s",
					Usage: "is second",
				},
			},
			Action: CommandTimestamp,
		},
	}

	if err := app.Run(os.Args); err != nil {
		logrus.Fatalln(err)
	}
}

const NotFound = ""

func CommandTime(c *cli.Context) (err error) {
	timeStr := c.Args().First()
	format := c.String("format")
	var t time.Time
	if format == NotFound {
		t, err = service.ParseTimeString(timeStr)
	} else {
		t, err = service.ParseTimeStringWithLayout(timeStr, format)
	}
	if err != nil {
		return
	}
	fmt.Println(t.Unix())
	return
}

func CommandTimeLast(c *cli.Context) (err error) {
	arg := c.Args().First()
	days, err := strconv.Atoi(arg)
	if err != nil {
		return
	}
	if days <= 0 {
		days = 7
	}

	ts := service.LastDays(days)
	for _, t := range ts {
		fmt.Printf("%v\t%v\n", t.Format(DefaultDayLayout), t.Unix())
	}
	return
}

func CommandTimestamp(c *cli.Context) (err error) {
	arg := c.Args().First()
	timestamp, err := strconv.ParseInt(arg, 10, 64)
	if err != nil {
		return
	}
	var t time.Time
	switch {
	case c.IsSet("ns"):
		logrus.Debugln("parse ns", timestamp)
		t = service.ParseTimestamp(timestamp)
		fmt.Println(t.Format(PreciseTimeLayout))
	case c.IsSet("s"):
		logrus.Debugln("parse s", timestamp)
		t = service.ParseTimestamp(timestamp * 1e9)
		fmt.Println(t.Format(DefaultTimeLayout))
	default:
		logrus.Debugln("parse adaptive", timestamp)
		t, err = service.ParseTimestampAdaptive(timestamp)
		if err != nil {
			return
		}
		fmt.Println(t.Format(PreciseTimeLayout))
	}
	return
}
