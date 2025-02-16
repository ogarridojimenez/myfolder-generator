/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Nombre del proyecto (se puede pasar como argumento)
		projectName := "myproject"
		if len(args) > 0 {
			projectName = args[0]
		}

		// Crear la estructura de carpetas
		folders := []string{
			"cmd",
			"internal",
			"internal/core",
			"internal/core/domain",
			"internal/core/services",
			"internal/core/services/domainEntity",
			"internal/core/ports",
			"internal/infrastructure",
			"internal/infrastructure/handlers",
			"internal/infrastructure/handlers/domainEntity",
			"internal/infrastructure/repositories",
			"internal/infrastructure/repositories/db",
			"internal/infrastructure/repositories/db/domainEntity",
			"internal/config",
			"pkg",
			"pkg/DB",
		}

		for _, folder := range folders {
			dir := filepath.Join(projectName, folder)
			if err := os.MkdirAll(dir, os.ModePerm); err != nil {
				fmt.Printf("Error creando la carpeta %s: %v\n", dir, err)
				return
			}
			fmt.Printf("Carpeta creada: %s\n", dir)
		}

		// Crear archivos base
		files := map[string]string{
			"cmd/main.go": `package main
 
 import "fmt"
 
 func main() {
	 fmt.Println("Hola, mundo!")
 }`,
			"internal/core/domain/domainEntity.go": `package domain
 
 type domainEntity struct {
	 ID    string
  }`,
			"internal/core/ports/domainEntity.go": `package ports
 
 import "github.com/yourusername/yourproject/internal/core/domain"
 
 type domainEntityService interface {
	 Create(domainEntity *domain.domainEntity) error
	 Get(id string) (*domain.domainEntity, error)
	 Update(domainEntity *domain.domainEntity) error
	 Delete(id string) error
 }
	 type domainEntityRepository interface {
	 Insert(domainEntity *domain.domainEntity) error
	 FindByID(id string) (*domain.domainEntity, error)
	 Update(domainEntity *domain.domainEntity) error
	 Delete(id string) error
 }`,
			"internal/core/services/domainEntity/service.go": `package domainEntity
 
 import (
	  "github.com/yourusername/yourproject/internal/core/ports"
 )
 
 type Service struct {
	 Repo ports.domainEntityRepository
 }
 
 func NewdomainEntityService(repo ports.domainEntityRepository) *Service {
	 return &Service{Repo: repo}
 }`,
			"internal/core/services/domainEntity/create.go": `package domainEntity
 
 import (
	 "github.com/yourusername/yourproject/internal/core/domain"
	 "fmt"
	"log"
 )
 
 func (s *Service) Create(domainEntity *domain.domainEntity) (err error) {
	
	//TODO: save the repo
	 err = s.Repo.Insert(domainEntity)
	 if err != nil {
		log.Println(err.Error())
		return fmt.Errorf("error creating domainEntity:%w", err)
	 }
	return nil
	}`,
			"internal/core/services/domainEntity/delete.go": `package domainEntity
 
 import (
	 "github.com/yourusername/yourproject/internal/core/domain"
	 "fmt"
	 "log"
	 
 )
 // Delete domainEntity by id
 func (s *Service) Delete(id string) (err error) {
	
	err = s.Repo.Delete(id)
	if err != nil {
	    log.Println(err.Error())	
	    return fmt.Errorf("Error deleting domainEntity:%w",err)
	}
	return nil
	}`,
			"internal/core/services/domainEntity/get.go": `package domainEntity
 
 import (
	 "github.com/yourusername/yourproject/internal/core/domain"
	 "fmt"
	 "log"
	 "errors"
 )
 // Get domainEntity by id
 func (s *Service) Get(id string) (domainEntity *domain.domainEntity, err error) {
	
	if id == "" {
		return nil, errors.New("id is required")
	}

	domainEntity, err = s.Repo.Get(id)
	if err != nil {
		log.Println(err.Error())
		return nil, fmt.Errorf("unexpected error getting domainEntity: %w", err)
	}

	return domainEntity, nil
	}`,
			"internal/infrastructure/handlers/domainEntity/handlers.go": `package domainEntity
 
 import "github.com/yourusername/yourproject/internal/core/ports"
	 
 
 type Handler struct {
	domainEntityService ports.domainEntityService
}`,
			"internal/infrastructure/handlers/domainEntity/create.go": `package domainEntity
 
import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/yourproject/internal/core/domain"
)
	

func (h Handler) CreatedomainEntity(c *gin.Context) {
	//TODO: funciones de un handlers
	//TODO: validación
	//TODO: 1- traducir un request
	//TODO: 2- consumir un servicio
	//TODO: 3- traducir un response

	//TODO: Paso # 1 - Traducir un request a un formato que puedas usar como input para tu servicio
	var domainEntityCreateParams domain.domainEntity
	if err := c.BindJSON(&pdomainEntityCreateParams); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//TODO: =============Paso # 2 Consumir el servicio =====================

	insertedId, err := h.domainEntityService.Create(domainEntityCreateParams)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ooops!"})
		return 
	}
	//TODO: ===============================

	//TODO: Paso # 3 traducción de la respuesta
	c.JSON(200, gin.H{"domainEntity_id": insertedId})

}`, "internal/infrastructure/handlers/domainEntity/delete.go": `package domainEntity
 
import (
	
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	
)
	

func (h Handler) DeletedomainEntity(c *gin.Context) {
	id := c.Param("id")

	err := h.domainEntityService.Delete(id)
	if err != nil {
		log.Println(err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}`, "internal/infrastructure/handlers/domainEntity/get.go": `package domainEntity
 
import (
	
	"log"
	"github.com/gin-gonic/gin"
	
)
	

func (h Handler) GetdomainEntity(c *gin.Context) {
	domainEntityIdParam := c.Param("id")
	domainEntity, err := h.domainEntityService.Get(domainEntityIdParam)
	if err != nil {
		log.Println(err.Error())
		return
	}

	c.JSON(200, domainEntity)
}`,
		}

		for filePath, content := range files {
			fullPath := filepath.Join(projectName, filePath)
			if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
				fmt.Printf("Error creando el archivo %s: %v\n", fullPath, err)
				return
			}
			fmt.Printf("Archivo creado: %s\n", fullPath)
		}

		fmt.Println("Estructura generada exitosamente!")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
