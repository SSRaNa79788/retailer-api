package Routes

import(
	"retailer-api/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	r:=gin.Default()
	grp := r.Group("/")
	{
		//Product Controller
		grp.POST("product",Controllers.AddProduct)
		grp.PATCH("product/:id",Controllers.UpdateProduct)
		grp.GET("product/:id",Controllers.GetProduct)
		grp.GET("products",Controllers.GetProducts)

		//Order Controller
		grp.POST("order",Controllers.PlaceOrder)
		grp.GET("order/:id",Controllers.GetOrder)
	}
	return r
}
