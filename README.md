# Blog
This repo contains a simple application using Go. This project uses a blog implementation as an example.

# Features
Post :
- Get Posts :
    Getting list of posts. Had several filter to search by such as title, content, tag, published date, status. 
- Get Post :
    Getting detailed data of specific post.
- Add Post :
    Adding new post. Will be a draft before later being published.
- Update Post :
    Change content of existing post.
- Delete Post :
    Remove an existing post.
- Publish Post :
    Publish a post.

Tag :
- Get Tags :
    Getting list of tags. Had a filter to search by such as label.
- Get Tag :
    Getting detailed data of specific tag.
- Add Tag :
    Adding new tag.
- Update Tag :
    Change label of existing tag.
- Delete Tag : 
    Remove an existing tag.

# Installation
1. Clone the repository
    ```bash
        git clone https://github.com/rez-a-put/blog.git
    ```
2. Change into project directory
    ```bash
        cd blog
    ```
3. Set up your .env file based on .env.example
4. Set up your vendor folder
    ```bash
        go mod vendor
    ```
5. Install postgres tags for migration if needed
    ```bash
        go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
    ```
6. Run migration files
    ```bash
        migrate -path database/migrations/ -database "postgres://postgres:postgres@localhost:5432/blog?sslmode=disable" -verbose up
    ```

# Run the project
1. Open terminal
2. Go to project folder
3. Build application
    ```bash
        go build
    ```
4. Run application from terminal or run using go command
    ```bash
        ./blog
    ```
    ```bash
        go run main.go
    ```

# Contributing
1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a merge request

# Disclaimer
There are several features that has not been added due to time constraint.
1. Users :
    Should have a User features so can have a user related features such as post created by user can only be read by its own user before being publish.
2. Authentication :
    Should have an Authentication feature for security so not everyone can read or change or even delete other user's post.
3. Validation :
    Should have a validation for several function. One that might related to user such as a user can only read its own unpublished post or maybe tags that has been assigned into post cannot be deleted before it is being free of any posts.
4. Testing :
    Should have a unit testing.
5. Not Deleting Data Directly :
    Removing a post or tag should not really deleting its data, just being hidden from public. So when it was needed maybe for historical purpose, the data can still be taken from database directly. 