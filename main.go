package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type data struct {
	Id    string `json:id`
	Title string `json:"title"`
	Disc  string `json:"disc"`
}

var datas = []data{
	{Id: "1", Title: "Blue Train", Disc: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. "},
	{Id: "2", Title: "Jeru", Disc: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. "},
	{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Disc: "Lorem Ipsum is simply dummy text of the printing and typesetting industry. "},
}

func homePage(c *gin.Context) {
	c.String(http.StatusOK, "has runing!")
}
func getDat(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, datas)
}

func addDat(c *gin.Context) {
	var newData data
	if err := c.BindJSON(&newData); err != nil {
		return
	}
	datas = append(datas, newData)
	c.IndentedJSON(http.StatusCreated, datas)
}
func getDataFromId(c *gin.Context) {
	id := c.Param("id")
	data, err := checkId((id))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "data not found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, data)
}
func checkId(id string) (*data, error) {
	for i, item := range datas {
		if item.Id == id {
			return &datas[i], nil
		}
	}
	return nil, errors.New("datas not found")
}
func update(c *gin.Context) {
	id := c.Param("id")
	data, err := checkId((id))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "data not found!"})
		return
	}

	data.Title = "this is title" //new data
	c.IndentedJSON(http.StatusOK, data)
}

func delete(c *gin.Context) {
	id := c.Param("id")
	data, err := checkId((id))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "not data to delete"})
		return
	}
	for i, item := range datas {
		if item.Id == id {
			datas = append(datas[:i], datas[i+1:]...)
			break
		}
	}

	c.IndentedJSON(http.StatusOK, data)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/", homePage)
	r.GET("/getDat", getDat)
	r.GET("/getDat/:id/", getDataFromId)
	r.PATCH("/updateDat/:id", update)
	r.DELETE("/delete/:id", delete)
	r.POST("/addDat", addDat)
	r.Run(":8080")
}
