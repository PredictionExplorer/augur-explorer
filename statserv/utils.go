package main
import (
	"fmt"
	"strconv"
	"net/http"
	"github.com/gin-gonic/gin"

)
func respond_error_html(c *gin.Context,error_text string) {

	c.HTML(http.StatusBadRequest, "error.html", gin.H{
		"title": "Augur Markets: Error",
		"ErrDescr": error_text,
	})
}
func respond_error_json(c *gin.Context,error_text string) {

	c.JSON(http.StatusBadRequest, gin.H{
		"status": 0,
		"error": error_text,
	})
}
func parse_integer_param_or_error(c *gin.Context,param_name string,is_json bool) (int64,bool) {

	p_param := c.Param(param_name)
	var param int64
	if len(p_param) > 0 {
		var err error
		param, err = strconv.ParseInt(p_param,10,64)
		if err != nil {
			if is_json {
				c.JSON(http.StatusBadRequest, gin.H{
					"status" : 0,
					"error": fmt.Sprintf("Can't parse integer parameter : ",err),
				})
			} else {
				c.HTML(http.StatusBadRequest, "error.html", gin.H{
					"title": "Error",
					"ErrDescr": fmt.Sprintf("Can't parse integer parameter : ",err),
				})
			}
			return 0,false
		}
	} else {
		if is_json {
			respond_error_json(c,"'status' parameter is not set")
		} else {
			respond_error_html(c,"'status' parameter is not set")
		}
		return 0,false
	}
	return param,true
}
