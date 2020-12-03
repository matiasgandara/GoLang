package motos

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// HTTPService ...
type HTTPService interface {
	Register(*gin.Engine)
}

type endpoint struct {
	method   string
	path     string
	function gin.HandlerFunc
}

type httpService struct {
	endpoints []*endpoint
}

func NewHTTPTransport(s Service) HTTPService {
	endpoints := makeEndpoints(s)
	return httpService{endpoints}
}

func makeEndpoints(s Service) []*endpoint {
	list := []*endpoint{}

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/motos",
		function: getAll(s),
	})

	list = append(list, &endpoint{
		method:   "GET",
		path:     "/motos/:id",
		function: getOne(s),
	})

	list = append(list, &endpoint{
		method:   "POST",
		path:     "/motos",
		function: create(s),
	})

	list = append(list, &endpoint{
		method:   "PUT",
		path:     "/motos/:id",
		function: update(s),
	})

	list = append(list, &endpoint{
		method:   "DELETE",
		path:     "/motos/:id",
		function: delete(s),
	})

	return list
}

func getAll(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		m, _ := s.FindAll()
		c.JSON(http.StatusOK, gin.H{
			"motos": m,
		})
	}
}

func getOne(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		f, err := s.FindByID(id)
		m := http.StatusOK
		if f == nil {
			m = http.StatusNotFound
		}
		c.JSON(m, gin.H{
			"motos": f,
		})
	}
}

func create(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var m Moto
		err := c.BindJSON(&m)
		if err != nil {
			panic(err)
		}
		s.AddMoto(m)
		c.JSON(http.StatusCreated, gin.H{
			"message": "Moto creada",
		})
	}
}

func update(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		var m Moto
		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		m.ID = id
		err = c.BindJSON(&m)
		if err != nil {
			panic(err)
		}
		s.Update(m)
		c.JSON(http.StatusOK, gin.H{
			"message": "ID " + i + " modificada",
		})
	}
}

func delete(s Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		i := c.Param("id")
		id, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		s.Delete(id)
		c.JSON(http.StatusOK, gin.H{
			"message": "ID " + i + " eliminada",
		})
	}
}

// Register ...
func (s httpService) Register(r *gin.Engine) {
	for _, e := range s.endpoints {
		r.Handle(e.method, e.path, e.function)
	}
}
