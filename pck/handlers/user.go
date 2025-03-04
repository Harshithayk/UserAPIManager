package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"todo/pck/models"

	"github.com/gin-gonic/gin"
)

func (h Handlers) UserSignup(c *gin.Context) {
	ctx := c.Request.Context()
	var sigupUser models.UsersModel
	err := json.NewDecoder(c.Request.Body).Decode(&sigupUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "please enter proper request",
		})
		return
	}
	users, err := h.Ser.UserSignup(ctx, sigupUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, users)

}

func (h Handlers) LoginPage(c *gin.Context) {
	ctx := c.Request.Context()
	var UserLogin models.UserLogin

	err := json.NewDecoder(c.Request.Body).Decode(&UserLogin)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "please enter proper request",
		})
		return
	}
	res, err := h.Ser.UserLogin(ctx, UserLogin)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "please enter the valid password",
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h Handlers) FetchAllUser(c *gin.Context) {
	ctx := c.Request.Context()
	fetch, err := h.Ser.FetchUser(ctx)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "please enter the valid password",
		})
		return

	}
	c.JSON(http.StatusOK, fetch)
}

func (h Handlers) FetchById(c *gin.Context) {
	var FetchId models.FetchByID
	ctx := c.Request.Context()
	err := json.NewDecoder(c.Request.Body).Decode(&FetchId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "please enter proper request",
		})
		return
	}
	id, err := h.Ser.FetchById(ctx, FetchId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "can not fetch the user by id",
		})
		return
	}
	c.JSON(http.StatusOK, id)
}

func (h Handlers) UpdateUserById(c *gin.Context) {
	var FetchId models.FetchUser
	ctx := c.Request.Context()
	err := json.NewDecoder(c.Request.Body).Decode(&FetchId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "please provid valid id",
		})
		return
	}
	id, err := h.Ser.UpdateUserById(ctx, FetchId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "can not fetch the user by id",
		})
		return
	}
	c.JSON(http.StatusOK, id)
}

func (h Handlers) DeleteById(c *gin.Context) {
	//	var FetchByID models.FetchByID
	ctx := c.Request.Context()
	id := c.Param("id")
	id1, _ := strconv.Atoi(id)
	// err := json.NewDecoder(c.Request.Body).Decode(&FetchByID)
	// if err != nil {
	//     c.AbortWithStatusJSON(http.StatusBadRequest,gin.H{
	// 		"Message": "please provide the valid request",
	// 	})
	// 	return
	// }
	resp, err := h.Ser.DeleteById(ctx, id1)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "can not fetch the user by id",
		})
		return
	}
	c.JSON(http.StatusOK, resp)

}
