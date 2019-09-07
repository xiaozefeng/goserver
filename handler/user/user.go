package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
	"github.com/xiaozefeng/goserver/handler"
	"github.com/xiaozefeng/goserver/model"
	"github.com/xiaozefeng/goserver/pkg/errno"
	"github.com/xiaozefeng/goserver/service"
	"github.com/xiaozefeng/goserver/util"
	"strconv"
	"time"
)

// Create creates a new uservo account
func Create(c *gin.Context) {
	log.Info("User create function called.", lager.Data{"X-Request-Id": util.GetSorterId()})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}
	u.CreatedTime = time.Now()
	u.UpdateTime= time.Now()
	u.DeletedTime= time.Now()

	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	if err := u.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, CreateResponse{
		Username: r.Username,
	})
}

// Delete delete an uservo by the uservo identifier.
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUserById(uint64(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}

func Get(c *gin.Context) {
	username := c.Param("username")
	// get the uservo by the `username` from the database
	userModel, err := model.GetUserByUsername(username)
	if err != nil {
		handler.SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	handler.SendResponse(c, nil, userModel)
}

// Update update a exist uservo account info.
func Update(c *gin.Context) {
	log.Info("update function called. ", lager.Data{"X-Request-Id": util.GetSorterId()})
	// Get the uservo id from the url parameter.
	userId, _ := strconv.Atoi(c.Param("id"))

	// Binding the uservo data
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// we update the record based on the uservo id .
	u.Id = uint64(userId)

	// Validate the data
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the uservo password.
	if err := u.Encrypt(); err != nil {
		handler.SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// save changed fields.
	if err := u.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, nil)
}

func List(c *gin.Context) {
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	infoList, total, err := service.ListUser(r.Username, r.Offset, r.Limit)
	if err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	res := ListResponse{
		TotalCount: total,
		UserList:   infoList,
	}
	handler.SendResponse(c, nil, res)
}
