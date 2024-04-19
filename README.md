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
|       |-- handlers
|       |   |-- helpers.go
|       |   |-- post_create_post_handler.go
|       |   |-- post_create_quote_handler.go
|       |   |-- post_create_repost_handler.go
|       |   |-- user_get_feed_handler.go
|       |   `-- user_get_info_handler.go
|       |-- api.go
|       `-- api_test.go
|-- docs
|   |-- img.png
|   `-- postman_collection.json
|-- internal
|   |-- application
|   |   |-- createpost.usecase.go
|   |   |-- createpost.usecase_test.go
|   |   |-- createquotepost.usecase.go
|   |   |-- createquotepost.usecase_test.go
|   |   |-- createrepost.usecase.go
|   |   |-- createrepost.usecase_test.go
|   |   |-- getuserfeed.usecase.go
|   |   |-- getuserfeed.usecase_test.go
|   |   |-- getuserinfo.usecase.go
|   |   `-- getuserinfo.usecase_test.go
|   |-- domain
|   |   |-- post.go
|   |   |-- post_test.go
|   |   |-- types.go
|   |   |-- user.go
|   |   `-- user_test.go
|   `-- infra
|       `-- repository
|           `-- in_memory
|               |-- post.go
|               |-- post_test.go
|               |-- user.go
|               `-- user_test.go
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

16 directories, 37 files

```


## Styleguide 

[uber go style guide](https://github.com/alcir-junior-caju/uber-go-style-guide-pt-br)
