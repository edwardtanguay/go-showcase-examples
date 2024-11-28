package main

import "fmt"

func main() {
    users := map[string]map[string]string{
        "u123": {"name": "Alice", "email": "alice@example.com"},
        "u456": {"name": "Bob", "email": "bob@example.com"},
    }

    users["u789"] = map[string]string{
        "name":  "Charlie",
        "email": "charlie@example.com",
    }

    userID := "u456"
    if user, exists := users[userID]; exists {
        fmt.Printf("User %s: Name = %s, Email = %s\n", userID, user["name"], user["email"])
    } else {
        fmt.Printf("User with ID %s not found.\n", userID)
    }

    fmt.Println("\nAll users:")
    for id, userInfo := range users {
        fmt.Printf("ID: %s, Name: %s, Email: %s\n", id, userInfo["name"], userInfo["email"])
    }

    delete(users, "u123")
    fmt.Println("\nAfter deleting user u123:")

    _, exists := users["u123"]
    fmt.Printf("Does user u123 exist? %v\n", exists)
}