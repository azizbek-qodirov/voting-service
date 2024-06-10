package handlers

import (
	"context"

	pb "api-gateway/proto-service/genprotos"

	"github.com/gin-gonic/gin"
)

// CreateCandidate godoc
// @ID create_candidate
// @Router /candidate [POST]
// @Summary Creates Candidate
// @Description Creates candidate by reading from body
// @Tags Candidate
// @Accept json
// @Produce json
// @Param Candiate body pb.CandidateCreate true "candiate body data"
// @Success 200 {object} string "Candiate created successfully"
// @Failure 500 {object} string "Failed to create candidate"
func (h *HTTPHandler) CreateCandidate(c *gin.Context) {
	var can pb.CandidateCreate
	if err := c.ShouldBindJSON(&can); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Candidate.CreateCandidate(context.Background(), &can)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Candiate created successfully"})
	h.Logger.INFO.Println("Candidate created")
}

// DeleteCandidate godoc
// @ID delete_candidate
// @Router /candidate/{id} [DELETE]
// @Summary deletes Candidate
// @Description Deletes candidate by reading id from body
// @Tags Candidate
// @Accept json
// @Produce json
// @Param id path string true "candiate id data"
// @Success 200 {object} string "Candiate deleted successfully"
// @Failure 500 {object} string "Failed to create candidate"
func (h *HTTPHandler) DeleteCandidate(c *gin.Context) {
	req := pb.ById{Id: c.Param("id")}
	_, err := h.Candidate.DeleteCandidate(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Candiate deleted successfully"})
	h.Logger.INFO.Println("Candidate deleted")
}

// UpdateCandidate godoc
// @ID updates_candidate
// @Summary updates Candidate
// @Description updates candidate by reading from body
// @Tags Candidate
// @Accept json
// @Produce json
// @Param candiate body  pb.Candidate true "candiate data"
// @Success 200 {object} string "Candiate updated successfully"
// @Failure 500 {object} string "Failed to create candidate"
// @Router /candidate/update [PUT]
func (h *HTTPHandler) UpdateCandidate(c *gin.Context) {
	var can pb.Candidate
	if err := c.ShouldBindJSON(&can); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := h.Candidate.UpdateCandidate(context.Background(), &can)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Candiate updated successfully"})
	h.Logger.INFO.Println("Candidate updated")
}

// GetByIdCandidate godoc
// @ID get_candidate_by_id
// @Router /candidate/{id} [GET]
// @Summary gets Candidate by id
// @Description Gets candidate by reading id from body
// @Tags Candidate
// @Accept json
// @Produce json
// @Param id path string true "candiate id"
// @Success 200 {object} pb.Candidate
// @Failure 500 {object} string "Failed to get candidate by id"
func (h *HTTPHandler) GetByIdCandidate(c *gin.Context) {
	req := pb.ById{Id: c.Param("id")}
	res, err := h.Candidate.GetByIdCandidate(context.Background(), &req)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}

// GetAllCandidates godoc
// @ID get_all_candidate
// @Router /candidates [GET]
// @Summary gets all Candidate
// @Description Gets all candidates
// @Tags Candidate
// @Accept json
// @Produce json
// @Success 200 {object} pb.GetAllCandidate
// @Failure 500 {object} string "Failed to get all candidates"
func (h *HTTPHandler) GetAllCandidates(c *gin.Context) {
	var can pb.Filter
	res, err := h.Candidate.GetAllCandidates(context.Background(), &can)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}

// GetAllCandidatesByParty godoc
// @ID get_candidates_by_party
// @Summary gets Candidate by the party id
// @Description Gets all candidates
// @Tags Candidate
// @Accept json
// @Produce json
// @Success 200 {object} pb.GetAllCandidate
// @Failure 500 {object} string "Failed to get candidates by party"
// @Router /candidate/by-party [GET]
func (h *HTTPHandler) GetAllCandidatesByParty(c *gin.Context) {
	var can pb.Filter
	party := c.Query("party")
	can.Party = party
	res, err := h.Candidate.GetAllCandidates(context.Background(), &can)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, res)
}
