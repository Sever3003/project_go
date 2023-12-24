package controler

import (
	"log"

	"driverserv/api/http"
	"driverserv/db"

	"github.com/gin-gonic/gin"
)

func ReturnModelOrder(name string, start string, end string, st int) db.Order {
	return db.Order{Clientname: name, Location_Start: start, Location_End: end, Status: st}

}

func ReturnModelDriver(name string) db.DriverInfo {
	return db.DriverInfo{}
}

func Add(c *gin.Context) {
	req := &http.RequstAdd{}
	_ = c.BindJSON(&req)

	o := ReturnModelOrder(req.Client_Name, req.Location_Start, req.Location_End, req.Status)
	err := db.AddDocument(o)

	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"status": "order added",
	})

}

func AddDriver(c *gin.Context) {
	req := (c.Param("name"))
	//_ = c.BindJSON(&req)

	o := ReturnModelDriver(req)
	err := db.AddDriver(o)

	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"status": "driver added",
	})
}

func GetDriver(c *gin.Context) {
	i, err := db.GetDriver()
	if err != nil {
		HandleError(c, err)
		return
	}

	a := []http.RequstDriverGet{}
	s := http.RequstDriverGet{}

	for _, d := range i {
		s.DriverName = d.DriverName
		s.Uuid = d.Uuid
		a = append(a, s)
	}

	c.JSON(200, http.ResponseDriverGet{Record: a})
}

func Update(c *gin.Context) {
	req := &http.RequstUpdate{}
	_ = c.BindJSON(&req)

	o := ReturnModelOrder(req.Client_Name, "", "", req.Status)
	err := db.UpdateDocument(o)

	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"status": "loctions change",
	})
}

func Delete(c *gin.Context) {

	n := (c.Param("name"))

	o := ReturnModelOrder(n, "", "", 0)
	err := db.DeleteDocument(o)

	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(200, gin.H{
		"status": "record delete",
	})

}

func Get(c *gin.Context) {

	i, err := db.GetDocument()
	if err != nil {
		HandleError(c, err)
		return
	}

	a := []http.RequstAdd{}
	s := http.RequstAdd{}

	for _, d := range i {

		s.Client_Name = d.Clientname
		s.Location_Start = d.Location_Start
		s.Location_End = d.Location_End
		s.Status = d.Status

		a = append(a, s)
	}

	c.JSON(200, http.ResponseGet{Record: a})
}

func HandleError(c *gin.Context, err error) {
	c.JSON(500, map[string]string{
		"error": err.Error(),
	})
	log.Println(err)
	return
}
