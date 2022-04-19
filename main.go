package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hyson007/GoSchoolAssignMent2/bst"
)

var (
	venues = []string{
		"East",
		"West",
		"South",
		"North",
		"DownTown",
	}
	movies = []string{
		"Spderman",
		"Avenger",
		"CaptainAmerica",
		"IronMan",
		"GolangMovie",
	}
	data = bst.Bst{nil, 0}
)

func init() {
	data.AddNode(2022041810, "North", "Spderman")
	data.AddNode(2022041710, "South", "IronMan")
	data.AddNode(2022041710, "East", "IronMan")
	data.AddNode(2022041910, "East", "Avenger")
	data.AddNode(2022042010, "DownTown", "CaptainAmerica")
	data.AddNode(2022041810, "DownTown", "CaptainAmerica")
	data.AddNode(2022041810, "South", "GolangMovie")
	data.AddNode(2022042110, "DownTown", "CaptainAmerica")
	data.AddNode(2022042210, "East", "CaptainAmerica")
	data.AddNode(2022042310, "South", "CaptainAmerica")
	data.AddNode(2022042410, "North", "Spderman")
	data.AddNode(2022042510, "South", "CaptainAmerica")
	data.AddNode(2022042610, "East", "Avenger")
	data.AddNode(2022042710, "North", "CaptainAmerica")
	data.AddNode(2022042810, "East", "IronMan")
	data.AddNode(2022042910, "North", "CaptainAmerica")
}

func checkInList(list []string, value string) bool {
	for _, v := range list {
		// fmt.Println(v, value)
		// fmt.Println(v == value)
		if v == value {
			return true
		}
	}
	return false
}

func main() {

	// data.PrintLevelOrder()
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Type", "ContentType"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))
	// r.Use(cors.Default())
	r.POST("/singledatehour", func(c *gin.Context) {
		var json struct {
			DateHour int    `json:"dateHour" binding:"required"`
			Venue    string `json:"venue" binding:"required"`
			Movie    string `json:"movie" binding:"required"`
		}
		if err := c.Bind(&json); err == nil {
			// fmt.Println(json)
			// check if venue and movies are in	 the list
			if json.Venue != "" {

				// fmt.Println(venues, json.Venue)
				res := checkInList(venues, json.Venue)
				if !res {
					c.JSON(400, gin.H{
						"error": "venue not found in list, please update venues list first",
					})
					return
				}
			}

			if json.Movie != "" {
				res := checkInList(movies, json.Movie)
				if !res {
					c.JSON(400, gin.H{
						"error": "movie not found in list, please update movies list first",
					})
					return
				}
			}

			err := data.AddNode(json.DateHour, json.Venue, json.Movie)
			if err != nil {
				c.JSON(400, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		} else {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		}
	})

	r.PUT("/singledatehour/:dh", func(c *gin.Context) {
		dateHour, err := strconv.Atoi(c.Param("dh"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "bad request, dateHour must be an integer"})
			return
		}

		venue := c.Query("venue")
		movie := c.Query("movie")
		if venue == "" || movie == "" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "bad request, must provide venue and movie to be modify"})
			return
		}

		var json struct {
			DateHour int    `json:"newdateHour" binding:"required"`
			Venue    string `json:"newvenue" binding:"required"`
			Movie    string `json:"newmovie" binding:"required"`
		}
		if err := c.Bind(&json); err == nil {
			//check if the new venue and movie are in the list
			if json.Venue != "" {

				// fmt.Println(venues, json.Venue)
				res := checkInList(venues, json.Venue)
				if !res {
					c.JSON(400, gin.H{
						"error": "venue not found in list, please update venues list first",
					})
					return
				}
			}

			if json.Movie != "" {
				res := checkInList(movies, json.Movie)
				if !res {
					c.JSON(400, gin.H{
						"error": "movie not found in list, please update movies list first",
					})
					return
				}
			}

			// check the new dateHour, new venue and new movie do not exist in the tree
			shouldNotExist := data.SearchSingleDateHour(json.DateHour).ByVenue(json.Venue).ByMovie(json.Movie)
			fmt.Println(*shouldNotExist, "hit")
			fmt.Println(*shouldNotExist != nil)
			if *shouldNotExist != nil {
				c.JSON(400, gin.H{
					"error": "the new venue and movie already exist in the target datehour",
				})
				return
			}

			// check if newdatehour is same as the old one
			if dateHour == json.DateHour {

				// same datehour, we proceed to update the venue or movies
				// first, let's check there should be no other item with the
				// same venue and movie already in the target datehour
				item := data.SearchSingleDateHour(dateHour).ByVenue(venue).ByMovie(movie)
				err := item.ModifyMovieOrVenue(json.Movie, json.Venue)
				if err != nil {
					log.Println(err)
					c.JSON(http.StatusBadRequest, gin.H{"status": "something wrong with the modify"})
					return
				}
				c.JSON(http.StatusOK, gin.H{"status": "ok"})
			} else {
				// new datehour, we proceed to delete the old one and add the new one
				err := data.ModifyDateHour(dateHour, json.DateHour, movie, json.Movie, venue, json.Venue)
				if err != nil {
					log.Println(err)
					c.JSON(http.StatusBadRequest, gin.H{"status": "something wrong with the modify"})
					return
				} else {
					c.JSON(http.StatusOK, gin.H{"status": "ok"})
				}
			}
		} else {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		}
	})

	r.DELETE("/singledatehour/:dh", func(c *gin.Context) {
		dateHour, err := strconv.Atoi(c.Param("dh"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "bad request"})
			return
		}
		venue := c.Query("venue")
		movie := c.Query("movie")

		if venue != "" && movie != "" {
			//calling from web should hit this
			// fmt.Println("hit from web", venue, movie)
			err := data.RemoveOneEntry(dateHour, venue, movie)
			if err != nil {
				log.Println(err)
				c.JSON(http.StatusBadRequest,
					gin.H{"status": "something wrong with the delete one entry"})
				return
			}

		} else {
			err = data.RemoveDateHour(dateHour)
		}
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": "dateHour not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.GET("/singledatehour/:dh", func(c *gin.Context) {
		datehour := c.Param("dh")
		datehourInt, err := strconv.Atoi(datehour)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		venue := c.Query("venue")
		movie := c.Query("movie")
		if venue != "" && movie != "" {
			result := data.SearchSingleDateHour(datehourInt).ByVenue(venue).ByMovie(movie)
			c.JSON(200, gin.H{
				"message": result,
			})
			return
		}
		if venue != "" {
			result := data.SearchSingleDateHour(datehourInt).ByVenue(venue)
			c.JSON(200, gin.H{
				"message": result,
			})
			return
		}
		if movie != "" {
			result := data.SearchSingleDateHour(datehourInt).ByMovie(movie)
			c.JSON(200, gin.H{
				"message": result,
			})
			return
		}
		if venue == "" && movie == "" {
			result := data.SearchSingleDateHour(datehourInt)
			c.JSON(200, gin.H{
				"message": result,
			})
			return
		}
	})

	r.GET("/rangedatehour/", func(c *gin.Context) {
		start := c.DefaultQuery("start", "2022010100")
		startInt, err := strconv.Atoi(start)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "start dateHour is not valid",
			})
			return
		}
		end := c.DefaultQuery("end", "2030010100")
		endInt, err := strconv.Atoi(end)
		if err != nil {
			c.JSON(400, gin.H{
				"error": "end dateHour is not valid",
			})
			return
		}
		venue := c.Query("venue")
		movie := c.Query("movie")

		if venue != "" && movie != "" {
			result := data.SearchRangeDateHour(startInt, endInt).ByVenue(venue).ByMovie(movie)
			c.JSON(200, gin.H{
				"message": result,
			})
			return
		} else if venue != "" {
			result := data.SearchRangeDateHour(startInt, endInt).ByVenue(venue)
			c.JSON(200, gin.H{
				"message": result,
			})
			return
		} else if movie != "" {
			result := data.SearchRangeDateHour(startInt, endInt).ByMovie(movie)
			c.JSON(200, gin.H{
				"message": result,
			})
			return
		} else {
			// fmt.Println("hit", startInt, endInt)
			result := data.SearchRangeDateHour(startInt, endInt)
			c.JSON(200, gin.H{
				"message": result,
			})
			return
		}
	})

	r.GET("/venues", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": venues,
		})
	})

	r.POST("/venues", func(c *gin.Context) {
		var json struct {
			Venue string `json:"venue" binding:"required"`
		}
		if err := c.Bind(&json); err == nil {
			fmt.Println(json)
			res := checkInList(venues, json.Venue)
			if !res {
				venues = append(venues, json.Venue)
				c.JSON(http.StatusOK, gin.H{"status": "ok"})
				return
			}
			c.JSON(400, gin.H{
				"error": "venue already exists",
			})
			return
		} else {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		}
	})

	r.DELETE("/venues/:venue", func(c *gin.Context) {
		venue := c.Param("venue")
		res := checkInList(venues, venue)
		if !res {
			c.JSON(400, gin.H{
				"error": "venue not found",
			})
			return
		}
		for i, v := range venues {
			if v == venue {
				venues = append(venues[:i], venues[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"status": "ok"})
				return
			}
		}
	})

	r.GET("/movies", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": movies,
		})
	})

	r.POST("/movies", func(c *gin.Context) {
		var json struct {
			Movie string `json:"movie" binding:"required"`
		}
		if err := c.Bind(&json); err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
		} else {

			fmt.Println(json.Movie)
			res := checkInList(movies, json.Movie)
			if !res {
				movies = append(movies, json.Movie)
				c.JSON(http.StatusOK, gin.H{"status": "ok"})
				return
			}
			c.JSON(400, gin.H{
				"error": "movie already exists",
			})
			return
		}
	})

	r.DELETE("/movies/:movie", func(c *gin.Context) {
		movie := c.Param("movie")
		res := checkInList(movies, movie)
		if !res {
			c.JSON(400, gin.H{
				"error": "movie not found",
			})
			return
		}
		for i, v := range movies {
			if v == movie {
				movies = append(movies[:i], movies[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"status": "ok"})
				return
			}
		}
	})

	r.GET("/balance", func(c *gin.Context) {
		log.Println("before balance")
		data.PrintLevelOrder()
		data.BalanceTree()
		log.Println("after balance")
		data.PrintLevelOrder()

		c.JSON(200, gin.H{
			"message": "done",
		})
	})

	r.Run()

	// it doesn;t make too much sense to do put for a list
	// as we can just add or delete elements

	// r.PUT("/movies/:movie", func(c *gin.Context) {
	// 	movie := c.Param("movie")
	// 	res := checkInList(movies, movie)
	// 	if !res {
	// 		c.JSON(400, gin.H{
	// 			"error": "old movie not found in movies list",
	// 		})
	// 		return
	// 	}
	// 	var json struct {
	// 		NewMovie string `json:"newmovie" binding:"required"`
	// 	}
	// 	if err := c.Bind(&json); err == nil {
	// 		fmt.Println(json)
	// 		res := checkInList(movies, json.NewMovie)
	// 		if !res {
	// 			c.JSON(400, gin.H{
	// 				"error": "new movie not found in movies list",
	// 			})
	// 			return
	// 		}
	// 		for i, v := range movies {
	// 			if v == movie {
	// 				movies[i] = json.NewMovie
	// 				c.JSON(http.StatusOK, gin.H{"status": "ok"})
	// 				return
	// 			}
	// 		}
	// 	} else {
	// 		c.JSON(400, gin.H{
	// 			"error": err.Error(),
	// 		})
	// 	}
	// })

	// err := data.SearchSingleDateHour(2022041810).ByVenue("Downtown").ModifyMovieOrVenue("Downtown3", "captain2")
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// err := data.RemoveOneEntry(2022041710, "East", "iron man3")
	// if err != nil {
	// 	log.Println(err.Error())
	// }

	// data.AddNode(2022042210, "South", "captain")
	// data.AddNode(2022042208, "	", "captain")
	// data.AddNode(2022042210, "East", "captain")
	// data.AddNode(2022042210, "Downtown", "captain")
	// data.AddNode(2022042220, "South", "captain2")
	// data.AddNode(2022042218, "South", "captain3")
	// data.AddNode(2022042010, "East", "wonder woman")
	// data.PrintLevelOrder()
	// data.BalanceTree()
	// data.PrintLevelOrder()

	// // data.RemoveDateHour(2022041810)
	// // data.RemoveOneEntry(2022041910, "avenger", "East")
	// fmt.Println("-----------")
	// data.PrintLevelOrder()
	// // data.PrintLevelOrder()
	// fmt.Println(data.SearchSingleDateHour(2022041810))
	// fmt.Println("-----------")
	// t := data.SearchSingleDateHour(2022041810).ByVenue("Downtown")
	// fmt.Println(t)
	// foo := *t
	// foo[0].Movie = "hit"
	// (*t)[0] = nil
	// fmt.Println(t)
	// fmt.Println(data.SearchSingleDateHour(2022041810).ByVenue("Downtown"))
	// fmt.Println(data.SearchSingleDateHour(2022041810).ByVenue("Downtown"))

	// fmt.Println(data.SearchSingleDateHour(2022041810).ByMovie("spider man"))
	// fmt.Println("--------------")
	// fmt.Println(data.SearchRangeDateHour(2022041700, 2022041810))
	// fmt.Println("--------------")
	// // fmt.Println(data.SearchRangeDateHour(2022041700, 2022042299).ByMovie("captain").ByVenue("South"))
	// fmt.Println(data.Modify(2022042210, 2022042210, "South", "South", "captain", "captain2"))
	// data.PrintLevelOrder()
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

}
