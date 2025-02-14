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
			"internal/core/ports",
			"internal/infrastructure",
			"internal/infrastructure/handlers",
			"internal/infrastructure/repositories",
			"internal/config",
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
			"internal/core/domain/user.go": `package domain
 
 type User struct {
	 ID    string
	 Name  string
	 Email string
 }`,
			"internal/core/ports/user_repository.go": `package ports
 
 import "github.com/yourusername/yourproject/internal/core/domain"
 
 type UserRepository interface {
	 Create(user *domain.User) error
	 FindByID(id string) (*domain.User, error)
	 Update(user *domain.User) error
	 Delete(id string) error
 }`,
			"internal/core/services/user_service.go": `package services
 
 import (
	 "github.com/yourusername/yourproject/internal/core/domain"
	 "github.com/yourusername/yourproject/internal/core/ports"
 )
 
 type UserService struct {
	 repo ports.UserRepository
 }
 
 func NewUserService(repo ports.UserRepository) *UserService {
	 return &UserService{repo: repo}
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
