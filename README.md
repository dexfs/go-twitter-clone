# Go Simple Twitter Clone

This project stems from my endeavor to learn Go. It mirrors another project created with NestJS, which you can find [here](https://github.com/dexfs/challenge-twitter-clone)

I'm actively learning and applying new knowledge as needed.

Please feel free to comment and offer suggestions for improvement on anything here.



## Database Model
https://dbdiagram.io/d/go-twitter-clone-660b4e8437b7e33fd741027f


![img.png](docs/img.png)


## Endpoints

### Fixed data

```javascript
// user data
const users = [
    {
        id: '4cfe67a9-defc-42b9-8410-cb5086bec2f5',
        username: 'alucard',
    },
    {
        id: 'b8903f77-5d16-4176-890f-f597594ff952',
        username: 'alexander',

    },
    {
        id: '75135a97-46be-405f-8948-0821290ca83e',
        username: 'seras_victoria',
    },
];
```

### Users
___
**GET** /users/*:username*/feed

**GET** /users/*:username*/info

### Posts
___
**POST** /posts - '{"content": "Post Content", "user_id": "uuid"}'

**POST** /posts/repost - '{"content": "Post Content", "user_id": "uuid", "post_id": "UUID"}'

**POST** /posts/quote - '{"quote": "Post Content", "user_id": "uuid", "post_id": "UUID"}'


### Project Structure
```bash
|-- cmd
|   `-- api
|       |-- api.go
|       `-- api_test.go
|-- docs
|   |-- img.png
|   `-- postman_collection.json
|-- internal
|   |-- application
|   |   |-- handlers
|   |   |   |-- http
|   |   |   |   |-- post.go
|   |   |   |   `-- user.go
|   |   |   `-- helpers.go
|   |   `-- usecases
|   |       |-- post_usecases
|   |       |   |-- create_post.go
|   |       |   |-- create_post_test.go
|   |       |   |-- create_quotepost.go
|   |       |   |-- create_quotepost_test.go
|   |       |   |-- create_repost.go
|   |       |   `-- create_repost_test.go
|   |       `-- user_usecases
|   |           |-- getuserfeed_usecase.go
|   |           |-- getuserfeed_usecase_test.go
|   |           |-- getuserinfo_usecase.go
|   |           `-- getuserinfo_usecase_test.go
|   |-- infra
|   |   |-- post_repo
|   |   |   |-- inmemory_repo.go
|   |   |   `-- inmemory_repo_test.go
|   |   `-- user_repo
|   |       |-- inmemory_repo.go
|   |       `-- inmemory_repo_test.go
|   |-- post
|   |   |-- entity.go
|   |   |-- entity_test.go
|   |   `-- types.go
|   `-- user
|       |-- entity.go
|       |-- entity_test.go
|       `-- types.go
|-- pkg
|   |-- database
|   |   |-- basedatabase.go
|   |   `-- inmemory_db.go
|   `-- shared
|       `-- helpers
|           `-- hdates.go
|-- tests
|   `-- mocks
|       `-- mocks.go
|-- go.mod
|-- go.sum
|-- Makefile
`-- README.md

```

## Styleguide 

[uber go style guide](https://github.com/alcir-junior-caju/uber-go-style-guide-pt-br)
