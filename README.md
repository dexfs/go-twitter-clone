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
// users data
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

### Project structure
 ```bash
 |-- cmd
|   `-- api
|       |-- api.go
|       `-- api_test.go
|-- docs
|   |-- img.png
|   `-- postman_collection.json
|-- internal
|   |-- infra
|   |   `-- repository
|   |       `-- inmemory
|   |           |-- post_inmemory_impl_repo.go
|   |           |-- post_inmemory_impl_repo_test.go
|   |           |-- user_inmemory_impl_repo.go
|   |           `-- user_inmemory_impl_repo_test.go
|   |-- post
|   |   |-- handler
|   |   |   `-- post.go
|   |   |-- usecase
|   |   |   |-- createpost.usecase.go
|   |   |   |-- createpost.usecase_test.go
|   |   |   |-- createquotepost.usecase.go
|   |   |   |-- createquotepost.usecase_test.go
|   |   |   |-- createrepost.usecase.go
|   |   |   `-- createrepost.usecase_test.go
|   |   |-- post.go
|   |   |-- post_test.go
|   |   `-- types.go
|   `-- user
|       |-- handler
|       |   `-- user.go
|       |-- usecase
|       |   |-- getuserfeed.usecase.go
|       |   |-- getuserfeed.usecase_test.go
|       |   |-- getuserinfo.usecase.go
|       |   `-- getuserinfo.usecase_test.go
|       |-- types.go
|       |-- user.go
|       `-- user_test.go
|-- mocks
|   `-- mocks.go
|-- pkg
|   |-- database
|   |   |-- basedatabase.go
|   |   `-- inmemory_db.go
|   `-- helpers
|       |-- hdates.go
|       `-- helpers.go
|-- tests
|-- go.mod
|-- go.sum
|-- Makefile
`-- README.md

18 directories, 35 files

 ```

## Styleguide 

[uber go style guide](https://github.com/alcir-junior-caju/uber-go-style-guide-pt-br)
