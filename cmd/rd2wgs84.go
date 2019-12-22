package main

import (
	"fmt"
	"github.com/arnobroekhof/rd2wgs84"
	"github.com/mitchellh/cli"
	"os"
	"strconv"
)

func main() {
	c := cli.NewCLI("rd2wgs84", "0.2")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"rd": func() (command cli.Command, e error) {
			return &RD{}, nil
		},
		"wgs84": func() (command cli.Command, e error) {
			return &WGS84{}, nil
		},
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(exitStatus)
}

type RD struct{}

func (*RD) Help() string {
	return "Convert to RD: rd2wgs84 rd <lat> <lon>"
}

func (r *RD) Run(args []string) int {
	if len(args) != 2 {
		fmt.Println(r.Help())
		return 1
	}

	cor, err := convertStringCoordinatesToFloat(args)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	wgs := rd2wgs84.NewWSG84Coordinates(cor[0], cor[1])

	rd := wgs.ToRD()
	fmt.Printf("X: %v Y: %v\n", rd.X, rd.Y)
	return 0
}

func (r *RD) Synopsis() string {
	return r.Help()
}

type WGS84 struct{}

func (*WGS84) Help() string {
	return "Convert to WGS84: rd2wgs84 wgs84 <X> <Y>"
}

func (r *WGS84) Run(args []string) int {
	if len(args) != 2 {
		fmt.Println(r.Help())
		return 1
	}

	cor, err := convertStringCoordinatesToFloat(args)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	rd := rd2wgs84.NewRDCoordinates(cor[0], cor[1])

	wgs := rd.ToWGS84()
	fmt.Printf("lat: %v lon: %v\n", wgs.Lat, wgs.Lon)
	return 0
}

func (r *WGS84) Synopsis() string {
	return r.Help()
}

func convertStringCoordinatesToFloat(coordinates []string) (cor []float64, err error) {
	cor1, err := strconv.ParseFloat(coordinates[0], 32)
	if err != nil {
		return nil, err
	}
	cor = append(cor, cor1)
	cor2, err := strconv.ParseFloat(coordinates[1], 32)
	if err != nil {
		return nil, err
	}
	cor = append(cor, cor2)

	return cor, nil
}
