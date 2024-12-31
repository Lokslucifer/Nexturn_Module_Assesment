package repository

import (
	"blog_server/model"
	"database/sql"
	"log"
)

type BlogRepository struct {
	DB *sql.DB
}

func NewBlogRepository(db *sql.DB) *BlogRepository {
	return &BlogRepository{DB: db}
}

func (repo *BlogRepository) CreateBlog(blog *model.Blog) (*model.Blog, error) {
	stmt, err := repo.DB.Prepare("INSERT INTO blogs (title,author,content,timestamp)VALUES(?,?,?,?)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(blog.Title, blog.Author, blog.Content, blog.Timestamp)
	if err != nil {
		return nil, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}
	blog.Id = int(id)
	return blog, nil

}

func (repo *BlogRepository) GetBlog(id int) (*model.Blog, error) {
	row := repo.DB.QueryRow("SELECT id,author,title,content FROM users WHERE id = ?", id)
	blog := &model.Blog{}
	err := row.Scan(&blog.Id, &blog.Author, &blog.Title, &blog.Content)
	if err != nil {
		return nil, err
	}
	return blog, nil
}

func (repo *BlogRepository) GetAllBlogs() ([]model.Blog, error) {
	rows, err := repo.DB.Query("SELECT id, title, author, content, timestamp FROM blogs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var blogs []model.Blog
	for rows.Next() {
		var blog model.Blog
		err := rows.Scan(&blog.Id, &blog.Title, &blog.Author, &blog.Content, &blog.Timestamp)
		if err != nil {
			log.Fatal(err)
		}
		blogs = append(blogs, blog)
	}
	return blogs, nil
}

func (repo *BlogRepository) UpdateBlog(blog *model.Blog) (*model.Blog, error) {
	stmt, err := repo.DB.Prepare("UPDATE blogs SET title = ?, author = ?, content = ?, timestamp = ? WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(blog.Title, blog.Author, blog.Content, blog.Timestamp, blog.Id)
	if err != nil {
		return nil, err
	}

	return blog, nil
}

func (repo *BlogRepository) DeleteBlog(id int) error {
	stmt, err := repo.DB.Prepare("DELETE FROM blogs WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
