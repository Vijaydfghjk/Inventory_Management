package controllers

import (
	models "inventory_management/Models"
	service "inventory_management/Service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	service service.OrderService
}

func Order_controll(service service.OrderService) *OrderController {

	return &OrderController{service: service}
}

func (o *OrderController) CreateOrder(c *gin.Context) {

	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": "In valid request"})
		return
	}

	if err := o.service.CreateOrder(&order); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (o *OrderController) GetOrder(c *gin.Context) {

	idparam := c.Param("id")

	id, err := strconv.Atoi(idparam)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": "invaild_id"})
		return
	}

	order, err := o.service.GetOrder(uint(id))

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": "order_not found"})
		return
	}

	c.JSON(http.StatusOK, order)

}

func (o *OrderController) GetAllOrders(c *gin.Context) {

	orders, err := o.service.GetAllOrders()

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Unable to fetch the data"})
		return
	}

	c.JSON(http.StatusOK, orders)
}

func (o *OrderController) UpdateOrder(c *gin.Context) {

	var order models.Order

	if err := c.ShouldBindJSON(&order); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": err.Error()})
		return
	}
	idparam := c.Param("id")

	id, err := strconv.Atoi(idparam)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": "invaild_id"})
		return
	}

	exist_order, err := o.service.GetOrder(uint(id))

	if err != nil {

		c.JSON(http.StatusNotFound, gin.H{"Message": "order not exist"})

		return
	}

	if order.CustomerName != "" {

		exist_order.CustomerName = order.CustomerName
	}

	if order.Status != "" {

		exist_order.Status = order.Status
	}

	if len(order.Items) == len(exist_order.Items) {
		for i, _ := range order.Items {

			exist_order.Items[i].Product = order.Items[i].Product
			exist_order.Items[i].Quantity = order.Items[i].Quantity
			exist_order.Items[i].UnitPrice = order.Items[i].UnitPrice
		}
	} else {

		c.JSON(http.StatusBadRequest, gin.H{"Message": "items count mismatch"})
		return
	}
	if err := o.service.UpdateOrder(exist_order); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": "Unable to update"})
		return
	}

	c.JSON(http.StatusOK, exist_order)
}

func (o *OrderController) DeleteOrder(c *gin.Context) {

	idparam := c.Param("id")

	id, err := strconv.Atoi(idparam)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"Message": "invaild_id"})
		return
	}

	if err := o.service.DeleteOrder(uint(id)); err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Recored has been deleted"})
}
