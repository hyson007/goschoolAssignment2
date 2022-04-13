package main

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"github.com/hyson007/GoSchoolAssignMent2/bst"
)

var (
	//venues (3digit)
	venues = []string{
		"East",
		"West",
		"South",
		"North",
		"DownTown",
	}
	//movies (4digit)
	movies = []string{
		"Spderman",
		"Avenger",
		"Captain America",
		"IronMan",
		"Golang Movie",
	}
)

const (
	sortByDate = iota
	sortByVenue
	sortByMovie
)

func getIndex(arr []string, target string) int {
	for idx, a := range arr {
		if a == target {
			return idx + 1
		}
	}
	log.Printf("unable to find target %s in array", target)
	return -1
}

func dateNodeWrapper(sortby int, inputs ...string) (int, error) {
	if len(inputs) != 3 {
		return -1, fmt.Errorf("please check the input of dateNodeWrapper, length is not equal to 3")
	}
	date, venue, movie := inputs[0], inputs[1], inputs[2]
	dateInt, err := strconv.Atoi(date)
	if err != nil {
		return -1, err
	}

	switch sortby {
	case sortByDate:
		return dateInt*int(math.Pow(10, 7)) +
			getIndex(venues, venue)*int(math.Pow(10, 4)) +
			getIndex(movies, movie), nil
	case sortByVenue:
		return getIndex(venues, venue)*int(math.Pow(10, 17)) +
			getIndex(movies, movie)*int(math.Pow(10, 14)) +
			dateInt, nil
	case sortByMovie:
		return getIndex(movies, movie)*int(math.Pow(10, 16)) +
			getIndex(venues, venue)*int(math.Pow(10, 13)) +
			dateInt, nil
	}

	return -1, fmt.Errorf("invalid sort by value")
}

// func init() {
// 	// populate the data into three bst
// 	rawData := settings.RawData
// 	// for dateBST, randomized based the date then add into tree

// 	rand.Seed(time.Now().UnixNano())
// 	rand.Shuffle(len(rawData), func(i, j int) {
// 		rawData[i][sortByDate], rawData[j][sortByDate] = rawData[j][sortByDate], rawData[i][sortByDate]
// 	})
// 	dateBst := bst.NodeBst{nil, 0}
// 	for _, rd := range rawData {
// 		input, err := dateNodeWrapper(sortByDate, rd...)
// 		if err != nil {
// 			log.Println(err.Error())
// 		}
// 		dateBst.AddNode(input)
// 	}
// 	// dateBst.PrintLevelOrder()

// 	// for venueBST, randomized based the venue then add into tree
// 	rand.Seed(time.Now().UnixNano())
// 	rand.Shuffle(len(rawData), func(i, j int) {
// 		rawData[i][sortByVenue], rawData[j][sortByVenue] = rawData[j][sortByVenue], rawData[i][sortByVenue]
// 	})
// 	venueBst := bst.NodeBst{nil, 0}
// 	for _, rd := range rawData {
// 		input, err := dateNodeWrapper(sortByVenue, rd...)
// 		if err != nil {
// 			log.Println(err.Error())
// 		}
// 		venueBst.AddNode(input)
// 	}
// 	// venueBst.PrintLevelOrder()

// 	// for movieBST, randomized based the movie then add into tree
// 	rand.Seed(time.Now().UnixNano())
// 	rand.Shuffle(len(rawData), func(i, j int) {
// 		rawData[i][sortByMovie], rawData[j][sortByMovie] = rawData[j][sortByMovie], rawData[i][sortByMovie]
// 	})
// 	movieBst := bst.NodeBst{nil, 0}
// 	for _, rd := range rawData {
// 		input, err := dateNodeWrapper(sortByMovie, rd...)
// 		if err != nil {
// 			log.Println(err.Error())
// 		}
// 		movieBst.AddNode(input)
// 	}
// 	movieBst.PrintLevelOrder()
// }

func main() {

	data := bst.NodeBst{nil, 0}
	data.AddNode(2022041810, "North", "spider man")
	data.AddNode(2022041710, "South", "iron man")
	data.AddNode(2022041710, "East", "iron man")
	data.AddNode(2022041910, "East", "avenger")
	data.AddNode(2022042010, "Downtown", "captain")
	// data.AddNode(2022041810, "Downtown", "captain2")
	// data.AddNode(2022041810, "South", "captain")
	// data.AddNode(2022042210, "South", "captain")
	// data.AddNode(2022042208, "South", "captain")
	// data.AddNode(2022042210, "East", "captain")
	// data.AddNode(2022042210, "Downtown", "captain")
	// data.AddNode(2022042220, "South", "captain2")
	// data.AddNode(2022042218, "South", "captain3")
	// data.AddNode(2022042010, "East", "wonder woman")
	data.PrintLevelOrder()
	data.RemoveDateHour(2022041810)
	fmt.Println("-----------")
	data.PrintLevelOrder()
	// fmt.Println(data.SearchSingleDateHour(2022041810))
	// fmt.Println(data.SearchSingleDateHour(2022041810).ByVenue("Downtown"))
	// fmt.Println(data.SearchSingleDateHour(2022041810).ByMovie("spider man"))
	// fmt.Println("--------------")
	// fmt.Println(data.SearchRangeDate(2022041700, 2022042299))
	// fmt.Println(data.SearchRangeDate(2022041700, 2022042299).ByVenue("Downtown"))

	// fmt.Println(data.ModifyByDate(2022041810).SubRangeSearchMovie("spider man").First())
	// fmt.Println(data.SearchDate(20220418).SubSearchVenue("Downtown"))
	// fmt.Println(data.SearchDate(20220418).SubSearchMovie("captain"))
	// fmt.Println(dateNodeWrapper("2022041708", "East", "Spderman"))

	//dateBst BST format
	//date + venue (3digit) + movies (4 digit)
	//2022041108 + xxx + yyyy

	// _ = dateBst
	// dateBst.AddNode()
	// dateBst.AddNode(20)
	// dateBst.AddNode(30)
	// dateBst.AddNode(120)
	// dateBst.AddNode(5)
	// dateBst.AddNode(6)
	// dateBst.AddNode(1)
	// dateBst.PrintLevelOrder()

	// fmt.Println()
	// venueBst := bst.NodeBst{nil, 0}
	// venueBst.AddNode(30)
	// venueBst.AddNode(50)
	// venueBst.AddNode(5)
	// venueBst.AddNode(7)
	// venueBst.AddNode(9)
	// venueBst.AddNode(20)
	// // venueBst.AddNode(50)
	// venueBst.PrintLevelOrder()

	// err := venueBst.Remove(50)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println()
	// 	venueBst.PrintLevelOrder()
	// }

}
