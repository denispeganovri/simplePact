package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type TransformedSku struct {
	Dims      Dimensions `json:"dimensions"`
	Lud       string     `json:"last_updated_date"`
	Mass      float64    `json:"mass"`
	MassUom   string     `json:"mass_uom"`
	SkuId     float64    `json:"sku_id"`
	Volume    float64    `json:"volume"`
	VolumeUom string     `json:"volume_uom"`
}

type Dimensions struct {
	Height float64 `json:"height"`
	Length float64 `json:"length"`
	Uom    string  `json:"uom"`
	Width  float64 `json:"width"`
}

func GetResponse(ctx *gin.Context) {
	dims := Dimensions{
		Height: 6.6,
		Length: 7.7,
		Uom:    "cm",
		Width:  8.8,
	}

	ctx.JSON(http.StatusOK, TransformedSku{
		Dims:      dims,
		Lud:       "2023-02-28T10:10:13Z",
		Mass:      5.5,
		MassUom:   "g",
		SkuId:     123666,
		Volume:    9.9,
		VolumeUom: "cm3",
	})
}
