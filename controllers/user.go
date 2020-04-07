package controllers

import (
	"encoding/json"
	"fmt"
	
	"github.com/gin-gonic/gin"
	"github.com/gin-lab/models"
)

func CreateUser(c *gin.Context) {

	userModel := models.User{}
	if err := c.ShouldBindJSON(&userModel); err != nil {
		resMap := map[string]interface{}{
			"success": false,
			"message": "Data invalid.",
		}

		resByte, _ := json.Marshal(resMap)

		c.Data(400, "application/json; charset=utf-8", resByte)
		c.Abort()

		return
	  }

	  if userModel.FirstName == "" || userModel.LastName == "" {
		resMap := map[string]interface{}{
			"success": false,
			"message": "Data invalid.",
		}

		resByte, _ := json.Marshal(resMap)

		c.Data(400, "application/json; charset=utf-8", resByte)
		c.Abort()

		return
	  }

		if createUserErr := models.DB.Create(&userModel).Error; createUserErr != nil {
			resMap := map[string]interface{}{
				"success": false,
				"message": "Create user error.",
			}
	
			resByte, _ := json.Marshal(resMap)
	
			c.Data(500, "application/json; charset=utf-8", resByte)
			c.Abort()
	
			return
		}
	
	resMap := map[string]interface{}{
		"success": true,
	}

	resByte, _ := json.Marshal(resMap)

	c.Data(200, "application/json; charset=utf-8", resByte)
	c.Abort()

	return
	
}

func GetUser(c *gin.Context) {
	id := c.Param("id")
	userModel := models.User{}

	if err := models.DB.Find(&userModel, id).Error; err != nil {
		resMap := map[string]interface{}{
			"success": false,
			"message": "User not found.",
		}
		resByte, _ := json.Marshal(resMap)

		c.Data(404, "application/json; charset=utf-8", resByte)
		c.Abort()

		return
	}

	resMap := userModel

	resByte, _ := json.Marshal(resMap)

	c.Data(200, "application/json; charset=utf-8", resByte)
	c.Abort()

	return
}

func GetAllUser(c *gin.Context) {
	userModel := []models.User{}

	if err := models.DB.Find(&userModel).Error; err != nil {
		resMap := map[string]interface{}{
			"success": false,
			"message": "No data.",
		}
		resByte, _ := json.Marshal(resMap)

		c.Data(404, "application/json; charset=utf-8", resByte)
		c.Abort()

		return
	}

	resMap := userModel

	resByte, _ := json.Marshal(resMap)

	c.Data(200, "application/json; charset=utf-8", resByte)
	c.Abort()

	return
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userModel := models.User{}

	if err := models.DB.Find(&userModel, id).Error; err != nil {
		resMap := map[string]interface{}{
			"success": false,
			"message": "User not found.",
		}
		resByte, _ := json.Marshal(resMap)

		c.Data(404, "application/json; charset=utf-8", resByte)
		c.Abort()

		return
	}

	if err := models.DB.Delete(&userModel).Error; err != nil {
		resMap := map[string]interface{}{
			"success": false,
			"message": "Delete error.",
		}
		resByte, _ := json.Marshal(resMap)

		c.Data(500, "application/json; charset=utf-8", resByte)
		c.Abort()

		return
	}

	resMap := map[string]interface{}{
		"success": true,
	}

	resByte, _ := json.Marshal(resMap)

	c.Data(200, "application/json; charset=utf-8", resByte)
	c.Abort()

	return
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")
	userModel := models.User{}

	if err := models.DB.First(&userModel, id).Error; err != nil {
		resMap := map[string]interface{}{
			"success": false,
			"message": "User not found.",
		}
		resByte, _ := json.Marshal(resMap)

		c.Data(404, "application/json; charset=utf-8", resByte)
		c.Abort()

		return
	}

	if err := c.ShouldBindJSON(&userModel); err != nil {
		resMap := map[string]interface{}{
			"success": false,
			"message": "No content.",
		}

		resByte, _ := json.Marshal(resMap)

		c.Data(400, "application/json; charset=utf-8", resByte)
		c.Abort()

		return
	  }

	  if userModel.FirstName == "" || userModel.LastName == "" {
		resMap := map[string]interface{}{
			"success": false,
			"message": "Data invalid.",
		}

		resByte, _ := json.Marshal(resMap)

		c.Data(400, "application/json; charset=utf-8", resByte)
		c.Abort()

		return
	  }
	 
	  if updateErr := models.DB.Save(&userModel).Error; updateErr != nil {
		resMap := map[string]interface{}{
			"success": false,
			"message": "Update failed.",
		}

		resByte, _ := json.Marshal(resMap)

		c.Data(500, "application/json; charset=utf-8", resByte)
		c.Abort()

		return
	  }

	resMap := map[string]interface{}{
		"success": true,
	}

	resByte, _ := json.Marshal(resMap)

	c.Data(200, "application/json; charset=utf-8", resByte)
	c.Abort()

	return
}

func RecoverUser(c *gin.Context) {
	id := c.Param("id")
	userModel := models.User{}

	if err := models.DB.Unscoped().Find(&userModel, id).Error; err != nil {
		resMap := map[string]interface{}{
			"success": false,
			"message": "Deleted User not found.",
		}
		resByte, _ := json.Marshal(resMap)

		c.Data(404, "application/json; charset=utf-8", resByte)
		c.Abort()

		return
	}

	userModel.DeletedAt = nil
	if deleteErr := models.DB.Unscoped().Save(&userModel).Error; deleteErr != nil {
		resMap := map[string]interface{}{
			"success": false,
			"message": "Recover failed",
		}

		resByte, _ := json.Marshal(resMap)

		c.Data(500, "application/json; charset=utf-8", resByte)
		c.Abort()

		return
	}

	resMap := map[string]interface{}{
		"success": true,
	}

	resByte, _ := json.Marshal(resMap)

	c.Data(200, "application/json; charset=utf-8", resByte)
	c.Abort()

	return
}

func commandExample() {
	/* Start Map Example */
	mapExample := map[string]interface{}{
		"stringja": "eueu",
		"intja":    2,
	}

	tmpString := mapExample["stringja"].(string)
	tmpInt := mapExample["intja"].(int)
	fmt.Println("tmpString: ", tmpString)

	fmt.Println("tmpInt: ", tmpInt)
	/* End Map Example */

	/* Start Struct Example */
	userModel := models.User{}

	userModel.FirstName = "Tananut"
	userModel.LastName = "Panyagosa"
	userModel.Email = "tongtananut@gmail.com"

	fmt.Println("userModel: ", userModel)
	/* End Struct Example */

	/* Start Marshal Example */
	mapExample2 := map[string]interface{}{
		"first_name": "Sam",
		"last_name":  "Za",
		"email":      "samza55plus@gmail.com",
	}

	tmpByte, marshalErr := json.Marshal(mapExample2)
	if marshalErr != nil {
		fmt.Println("marshal err: ", marshalErr)
	}
	/* End Marshal Example */

	/* Start Unmarshal Example */
	inputMap := map[string]interface{}{}
	json.Unmarshal(tmpByte, &inputMap)

	userModel2 := models.User{}
	json.Unmarshal(tmpByte, &userModel2)

	fmt.Println("tmpByte: ", string(tmpByte))
	fmt.Println("inputMap: ", inputMap)
	fmt.Println("inputMap: ", userModel2)
	/* End Unmarshal Example */
}
