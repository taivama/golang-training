package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/taivama/golang-training/controllers"
	"github.com/taivama/golang-training/middleware"
)

func AddUnSecuredRoutes(g *gin.Engine, u *controllers.UserController) {
	users := g.Group("/api/users")
	users.POST("/", u.RegisterUser)
	users.POST("/login", u.Login)
}

func AddSecuredRoutes(g *gin.Engine, u *controllers.UserController) {
	g.Use(middleware.Authenticate())
	g.POST("/api/users/logout", u.Logout)
}

func AddProductRoutes(g *gin.Engine, p *controllers.ProductController) {
	product := g.Group("/api/products")
	product.POST("/", p.AddProduct)
	product.GET("/:id", p.GetProductById)
	product.GET("/search/:name", p.SearchProducts)
}
