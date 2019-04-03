package main

import (
	"net/http"

	"github.com/brettman/nyc_footprints/dao"
	"github.com/gin-gonic/gin"
)

func testTheAPI(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "wordup"})
}

func topTenNames(c *gin.Context) {
	repo := dao.NewGeoJSONRepo()
	repo.TopTenNames()
	c.JSON(http.StatusOK, gin.H{"ok": "not ok?"})
}
