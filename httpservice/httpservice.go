package httpservice

import (
	pwc "grab/week5/GrabGoTrainingWeek5Assignment/postwithcomment"
	"grab/week5/GrabGoTrainingWeek5Assignment/renderer"
	"log"
	"net/http"
)

type PostWithCommentsResponse struct {
	Posts []pwc.PostWithComments `json:"posts"`
}

type PostWithCommentHttpService struct {
	postWithCommentService pwc.PostWithCommentsInterface
	renderService          renderer.Renderer
}

func NewPostWithCommentHttpService(pwcService pwc.PostWithCommentsInterface, renderService renderer.Renderer) *PostWithCommentHttpService {
	service := &PostWithCommentHttpService{pwcService, renderService}
	return service
}

func (httpService *PostWithCommentHttpService) StartServer() {
	http.HandleFunc("/postWithComments", func(writer http.ResponseWriter, request *http.Request) {
		postWithComments, err := httpService.postWithCommentService.GetPostWithComments()
		if err != nil {
			log.Println("unable to get post with comments: ", err)
			writer.WriteHeader(500)
			return
		}
		resp := PostWithCommentsResponse{Posts: postWithComments}
		buf, contentType, err := httpService.renderService.Render(resp)
		if err != nil {
			log.Println("unable to render response: ", err)
			writer.WriteHeader(500)
			return
		}

		writer.Header().Set("Content-Type", contentType)
		_, err = writer.Write(buf)
	})

	log.Println("httpServer starts ListenAndServe at 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
