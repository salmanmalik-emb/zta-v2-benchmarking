package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HopResult struct {
	BytesOut   int    `json:"bytes_out"`
	PacketsOut int    `json:"packets_out"`
	BytesIn    int    `json:"bytes_in"`
	PacketsIn  int    `json:"packets_in"`
	Id         string `json:"id"`
}

type HopResultHandler struct {
	results []*HopResult
}

func NewHopResultHandler() *HopResultHandler {

	results := []*HopResult{}
	handler := &HopResultHandler{
		results: results,
	}
	return handler
}

func (h *HopResultHandler) SaveResult(c *gin.Context) {
	var data HopResult
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.results = append(h.results, &data)

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (h *HopResultHandler) GetResults(c *gin.Context) {
	total := HopResult{}
	for _, result := range h.results {
		total.BytesOut += result.BytesOut
		total.BytesIn += result.BytesIn
		total.PacketsIn += result.PacketsIn
		total.PacketsOut += result.PacketsOut
	}

	c.JSON(http.StatusOK, gin.H{"results": h.results, "total": total})
}

func (h *HopResultHandler) ClearResults(c *gin.Context) {
	h.results = []*HopResult{}
}
