package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type CsvLine struct {
	TrainNo                string
	TrainN                 string
	SEQ                    string
	StationCode            string
	StationName            string
	ArrivalTime            string
	DeapartureTime         string
	Distance               string
	SourceStation          string
	SourceStationName      string
	DestinationStation     string
	DestinationStationName string
}

func main() {

	lines, err := ReadCsv("Indian_railway1.csv")
	if err != nil {
		panic(err)
	}

	// Loop through lines & turn into object
	for _, line := range lines {
		data := CsvLine{
			TrainNo:                line[0],
			TrainN:                 line[1],
			SEQ:                    line[2],
			StationCode:            line[3],
			StationName:            line[4],
			ArrivalTime:            line[5],
			DeapartureTime:         line[6],
			Distance:               line[7],
			SourceStation:          line[8],
			SourceStationName:      line[9],
			DestinationStation:     line[10],
			DestinationStationName: line[11],
		}
		fmt.Println(data.TrainNo + "|" + data.TrainN + "|" + data.SEQ + "|" + data.StationCode + "|" + data.StationName + "|" + data.ArrivalTime + "|" + data.DeapartureTime + "|" + data.Distance + "|" + data.SourceStation + "|" + data.SourceStationName + "|" + data.DestinationStation + "|" + data.DestinationStationName)
	}
}

// ReadCsv accepts a file and returns its content as a multi-dimentional type
// with lines and each column. Only parses to string type.
func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}
