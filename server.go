package main

import (
	"fmt"
	"net/http"
	"os"

	"goTut/controller"
	router "goTut/http"
	"goTut/repository"
	"goTut/service"
)

var (
	carDetailsService    service.CarDetailsService       = service.NewCarDetailsService()
	carDetailsController controller.CarDetailsController = controller.NewCarDetailsConroller(carDetailsService)
	postRepository       repository.PostRepository       = repository.NewFirestoreRepository()
	postService          service.PostService             = service.NewPostService(postRepository)
	postController       controller.PostController       = controller.NewPostController(postService)
	httpRouter           router.Router                   = router.NewChiRouter()
)

func main() {
	var port = ":8000"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Running")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.GET("/carDetails", carDetailsController.GetCarDetails)
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	httpRouter.Serve(port)
}
