package handlers

import (
	"context"

	pb "api-gateway/proto-service/genprotos"

	"github.com/gin-gonic/gin"
)

// CreateElection godoc
// @ID create_election
// @Router /election [POST]
// @Summary Creates Election
// @Description Create election
// @Tags Election
// @Accept json
// @Produce json
// @Param Election body pb.Election true "election body data"
// @Success 200 {object} pb.Void
// @Failure 500 {object} string "Failed to get product by id"
func (h *HTTPHandler) CreateElection(c *gin.Context) {
	var election pb.Election
	if err := c.ShouldBindJSON(&election); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := h.Election.CreateElection(context.Background(), &election)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
	h.Logger.INFO.Println("Election created")
}

// GetElectionByID godoc
// @ID get_election_by_id
// @Router /election/{id} [GET]
// @Summary Gets Election by id
// @Description gets election by reading id  from path
// @Tags Election
// @Accept json
// @Produce json
// @Param id path string true "election ID"
// @Success 200 {object} pb.Election
// @Failure 500 {object} string "Failed to get product by id"
func (h *HTTPHandler) GetElectionByID(c *gin.Context) {
	id := c.Param("id")
	var election pb.ById
	election.Id = id
	res, err := h.Election.GetByIdElection(context.Background(), &election)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}

// UpdateElection godoc
// @ID update_election
// @Summary Update Election
// @Description Updates election by reading election from body
// @Tags Election
// @Accept json
// @Produce json
// @Param election body pb.Election true "election data"
// @Success 200 {object} pb.Void
// @Failure 500 {object} string "Failed to update election data"
// @Router /election/update [PUT]
func (h *HTTPHandler) UpdateElection(c *gin.Context) {
	var election pb.Election
	if err := c.ShouldBindJSON(&election); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	res, err := h.Election.UpdateElection(context.Background(), &election)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
	h.Logger.INFO.Println("Election updated")
}

// DeleteElection godoc
// @ID delete_election
// @Router /election/{id} [DELETE]
// @Summary Deletes Election
// @Description Deletes election by reading id
// @Tags Election
// @Accept json
// @Produce json
// @Param id path string true "election ID"
// @Success 200 {object} string "Deleted successfully"
// @Failure 500 {object} string "Failed to delete election data"
func (h *HTTPHandler) DeleteElection(c *gin.Context) {
	id := c.Param("id")
	var election pb.ById
	election.Id = id
	_, err := h.Election.DeleteElection(context.Background(), &election)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"success": "Deleted successfully"})
	h.Logger.INFO.Println("Election deleted")
}

// GetElectionsByDate godoc
// @ID get_elections_by_date
// @Router /election/by-date [GET]
// @Summary Getss Elections by date
// @Description Gets Elections by date
// @Tags Election
// @Accept json
// @Produce json
// @Param date query string true "election date"
// @Success 200 {object} pb.GetAllElection
// @Failure 500 {object} string "Failed to Get elections data"
func (h *HTTPHandler) GetElectionsByDate(c *gin.Context) {
	var election pb.Filter
	date := c.Query("date")
	election.Date = date
	res, err := h.Election.GetAllElections(context.Background(), &election)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}

// GetAllElections godoc
// @ID get_all_elections
// @Router /elections [GET]
// @Summary Gets all Elections
// @Description Gets all Elections
// @Tags Election
// @Accept json
// @Produce json
// @Success 200 {object} pb.GetAllElection
// @Failure 500 {object} string "Failed to Get all elections"
func (h *HTTPHandler) GetAllElections(c *gin.Context) {
	var election pb.Filter
	res, err := h.Election.GetAllElections(context.Background(), &election)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"object": res})
}
