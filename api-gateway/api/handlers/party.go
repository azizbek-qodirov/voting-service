package handlers

import (
	"context"
	"fmt"
	"net/http"

	_ "api-gateway/api/docs"
	pb "api-gateway/proto-service/genprotos"

	"github.com/gin-gonic/gin"
)

// GetPartyID godoc
// @ID getparty
// @Router /party/{id} [GET]
// @Summary Get Party By ID
// @Description Get party by ID
// @Tags party
// @Accept json
// @Produce json
// @Param id path string true "party ID"
// @Success 200 {object} pb.PartyGetById "party data"
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "party not found"
// @Failure 500 {object} string "Server error"
func (h *HTTPHandler) GetPartyByID(c *gin.Context) {
	id := &pb.PartyGetById{Id: c.Param("id")}
	res, err := h.Party.GetById(context.Background(), id)
	if err != nil {
		c.String(500, "Couldn't get the party: "+err.Error())
		h.Logger.ERROR.Println("Couldn't get the party: " + err.Error())
		return
	}
	c.JSON(200, res)
}

// GetAllparties godoc
// @ID getallparties
// @Router /parties [GET]
// @Summary Get all parties
// @Description Get all parties
// @Tags party
// @Accept json
// @Produce json
// @Success 200 {object} pb.PartyGetAllReq "party data"
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "party not found"
// @Failure 500 {object} string "Server error"
func (h *HTTPHandler) GetAllParty(c *gin.Context) {
	req := pb.PartyGetAllReq{Name: c.Query("name"), OpenedDate: c.Query("opened_date")}
	res, err := h.Party.GetAll(context.Background(), &req)
	if err != nil {
		c.String(500, "Couldn't get parties: "+err.Error())
		h.Logger.ERROR.Println("Couldn't get parties: " + err.Error())
		return
	}
	c.JSON(200, res)
}

// Create party godoc
// @ID createparty
// @Router /party [POST]
// @Summary Create parties
// @Description Create parties
// @Tags party
// @Accept json
// @Produce json
// @Param party body pb.PartyCreate true "Party data"
// @Success 200 {object} string "party created"
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "party not found"
// @Failure 500 {object} string "Server error"
func (h *HTTPHandler) CreateParty(c *gin.Context) {
	req := pb.PartyCreate{}
	err := c.BindJSON(&req)
	if err != nil {
		c.String(500, "Couldn't bind the data: "+err.Error())
		h.Logger.ERROR.Println("Couldn't bind the data: " + err.Error())
		return
	}
	_, err = h.Party.Create(context.Background(), &req)
	if err != nil {
		c.String(500, "Couldn't create the party: "+err.Error())
		h.Logger.ERROR.Println("Couldn't create the party: " + err.Error())
		return
	}
	c.String(http.StatusCreated, "party created")
	h.Logger.INFO.Println("party created")
}

// Update party godoc
// @ID updateparty
// @Router /party/{id} [PUT]
// @Summary Update party
// @Description Update party
// @Tags party
// @Accept json
// @Produce json
// @Param id path string true "Party id"
// @Param public body pb.PartyUpdate true "Party data"
// @Success 200 {object} pb.PartyUpdate "party data"
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "party not found"
// @Failure 500 {object} string "Server error"
func (h *HTTPHandler) UpdateParty(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	req := pb.PartyUpdate{}
	err := c.BindJSON(&req)
	if err != nil {
		c.String(500, "Couldn't bind the data: "+err.Error())
		h.Logger.ERROR.Println("Couldn't bind the data: " + err.Error())
		return
	}
	req.Id = id
	_, err = h.Party.Update(context.Background(), &req)
	if err != nil {
		c.String(500, "Couldn't update the party: "+err.Error())
		h.Logger.ERROR.Println("Couldn't update the party: " + err.Error())
		return
	}
	c.String(http.StatusOK, "party updated")
	h.Logger.INFO.Printf("party updated: %s", req.GetId())
}

// Delete Party godoc
// @ID deleteparty
// @Router /party/{id} [DELETE]
// @Summary Delete party
// @Description Delete party
// @Tags party
// @Accept json
// @Produce json
// @Param id path string true "party ID"
// @Success 200 {object} pb.PartyDelete "party data"
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "party not found"
// @Failure 500 {object} string "Server error"
func (h *HTTPHandler) DeleteParty(c *gin.Context) {
	req := pb.PartyDelete{Id: c.Param("id")}
	_, err := h.Party.Delete(context.Background(), &req)
	if err != nil {
		c.String(500, "Couldn't delete the party: "+err.Error())
		h.Logger.ERROR.Println("Couldn't delete the party: " + err.Error())
		return
	}
	c.String(http.StatusOK, "party deleted")
	h.Logger.INFO.Printf("party deleted: %s", req.GetId())
}
