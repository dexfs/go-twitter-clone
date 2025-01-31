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
        id: '0194bd04-66e2-7cd8-b3d9-66eda709f2ee',
        username: 'alucard',
    },
    {
        id: '0194bd04-8eac-7e70-97cd-c526cdda3d6a',
        username: 'alexander',

    },
    {
        id: '0194bdb1-0588-7181-809e-a825badac714',
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