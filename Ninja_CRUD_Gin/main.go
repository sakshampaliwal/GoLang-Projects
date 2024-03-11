package main

import (
	"github.com/gin-gonic/gin"
)

// ninjaweapons is a map to store ninja weapons, where the key is the weapon type and the value is its name.
var ninjaweapons = map[string]string{
	"ninjaStar": "Beginner Ninja Star - Damage 1",
}

// GetWeapon handles GET requests to retrieve information about a specific ninja weapon.
func GetWeapon(c *gin.Context) {
	// Retrieve the weapon type from the query parameters.
	weapontype := c.Query("type")
	// Check if the weapon type exists in the ninjaweapons map.
	weaponname, ok := ninjaweapons[weapontype]
	if !ok {
		// If the weapon type doesn't exist, return a JSON response with a 404 status indicating that the weapon was not found.
		c.JSON(404, gin.H{
			"error": "weapon not found",
		})
		return
	}
	// If the weapon type exists, return a JSON response with a 200 status containing information about the weapon.
	c.JSON(200, gin.H{
		"weaponType": weaponname,
	})
}

// PostWeapon handles POST requests to add a new ninja weapon to the ninjaweapons map.
func PostWeapon(c *gin.Context) {
	// Retrieve the weapon type and name from the query parameters.
	weapontype := c.Query("type")
	weaponname := c.Query("name")
	// Check if both weapon type and name are provided.
	if len(weapontype) == 0 || len(weaponname) == 0 {
		// If either the weapon type or name is missing, return a JSON response with a 400 status indicating the error.
		c.JSON(400, gin.H{
			"error": "both weapon type and name are required",
		})
		return
	}
	// Check if the weapon type already exists in the ninjaweapons map.
	if _, ok := ninjaweapons[weapontype]; ok {
		// If the weapon type already exists, return a JSON response with a 409 status indicating that the weapon already exists.
		c.JSON(409, gin.H{
			"error": "weapon already exists",
		})
		return
	}
	// Add the new weapon to the ninjaweapons map.
	ninjaweapons[weapontype] = weaponname
	// Return a JSON response with a 201 status indicating that the weapon was added successfully.
	c.JSON(201, gin.H{
		"message": "weapon added successfully",
	})
}

func PatchWeapon(c *gin.Context) {
	// Retrieve the weapon type and new name from the query parameters.
	weapontype := c.Query("type")
	newWeaponName := c.Query("newName")
	
	// Check if both weapon type and new name are provided.
	if len(weapontype) == 0 || len(newWeaponName) == 0 {
		// If either the weapon type or new name is missing, return a JSON response with a 400 status indicating the error.
		c.JSON(400, gin.H{
			"error": "both weapon type and new name are required",
		})
		return
	}
	
	// Check if the weapon type exists in the ninjaweapons map.
	_, ok := ninjaweapons[weapontype]
	if !ok {
		// If the weapon type doesn't exist, return a JSON response with a 404 status indicating that the weapon was not found.
		c.JSON(404, gin.H{
			"error": "weapon not found",
		})
		return
	}
	
	// Update the name of the weapon with the provided new name.
	ninjaweapons[weapontype] = newWeaponName
	
	// Return a JSON response with a 200 status indicating that the weapon was updated successfully.
	c.JSON(200, gin.H{
		"message": "weapon updated successfully",
	})
}

// DeleteWeapon handles DELETE requests to remove a ninja weapon from the ninjaweapons map.
func DeleteWeapon(c *gin.Context) {
	// Retrieve the weapon type from the query parameters.
	weapontype := c.Query("type")
	// Check if the weapon type exists in the ninjaweapons map.
	weaponname, ok := ninjaweapons[weapontype]
	if !ok {
		// If the weapon type doesn't exist, return a JSON response with a 404 status indicating that the weapon was not found.
		c.JSON(404, gin.H{
			"error": "weapon not found",
		})
		return 
	}
	// If the weapon type exists, delete it from the ninjaweapons map.
	delete(ninjaweapons, weapontype)
	// Return a JSON response with a 200 status indicating that the weapon was deleted successfully, along with information about the deleted weapon.
	c.JSON(200, gin.H{
		"message": "weapon deleted successfully",
		"deleted": weaponname,
	})
}

func main() {
	// Create a new Gin router.
	r := gin.Default()

	// Define a route for handling GET requests to check if the server is running.
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Define routes for handling GET, POST, and DELETE requests related to ninja weapons.
	r.GET("/weapon", GetWeapon)
	r.POST("/weapon", PostWeapon)
	r.DELETE("/weapon", DeleteWeapon)

	// Start the Gin server on port 8080.
	r.Run(":8080")
}
