package shopController

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tmsshopping/dao"
)

func ShopUpdate(ctx *gin.Context) {
	ctx.Header("Content-Type", "text/html;charset=utf-8")
	var (
		paramPid    string
		paramAction string
		//paramGetValue string
		pid int
	)

	paramPid = ctx.Query("pid")
	paramAction = ctx.Query("action")
	//paramGetValue = ctx.Query("getvalue")

	pid, err := strconv.Atoi(paramPid)
	if err != nil {
		ctx.Redirect(http.StatusFound, "/ShopSelect")
		return
	}

	switch paramAction {
	case "jia":
		_, _ = dao.PlusItem(pid)
	case "jian":
	case "closeText":
	case "delText":
	}

	ctx.Redirect(http.StatusFound, "/ShopSelect")
}
