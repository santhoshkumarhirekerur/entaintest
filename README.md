### ### NOTE:

I am new to golang.  Used my Java backgroud technique to resolve all the problems. Its great learning experience for me.

I have created indivisual branch for every tasks. branch links are below.  Every task details is documented in 'readme' folder. read me file having all details about the task. The file also contains manual integration test details.  
 
 

BRANCH LINKS:

https://github.com/santhoshkumarhirekerur/entaintest/tree/task-1
https://github.com/santhoshkumarhirekerur/entaintest/tree/task-2
https://github.com/santhoshkumarhirekerur/entaintest/tree/task-3
https://github.com/santhoshkumarhirekerur/entaintest/tree/task-5
https://github.com/santhoshkumarhirekerur/entaintest/tree/task-6
https://github.com/santhoshkumarhirekerur/entaintest/tree/task-4

PULL REQUEST
------------------

https://github.com/santhoshkumarhirekerur/entaintest/pulls

### Changes/Updates Required

Ideally, we'd like to see you push this repository up to Github/Gitlab/Bitbucket and lodge a Pull/Merge Request for each of the following tasks. This means, we'd end up with 7x PR's in total. Each PR should target the previous so they build on one-another. This will allow us to review your changes as best as we possibly can.

... and now to the test! Please complete the following tasks.

1. Add another filter to the existing RPC, so we can call `ListRaces` asking for races that are visible only.
2. We'd like to see the races returned, ordered by their `advertised_start_time`
3. Our races require a new `status` field that is derived based on their `advertised_start_time`'s. The status is simply, `OPEN` or `CLOSED`. All races that have an `advertised_start_time` in the passed should reflect `CLOSED`. 
4. Using the concept of [Resource Views](https://cloud.google.com/apis/design/design_patterns#resource_view), introduce a view that can allow us to return the race name and number only when listing races.
5. Introduce a new RPC, that allows us to fetch a single race by its ID.
6. Create a `sports` service that, for sake of simplicity, implements a similar API to racing, lets call this `ListEvents`. We'll leave it up to you to determine what you might think a sports event is made up off, but it should at minimum have an ID, a name and an advertised start. Note, this should be a separate service, not bolted onto the existing racing service.
7. Document and comment! Please make sure your work is appropriately documented/commented, so fellow developers know whats going on.

**Note:**

> To aid in proto generation following any changes, you can run `go generate ./...` from `api` and `racing` directories.

### Good Reading

- [Protocol Buffers](https://developers.google.com/protocol-buffers)
- [Google API Design](https://cloud.google.com/apis/design)
- [Go Modules](https://golang.org/ref/mod)
- [Ubers Go Style Guide](https://github.com/uber-go/guide/blob/2910ce2e11d0e0cba2cece2c60ae45e3a984ffe5/style.md)
