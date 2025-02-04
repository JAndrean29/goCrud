package main

import (
	_ "database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID     int64  `db:"id"`
	Name   string `db:"name"`
	Age    int    `db:"age"`
	Gender string `db:"gender"`
}

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *User) (*User, error) {
	query := `insert into user (name,age,gender) values (?,?,?)`
	result, err := r.db.Exec(query, user.Name, user.Age, user.Gender)
	if err != nil {
		return nil, fmt.Errorf("Failed to create user: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("Error hit: Create")
	}

	user.ID = id
	return user, nil
}

func (r *UserRepo) FindById(id int64) (*User, error) {
	var user User
	query := `select * from user where id = ?`
	err := r.db.Get(&user, query, id)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch user: %w", err)
	}

	return &user, nil
}

func (r *UserRepo) FindByName(s string) (*User, error) {
	var user User
	query := `select * from user where lower(name) like '%' || lower(?) || '%'`
	err := r.db.Get(&user, query, s)
	if err != nil {
		fmt.Println(query)
		return nil, fmt.Errorf("Failed to fetch user: %w", err)
	}

	return &user, nil
}

func (r *UserRepo) GetByAgeSort(s string) ([]User, error) {
	var users []User
	query := `select * from user order by age ||?||`
	err := r.db.Select(&users, query, s)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch users data: %w", err)
	}

	return users, nil
}

func (r *UserRepo) ShowAll() ([]User, error) {
	var users []User
	query := `select * from user`
	err := r.db.Select(&users, query)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch users data: %w", err)
	}

	return users, nil
}

func (r *UserRepo) Update(user *User) (*User, error) {
	query := `update user set name = ?, age = ?, gender = ? where id = ?`
	result, err := r.db.Exec(query, user.Name, user.Age, user.Gender, user.ID)
	if err != nil {
		return nil, fmt.Errorf("Failed to update user data: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("Error Hit: Update")
	}
	if rows == 0 {
		return nil, fmt.Errorf("The following user ID does not exists / found: %d", user.ID)
	}
	return user, nil
}

func (r *UserRepo) Delete(id int64) error {
	query := `delete from user where id = ?`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("Failed to execute delete: %w", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Error hit: Delete")
	}
	if rows == 0 {
		return fmt.Errorf("The following ID does not exists / found: %d", id)
	}

	return nil
}

func main() {
	r := gin.Default()
	dbPath := "./users.db"

	db, err := sqlx.Connect("sqlite3", dbPath)
	if err != nil {
		log.Fatalf("Failed to connect to DB: ", err)
	}
	defer db.Close()

	UserRepo := NewUserRepo(db)

	//CREATE NEW USER
	r.POST("/users/create", func(c *gin.Context) {
		var req User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		createdUser, err := UserRepo.Create(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return
		}

		log.Println(createdUser)
		c.JSON(http.StatusCreated, gin.H{"message": "User Successfully created"})
	})

	//GET ALL USER
	r.GET("/users", func(c *gin.Context) {
		users, err := UserRepo.ShowAll()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, users)
	})

	//GET USER AGE SORT
	r.GET("/users/age/:sort", func(c *gin.Context) {
		sort := c.Param("sort")
		users, err := UserRepo.GetByAgeSort(sort)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, users)
	})

	//GET USER BY ID
	r.GET("/users/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, e := strconv.ParseInt(idParam, 10, 64)
		if e != nil {
			panic(e)
		}
		user, err := UserRepo.FindById(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		}

		c.JSON(http.StatusOK, user)
	})

	//GET USER BY NAME
	r.GET("/users/find/:name", func(c *gin.Context) {
		name := c.Param("name")

		user, err := UserRepo.FindByName(name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		}

		c.JSON(http.StatusOK, user)
	})

	//UPDATE USER
	r.POST("/users/edit", func(c *gin.Context) {
		var req User
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		}

		targetUser, err := UserRepo.Update(&req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		}

		log.Println(targetUser)
		c.JSON(http.StatusCreated, gin.H{"message": "User updated!"})
	})

	//DELETE USER
	r.DELETE("/users/delete/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			panic(err)
		}

		err = UserRepo.Delete(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{"message": "User successfully removed"})
	})

	//Run the localhost server
	r.Run()
}
