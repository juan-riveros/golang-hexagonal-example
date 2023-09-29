> NB: Originally constructed to explain how some design patterns for interns. I wouldn't recommend this type of design for anything but the most enterprise of applications. Additionally the application/core/infrastructure/utils directories are serious anti-patterns for discoverability in the code base.
> 
> Still leaving this here as a good example of hexagonal design and it still is a good place to teach more useful design patterns like the Strategy and Proxy patterns.

# golang-hexagonal-example
Hexagonal Architecture in Go

## Reasoning

Hexagonal Architecture is cool. It lets you break down a design between external user facing interfaces (e.g. CLIs, REST APIs, GUIs, etc), the logic that resolves the core problem that we are addressing with this software, and the underlying technologies required to resolve the core problem (e.g. MySQL for a persistent datastore, Redis for a caching layer, RabbitMQ for a message queue). 

Since it decouples application logic from the underlying tech and at the same times ensures that the UIs are decoupled from the internals of the Core logic, we can mix and match between UI and backend. Additionally, adding new functionality is fairly straight forward since its so removed from the core logic. 

## Debugging

For debugging issues, there are the `cmd` UI and the `mock` DB. 

```
./SceneGenerator -database mock -ui cmd
```

## Deployment

Just compile then copy it out to where ever you want.

```bash
go build -o SceneGenerator
```

## Built With

* [Go](https://golang.org/)
* [SQLite3](https://www.sqlite.org/index.html)


## Authors

* **Juan Riveros** - *Initial work* - [Juan Riveros](https://github.com/juan-riveros)

See also the list of [contributors](https://github.com/your/project/contributors) who participated in this project.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* [Hexagonal Architecture](https://blog.octo.com/en/hexagonal-architecture-three-principles-and-an-implementation-example/)
* [Ports and Adapters Pattern](https://softwarecampament.wordpress.com/portsadapters/)
