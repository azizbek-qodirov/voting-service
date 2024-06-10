package handlers

import (
	"context"
	"net/http"

	pb "api-gateway/proto-service/genprotos"

	_ "api-gateway/api/docs"

	"github.com/gin-gonic/gin"
)

// GetPublicByID godoc
// @ID getpublicbyid
// @Router /public/{id} [GET]
// @Summary Get Public By ID
// @Description Get Public by ID
// @Tags Public
// @Accept json
// @Produce json
// @Param id path string true "Public ID"
// @Success 200 {object} pb.PublicGetById "Public data"
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Public not found"
// @Failure 500 {object} string "Server error"
func (h *HTTPHandler) GetPublicByID(c *gin.Context) {
	id := pb.PublicGetById{Id: c.Param("id")}
	res, err := h.Public.GetById(context.Background(), &id)
	if err != nil {
		c.String(500, "Couldn't get the public: "+err.Error())
		h.Logger.ERROR.Println("Couldn't get the public: " + err.Error())
		return
	}
	c.JSON(200, res)
}

// GetAllPublics godoc
// @ID getallpublics
// @Router /publics [GET]
// @Summary Get all publics
// @Description Get all publics
// @Tags Public
// @Accept json
// @Produce json
// @Success 200 {object} pb.PublicGetAllReq "Public data"
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Public not found"
// @Failure 500 {object} string "Server error"
func (h *HTTPHandler) GetAllPublics(c *gin.Context) {
	req := pb.PublicGetAllReq{
		PartyId:  c.Query("party_id"),
		Name:     c.Query("name"),
		LastName: c.Query("last_name"),
		Phone:    c.Query("phone"),
		Email:    c.Query("email"),
		Birthday: c.Query("birthday"),
		Gender:   c.Query("gender"),
	}
	res, err := h.Public.GetAll(context.Background(), &req)
	if err != nil {
		c.String(500, "Couldn't get publics: "+err.Error())
		h.Logger.ERROR.Println("Couldn't get publics: " + err.Error())
		return
	}
	c.JSON(200, res)
}

// CreatePublic godoc
// @ID createpublics
// @Router /public/ [POST]
// @Summary create public
// @Description create public
// @Tags Public
// @Accept json
// @Produce json
// @Param public body pb.PublicCreate true "Public data"
// @Success 200 {object} pb.PublicCreate "Public data"
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Public not found"
// @Failure 500 {object} string "Server error"
func (h *HTTPHandler) CreatePublic(c *gin.Context) {
	req := pb.PublicCreate{}
	err := c.BindJSON(&req)
	if err != nil {
		c.String(500, "Couldn't bind the data: "+err.Error())
		h.Logger.ERROR.Println("Couldn't bind the data: " + err.Error())
		return
	}
	_, err = h.Public.Create(context.Background(), &req)
	if err != nil {
		c.String(500, "Couldn't create the public: "+err.Error())
		h.Logger.ERROR.Println("Couldn't create the public: " + err.Error())
		return
	}
	c.String(http.StatusCreated, "Public created")
	h.Logger.INFO.Printf("Public created")
}

// UpdatePublic godoc
// @ID updatepublic
// @Router /public/{id} [PUT]
// @Summary Update public
// @Description Update public
// @Tags Public
// @Accept json
// @Produce json
// @Param id path string true "Public ID"
// @Param public body pb.PublicUpdate true "Public data"
// @Success 200 {object} pb.PublicUpdate "Public data"
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Public not found"
// @Failure 500 {object} string "Server error"
func (h *HTTPHandler) UpdatePublic(c *gin.Context) {
	id := c.Param("id")
	req := pb.PublicUpdate{}
	err := c.BindJSON(&req)
	if err != nil {
		c.String(500, "Couldn't bind the data: "+err.Error())
		h.Logger.ERROR.Println("Couldn't bind the data: " + err.Error())
		return
	}
	req.Id = id
	_, err = h.Public.Update(context.Background(), &req)
	if err != nil {
		c.String(500, "Couldn't update the public: "+err.Error())
		h.Logger.ERROR.Println("Couldn't update the public: " + err.Error())
		return
	}
	c.String(http.StatusCreated, "Public updated")
	h.Logger.INFO.Printf("Public updated: %s", req.GetId())
}

// Deletepublic godoc
// @ID deletepublic
// @Router /public/{id} [DELETE]
// @Summary Delete public
// @Description Delete public
// @Tags Public
// @Accept json
// @Produce json
// @Param id path string true "Public ID"
// @Success 200 {object} pb.PublicDelete "Public data"
// @Failure 400 {object} string "Bad Request"
// @Failure 404 {object} string "Public not found"
// @Failure 500 {object} string "Server error"
func (h *HTTPHandler) DeletePublic(c *gin.Context) {
	req := pb.PublicDelete{Id: c.Param("id")}
	_, err := h.Public.Delete(context.Background(), &req)
	if err != nil {
		c.String(500, "Couldn't delete the public: "+err.Error())
		h.Logger.ERROR.Println("Couldn't delete the public: " + err.Error())
		return
	}
	c.String(http.StatusCreated, "Public deleted")
	h.Logger.INFO.Printf("Public deleted: %s", req.GetId())
}
