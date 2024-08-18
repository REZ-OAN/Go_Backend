package api

import (
	"database/sql"
	"net/http"
	db "simple-bank/database/sqlc"

	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD BDT"`
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: db.Currency(req.Currency),
		Balance:  0,
	}
	account, err := server.store.CreateAccount(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest

	if err := ctx.ShouldBindUri(&req); err != nil {

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)

	if err != nil {

		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}

type listAccountsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var req listAccountsRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {

		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	accounts, err := server.store.ListAccounts(ctx, arg)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, accounts)
}

type updateRequestBody struct {
	Balance int64 `json:"balance" binding:"required,min=1"`
}

type updateRequestUri struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) updateAccount(ctx *gin.Context) {
	var reqBody updateRequestBody
	var reqUri updateRequestUri

	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindJSON(&reqBody); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateAccountParams{
		ID:      reqUri.ID,
		Balance: reqBody.Balance,
	}
	account, err := server.store.UpdateAccount(ctx, arg)

	if err != nil {

		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}
