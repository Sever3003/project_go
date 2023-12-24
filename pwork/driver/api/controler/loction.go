package controler

import (
	"log"

	"github.com/Gwoop/Driver/api/http"
	"github.com/Gwoop/Driver/db"
	"github.com/gin-gonic/gin"
)

func ReturnModelLocation(l string, u string) *db.Location {
	return &db.Location{Locations: l, DriverUuid: u}

}

func Add(c *gin.Context) {
	req := &http.RequstAdd{}
	_ = c.BindJSON(&req)
	result := db.DB.Create(ReturnModelLocation(req.Locations, req.DriverUuid))
	if result.Error != nil {
		HandleError(c, result.Error)
		return
	}

	c.JSON(200, gin.H{
		"status": "loctions added",
	})

}

func Update(c *gin.Context) {
	uuid := (c.Param("uuid"))
	req := &http.RequstUpdate{}
	_ = c.BindJSON(&req)

	l := db.Location{}
	rows := db.DB.Find(&l).Where("driver_uuid = ?", uuid)
	if rows.Error != nil {
		HandleError(c, rows.Error)
		return
	}

	l.Locations = req.Locations

	rowsup := db.DB.Save(l)

	if rowsup.Error != nil {
		HandleError(c, rows.Error)
		return
	}

	c.JSON(200, gin.H{
		"status": "loctions change",
	})
}

func Delete(c *gin.Context) {
	uuid := (c.Param("uuid"))

	//rows := db.DB.Delete(db.Location{DriverUuid: uuid}) {"text":"driver_uuid","cur":{"from":11,"to":11}}

	rows := db.DB.Where("driver_uuid = ?", uuid).Delete(&db.Location{})

	if rows.Error != nil {
		HandleError(c, rows.Error)
		return
	}

	c.JSON(200, gin.H{
		"status": "record delete",
	})

}

func Get(c *gin.Context) {
	l := []db.Location{}
	rows := db.DB.Find(&l)
	if rows.Error != nil {
		HandleError(c, rows.Error)
		return
	}
	a := []http.ResponesGet{}
	s := http.ResponesGet{}

	for _, d := range l {
		s.DriverUuid = d.DriverUuid
		s.Locations = d.Locations

		a = append(a, s)
	}

	c.JSON(200, a)
}

func HandleError(c *gin.Context, err error) {
	c.JSON(500, map[string]string{
		"error": err.Error(),
	})
	log.Println(err)
	return
}
