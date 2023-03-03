package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Result struct {
	Status          string  `json:"status"`
	ReceivedPackets int     `json:"received_packets"`
	MinMaxWindow    int     `json:"min_max_window"`
	SentPackets     int     `json:"sent_packets"`
	Loss            float64 `json:"loss"`
	BadLoss         float64 `json:"bad_loss"`
	Rtt4            struct {
		Zero49     float64 `json:"0_49"`
		Five0179   float64 `json:"50_179"`
		One80399   float64 `json:"180_399"`
		Four00999  float64 `json:"400_999"`
		One0001999 float64 `json:"1000_1999"`
		Two000     float64 `json:"2000+"`
	} `json:"rtt4"`
	Score float64 `json:"score"`
}

type ResultHandler struct {
	results []*Result
}

func NewResultHandler() *ResultHandler {

	results := []*Result{}
	handler := &ResultHandler{
		results: results,
	}
	return handler
}

func (h *ResultHandler) SaveResult(c *gin.Context) {
	var data Result
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.results = append(h.results, &data)

	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func (h *ResultHandler) GetResults(c *gin.Context) {
	total := Result{}
	titem := len(h.results)
	if titem == 0 {
		c.JSON(http.StatusOK, gin.H{"results": h.results, "total": total})
		return
	}
	for _, result := range h.results {
		total.ReceivedPackets += result.ReceivedPackets
		total.MinMaxWindow += result.MinMaxWindow
		total.SentPackets += result.SentPackets
		total.Loss += result.Loss
		total.BadLoss += result.BadLoss
		total.Rtt4.Zero49 += result.Rtt4.Zero49
		total.Rtt4.Five0179 += result.Rtt4.Five0179
		total.Rtt4.One80399 += result.Rtt4.One80399
		total.Rtt4.Four00999 += result.Rtt4.Four00999
		total.Rtt4.One0001999 += result.Rtt4.One0001999
		total.Rtt4.Two000 += result.Rtt4.Two000
		total.Score += result.Score
	}

	total.ReceivedPackets = total.ReceivedPackets / titem
	total.MinMaxWindow = total.MinMaxWindow / titem
	total.SentPackets = total.SentPackets / titem
	total.Loss = total.Loss / float64(titem)
	total.BadLoss = total.BadLoss / float64(titem)
	total.Rtt4.Zero49 = total.Rtt4.Zero49 / float64(titem)
	total.Rtt4.Five0179 = total.Rtt4.Five0179 / float64(titem)
	total.Rtt4.One80399 = total.Rtt4.One80399 / float64(titem)
	total.Rtt4.Four00999 = total.Rtt4.Four00999 / float64(titem)
	total.Rtt4.One0001999 = total.Rtt4.One0001999 / float64(titem)
	total.Rtt4.Two000 = total.Rtt4.Two000 / float64(titem)
	total.Score = total.Score / float64(titem)

	c.JSON(http.StatusOK, gin.H{"results": h.results, "total": total})
}

func (h *ResultHandler) ClearResults(c *gin.Context) {
	h.results = []*Result{}
}
