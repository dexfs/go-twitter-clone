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



## Styleguide 

[uber go style guide](https://github.com/alcir-junior-caju/uber-go-style-guide-pt-br)
