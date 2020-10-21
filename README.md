# simplebank
Just a simplebank app based on https://github.com/techschool/simplebank. Kicking some rust (pun intended) off my golang skills ahead of gophercon.

# database design
Based entirely on the lecture:
<iframe width="560" height="315" src='https://dbdiagram.io/embed/5f8ea4363a78976d7b785bdb'> </iframe>

The site dbdiagram.io is certainly interesting. Not sure I agree that creating the database here is more intuitive than raw sql, but the diagram is nice enough I suppose.  

# database init
Pretty standard docker run of postgres. Documenting for the giggles. For what it's worth.

```
➜  simplebank git:(2) ✗ docker pull postgres:13-alpine
...
➜  simplebank git:(2) ✗ docker images
REPOSITORY                                  TAG                 IMAGE ID            CREATED             SIZE
...
postgres                                    13-alpine           f9dc9f9f4f4d        3 weeks ago         160MB
...
➜  simplebank git:(2) ✗ docker run --name=postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:13-alpine
e95862a05016be11951fb627ffabdbbfedce36975efc5e492abf71a9299b66b9
➜  simplebank git:(2) ✗ docker ps
CONTAINER ID        IMAGE                COMMAND                  CREATED             STATUS              PORTS                    NAMES
e95862a05016        postgres:13-alpine   "docker-entrypoint.s…"   4 seconds ago       Up 3 seconds        0.0.0.0:5432->5432/tcp   postgres13
➜  simplebank git:(2) ✗ docker exec -it postgres13 psql -U root                                                                          
psql (13.0)
Type "help" for help.

root=# select now();
             now             
-----------------------------
 2020-10-20 09:22:56.6168+00
(1 row)

root=# exit
➜  simplebank git:(2) ✗ docker logs postgres13
...
```

IntelliJ's (GoLand's) database ui doesn't seem to like using the `root` user, complains about not having a `root` role. Using user=postgres actually works though. I guess postgres sets up a postgres user no matter what? I may look into that more later. Moving on for now.

# Using migrate & make to init and teardown db

Pretty strait forward. This is a fairly rudimentary Makefile, and could really use some improving, but whateves, it's fine for now.
I do feel like migrate is probably underutilizied here. Should probably look into that more too.

# SQL -> CRUD code
Good to know how to use the generator. Didn't make Update or Delete for entry and transaction, as it's not clear how we're protecting this data. Presumably some sort of deletion, or preferably pii fuzzing, should happen when an account gets deleted, but that's a whole can of worms I suspect we're not dealing with here.
Checked authors code, and they are indeed not dealing with it. Interestingly, their transaction listing matches on both the `from_account` and the `to_account`, but with an OR, which seems... peculiar. I have a hard time imagining a use case where that's actually what you want. Leaving mine without filtering for now, may adjust when I better understand the authors hypothetical business domain.
Definitely appreciate the pagination example, that was not SUPER obvious and is a necessity.