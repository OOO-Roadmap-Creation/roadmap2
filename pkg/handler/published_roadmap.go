package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

// @Summary Получение одного роадмапа
// @Tags published-roadmap-controller
// @Description ...
// @ID one-roadmap
// @Accept json
// @Produce  json
// @Param        id   path      int  true  "Roadmap ID"
// @Success 200 {object} publishedRoadmapResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /published-roadmap/{id}  [GET]
func (h *Handler) oneRoadmap(c *gin.Context, rmIdStr string) {
	rmId, err := strconv.Atoi(rmIdStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid roadmap id param")
		return
	}

	rm, err := h.repo.PR.GetById(rmId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(rm)

	// TODO read from db
	c.JSON(http.StatusOK, publishedRoadmapResponse{
		Id:            rm.Id,
		Version:       rm.Version,
		Visible:       rm.Visible,
		Title:         rm.Title,
		Description:   rm.Description,
		DateOfPublish: rm.DateOfPublish,
	})
}

// @Summary Получение списка роадмапов
// @Tags published-roadmap-controller
// @Description ...
// @ID list-roadmap
// @Accept json
// @Produce  json
// @Success 200 {object} map[string][]publishedRoadmapResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /published-roadmap  [GET]
func (h *Handler) list(c *gin.Context) {
	all, err := h.repo.PR.GetAll()
	if err != nil {
		fmt.Println(err)
		//return
	}
	fmt.Println(all)
	ans := map[string][]publishedRoadmapResponse{}
	for _, r := range all {
		ans[string(r.Id)] = []publishedRoadmapResponse{publishedRoadmapResponse{
			Id:            r.Id,
			Version:       r.Version,
			Visible:       r.Visible,
			Title:         r.Title,
			Description:   r.Description,
			DateOfPublish: r.DateOfPublish,
		},
		}
	}
	c.JSON(http.StatusOK, ans)
	//c.JSON(http.StatusOK, map[string][]publishedRoadmapResponse{
	//	"additionalProp1": {publishedRoadmapResponse{
	//		Id:            9,
	//		Version:       3,
	//		Visible:       true,
	//		Title:         "EXAMPLE",
	//		Description:   "LA LA LA",
	//		DateOfPublish: time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC),
	//	}},
	//	"additionalProp2": {
	//		publishedRoadmapResponse{
	//			Id:            34,
	//			Version:       3,
	//			Visible:       true,
	//			Title:         "EXAMPLE",
	//			Description:   "LA LA LA",
	//			DateOfPublish: time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC),
	//		}},
	//	"original": {
	//		publishedRoadmapResponse{
	//			Id:            all[0].Id,
	//			Version:       all[0].Version,
	//			Visible:       all[0].Visible,
	//			Title:         all[0].Title,
	//			Description:   all[0].Description,
	//			DateOfPublish: all[0].DateOfPublish,
	//		},
	//	},
	//},
	//)
}

// @Summary Изменение видимости роадмапа
// @Tags published-roadmap-controller
// @Description ...
// @ID visibility
// @Accept json
// @Produce  json
// @Param        id   path      int  true  "Roadmap ID"
// @Param        visibility   path      bool  true  "new visible status"
// @Success 200 {object} publishedRoadmapResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /published-roadmap/{id}/{visibility}  [PUT]
func (h *Handler) visibility(c *gin.Context) {
	rmId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid roadmap id param")
		return
	}

	v, err := strconv.ParseBool(c.Param("v"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid roadmap id param")
		return
	}

	// TODO read from db
	c.JSON(http.StatusOK, publishedRoadmapResponse{
		Id:            rmId,
		Version:       3,
		Visible:       v,
		Title:         "EXAMPLE",
		Description:   "LA LA LA",
		DateOfPublish: time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC),
	})
}

func (h *Handler) common(c *gin.Context) {
	p1 := c.Param("p1")
	p2 := c.Param("p2")
	if p1 == "roadmap" && p2 == "personal" {
		h.list(c)
	} else {
		h.oneRoadmap(c, p1)
	}
}

// @Summary Получение списка нод роадмапа
// @Tags published-node-controller
// @Description ...
// @ID nodes
// @Accept json
// @Produce  json
// @Param        id   path      int  true  "Roadmap ID"
// @Success 200 {object} []publishedNodeResponse
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /published-node/roadmap/  [GET]
func (h *Handler) nodes(c *gin.Context) {
	rmId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid roadmap id param")
		return
	}
	// TODO read from db
	c.JSON(http.StatusOK, []publishedNodeResponse{
		publishedNodeResponse{
			Id:          4,
			Title:       "NODE 959",
			Description: "LA LA LA",
			Priority:    1,
			//ParentId:    nil,
		},
		publishedNodeResponse{
			Id:          rmId,
			Title:       "EXAMPLE",
			Description: "LA LA LA",
			Priority:    0,
			ParentId:    4,
		},
	},
	)
}

// @Summary Получение рейтинга роадмапа
// @Tags rating-controller
// @Description ...
// @ID get-rating
// @Accept json
// @Produce  json
// @Param        roadmap_id   path      int  true  "Roadmap ID"
// @Success 200 {int} 0
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /rating/roadmap/{roadmap_id} [GET]
func (h *Handler) getRating(c *gin.Context) {
	rmId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid roadmap id param")
		return
	}
	// TODO read from db
	c.JSON(http.StatusOK, rmId*2)
}

// @Summary Получение оценки роадмапа пользователем
// @Tags rating-controller
// @Description ...
// @ID get-rating-user
// @Accept json
// @Produce  json
// @Param        roadmap_id   path      int  true  "Roadmap ID"
// @Success 200 {bool} true
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /rating/roadmap/{roadmap_id}/user [GET]
func (h *Handler) getRatingUser(c *gin.Context) {
	_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid roadmap id param")
		return
	}
	// TODO read from db
	c.JSON(http.StatusOK, true)
}

// @Summary Лайк роадмапа (установка\снятие)
// @Tags rating-controller
// @Description ...
// @ID set-rating
// @Accept json
// @Produce  json
// @Param        roadmap_id   path      int  true  "Roadmap ID"
// @Success 200 {bool} true
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /rating/roadmap/{roadmap_id} [POST]
func (h *Handler) setRating(c *gin.Context) {
	_, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid roadmap id param")
		return
	}
	// TODO read from db
	c.JSON(http.StatusOK, true)
}
