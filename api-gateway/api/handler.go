package api

import (
	_ "api-gateway/api/docs"
	"api-gateway/api/handlers"
	"api-gateway/config/logger"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
)

func NewGin(ConnP, ConnV *grpc.ClientConn, Logger *logger.Logger) *gin.Engine {
	handler := handlers.NewHandler(ConnP, ConnV, *Logger)
	r := gin.Default()
	r.GET("/api/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	public := r.Group("/public")
	public.GET("/:id", handler.GetPublicByID)
	public.POST("/", handler.CreatePublic)
	public.PUT("/:id", handler.UpdatePublic)
	public.DELETE("/:id", handler.DeletePublic)
	r.GET("/publics", handler.GetAllPublics)

	party := r.Group("/party")
	party.GET("/:id", handler.GetPartyByID)
	party.POST("/", handler.CreateParty)
	party.PUT("/:id", handler.UpdateParty)
	party.DELETE("/:id", handler.DeleteParty)
	r.GET("/parties", handler.GetAllParty)

	election := r.Group("/election")
	election.GET("/:id", handler.GetElectionByID)
	election.POST("/", handler.CreateElection)
	election.PUT("/:id", handler.UpdateElection)
	election.DELETE("/:id", handler.DeleteElection)
	election.GET("/by-date", handler.GetElectionsByDate)
	r.GET("/elections", handler.GetAllElections)

	candidate := r.Group("/candidate")
	candidate.GET("/:id", handler.GetByIdCandidate)
	candidate.POST("/", handler.CreateCandidate)
	candidate.PUT("/:id", handler.UpdateCandidate)
	candidate.DELETE("/:id", handler.DeleteCandidate)
	candidate.GET("/by-party", handler.GetAllCandidatesByParty)
	r.GET("/candidates", handler.GetAllCandidates)

	publicVote := r.Group("/public-vote")
	publicVote.POST("/create", handler.CreatePublicVote)
	publicVote.GET("/:id", handler.GetByIdPublicVotes)
	publicVote.PUT("/update", handler.UpdatePublicVotes)
	publicVote.DELETE("/:id", handler.DeletePublicVotes)
	publicVote.GET("/find-winner", handler.FindWinner)
	publicVote.GET("/by-election", handler.GetAllPublicVotesByElection)
	r.GET("/public-votes", handler.GetAllPublicVotes)
	return r
}
